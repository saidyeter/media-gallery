package content

import (
	"os"
	"path/filepath"
	"strings"
)

func (f *ContentService) IsVideoFile(filePath string) bool {
	ext := filepath.Ext(filePath)
	return doesExistInCollection(supportedVideoFileExtensions, ext)
}
func (f *ContentService) IsImageFile(filePath string) bool {
	ext := filepath.Ext(filePath)
	return doesExistInCollection(supportedImageFileExtensions, ext)
}

func (f *ContentService) GetTempPath(path string) string {

	hiearchy := path
	if strings.Contains(hiearchy, ":") {
		splitted := strings.Split(hiearchy, ":")
		hiearchy = splitted[len(splitted)-1]
	}
	outputPath := filepath.Join(os.TempDir(), "media-gallery", hiearchy) + "_resized.png"
	return outputPath
}

func doesExistInCollection(slice []string, val string) bool {
	for _, item := range slice {
		if item == val || item == strings.ToLower(val) {
			return true
		}
	}
	return false
}

func (f *ContentService) FileExists(filePath string) bool {
	return doesExist(filePath)
}
func (f *ContentService) FolderExists(folderPath string) bool {
	return doesExist(folderPath)
}
func doesExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

var supportedVideoFileExtensions []string = []string{".mp4", ".avi", ".mov", ".wmv", ".flv", ".mkv"}
var supportedImageFileExtensions []string = []string{".png", ".jpg", ".jpeg", ".webp"}
