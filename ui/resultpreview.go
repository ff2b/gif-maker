package ui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

type ResultPreView struct {
	ctx      *UIContext
	events   map[EventType]func()
	gif      *xwidget.AnimatedGif
	distPath fyne.URI
}

func NewResultPreView(ctx *UIContext) *ResultPreView {
	view := &ResultPreView{ctx: ctx, events: nil}
	view.events = map[EventType]func(){
		"play":         view.onStartGIF,
		"pause":        view.onPauseGIF,
		"back":         view.onBackMain,
		"save":         view.onSave,
		"save-success": view.onSaveSuccess,
	}
	return view
}

func (v *ResultPreView) ShowUI() {
	v.Refresh()
}

func (v *ResultPreView) GetViewType() ViewType {
	return "result-preview"
}

func (v *ResultPreView) Refresh() {
	v.ctx.win.SetContent(v.createComponents())
}

func (v *ResultPreView) createComponents() *fyne.Container {
	// Preview GIF Image
	gif, err := xwidget.NewAnimatedGif(v.ctx.tempGIF)
	if err != nil {
		log.Printf("NewAnimatedGif failed\n%s", err)
	}
	v.gif = gif
	v.gif.SetMinSize(fyne.NewSize(350, 350))
	v.gif.Start()
	// Buttons: Play, Pause
	playButton := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
		On("play", v.events)
	})
	pauseButton := widget.NewButtonWithIcon("", theme.MediaPauseIcon(), func() {
		On("pause", v.events)
	})
	// Footer buttons: Back, Save GIF
	backMainButton := widget.NewButtonWithIcon("Back Home", theme.NavigateBackIcon(), func() {
		On("back", v.events)
	})
	saveButton := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(),
		func() {
			On("save", v.events)
		})

	return container.NewVBox(
		gif,
		container.NewCenter(container.NewHBox(playButton, pauseButton)),
		container.NewHBox(backMainButton, layout.NewSpacer(), saveButton),
	)
}

// Event handler functions
func (v *ResultPreView) onStartGIF() {
	v.gif.Start()
	log.Print("start clicked")
}

func (v *ResultPreView) onPauseGIF() {
	v.gif.Stop()
	log.Print("pause clicked")
}

func (v *ResultPreView) onBackMain() {
	log.Print("back clicked")
	v.ctx.ClearTempData()
	v.ctx.SetState(NewMainView(v.ctx))
}

func (v *ResultPreView) onSave() {
	fileSave := dialog.NewFileSave(func(chosen fyne.URIWriteCloser, err error) {
		if chosen == nil {
			log.Println("FileSave dialog canceled.")
			return
		}
		if chosen.URI().Extension() == "" {
			v.distPath = storage.NewFileURI(chosen.URI().Path() + ".gif")
		} else {
			v.distPath = chosen.URI()
		}

		if err != nil {
			log.Fatal(err)
		}
		log.Println("dist: ", v.distPath)

		err = storage.Copy(v.ctx.tempGIF, v.distPath)
		if err != nil {
			log.Println(err)
		}

		On("save-success", v.events)
	}, v.ctx.win)

	fileSave.Show()
	log.Print("save clicked")
}

func (v *ResultPreView) onSaveSuccess() {
	dialog.ShowInformation("Save GIF file was successed", fmt.Sprint("Save to\n", v.distPath.Path()), v.ctx.win)
}
