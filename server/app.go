package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

var allowedDirs []Dir = []Dir{
	Dir{Id: "1", Path: "/Users/saidyeter/Pictures/meyve"},
	Dir{Id: "2", Path: "/Users/saidyeter/Pictures/kus"},
	Dir{Id: "3", Path: "/Users/saidyeter/Pictures/cicek"},
}

type App struct {
	Router *mux.Router
}

type Directories struct {
	DirList []Dir
}

type File struct {
	Name  string
	Path  string
	Thumb string
}

type Dir struct {
	Id   string
	Path string
}

func (a *App) Init() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/", home)
	a.Router.HandleFunc("/dirs", dirs)
	a.Router.HandleFunc("/files/{dir}", files)
}

func (a *App) Run(addr string) {
	fmt.Println("listening on " + addr)
	http.ListenAndServe(addr, a.Router)
}

// func Run(addr string) {
// 	fmt.Println("listening on " + addr)
// 	http.ListenAndServe(addr, a.Router)
// }

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func home(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, "welcome to home gallery")
	// fmt.Fprintf(w, "welcome to home gallery")
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
	var folderPath string
	for i := range allowedDirs {
		if allowedDirs[i].Id == id {
			folderPath = allowedDirs[i].Path
		}
	}

	respondWithJSON(w, 200, filesFromDir(folderPath))
}

func filesFromDir(dir string) []File {
	var files []File

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		if strings.Contains(path, ".DS_Store") {
			return nil
		}

		files = append(files, File{
			Name:  info.Name(),
			Path:  path,
			Thumb: path, //encoded,
		})
		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}

func getBase64(path string) string {

	f, _ := os.Open(path)
	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	return encoded
}
