package content

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"
)

func (f *ContentService) CreateThumbnail(path string, outputPath string) {
	if !f.IsImageFile(path) && !f.IsVideoFile(path) {
		return
	}
	if f.IsImageFile(path) {
		createImgThumbnail(path, outputPath)
	} else {
		createVideoThumbnail(path, outputPath, false)
	}

}
func createImgThumbnail(path string, outputPath string) {

	if !createParentFolder(outputPath) {
		return
	}

	input, err := os.Open(path)
	if err != nil {
		fmt.Println("createThumbnail Open input error:", err)
	}
	defer input.Close()

	output, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("createThumbnail, Create output error:", err)
		return
	}
	defer output.Close()

	src, _, err := image.Decode(input)
	if err != nil {
		fmt.Println("createThumbnail Decode input error:", err)
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

// convert video to gif using ffmpeg
func createVideoThumbnail(path string, outputPath string, all bool) {

}

func createParentFolder(path string) bool {
	parent := filepath.Dir(path)

	err := os.MkdirAll(parent, os.ModePerm)
	if err != nil {
		fmt.Println("create parent folder error:", err)
		return false
	}
	return true
}
