package ui

import (
	"log"

	"fyne.io/fyne/v2"
)

type EventType string
type ViewType string

type IView interface {
	ShowUI()
	GetViewType() ViewType
	Refresh()
	createComponents() *fyne.Container
}

// Common
// handle event functions.
func On(e EventType, events map[EventType]func()) {
	// Execute event function which is mapped EventType.
	if f, ok := events[e]; !ok {
		log.Fatal("Error: Invalid event fired.", e, ok)
	} else {
		log.Println("Event handled.", e)
		f()
	}
}
