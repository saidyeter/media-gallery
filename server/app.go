package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/image/draw"
)

var allowedDirs []Dir

type App struct {
	Router *mux.Router
}

func (a *App) Init() {

	var varsConfig VarsConfig
	vars, err := os.ReadFile("vars.json")
	if err != nil {
		fmt.Println("could not read vars.json :", err)
	}
	err = json.Unmarshal(vars, &varsConfig)
	if err != nil {
		fmt.Println("could not deserialize vars.json :", err)
	}
	for k, v := range varsConfig.Dirs {

		id := strconv.Itoa(k)
		allowedDirs = append(allowedDirs, Dir{
			Id:   id,
			Path: v,
		})
	}

	//path dönülecek  image/jpeg ile
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/", home)
	a.Router.HandleFunc("/dirs", dirs)
	a.Router.HandleFunc("/files/{dir}", files)
	a.Router.HandleFunc("/file/{dir}", file).Methods("GET", "OPTIONS")
}

func (a *App) Run(addr string) {
	fmt.Println("listening on " + addr)

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling
	// log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(router)))
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(originsOk, headersOk, methodsOk)(a.Router)))

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func downloadResponse(w http.ResponseWriter, f *os.File) {
	contentType, err := GetFileContentType(f)
	if err != nil {
		panic(err)
	}

	f.Seek(0, 0)

	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}

	// Openfile, err := os.Open(Filename)
	// FileSize := strconv.FormatInt(FileStat.Size(), 10)
	// FileContentType := http.DetectContentType(FileHeader)
	//Send the headers before sending the file
	// writer.Header().Set("Content-Disposition", "attachment; filename="+Filename)
	// writer.Header().Set("Content-Type", FileContentType)
	// writer.Header().Set("Content-Length", FileSize)
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(f.Name()))
	w.Header().Set("Content-Length", strconv.FormatInt(fi.Size(), 10))
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(200)

	_, err = io.Copy(w, f)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetFileContentType(out *os.File) (string, error) {
	//https://golangcode.com/get-the-content-type-of-file/

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func home(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, `
- "/" : Home endpoint. Returns other endpoint details.
- "/dirs" : Directories endpoint. Returns available root directories that specified in "vars.json". Each root directory returns with numeric id. That id can be used in "/files/{id}" as id to retrieve contents in the directory.
- "/files/{id}" : Files endpoint. Returns directory paths and image file paths, also thumbnail image content as base64. That file path can be used in "/file/{path}" as path to retrieve actual image content. This endpoint also has paging functionality. To do that, start index("s") and end index("e") need to be specified as query parameters. Eg. "http://localhost:8080/files/3?s=3&e=5". Otherwise, the endpoint will return from 0 (zero) to limit.
- "/file/{path}" : File endpoint. Returns actual image content as base64.
`)

}
func dirs(w http.ResponseWriter, r *http.Request) {

	direstoryList := Directories{
		DirList: allowedDirs,
	}
	respondWithJSON(w, 200, direstoryList)
}

func files(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["dir"]

	v := r.URL.Query()

	s := v.Get("s")
	e := v.Get("e")

	intS, err := strconv.Atoi(s)
	if err != nil {
		intS = 0
	}
	intE, err := strconv.Atoi(e)
	if err != nil {
		intE = 0
	}

	var folderPath string
	for i := range allowedDirs {
		if allowedDirs[i].Id == id {
			folderPath = allowedDirs[i].Path
		}
	}

	respondWithJSON(w, 200, filesFromDir(folderPath, intS, intE))
}

func file(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	vars := mux.Vars(r)
	path := vars["dir"]

	decodedValue, _ := url.QueryUnescape(path)

	input, err := os.Open(decodedValue)
	if err != nil {
		fmt.Println("read file error:", err)
	}
	defer input.Close()

	downloadResponse(w, input)
}

func filesFromDir(dir string, start int, end int) FilesResponse {

	limit := 5
	var files []File
	if end-start > limit || end <= start {
		end = start + limit
	}
	counter := 0
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		if strings.Contains(path, ".DS_Store") {
			return nil
		}

		if start+counter <= end {

			tempPath := getTempPath(path)
			createThumbnailToTemp(path, tempPath)

			files = append(files, File{
				Name:  info.Name(),
				Path:  path,
				Thumb: getBase64(tempPath),
			})
			counter++
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return FilesResponse{
		Files: files,
		Start: start,
		End:   end,
	}
}

func getBase64(path string) string {

	f, err := os.Open(path)
	if err != nil {
		fmt.Println("getBase64 error:", err)
		return err.Error()
	}
	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	return encoded
}
func getTempPath(path string) string {

	hiearchy := path
	if strings.Contains(hiearchy, ":") {
		splitted := strings.Split(hiearchy, ":")
		hiearchy = splitted[len(splitted)-1]
	}
	outputPath := filepath.Join(os.TempDir(), "media-gallery", hiearchy) + "_resized.png"
	return outputPath
}

func createThumbnailToTemp(path string, outputPath string) {
	parent := filepath.Dir(outputPath)

	err := os.MkdirAll(parent, os.ModePerm)
	if err != nil {
		fmt.Println("createThumbnail error:", err)
	}

	input, err := os.Open(path)
	if err != nil {
		fmt.Println("createThumbnail error:", err)
	}
	defer input.Close()

	output, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("createThumbnail error:", err)
		return
	}
	defer output.Close()

	src, _, err := image.Decode(input)
	if err != nil {
		fmt.Println("createThumbnail error:", err)
		return
	}

	currentX := src.Bounds().Max.X
	currentY := src.Bounds().Max.Y
	xBigger := currentX > currentY

	if currentX < 600 && currentY < 600 {
		png.Encode(output, src)
	} else {
		if xBigger {
			currentY = 600 * currentY / currentX
			currentX = 600
		} else {
			currentX = 600 * currentX / currentY
			currentY = 600
		}

		expectedSize := image.NewRGBA(image.Rect(0, 0, currentX, currentY))

		draw.NearestNeighbor.Scale(expectedSize, expectedSize.Rect, src, src.Bounds(), draw.Over, nil)

		png.Encode(output, expectedSize)
	}

}
