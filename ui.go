package main

import (
	"github.com/dangodai/demoviewer/demos"
	"github.com/visualfc/goqt/ui"
)

//MainWindow creates main gui
type MainWindow struct {
	*ui.QMainWindow
}

func NewMainWindow() *MainWindow {
	window := &MainWindow{}
	window.QMainWindow = ui.NewMainWindow()
	window.SetWindowTitle("Demo Viewer v0.0.1")
	window.SetMinimumSizeWithMinwMinh(500, 300)
	window.createMenuBar()

	return window
}

func (w *MainWindow) createMenuBar() {
	selectFolder := ui.NewActionWithTextParent(w.Tr("Select Demo Folder"), w)
	selectFolder.OnTriggered(w.selectDemoFolder)

	fileMenu := w.MenuBar().AddMenuWithTitle("File")
	fileMenu.AddAction(selectFolder)
}

func (w *MainWindow) selectDemoFolder() {
	w.StatusBar().ShowMessage("Finding demo files...")
	setDemoFolder()
	w.StatusBar().ShowMessage("Demo files loaded")
	w.displayDemos()
}

func (w *MainWindow) displayDemos() {

}

func setDemoFolder() {
	demos.SetPath(ui.QFileDialogGetExistingDirectory())
}
