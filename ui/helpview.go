package ui

import (
	"os"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const (
	HELP_WIDTH  = float32(400)
	HELP_HEIGHT = float32(500)
)

type HelpView struct {
	subwin fyne.Window
}

func NewHelpView() *HelpView {
	return &HelpView{subwin: fyne.CurrentApp().NewWindow("Help")}
}

func (v *HelpView) ShowUI() {
	v.subwin.Resize(fyne.NewSize(HELP_WIDTH, HELP_HEIGHT))
	v.subwin.SetContent(container.NewScroll(v.createComponents()))
	v.subwin.CenterOnScreen()
	v.subwin.Show()
}

func (v *HelpView) GetViewType() ViewType {
	return "help"
}

func (v *HelpView) Refresh() {
	// Nothing to do
}

func (v *HelpView) createComponents() *fyne.Container {
	return container.NewVBox(
		container.NewPadded(
			widget.NewRichTextFromMarkdown(v.loadHelp()),
		),
	)
}

func (v *HelpView) loadHelp() string {
	markdowns, err := os.ReadFile(path.Join("readme.md"))
	if err != nil {
		dialog.ShowError(err, v.subwin)
	}
	return string(markdowns)
}
