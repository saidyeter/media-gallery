package main

import "github.com/kordiseps/media-gallery/internal/app"

func main() {

	app := app.App{}
	app.Init()
	app.Run(":8080")

}
