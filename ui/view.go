package ui

import "fyne.io/fyne/v2"

type EventType string
type ViewType string

type IView interface {
	ShowUI()
	On(e EventType)
	GetViewType() ViewType
	Refresh()
	createComponents() *fyne.Container
	next()
}
