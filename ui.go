package main

import (
	"fmt"

	"github.com/dangodai/demoviewer/demos"
	"github.com/visualfc/goqt/ui"
)

var (
	demolist []demos.Demo
)

//MainWindow creates main gui
type MainWindow struct {
	*ui.QMainWindow
	list    *ui.QListWidget
	details *ui.QPlainTextEdit
	events  *ui.QPlainTextEdit
}

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

func (w *MainWindow) setupWidgets() {
	w.list = ui.NewListWidget()
	w.list.OnItemSelectionChanged(w.displayDemoDetails)
	w.list.OnItemDoubleClicked(func(widget *ui.QListWidgetItem) { w.playSelectedDemo() })

	w.details = ui.NewPlainTextEdit()
	w.details.SetReadOnly(true)

	w.events = ui.NewPlainTextEdit()
	w.events.SetReadOnly(true)
}

func (w *MainWindow) createMenuBar() {
	selectFolder := ui.NewActionWithTextParent(w.Tr("Select Demo Folder"), w)
	selectFolder.OnTriggered(w.selectDemoFolder)

	fileMenu := w.MenuBar().AddMenuWithTitle("File")
	fileMenu.AddAction(selectFolder)

	playDemo := ui.NewActionWithTextParent(w.Tr("Play Demo"), w)
	playDemo.OnTriggered(w.playSelectedDemo)
	deleteDemo := ui.NewActionWithTextParent(w.Tr("Delete Demo"), w)
	deleteDemo.OnTriggered(w.deleteSelectedDemo)

	demoMenu := w.MenuBar().AddMenuWithTitle("Demo")
	demoMenu.AddAction(playDemo)
	demoMenu.AddSeparator()
	demoMenu.AddAction(deleteDemo)
}

func (w *MainWindow) selectDemoFolder() {
	w.StatusBar().ShowMessage("Finding demo files...")
	setDemoFolder()
	w.StatusBar().ShowMessage("Demo files loaded")
	w.displayDemos()
}

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

func (w *MainWindow) displayDemoDetails() {
	if len(demolist) == 0 {
		return
	}

	row := w.list.CurrentRow()

	//Display the file details
	w.details.SetPlainText(fmt.Sprintf("User: %v\nMap: %v\nCommand: playdemo %v\nDate: %v\n",
		demolist[row].ClientName(),
		demolist[row].MapName(),
		demolist[row].PathInTFFolder(),
		demolist[row].Date().Format("Jan 2 15:04:05, 2006")))

	//Display the event details
	var s string
	for _, e := range demolist[row].Events() {
		s += fmt.Sprintf("[%v] %v (%v)\n", e.Name, e.Value, e.Tick)
	}
	w.events.SetPlainText(s)
}

func (w *MainWindow) playSelectedDemo() {
	demolist[w.list.CurrentRow()].Play()
}

func (w *MainWindow) deleteSelectedDemo() {
	demolist[w.list.CurrentRow()].Delete()
	w.list.CurrentItem().Delete()
}

func setDemoFolder() {
	demos.SetPath(ui.QFileDialogGetExistingDirectory())
}
