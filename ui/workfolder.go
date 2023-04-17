package ui

import (
	"log"
	"sync"

	"fyne.io/fyne/v2"
)

type WorkFolder struct {
	UriList         []fyne.URI
	IsSelectedFlags []bool
}

// Avoid multi thread processes create multiple instances.
var lock = &sync.Mutex{}

// Singleton
var singleWorkFolder *WorkFolder

func getWorkFolderInstance() *WorkFolder {
	if singleWorkFolder != nil {
		return singleWorkFolder
	}

	lock.Lock()
	defer lock.Unlock()
	singleWorkFolder = &WorkFolder{}

	return singleWorkFolder
}

func NewWorkFolder(uriList []fyne.URI) *WorkFolder {
	wf := getWorkFolderInstance()
	wf.UriList = filterJPGorPNG(uriList)
	wf.IsSelectedFlags = make([]bool, len(wf.UriList))
	return wf
}

func filterJPGorPNG(uriList []fyne.URI) []fyne.URI {
	var filteredURIs []fyne.URI

	for _, u := range uriList {
		if u.Extension() == ".jpg" || u.Extension() == ".png" {
			filteredURIs = append(filteredURIs, u)
		}
	}

	return filteredURIs
}

func (wf *WorkFolder) UpdateURIListItem(id int, path fyne.URI, isSelect bool) {
	if id < 0 || id > len(wf.UriList) {
		// invalid id, panic.
		log.Fatalf("Invalid argument [id]:%v.\n", id)
	}
	wf.UriList[id] = path
	wf.IsSelectedFlags[id] = isSelect
}

func (wf *WorkFolder) GetSelectedURIs() []fyne.URI {
	var selectedList []fyne.URI
	for i, isSelected := range wf.IsSelectedFlags {
		if isSelected {
			selectedList = append(selectedList, wf.UriList[i])
		}
	}
	return selectedList
}

func (wf *WorkFolder) SelectFlagsAll(isSelect bool) {
	for i := range wf.IsSelectedFlags {
		wf.IsSelectedFlags[i] = isSelect
	}
}

func (wf *WorkFolder) QueryListItemIsSelected(uri fyne.URI) bool {
	id := wf.whereURIIndex(uri)
	return wf.IsSelectedFlags[id]
}

func (wf *WorkFolder) whereURIIndex(uri fyne.URI) int {
	for i, v := range wf.UriList {
		if uri.Name() == v.Name() {
			return i
		}
	}
	log.Fatal("Same uri not found.")
	return -1
}
