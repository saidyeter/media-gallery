package app

import (
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "golang.org/x/image/webp"

	"golang.org/x/image/draw"

	"github.com/kordiseps/media-gallery/model"
)

var supportedImageFileExtensions []string = []string{".png", ".jpg", ".jpeg", ".webp"}

func getFileContentType(out *os.File) (string, error) {
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
	out.Seek(0, 0)
	return contentType, nil
}

func filesFromDir(dir string, start int, end int) model.FilesResponse {

	limit := 5
	var files []model.File
	if end-start > limit || end <= start {
		end = start + limit
	}
	index := 0
	dir = strings.ReplaceAll(dir, "\\", "/")
	dir = filepath.FromSlash(dir)

	doesExist := false

	_, err := os.Stat(dir)

	if err != nil {
		st1 := !os.IsNotExist(err)
		st2 := os.IsExist(err)
		doesExist = st1
		fmt.Println(dir)
		fmt.Println(st1)
		fmt.Println(st2)

	} else {
		doesExist = true
	}

	if doesExist {

		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

			if strings.Contains(path, ".DS_Store") {
				return nil
			}

			if path == dir {
				return nil
			}
			encodedPath := base64.StdEncoding.EncodeToString([]byte(path))

			if start == 0 && info.IsDir() {
				files = append(files, model.File{
					Name:       filepath.Base(dir),
					ActualPath: encodedPath,
					ThumbPath:  "",
					IsDir:      true,
				})
				return nil
			}

			if start < index && index <= end {

				tempPath := getTempPath(path)
				encodedTempPath := base64.StdEncoding.EncodeToString([]byte(tempPath))

				//https://stackoverflow.com/a/12518877
				_, err := os.Stat(tempPath)
				if err != nil && errors.Is(err, os.ErrNotExist) {
					// path/to/whatever does *not* exist
					createThumbnailToTemp(path, tempPath)
				}

				files = append(files, model.File{
					Name:       info.Name(),
					ActualPath: encodedPath,
					ThumbPath:  encodedTempPath,
					IsDir:      false,
				})
			}

			if !info.IsDir() {
				index++
			}

			return nil
		})
		if err != nil {
			panic(err)
		}
	}
	next := "?s=" + strconv.Itoa(end) + "&&e=" + strconv.Itoa(end+5)
	return model.FilesResponse{
		Files: files,
		Next:  next,
	}
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
func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val || item == strings.ToLower(val) {
			return true
		}
	}
	return false
}

func createThumbnailToTemp(path string, outputPath string) {

	ext := filepath.Ext(path)

	if !find(supportedImageFileExtensions, ext) {
		return
	}

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
