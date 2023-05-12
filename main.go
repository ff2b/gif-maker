package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/ff2b/gif-maker/ui"
)

const (
	APP_TITLE = "** GIF file maker **"
	APP_WIDTH = float32(640)
	APP_HIGHT = float32(480)
)

type GifMakerGUI struct {
	app     fyne.App
	context ui.UIContext
}

// Initialize ImagicGUI struct.
func (g *GifMakerGUI) init() {
	g.app = app.New()
	win := g.app.NewWindow(APP_TITLE)
	win.Resize(fyne.NewSize(APP_WIDTH, APP_HIGHT))
	win.SetMaster()
	ctx := ui.NewUIContext(&win)

	g.context = *ctx
	win.SetOnClosed(g.context.ClearTempData)
	g.app.Run()
}

// Application initialize.
func initializeApp() {
	// create main UI. returned &fyne.Window
	gui := new(GifMakerGUI)
	gui.init()
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("App Initialize")
	initializeApp()
	log.Println("App Exited")
}
