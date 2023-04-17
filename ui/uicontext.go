package ui

import (
	"log"

	"fyne.io/fyne/v2"
)

type UIContext struct {
	win  fyne.Window
	view IView
}

func NewUIContext(win *fyne.Window) *UIContext {
	ctx := &UIContext{win: (*win), view: nil}
	ctx.SetState(NewMainView(ctx))
	log.Println("App initialized. Now state:", ctx.view.GetViewType())
	ctx.win.Show()
	ctx.win.Content().Refresh()
	return ctx
}

func (u *UIContext) SetState(view IView) {
	u.view = view
	u.view.ShowUI()
	log.Println("App State Changed. Transition to:", u.view.GetViewType())
}
