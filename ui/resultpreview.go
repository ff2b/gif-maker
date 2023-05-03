package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

type ResultPreView struct {
	ctx    *UIContext
	events map[EventType]func()
}

func NewResultPreView(ctx *UIContext) *ResultPreView {
	view := &ResultPreView{ctx: ctx, events: nil}
	view.events = map[EventType]func(){
		"play":   view.onStartGIF,
		"pause":  view.onPauseGIF,
		"replay": view.onReplayGIF,
		"back":   view.onBackMain,
		"save":   view.onSave,
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
	gif.SetMinSize(fyne.NewSize(350, 350))
	gif.Start()
	// Buttons: Play, Pause, Replay
	playButton := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
		On("play", v.events)
	})
	pauseButton := widget.NewButtonWithIcon("", theme.MediaPauseIcon(), func() {
		On("pause", v.events)
	})
	replayButton := widget.NewButtonWithIcon("", theme.MediaReplayIcon(), func() {
		On("replay", v.events)
	})
	// Footer buttons: Back, Save GIF
	backMainButton := widget.NewButtonWithIcon("Back", theme.NavigateBackIcon(), func() {
		On("back", v.events)
	})
	saveButton := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(),
		func() {
			On("save", v.events)
		})

	return container.NewVBox(
		gif,
		container.NewHBox(playButton, pauseButton, replayButton),
		container.NewHBox(backMainButton, saveButton),
	)
}

// Event handler functions
func (v *ResultPreView) onStartGIF() {
	log.Print("start clicked")
}

func (v *ResultPreView) onPauseGIF() {
	log.Print("pause clicked")
}

func (v *ResultPreView) onReplayGIF() {
	log.Print("replay clicked")
}

func (v *ResultPreView) onBackMain() {
	log.Print("back clicked")
	v.ctx.ClearTempData()
	v.ctx.SetState(NewMainView(v.ctx))
}

func (v *ResultPreView) onSave() {
	log.Print("save clicked")
}
