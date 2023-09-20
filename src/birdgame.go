package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	gameView := NewGameView(100, 100)
	gameView.SetChangedFunc(func() { go app.Draw() })

	bird := NewBird(10, 10)
	gravityAnimation := NewGravityAnimation(bird)
	bird.AddAnimation(gravityAnimation)
	gameView.AddObject(bird)

	spaceKeyEh := NewSpaceKeyEventHandler()
	spaceKeyEh.AddListener(bird)
	gameView.RegisterSpaceKeyEventHandler(spaceKeyEh)
	go gameView.RunRenderLoop(60)

	if err := app.SetRoot(gameView, true).Run(); err != nil {
		panic(err)
	}
}
