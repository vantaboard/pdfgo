package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/vantaboard/pdfgo/gui"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello Person")

	w.SetContent(container.NewVBox(gui.MakeUI()))
	w.ShowAndRun()
}
