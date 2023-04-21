package ui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const (
	CONFIRM_TXT   = "OK"
	DISMISS_TXT   = "Cancel"
	CONFIRM_WIDTH = float32(250)
	CONFIRM_HIGHT = float32(250)
)

type ConfirmView struct {
	ctx    *UIContext
	events map[EventType]func()
}

func NewConfirmView(ctx *UIContext) *ConfirmView {
	view := &ConfirmView{ctx: ctx, events: nil}
	view.events = map[EventType]func(){
		"ok":     view.onOK,
		"cancel": view.onCancel,
	}
	return view
}

func (v *ConfirmView) ShowUI() {
	workfolder := GetWorkFolder()
	selectedItems := workfolder.GetSelectedURIs()
	msg := fmt.Sprintf("%d/%d files Selected. Create GIF file?", len(selectedItems), len(workfolder.UriList))
	for i, isSelect := range workfolder.IsSelectedFlags {
		if isSelect {
			msg = fmt.Sprintf("%s\n%s", msg, workfolder.UriList[i].Name())
		}
	}
	content := container.NewVScroll(widget.NewLabel(msg))
	content.SetMinSize(fyne.NewSize(CONFIRM_WIDTH, CONFIRM_HIGHT))
	dialog := dialog.NewCustomConfirm("Confirm", CONFIRM_TXT, DISMISS_TXT, content, func(result bool) {
		log.Printf("confirm closed: %v", result)
		if result {
			On("ok", v.events)
			return
		}
		On("cancel", v.events)
	}, v.ctx.win)

	dialog.Show()
}

func (v *ConfirmView) GetViewType() ViewType {
	return "confirm"
}

func (v *ConfirmView) Refresh() {
	// Nothing to do
}

func (v *ConfirmView) createComponents() *fyne.Container {
	return nil
}

func (v *ConfirmView) onOK() {
	log.Print("OK clicked, TODO: exec GIF create function")
	v.ctx.SetState(NewMainView(v.ctx))
}

func (v *ConfirmView) onCancel() {
	v.ctx.SetState(NewMainView(v.ctx))
}
