package ui

import (
	"log"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/storage"
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
	wf.SetSelectFlagsAll(false)
	return wf
}

func GetWorkFolder() *WorkFolder {
	return getWorkFolderInstance()
}

func GetWorkFolderSpecifyPath(folder fyne.URI) *WorkFolder {
	wf := getWorkFolderInstance()
	if ok, err := storage.CanList(folder); !ok || err != nil {
		return nil
	}
	list, _ := storage.List(folder)
	wf.UriList = filterJPGorPNG(list)
	wf.IsSelectedFlags = make([]bool, len(wf.UriList))
	wf.SetSelectFlagsAll(false)
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

func (wf *WorkFolder) UpdateSelectedURIListItem(id int) {
	if id < 0 || id > len(wf.UriList) {
		// invalid id, panic.
		log.Fatalf("Invalid argument [id]:%v.\n", id)
	}
	// reverse isSelectedFlag
	old := wf.IsSelectedFlags[id]
	wf.IsSelectedFlags[id] = !old
	log.Printf("[%s]: %v -> %v", wf.UriList[id], old, wf.IsSelectedFlags[id])
	log.Printf("%v", wf.IsSelectedFlags)
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

func (wf *WorkFolder) SetSelectFlagsAll(isSelect bool) {
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

func (wf *WorkFolder) CreateBindingURIList() binding.URIList {
	bindList := binding.NewURIList()
	if len(wf.UriList) == 0 {
		// return empty bindList
		return bindList
	}
	// Binding list
	err := bindList.Set(wf.UriList)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return bindList
}
