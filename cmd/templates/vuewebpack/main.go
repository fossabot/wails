package main

import (
	"github.com/gobuffalo/packr"
	"github.com/wailsapp/wails"
)

func main() {

	assets := packr.NewBox("./frontend/dist")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "Vue Simple",
		JS:     assets.String("app.js"),
		CSS:    assets.String("app.css"),
	})
	app.Bind(newQuotesCollection())
	app.Run()
}
