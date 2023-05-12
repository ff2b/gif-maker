package ui

import (
	"errors"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	MSG_INIT        = "Target file not found. Open workspace folder"
	MSG_SELECT_FILE = "Select target images"
)

type MainView struct {
	ctx            *UIContext
	events         map[EventType]func()
	message        *widget.Label
	pathListWidget *widget.List
	bindURIList    binding.DataList
	preview        *fyne.Container
	workfolder     *WorkFolder
}

func NewMainView(ctx *UIContext) *MainView {
	view := &MainView{ctx: ctx, events: nil, message: widget.NewLabel(MSG_INIT), pathListWidget: nil, bindURIList: nil, preview: nil, workfolder: nil}
	// Register events
	view.events = map[EventType]func(){
		"open":    view.onOpenFolder,
		"confirm": view.onOpenConfirm,
	}
	view.workfolder = GetWorkFolder()
	view.bindURIList = view.workfolder.CreateBindingURIList()
	return view
}

func (v *MainView) ShowUI() {
	v.ctx.win.SetContent(v.createComponents())
	v.Refresh()
}

func (v *MainView) GetViewType() ViewType {
	return "main"
}

func (v *MainView) Refresh() {
	v.bindURIList = v.workfolder.CreateBindingURIList()
	if v.bindURIList.Length() > 0 {
		v.message.Text = MSG_SELECT_FILE
	}
	v.ctx.win.Content().Refresh()
}

func (v *MainView) createComponents() *fyne.Container {
	return container.NewBorder(v.makeHeader(), v.makeFooter(), nil, nil, v.makeBody())
}

func (v *MainView) makeHeader() *fyne.Container {
	return container.NewVBox(
		container.NewHBox(
			widget.NewButtonWithIcon("Open", theme.FolderOpenIcon(), func() {
				On("open", v.events)
			}),
			widget.NewButtonWithIcon("Help", theme.HelpIcon(), func() {
				// v.notifyEvent("openhelp")
			}),
		),
		canvas.NewLine(color.RGBA{117, 117, 117, 255}),
		container.NewHBox(widget.NewIcon(theme.InfoIcon()), v.message),
		container.NewHBox(
			widget.NewButtonWithIcon("Check All", theme.CheckButtonCheckedIcon(), func() {
				v.workfolder.SetSelectFlagsAll(true)
				v.Refresh()
			}),
			widget.NewButtonWithIcon("Clear All", theme.CheckButtonIcon(), func() {
				v.workfolder.SetSelectFlagsAll(false)
				v.Refresh()
			}),
			widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
				v.Refresh()
			}),
		),
	)
}

func (v *MainView) makeBody() *fyne.Container {
	v.pathListWidget = widget.NewListWithData(v.bindURIList,
		func() fyne.CanvasObject {
			check := widget.NewCheck("", func(value bool) {})
			check.Disable()
			return container.NewHBox(
				check,
				widget.NewIcon(theme.FileImageIcon()),
				widget.NewLabel("Template Object"),
			)
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			uri, _ := i.(binding.URI).Get()
			name := uri.Name()
			isSelect := GetWorkFolder().QueryListItemIsSelected(uri)
			// Bind Checkbox
			o.(*fyne.Container).Objects[0].(*widget.Check).Bind(binding.BindBool(&isSelect))
			// Bind URI List
			o.(*fyne.Container).Objects[2].(*widget.Label).Bind(binding.BindString(&name))
		})

	image := canvas.NewImageFromResource(theme.FileImageIcon())
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(50, 50))
	v.preview = container.NewMax(canvas.NewRectangle(theme.BackgroundColor()), image)

	v.pathListWidget.OnSelected = func(id widget.ListItemID) {
		v.workfolder = GetWorkFolder()
		// Set Select Checkbox state.
		v.workfolder.UpdateSelectedURIListItem(id)
		// Get latest URI list.
		v.bindURIList = v.workfolder.CreateBindingURIList()

		// Image Preview Update asyncronous, because image load process may heavy.
		go func() {
			newImage := canvas.NewImageFromFile(v.workfolder.UriList[id].Path())
			// log.Printf("Preview newImage: %#v\n", newImage)
			newImage.FillMode = canvas.ImageFillContain
			newImage.SetMinSize(fyne.NewSize(50, 50))
			// Replace Image
			v.preview.Remove(v.preview.Objects[1])
			v.preview.Add(newImage)
			v.Refresh()
		}()
	}

	split := container.NewHSplit(v.pathListWidget, v.preview)
	split.SetOffset(0.7)

	return container.NewMax(split)
}

func (v *MainView) makeFooter() *fyne.Container {
	createGIFButton := widget.NewButton("Create GIF", func() {
		On("confirm", v.events)
	})
	createGIFButton.Importance = widget.HighImportance

	return container.NewHBox(
		layout.NewSpacer(),
		createGIFButton,
	)
}

// Event functions
func (v *MainView) onOpenConfirm() {
	if len(GetWorkFolder().GetSelectedURIs()) == 0 {
		dialog.ShowError(errors.New("please select image file at least one"), v.ctx.win)
		return
	}
	v.ctx.SetState(NewConfirmView(v.ctx))
}

func (v *MainView) onOpenFolder() {
	v.ctx.SetState(NewOpenFolderView(v.ctx))
}
