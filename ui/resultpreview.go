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
	"github.com/ff2b/gif-maker/config"
)

const (
	DEFAULT_SAVE_PATH = "data"
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
		"play":         view.onStartGIF,
		"stop":         view.onStopGIF,
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
}

func (v *ResultPreView) onStopGIF() {
	v.gif.Stop()
}

func (v *ResultPreView) onBackMain() {
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
			// fyne.URIWriteCloser create empty file.
			// So delete that unnecessary file, bcause destination path was changed.
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

	// Load conf and set default save destination folder
	lister, err := storage.ListerForURI(loadSavePathConfig())
	if err != nil {
		log.Fatal("Unexpected Error: load config was failed, Default Save Path is invalid uri")
	}
	fileSave.SetLocation(lister)
	fileSave.Show()
}

func (v *ResultPreView) onSaveSuccess() {
	err := storage.Copy(v.ctx.tempGIF, v.destPath)
	if err != nil {
		log.Println(err)
	}
	dialog.ShowInformation("Save GIF file was successed", fmt.Sprint("Save to\n", v.destPath.Path()), v.ctx.win)
}

// private functions
func loadSavePathConfig() fyne.URI {
	conf := config.NewConfig()
	conf.Load()
	if conf.DefaultSavePath != DEFAULT_SAVE_PATH {
		parsed, err := storage.ParseURI(conf.DefaultSavePath)
		if err != nil {
			log.Println(err)
			return storage.NewFileURI(DEFAULT_SAVE_PATH)
		}
		return parsed
	}
	return storage.NewFileURI(DEFAULT_SAVE_PATH)
}
