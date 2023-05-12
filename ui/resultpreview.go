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
	destPath fyne.URI
}

func NewResultPreView(ctx *UIContext) *ResultPreView {
	view := &ResultPreView{ctx: ctx, events: nil}
	view.events = map[EventType]func(){
		"play":                   view.onStartGIF,
		"stop":                   view.onStopGIF,
		"back":                   view.onBackMain,
		"save":                   view.onSave,
		"save-success":           view.onSaveSuccess,
		"invalid-file-extension": view.onInvalidFileExtension,
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

	playTool := container.NewCenter(
		widget.NewToolbar(
			widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
				On("play", v.events)
			}),
			widget.NewToolbarAction(theme.MediaStopIcon(), func() {
				On("stop", v.events)
			}),
		),
	)

	backMainButton := widget.NewButtonWithIcon("Back Home", theme.NavigateBackIcon(), func() {
		On("back", v.events)
	})
	saveButton := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(),
		func() {
			On("save", v.events)
		})

	return container.NewVBox(
		widget.NewLabelWithStyle("Preview", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		gif,
		playTool,
		container.NewHBox(backMainButton, layout.NewSpacer(), saveButton),
	)
}

// Event handler functions
func (v *ResultPreView) onStartGIF() {
	v.gif.Start()
	log.Print("start clicked")
}

func (v *ResultPreView) onStopGIF() {
	v.gif.Stop()
	log.Print("stop clicked")
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

		v.destPath = chosen.URI()

		// .gif prefix validation
		if v.destPath.Extension() == "" {
			// chosen was unnessesary file open. So delete that file.
			err := storage.Delete(chosen.URI())
			if err != nil {
				log.Fatal(err)
			}
			v.destPath = storage.NewFileURI(fmt.Sprint(v.destPath.Path(), ".gif"))
		}

		if err != nil {
			log.Fatal(err)
		}

		On("save-success", v.events)
	}, v.ctx.win)

	fileSave.Show()
	log.Print("save clicked")
}

func (v *ResultPreView) onSaveSuccess() {
	err := storage.Copy(v.ctx.tempGIF, v.destPath)
	if err != nil {
		log.Println(err)
	}
	dialog.ShowInformation("Save GIF file was successed", fmt.Sprint("Save to\n", v.destPath.Path()), v.ctx.win)
}

func (v *ResultPreView) onInvalidFileExtension() {
	dialog.ShowInformation("Caution", fmt.Sprintf("Please specify .gif prefix. %s is invalid filename", v.destPath.Name()), v.ctx.win)
}
