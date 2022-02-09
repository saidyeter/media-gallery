package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/image/draw"
)

var allowedDirs []Dir

// = []Dir{
// 	{Id: "1", Path: "/Users/saidyeter/Pictures/meyve"},
// 	{Id: "2", Path: "/Users/saidyeter/Pictures/kus"},
// 	{Id: "3", Path: "/Users/saidyeter/Pictures/cicek"},
// 	{Id: "4", Path: "C:/Users/said.yeter/Desktop/fotolar"},
// }

type App struct {
	Router *mux.Router
}

type Directories struct {
	DirList []Dir
}

type File struct {
	Name   string
	Path   string
	Thumb  string
	Actual string
}

type FilesResponse struct {
	Files []File
	Start int
	End   int
}

type Dir struct {
	Id   string
	Path string
}

type VarsConfig struct {
	Dirs []string `json:"dirs"`
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

	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/", home)
	a.Router.HandleFunc("/dirs", dirs)
	a.Router.HandleFunc("/files/{dir}", files)
	a.Router.HandleFunc("/file/{dir}", file)
}

func (a *App) Run(addr string) {
	fmt.Println("listening on " + addr)
	http.ListenAndServe(addr, a.Router)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
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
	vars := mux.Vars(r)
	id := vars["dir"]

	decodedValue, _ := url.QueryUnescape(id)

	file := File{
		Name:   decodedValue,
		Path:   decodedValue,
		Actual: getBase64(decodedValue),
	}
	respondWithJSON(w, 200, file)
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
