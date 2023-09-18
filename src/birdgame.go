package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	gameView := NewGameView(100, 100)
	gameView.SetChangedFunc(func() { go app.Draw() })

	bird := NewBird(10, 10)
	gameView.AddObject(bird)

	spaceKeyEh := NewSpaceKeyEventHandler()
	spaceKeyEh.AddListener(bird)
	gameView.RegisterSpaceKeyEventHandler(spaceKeyEh)
	gameView.Render()

	if err := app.SetRoot(gameView, true).Run(); err != nil {
		panic(err)
	}
}
