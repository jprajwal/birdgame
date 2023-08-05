package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView().SetChangedFunc(func() {
		app.Draw()
	})
	textView.SetTitle("Birdgame")
	textView.SetBorder(true)
	rt := NewGameRuntime()
	go rt.Run(textView)
	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}

}
