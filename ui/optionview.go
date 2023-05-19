package ui

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ff2b/gif-maker/config"
	_ "github.com/ff2b/gif-maker/config"
)

const (
	OPTION_WIN_WIDTH = float32(400)
	OPTION_WIN_HIGHT = float32(200)
)

type OptionView struct {
	subwin  fyne.Window
	events  map[EventType]func()
	config  *config.Config
	binding *OptionBindings
	entry1  *widget.Entry
	entry2  *widget.Entry
	entry3  *widget.Entry
}

type OptionBindings struct {
	gifRate         binding.String
	gifLoop         binding.Bool
	workspace       binding.String
	defaultSavePath binding.String
}

func newOptionBindings(config config.Config) *OptionBindings {
	ob := &OptionBindings{
		gifRate:         binding.NewString(),
		gifLoop:         binding.NewBool(),
		workspace:       binding.NewString(),
		defaultSavePath: binding.NewString(),
	}
	// sync config data
	config.Load()
	ob.gifRate.Set(fmt.Sprint(config.GIFRate))
	ob.gifLoop.Set(config.GIFLoop)
	ob.workspace.Set(config.WorkSpace)
	ob.defaultSavePath.Set(config.DefaultSavePath)
	return ob
}

func NewOptionView() *OptionView {
	conf := config.NewConfig()
	view := &OptionView{config: conf, binding: newOptionBindings(*conf)}
	events := map[EventType]func(){
		"apply":  view.OnApply,
		"cancel": view.OnCancel,
	}
	view.events = events
	return view
}

func (v *OptionView) ShowUI() {
	v.subwin = fyne.CurrentApp().NewWindow("Options")
	v.subwin.Resize(fyne.NewSize(OPTION_WIN_WIDTH, OPTION_WIN_HIGHT))
	v.subwin.SetContent(v.createComponents())
	v.subwin.CenterOnScreen()
	v.subwin.Show()
}
func (v *OptionView) GetViewType() ViewType {
	return "options"
}

func (v *OptionView) Refresh() {
	// Nothing to do
}

func (v *OptionView) createComponents() *fyne.Container {
	v.entry1 = widget.NewEntryWithData(v.binding.gifRate)
	v.entry2 = widget.NewEntryWithData(v.binding.workspace)
	v.entry3 = widget.NewEntryWithData(v.binding.defaultSavePath)

	return container.NewVBox(
		widget.NewLabel("GIF encoding options"),
		widget.NewForm(
			&widget.FormItem{
				Text:     "GIF frame rate[ms]",
				HintText: "Range: 20-9999",
				Widget:   v.entry1,
			},
		),
		widget.NewForm(
			widget.NewFormItem(
				"GIF loop",
				widget.NewCheckWithData("", v.binding.gifLoop),
			)),
		widget.NewLabel("Application defaults"),
		widget.NewForm(
			widget.NewFormItem(
				"Default workspace folder",
				v.entry2)),
		widget.NewForm(
			widget.NewFormItem(
				"Default save folder",
				v.entry3)),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButton("Cancel", func() { On("cancel", v.events) }),
			widget.NewButton("Apply", func() { On("apply", v.events) }),
		),
	)
}

// event handler
func (v *OptionView) OnCancel() {
	v.subwin.Close()
}

func (v *OptionView) OnApply() {
	if err := v.binding.validation(); err != nil {
		dialog.ShowError(err, v.subwin)
		return
	}

	v.config = v.binding.convertConfig()
	log.Println(v.config)
	v.config.Save()

	v.subwin.Close()
}

// OptionBindings private functions
func (ob *OptionBindings) validation() error {
	rate := ob.gifRateToInt()
	if rate < 20 || 9999 < rate {
		return fmt.Errorf("%d ms is out of range", rate)
	}

	invalidChar := `[:\"\'\?<>|]`
	regexp, _ := regexp.Compile(invalidChar)
	workspace, _ := ob.workspace.Get()
	savepath, _ := ob.defaultSavePath.Get()
	if regexp.MatchString(workspace) {
		return fmt.Errorf("workspace path entry use invalid symbol character\nYour Input: %s", workspace)
	}
	if regexp.MatchString(savepath) {
		return fmt.Errorf("save path entry use invalid symbol character\nYour Input: %s", savepath)
	}

	return nil
}

func (ob *OptionBindings) convertConfig() *config.Config {
	rate := ob.gifRateToInt()
	loop, _ := ob.gifLoop.Get()
	workspace, _ := ob.workspace.Get()
	savepath, _ := ob.defaultSavePath.Get()
	return &config.Config{
		GIFRate:         rate,
		GIFLoop:         loop,
		WorkSpace:       workspace,
		DefaultSavePath: savepath,
	}
}

func (ob *OptionBindings) gifRateToInt() int {
	gifrate, _ := ob.gifRate.Get()
	log.Println(gifrate)
	rate, _ := strconv.ParseInt(gifrate, 0, 0)
	return int(rate)
}
