package app

import (
	"bufio"
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

	"github.com/kordiseps/media-gallery/internal/content"
	"github.com/kordiseps/media-gallery/internal/util"
	"github.com/kordiseps/media-gallery/model"
)

var rootDirs []string

var contentservice content.ContentService

type App struct {
	Router *mux.Router
}

func (a *App) Init() {

	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/api", home)
	a.Router.HandleFunc("/content", rootContent).Methods("GET", "OPTIONS")
	a.Router.HandleFunc("/content/{dir}", contents).Methods("GET", "OPTIONS")
	a.Router.HandleFunc("/file/{dir}", file).Methods("GET", "OPTIONS")
	a.Router.HandleFunc("/test/{dir}", test).Methods("GET", "OPTIONS")
	a.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("../client/")))
	contentservice = content.ContentService{}
	loadDirs()
}

func loadDirs() {

	vars, err := readFileLineByLine("dirs.txt")
	if err != nil {
		vars, err = readFileLineByLine("../../dirs.txt")
	}
	if err == nil {
		for _, val := range vars {
			if contentservice.FolderExists(val) {
				rootDirs = append(rootDirs, val)
			}
		}
	}

	if contentservice.FolderExists("../content") {
		content := contentservice.DirsFromDir("../content")
		for _, val := range content {
			decoded, err := base64.StdEncoding.DecodeString(val.ActualPath)
			if err != nil {
				continue
			}
			rootDirs = append(rootDirs, string(decoded))
		}
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

func contents(w http.ResponseWriter, r *http.Request) {

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
	intS := util.ToIntSafely(s)
	intE := util.ToIntSafely(e)
	response := contentservice.FilesFromDir(dir, intS, intE)
	response.Next = folderPath + response.Next
	jsonResponse(w, 200, response)
}

func file(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	path := vars["dir"]

	decoded, err := base64.StdEncoding.DecodeString(path)
	if err != nil {
		fmt.Println("read file error", err)
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
		fmt.Println("read file error", err)
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

func getFileContentType(out *os.File) (string, error) {
	//https://golangcode.com/get-the-content-type-of-file/

	fi, err := out.Stat()
	if err != nil {
		// Could not obtain stat, handle error
		return "", err
	}

	size := fi.Size()
	if size > 512 {
		size = 512
	}

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, size)

	_, err = out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)
	out.Seek(0, 0)
	return contentType, nil
}

// read file line by line
func readFileLineByLine(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func test(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := vars["dir"]

	decoded, err := base64.StdEncoding.DecodeString(path)
	if err != nil {
		fmt.Println("read file error", err)
		jsonResponse(w, 200, err)
		return
	}
	decodedValue := string(decoded)

	contentservice.Test((decodedValue))

	jsonResponse(w, 200, `test run`)

}
