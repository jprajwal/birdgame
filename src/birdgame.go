package main

import (
	"github.com/rivo/tview"
)

func main() {
	InitLogger(true)
	defer FiniLogger()

	app := tview.NewApplication()
	gameView := NewGameView(100, 100)
	gameView.SetChangedFunc(func() { go app.Draw() })

	bird := NewBird(10, 10)
	gravityAnimation := NewGravityAnimation(bird)
	bird.AddAnimation(gravityAnimation)
	gameView.AddObject(bird)

	gameView.AddObject(NewPillar(25, 20, 0, 20))
	gameView.AddObject(NewPillar(25, 30, 0, 50))
	gameView.AddObject(NewPillar(25, 25, 0, 80))
	gameView.AddObject(NewPillar(25, 20, 0, 110))
	gameView.AddObject(NewPillar(25, 20, 0, 140))

	spaceKeyEh := NewSpaceKeyEventHandler()
	spaceKeyEh.AddListener(bird)
	gameView.RegisterSpaceKeyEventHandler(spaceKeyEh)
	go gameView.RunRenderLoop(60)

	if err := app.SetRoot(gameView, true).Run(); err != nil {
		panic(err)
	}
}
