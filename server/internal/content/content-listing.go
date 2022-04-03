package content

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kordiseps/media-gallery/model"
)

var limit int = 15

func (f *ContentService) FilesFromDir(dir string, start int, end int) model.ContentsResponse {
	var contents []model.Content
	if end-start > limit || end <= start {
		end = start + limit
	}
	index := 0
	dir = strings.ReplaceAll(dir, "\\", "/")
	dir = filepath.FromSlash(dir)

	if f.FileExists(dir) {
		if start == 0 {
			dirs := f.DirsFromDir(dir)
			contents = append(contents, dirs...)
		}
		contentsFromDir, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}

		for _, content := range contentsFromDir {

			path := filepath.Join(dir, content.Name())
			if path == dir || strings.Contains(path, ".DS_Store") || strings.Contains(path, ".localized") {
				continue
			}

			if content.IsDir() {
				continue
			} else {
				index++
			}

			if start < index && index <= end {

				encodedPath := base64.StdEncoding.EncodeToString([]byte(path))
				tempPath := f.GetTempPath(path)
				encodedTempPath := base64.StdEncoding.EncodeToString([]byte(tempPath))

				//https://stackoverflow.com/a/12518877
				_, err := os.Stat(tempPath)
				if err != nil && errors.Is(err, os.ErrNotExist) {
					// path/to/whatever does *not* exist
					f.CreateThumbnail(path, tempPath)
				}

				contents = append(contents, model.Content{
					Name:       filepath.Base(path),
					ActualPath: encodedPath,
					ThumbPath:  encodedTempPath,
					IsDir:      false,
				})
			}
		}
	}
	next := "?s=" + strconv.Itoa(end) + "&&e=" + strconv.Itoa(end+limit)
	return model.ContentsResponse{
		Contents: contents,
		Next:     next,
	}
}

func (f *ContentService) DirsFromDir(dir string) []model.Content {

	var dirs []model.Content
	dir = strings.ReplaceAll(dir, "\\", "/")
	dir = filepath.FromSlash(dir)

	contents, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return []model.Content{}
	}

	for _, content := range contents {
		// fmt.Println(content.Name(), content.IsDir())
		path := filepath.Join(dir, content.Name())
		if strings.Contains(path, ".DS_Store") || strings.Contains(path, ".localized") {
			continue
		}

		if path == dir {
			continue
		}

		if content.IsDir() {
			encodedPath := base64.StdEncoding.EncodeToString([]byte(path))
			dirs = append(dirs, model.Content{
				Name:       filepath.Base(path),
				ActualPath: encodedPath,
				ThumbPath:  "",
				IsDir:      true,
			})
			continue
		}
	}
	return dirs
}
