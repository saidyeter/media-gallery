package model

type Directories struct {
	DirList []Dir
}

type Content struct {
	Name       string
	ThumbPath  string
	ActualPath string
	IsDir      bool
}

type ContentsResponse struct {
	Contents []Content
	Next     string
}

type Dir struct {
	Id   string
	Path string
}

type VarsConfig struct {
	Dirs []string `json:"dirs"`
}
