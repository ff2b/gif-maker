package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MainView struct {
	ctx    *UIContext
	events map[EventType]func()
}

func NewMainView(ctx *UIContext) *MainView {
	event_routes := map[EventType]func(){
		"click": onClick,
	}
	return &MainView{ctx: ctx, events: event_routes}
}

func (v *MainView) ShowUI() {
	v.ctx.win.SetContent(v.createComponents())
}

func (v *MainView) On(e EventType) {
	// Execute event function which is mapped EventType.
	if f, ok := v.events[e]; !ok {
		log.Fatal("Error: Invalid event fired.", e, ok)
	} else {
		f()
	}
}

func (v *MainView) GetViewType() ViewType {
	return "main"
}

func (v *MainView) Refresh() {
	v.ctx.win.Content().Refresh()
}

func (v *MainView) createComponents() *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Hello World!"),
		v.makefooter(),
	)
}

func (v *MainView) next() {
	v.ctx.SetState(NewNextView(v.ctx))
}

func (v *MainView) makeHeader() *fyne.Container {
	return container.NewVBox(
		container.NewHBox(
			widget.NewButton("Open", func() {
				// v.notifyEvent("openfolder")
			}),
			widget.NewButton("Help", func() {
				// v.notifyEvent("openhelp")
			}),
		),
		container.NewHBox(
			widget.NewButtonWithIcon("Check All", theme.CheckButtonCheckedIcon(), func() {
				// v.notifyEvent("onselectall")
			}),
			widget.NewButtonWithIcon("Clear All", theme.CheckButtonIcon(), func() {
				// v.notifyEvent("onunselectall")
			}),
			widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
				v.Refresh()
			}),
		),
	)
}

func (v *MainView) makeBody() *fyne.Container {
	// v.pathListWidget = widget.NewListWithData(v.bindURIList,
	// 	func() fyne.CanvasObject {
	// 		check := widget.NewCheck("", func(value bool) {})
	// 		check.Disable()
	// 		return container.NewHBox(
	// 			check,
	// 			widget.NewIcon(theme.FileImageIcon()),
	// 			widget.NewLabel("Template Object"),
	// 		)
	// 	},
	// 	func(i binding.DataItem, o fyne.CanvasObject) {
	// 		uri, _ := i.(binding.URI).Get()
	// 		name := uri.Name()
	// 		isSelect := model.GetWorkFolderModel().QueryListItemIsSelected(uri)
	// 		o.(*fyne.Container).Objects[0].(*widget.Check).Bind(binding.BindBool(&isSelect))
	// 		o.(*fyne.Container).Objects[2].(*widget.Label).Bind(binding.BindString(&name))
	// 	})

	// image := canvas.NewImageFromResource(theme.FileImageIcon())
	// image.FillMode = canvas.ImageFillContain
	// image.SetMinSize(fyne.NewSize(50, 50))
	// v.preview = container.NewMax(canvas.NewRectangle(theme.BackgroundColor()), image)

	// v.pathListWidget.OnSelected = func(id widget.ListItemID) {
	// 	workFolder := model.GetWorkFolderModel()
	// 	// Set Select Checkbox state.
	// 	selectedURI, selectedFlag := workFolder.GetListItem(id)
	// 	workFolder.UpdateURIListItem(id, selectedURI, !selectedFlag)
	// 	log.Printf("[%s, %v] -> [%s, %v]", selectedURI, selectedFlag, selectedURI, !selectedFlag)
	// 	// Get recently list again.
	// 	v.bindURIList = workFolder.CreateBindingURIList()

	// 	// Image Preview Update asyncronous, because image load process may heavy.
	// 	go func() {
	// 		newImage := canvas.NewImageFromFile(selectedURI.Path())
	// 		log.Printf("newImage: %#v\n", newImage)
	// 		newImage.FillMode = canvas.ImageFillContain
	// 		newImage.SetMinSize(fyne.NewSize(50, 50))
	// 		// Delete Image object
	// 		v.preview.Remove(v.preview.Objects[1])
	// 		v.preview.Add(newImage)
	// 		v.Refresh()
	// 	}()
	// }

	// return container.NewMax(container.NewHSplit(v.pathListWidget, v.preview))
	return container.NewWithoutLayout()
}

func (v *MainView) makefooter() *fyne.Container {
	return container.NewHBox(
		widget.NewButton("event test", func() {
			v.On("click")
		}),
		widget.NewButton("change state", func() { v.next() }),
	)
}

// Event functions
func onClick() {
	log.Println("onclick event hudled!")
}
