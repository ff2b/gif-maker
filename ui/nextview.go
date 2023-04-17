package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type NextView struct {
	ctx *UIContext
}

func NewNextView(ctx *UIContext) *NextView {
	return &NextView{ctx: ctx}
}

func (v *NextView) ShowUI() {
	v.ctx.win.SetContent(v.createComponents())
	v.Refresh()
}

func (v *NextView) On(e EventType) {

}

func (v *NextView) GetViewType() ViewType {
	return "next-demo"
}

func (v *NextView) Refresh() {
	// clear_layout := container.NewMax(
	// 	canvas.NewRectangle(theme.BackgroundColor()),
	// )
	// v.ctx.win.SetContent(clear_layout)
	v.ctx.win.SetContent(v.createComponents())
	v.ctx.win.Content().Refresh()
}

func (v *NextView) createComponents() *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Next view transitioned. Button click will back Mainview."),
		widget.NewButton("Back", func() { v.next() }),
	)
}

func (v *NextView) next() {
	v.ctx.SetState(&MainView{ctx: v.ctx})
}
