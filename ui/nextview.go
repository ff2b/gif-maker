package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type NextView struct {
	ctx    *UIContext
	events map[EventType]func()
}

func NewNextView(ctx *UIContext) *NextView {
	view := &NextView{ctx: ctx, events: nil}
	// Register events
	view.events = map[EventType]func(){
		"back": view.onBack,
	}
	return view
}

func (v *NextView) ShowUI() {
	v.ctx.win.SetContent(v.createComponents())
	v.Refresh()
}

func (v *NextView) On(e EventType) {
	// Execute event function which is mapped EventType.
	if f, ok := v.events[e]; !ok {
		log.Fatal("Error: Invalid event fired.", e, ok)
	} else {
		log.Println("Event hundled.", e)
		f()
	}
}

func (v *NextView) GetViewType() ViewType {
	return "next-demo"
}

func (v *NextView) Refresh() {
	v.ctx.win.SetContent(v.createComponents())
	v.ctx.win.Content().Refresh()
}

func (v *NextView) createComponents() *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Next view has transitioned. Button click will back Mainview."),
		widget.NewButton("Back", func() { v.On("back") }),
	)
}

func (v *NextView) next() {
	v.ctx.SetState(NewMainView(v.ctx))
}

func (v *NextView) onBack() {
	v.next()
}
