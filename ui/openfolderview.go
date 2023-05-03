package ui

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

type OpenFolderView struct {
	ctx        *UIContext
	events     map[EventType]func()
	chosenURIs []fyne.URI
}

func NewOpenFolderView(ctx *UIContext) *OpenFolderView {
	view := &OpenFolderView{ctx: ctx, events: nil}
	// Register events
	view.events = map[EventType]func(){
		"cancel":   view.onCancel,
		"selected": view.onFolderSelect,
		"closed":   view.onCloseDialog,
	}
	return view
}

func (v *OpenFolderView) ShowUI() {
	// Make dialog.
	dialog := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
		// set model filepath list and binding.
		if err != nil {
			log.Fatal(err)
		}
		if lu == nil {
			log.Println("OpenFolder dialog Canceled")
			On("cancel", v.events)
			return
		}
		v.chosenURIs, err = lu.List()
		if err != nil {
			log.Fatal(err)
		}
		On("selected", v.events)
	}, v.ctx.win)

	dialog.SetOnClosed(func() {
		On("closed", v.events)
	})

	// Set Location to current working directory.
	path, _ := os.Getwd()
	listerForUri, _ := storage.ListerForURI(storage.NewFileURI(path))
	dialog.SetLocation(listerForUri)

	dialog.Show()
}

func (v *OpenFolderView) GetViewType() ViewType {
	return "openfolder"
}

func (v *OpenFolderView) Refresh() {
	// Nothing to do
}

func (v *OpenFolderView) createComponents() *fyne.Container {
	// Don't use because dialog object type is not fyne.Container.
	return nil
}

// Event handler functions
func (v *OpenFolderView) onCancel() {
	v.ctx.SetState(NewMainView(v.ctx))
}

func (v *OpenFolderView) onFolderSelect() {
	_ = NewWorkFolder(v.chosenURIs)
	// log.Printf("Event: Folder selected. %d items generated. %v", len(workFolder.UriList), workFolder.UriList)
	v.ctx.SetState(NewMainView(v.ctx))
}

func (v *OpenFolderView) onCloseDialog() {
	// log.Println("Event: Folder dialog closed.")
	v.ctx.SetState(NewMainView(v.ctx))
}
