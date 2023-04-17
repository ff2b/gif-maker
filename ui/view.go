package ui

import "fyne.io/fyne/v2"

type IView interface {
	ShowUI()
	On(e EventType)
	GetViewType() ViewType
	Refresh()
	createComponents() *fyne.Container
	next()
}

type EventType string
type ViewType string
