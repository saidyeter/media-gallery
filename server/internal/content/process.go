package content

import (
	"fmt"
	_ "image/jpeg"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kordiseps/media-gallery/internal/util"
)

func (f *ContentService) CreateThumbnail(path string, outputPath string) {

	if !f.FileExists(path) {
		fmt.Println("File doesn't exists", path)
		return
	}

	if !f.IsImageFile(path) && !f.IsVideoFile(path) {
		fmt.Println("File isn't video or image", path)
		return
	}

	if !createParentFolder(outputPath) {
		fmt.Println("Couldn't create parent folder", outputPath)
		return
	}

	// input, err := os.Open(path)
	// if err != nil {
	// 	fmt.Println("createThumbnail Open input error:", err)
	// }
	// defer input.Close()

	// output, err := os.Create(outputPath)
	// if err != nil {
	// 	fmt.Println("createThumbnail, Create output error:", err)
	// 	return
	// }
	// defer output.Close()

	x, y := getCorrectDimensions(getFileDimensions(path))
	strx := strconv.Itoa(x)
	stry := strconv.Itoa(y)
	if f.IsImageFile(path) {
		processImg(path, outputPath, strx, stry)
	} else {
		processVideo(path, outputPath, strx, stry)
	}

}

func processImg(path string, outputPath string, x string, y string) {

	parent := filepath.Dir(outputPath)

	err := os.MkdirAll(parent, os.ModePerm)
	if err != nil {
		fmt.Println("processImg error:", err)
	}

	cmd := exec.Command("ffmpeg", "-i", path, "-vf", "scale="+string(x)+":"+string(y), outputPath)

	err = cmd.Run()
	if err != nil {
		fmt.Println("processImg error:", err, cmd.Args)
	}

}

func processVideo(path string, outputPath string, x string, y string) {
	parent := filepath.Dir(outputPath)

	err := os.MkdirAll(parent, os.ModePerm)
	if err != nil {
		fmt.Println("createThumbnail error:", err)
	}

	cmd := exec.Command("ffmpeg", "-t", "3", "-i", path, "-vf", "scale="+string(x)+":"+string(y), outputPath)

	err = cmd.Run()
	if err != nil {
		cmd = exec.Command("ffmpeg", "-i", path, "-vf", "scale="+string(x)+":"+string(y), outputPath)
		err = cmd.Run()
		if err != nil {
			fmt.Println("createThumbnail error:", err, cmd.Args)
		}

	}

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

func (f *ContentService) Test(input string) {

	out, err := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=width,height", "-of", "csv=s=x:p=0", input).Output()
	if err != nil {
		log.Fatal(err)
	} else {
		split := strings.Split(string(out), "x")
		fmt.Println(split[0], split[1])
	}

}

func getFileDimensions(path string) (int, int) {
	out, err := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=width,height", "-of", "csv=s=x:p=0", path).Output()
	if err != nil {
		log.Fatal(err)
	} else {
		outstr := strings.ReplaceAll(string(out), "\n", "")
		outstr = strings.ReplaceAll(outstr, "\r", "")
		outstr = strings.ReplaceAll(outstr, " ", "")
		split := strings.Split(outstr, "x")
		x := util.ToIntSafely(split[0])
		y := util.ToIntSafely(split[1])
		return x, y
	}
	return 0, 0
}

func getCorrectDimensions(x int, y int) (int, int) {

	xBigger := x > y

	if x > 600 || y > 600 {
		if xBigger {
			y = 600 * y / x
			x = 600
		} else {
			x = 600 * x / y
			y = 600
		}
	}
	return x, y
}
