package main

import "github.com/visualfc/goqt/ui"

func main() {
	ui.Run(func() {
		window := NewMainWindow()
		window.Show()
	})
}
