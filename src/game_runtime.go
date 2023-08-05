package main

import (
	"github.com/jprajwal/birdgame/animation"
	"github.com/jprajwal/birdgame/rendering"
	"github.com/rivo/tview"
	"time"
)

const FPS = 30

type GameRuntime struct {
	space    *Space
	curFrame *Frame
}

func NewGameRuntime() *GameRuntime {
	return &GameRuntime{space: NewSpace(10, 180), curFrame: nil}
}

func (rt *GameRuntime) RenderInSpace() {
	rt.space.Clear()
	for _, renderable := range rt.curFrame.Objects() {
		pos := renderable.GetPosition()
		data := renderable.GetRenderData()

		for i := 0; i < len(data); i++ {
			for j := 0; j < len(data[i]); j++ {
				rt.space.SetDataAt(i+pos.GetX(), j+pos.GetY(), data[i][j])
			}
		}
	}
}

func createRenderables() []rendering.Renderable {
	renderables := make([]rendering.Renderable, 0)
	renderable := rendering.NewText("I am prajwal", *rendering.NewPosition(1, 179))
	renderable.AddAnimation(animation.NewSlideLeftAnimation(180, 5000))
	renderables = append(renderables, renderable)
	return renderables
}

func (rt *GameRuntime) Run(textView *tview.TextView) {
	firstTime := true
	for {
		if rt.curFrame == nil {
			rt.curFrame = NewFrame()
		}
		newframe := rt.curFrame.Clone()
		if firstTime {
			firstTime = false
			renderables := createRenderables()
			for _, renderable := range renderables {
				newframe.AddRenderable(renderable)
			}
		}
		for _, renderable := range newframe.GetRenderables() {
			animations := renderable.GetAnimations()
			for _, anim := range animations {
				anim.Animate(renderable.GetPosition())
			}
		}
		rt.curFrame = newframe
		rt.RenderInSpace()
		str := rt.space.GetDataAsString()
		textView.Clear()
		textView.SetText(str)
		time.Sleep(time.Second / time.Duration(FPS))
	}
}
