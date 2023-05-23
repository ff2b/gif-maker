package ui

import (
	"os"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type VersionView struct {
	subwin fyne.Window
}

func NewVersionView() *VersionView {
	return &VersionView{subwin: fyne.CurrentApp().NewWindow("Version Infomation")}
}

func (v *VersionView) ShowUI() {
	v.subwin.SetContent(v.createComponents())
	v.subwin.CenterOnScreen()
	v.subwin.Show()
}

func (v *VersionView) GetViewType() ViewType {
	return "version"
}

func (v *VersionView) Refresh() {
	// Nothing to do
}

func (v *VersionView) createComponents() *fyne.Container {
	return container.NewVBox(
		container.NewPadded(
			widget.NewRichTextFromMarkdown(v.loadHelp()),
		),
	)
}

func (v *VersionView) loadHelp() string {
	markdowns, err := os.ReadFile(path.Join("version.md"))
	if err != nil {
		dialog.ShowError(err, v.subwin)
	}
	return string(markdowns)
}
