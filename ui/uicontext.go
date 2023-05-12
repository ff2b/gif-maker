package ui

import (
	"log"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

type UIContext struct {
	win          fyne.Window
	view         IView
	tempGIF      fyne.URI
	encodeStatus chan string
}

func NewUIContext(win *fyne.Window) *UIContext {
	ctx := &UIContext{win: (*win), view: nil, tempGIF: nil, encodeStatus: make(chan string)}
	ctx.SetState(NewMainView(ctx))
	ctx.win.SetMainMenu(GetMainMenu())
	log.Println("App initialized. Now state:", ctx.view.GetViewType())
	ctx.win.Show()
	ctx.win.Content().Refresh()
	return ctx
}

func (u *UIContext) SetState(view IView) {
	u.view = view
	u.view.ShowUI()
	log.Printf("App State Changed. Transition to: %s", u.view.GetViewType())
}

func (u *UIContext) UpdateGIFEncodeStatus(status string) {
	// ***Status Description***
	// "encoding", It is still GIF Encoding.
	// "complete", GIF Encoding has completed.
	if status != "encoding" && status != "complete" {
		log.Fatal("App internal error, Invalid status specified: ", status)
	}
	u.encodeStatus <- status
}

func (u *UIContext) SetTempGIF(path string) {
	parsedURI, error := storage.ParseURI("file:" + path)
	if error != nil {
		log.Fatalf("ERROR: ParseURI(%s) failed.", path)
	}
	u.tempGIF = parsedURI
}

func (u *UIContext) ClearTempData() {
	tmp := "tmp"
	files, err := os.ReadDir(tmp)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		err = os.Remove(filepath.Join(tmp, file.Name()))
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Deleted tmp GIF files.")
}
