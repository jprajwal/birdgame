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

	pillarContainer := NewBasicPillarContainer()
	pillarGenerator := NewUniformPillarGenerator(30, 20)
	pillarContainer.RegisterPillarGenerator(pillarGenerator)
	pillarContainer.AddAnimation(NewLeftSlidingAnimation(pillarContainer, MetrePerSecond(30)))
	gameView.AddObject(pillarContainer)

	bird := NewBird(10, 10)
	gravityAnimation := NewGravityAnimation(bird)
	bird.AddAnimation(gravityAnimation)
	gameView.AddObject(bird)
	spaceKeyEh := NewSpaceKeyEventHandler()
	spaceKeyEh.AddListener(bird)
	gameView.RegisterSpaceKeyEventHandler(spaceKeyEh)

	go gameView.RunRenderLoop(10)

	if err := app.SetRoot(gameView, true).Run(); err != nil {
		panic(err)
	}
}
