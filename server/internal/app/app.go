package app

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/kordiseps/media-gallery/model"
)

var rootDirs []string

type App struct {
	Router *mux.Router
}

func (a *App) Init() {

	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/api", home)
	// a.Router.HandleFunc("/dirs", dirs)
	a.Router.HandleFunc("/content", rootContent).Methods("GET", "OPTIONS")
	a.Router.HandleFunc("/content/{dir}", content).Methods("GET", "OPTIONS")
	a.Router.HandleFunc("/file/{dir}", file).Methods("GET", "OPTIONS")
	a.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("../client/")))
	// a.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("../../../client/")))
	loadDirs()
}

func loadDirs() {
	var varsConfig model.VarsConfig

	vars, err := os.ReadFile("vars.json")
	if err != nil {
		vars, err = os.ReadFile("../../vars.json")
		if err != nil {
			fmt.Println("could not read vars.json :", err)
			return
		}
	}
	err = json.Unmarshal(vars, &varsConfig)
	if err != nil {
		fmt.Println("could not deserialize vars.json :", err)
	}
	for _, val := range varsConfig.Dirs {

		if _, err := os.Stat(val); os.IsNotExist(err) {
			continue
		}
		rootDirs = append(rootDirs, val)
	}

}

func (a *App) Run(addr string) {
	fmt.Println("listening on " + addr)

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "OPTIONS"})

	log.Fatal(http.ListenAndServe(addr, handlers.CORS(originsOk, headersOk, methodsOk)(a.Router)))

}

func jsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func downloadResponse(w http.ResponseWriter, f *os.File) {
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(f.Name()))
	contentType, err := getFileContentType(f)
	if err != nil {
		fmt.Println("file type cannot find", err)
	} else {
		w.Header().Set("Content-Type", contentType)
	}

	fi, err := f.Stat()
	if err != nil {
		fmt.Println("file size cannot get", err)
	} else {
		w.Header().Set("Content-Length", strconv.FormatInt(fi.Size(), 10))
	}

	w.WriteHeader(200)
	_, err = io.Copy(w, f)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, 200, `
- '/' : Home endpoint. Returns endpoint details.  
- '/content' : Root directories endpoint. Returns available root directories that specified in 'vars.json'. For content, there is 'actualPath' property.   
- '/content/{dir}' : Content endpoint. Returns directory paths and image file paths, also thumbnail image path. There is 'actualPath' to retrieve actual image content or folder content. This endpoint also has paging functionality. To do that, start index('s') and end index('e') need to be specified as query parameters. Eg. 'http://localhost:8080/files/3?s=3&&e=5'. Otherwise, the endpoint will return from 0 (zero) to limit.
- '/file/{path}' : File endpoint. Returns actual image content as base64.
`)

}

func rootContent(w http.ResponseWriter, r *http.Request) {

	var files []model.Content
	for _, val := range rootDirs {

		if _, err := os.Stat(val); os.IsNotExist(err) {
			continue
		}
		encodedPath := base64.StdEncoding.EncodeToString([]byte(val))

		files = append(files, model.Content{
			Name:       filepath.Base(val),
			ActualPath: encodedPath,
			ThumbPath:  "",
			IsDir:      true,
		})
	}

	response := model.ContentsResponse{}
	response.Contents = files
	jsonResponse(w, 200, response)
}

func content(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	folderPath := vars["dir"]

	if len(folderPath) == 0 {
		rootContent(w, r)
		return
	}
	decoded, err := base64.StdEncoding.DecodeString(folderPath)
	if err != nil {
		jsonResponse(w, 404, err)
	}
	dir := string(decoded)

	if !isPathUnderRoot(dir) {
		fmt.Println("folder cannot find:", dir)
		jsonResponse(w, 404, "")
	}

	v := r.URL.Query()

	s := v.Get("s")
	e := v.Get("e")

	intS, err := strconv.Atoi(s)
	if err != nil {
		intS = 0
	}
	intE, err := strconv.Atoi(e)
	if err != nil {
		intE = 5
	}
	response := filesFromDir(dir, intS, intE)
	response.Next = folderPath + response.Next
	jsonResponse(w, 200, response)
}

func file(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	path := vars["dir"]

	decoded, err := base64.StdEncoding.DecodeString(path)
	if err != nil {
		fmt.Println("read file error")
		jsonResponse(w, 200, err)
		return
	}
	decodedValue := string(decoded)

	if !isPathUnderRoot(decodedValue) {
		fmt.Println("file cannot find:", decodedValue)
		jsonResponse(w, 404, "")
		return
	}

	input, err := os.Open(decodedValue)
	if err != nil {
		fmt.Println("read file error")
		jsonResponse(w, 404, err)
		return
	}
	defer input.Close()

	downloadResponse(w, input)
}

func isPathUnderRoot(path string) bool {
	if strings.HasPrefix(path, os.TempDir()) {
		return true
	}
	for _, val := range rootDirs {
		cleanPath := strings.ReplaceAll(path, "\\", "")
		cleanPath = strings.ReplaceAll(cleanPath, "/", "")
		cleanVal := strings.ReplaceAll(val, "\\", "")
		cleanVal = strings.ReplaceAll(cleanVal, "/", "")
		if strings.HasPrefix(cleanPath, cleanVal) {
			return true
		}
	}
	return false
}
