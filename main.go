package main

import "github.com/visualfc/goqt/ui"

func main() {
	//Run the Qt UI
	ui.Run(func() {
		window := NewMainWindow()
		window.Show()
	})
}
