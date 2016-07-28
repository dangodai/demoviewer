package main

import (
	"fmt"

	"github.com/dangodai/demoviewer/demos"
	"github.com/visualfc/goqt/ui"
)

var (
	demolist []demos.Demo
)

//MainWindow expands on the QMainWindow struct to hold a few key components
//of our UI
type MainWindow struct {
	*ui.QMainWindow
	list    *ui.QListWidget
	details *ui.QPlainTextEdit
	events  *ui.QPlainTextEdit
}

//NewMainWindow builds the program's main window and returns the created object
func NewMainWindow() *MainWindow {
	window := &MainWindow{}
	window.QMainWindow = ui.NewMainWindow()
	window.SetWindowTitle("Demo Viewer v0.1")
	window.SetMinimumSizeWithMinwMinh(500, 300)
	window.createMenuBar()
	window.setupWidgets()

	detailsLayout := ui.NewBoxLayout(ui.QBoxLayout_LeftToRight, nil)
	detailsLayout.SetContentsMarginsWithLeftTopRightBottom(0, 0, 0, 0)
	detailsLayout.AddWidget(window.details)
	detailsLayout.AddWidget(window.events)

	detailsContainer := ui.NewWidget()
	detailsContainer.SetLayout(detailsLayout)

	containerLayout := ui.NewBoxLayout(ui.QBoxLayout_TopToBottom, nil)
	containerLayout.AddWidget(window.list)
	containerLayout.AddWidget(detailsContainer)

	container := ui.NewWidget()
	container.SetLayout(containerLayout)

	window.SetCentralWidget(container)

	//window.displayDemos()
	return window
}

//setupWidgets adds the widgets that display dynamic information (the demo list
//and the demo details) to the main window
func (w *MainWindow) setupWidgets() {
	w.list = ui.NewListWidget()
	w.list.OnItemSelectionChanged(w.displayDemoDetails)
	w.list.OnItemDoubleClicked(func(widget *ui.QListWidgetItem) { w.playSelectedDemo() })

	w.details = ui.NewPlainTextEdit()
	w.details.SetReadOnly(true)

	w.events = ui.NewPlainTextEdit()
	w.events.SetReadOnly(true)
}

//createMenuBar creates and adds the menu bars to the main window
func (w *MainWindow) createMenuBar() {
	//Set up File menu
	selectFolder := ui.NewActionWithTextParent(w.Tr("Select Demo Folder"), w)
	selectFolder.OnTriggered(w.selectDemoFolder)

	fileMenu := w.MenuBar().AddMenuWithTitle("File")
	fileMenu.AddAction(selectFolder)

	//Set up Demo menu
	playDemo := ui.NewActionWithTextParent(w.Tr("Play Demo"), w)
	playDemo.OnTriggered(w.playSelectedDemo)
	deleteDemo := ui.NewActionWithTextParent(w.Tr("Delete Demo"), w)
	deleteDemo.OnTriggered(w.deleteSelectedDemo)

	demoMenu := w.MenuBar().AddMenuWithTitle("Demo")
	demoMenu.AddAction(playDemo)
	demoMenu.AddSeparator()
	demoMenu.AddAction(deleteDemo)

	//Set up Sort menu
	sortNameAZ := ui.NewActionWithTextParent(w.Tr("Sort by name (A->Z)"), w)
	sortNameAZ.OnTriggered(func() { w.sort("NameAZ") })
	sortNameZA := ui.NewActionWithTextParent(w.Tr("Sort by name (Z->A)"), w)
	sortNameZA.OnTriggered(func() { w.sort("NameZA") })

	sortMapAZ := ui.NewActionWithTextParent(w.Tr("Sort by map (A->Z)"), w)
	sortMapAZ.OnTriggered(func() { w.sort("MapAZ") })
	sortMapZA := ui.NewActionWithTextParent(w.Tr("Sort by map (Z->A)"), w)
	sortMapZA.OnTriggered(func() { w.sort("MapZA") })

	sortUserAZ := ui.NewActionWithTextParent(w.Tr("Sort by user (A->Z)"), w)
	sortUserAZ.OnTriggered(func() { w.sort("UserAZ") })
	sortUserZA := ui.NewActionWithTextParent(w.Tr("Sort by user (Z->A)"), w)
	sortUserZA.OnTriggered(func() { w.sort("UserZA") })

	sortDateNew := ui.NewActionWithTextParent(w.Tr("Sort by date (Newest)"), w)
	sortDateNew.OnTriggered(func() { w.sort("DateNew") })
	sortDateOld := ui.NewActionWithTextParent(w.Tr("Sort by date (Oldest)"), w)
	sortDateOld.OnTriggered(func() { w.sort("DateOld") })

	sortMenu := w.MenuBar().AddMenuWithTitle("Sort")
	sortMenu.AddAction(sortNameAZ)
	sortMenu.AddAction(sortNameZA)
	sortMenu.AddSeparator()
	sortMenu.AddAction(sortMapAZ)
	sortMenu.AddAction(sortMapZA)
	sortMenu.AddSeparator()
	sortMenu.AddAction(sortUserAZ)
	sortMenu.AddAction(sortUserZA)
	sortMenu.AddSeparator()
	sortMenu.AddAction(sortDateNew)
	sortMenu.AddAction(sortDateOld)
}

//selectDemoFolder updates the status bar, and controls the flow for opening
//the folder select prompt, and then calling displayDemos on the main window
//once a folder has been selected
func (w *MainWindow) selectDemoFolder() {
	w.StatusBar().ShowMessage("Finding demo files...")
	setDemoFolder()
	w.StatusBar().ShowMessage("Demo files loaded")
	w.displayDemos()
}

func (w *MainWindow) sort(t string) {
	switch t {
	case "NameAZ":
		demos.SortByNameAZ()
	case "NameZA":
		demos.SortByNameZA()
	case "MapAZ":
		demos.SortByMapAZ()
	case "MapZA":
		demos.SortByMapZA()
	case "UserAZ":
		demos.SortByUserAZ()
	case "UserZA":
		demos.SortByUserZA()
	case "DateNew":
		demos.SortByDateNewest()
	case "DateOld":
		demos.SortByDateOldest()
	}
	w.displayDemos()
}

//displayDemos loops over the slice of demos, and displays their names in the
//main window's list.
func (w *MainWindow) displayDemos() {
	w.list.Clear()
	demolist = demos.GetDemos()
	for _, d := range demolist {
		temp := ui.NewListWidgetItem()
		temp.SetText(d.Name())

		if d.Events() != nil {
			temp.SetTextColor(ui.NewColorWithInt32Int32Int32Int32(250, 117, 50, 255))
		}

		w.list.AddItem(temp)
	}
	w.list.SetCurrentRow(0)
}

//displayDemoDetails updates the details and events fields on the UI,
//showing the information corresponding to the currently selected demo
func (w *MainWindow) displayDemoDetails() {
	if len(demolist) == 0 {
		return
	}

	row := w.list.CurrentRow()

	//Display the file details
	w.details.SetPlainText(fmt.Sprintf("User: %v\nMap: %v\nCommand: playdemo %v\nType: %v\nDate: %v\n",
		demolist[row].ClientName(),
		demolist[row].MapName(),
		demolist[row].PathInTFFolder(),
		demolist[row].Type(),
		demolist[row].Date().Format("Jan 2 15:04:05, 2006")))

	//Display the event details
	var s string
	for _, e := range demolist[row].Events() {
		s += fmt.Sprintf("[%v] %v (%v)\n", e.Name, e.Value, e.Tick)
	}
	w.events.SetPlainText(s)
}

//playSelectedDemo calls Play() on whatever demo is selected
func (w *MainWindow) playSelectedDemo() {
	demolist[w.list.CurrentRow()].Play()
}

//deleteSelectedDemo calls Delete() on whatever demo is selected
func (w *MainWindow) deleteSelectedDemo() {
	demolist[w.list.CurrentRow()].Delete()
	w.list.CurrentItem().Delete()
}

//setDemoFolder opens the dialog to select a directory, and sets the demo search
//path to the chosen directory.
func setDemoFolder() {
	demos.SetPath(ui.QFileDialogGetExistingDirectory())
}
