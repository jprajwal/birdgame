package main

import (
	"fmt"
	"time"

	"github.com/jprajwal/birdgame/animation"
	"github.com/jprajwal/birdgame/rendering"
	"github.com/rivo/tview"
)

type Text struct {
	text       string
	pos        rendering.Position
	animations []animation.Animation
}

func NewText(text string, pos rendering.Position) *Text {
	return &Text{
		text:       text,
		pos:        pos,
		animations: make([]animation.Animation, 0),
	}
}

func (t *Text) AddAnimation(a animation.Animation) {
	t.animations = append(t.animations, a)
}

func (t *Text) GetPosition() *rendering.Position {
	return &t.pos
}

func (t *Text) GetAnimations() []animation.Animation {
	return t.animations
}

func (t *Text) GetRenderData() [][]rune {
	data := make([][]rune, 0)
	firstRow := []rune(t.text)
	data = append(data, firstRow)
	return data
}

type Frame struct {
	objects []rendering.Renderable
}

func NewFrame() *Frame {
	return &Frame{objects: make([]rendering.Renderable, 0)}
}

func (f *Frame) AddRenderable(r rendering.Renderable) {
	f.objects = append(f.objects, r)
}

func (f *Frame) Clone() *Frame {
	objectsCopy := make([]rendering.Renderable, len(f.objects))
	for i, obj := range f.objects {
		objectsCopy[i] = obj
	}
	return &Frame{objects: objectsCopy}
}

type GameRuntime struct {
	space    *[10][100]rune
	curFrame *Frame
}

func NewGameRuntime() *GameRuntime {
	space := new([10][100]rune)
	for i := 0; i < 10; i++ {
		for j := 0; j < 100; j++ {
			space[i][j] = ' '
		}
	}
	return &GameRuntime{space: space, curFrame: nil}
}

func (rt *GameRuntime) RenderInSpace() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 100; j++ {
			rt.space[i][j] = ' '
		}
	}
	for _, renderable := range rt.curFrame.objects {
		pos := renderable.GetPosition()
		data := renderable.GetRenderData()

		for i, r := range data {
			// if pos.GetX() < 0 {
			// 	continue
			// }
			if pos.GetX()+i >= 10 {
				break
			}
			for j, c := range r {
				// if pos.GetY() < 0 {
				// 	continue
				// }
				// fmt.Printf("x: %v, y: %v, j: %v\n", pos.GetX(), pos.GetY(), j)
				rt.space[pos.GetX()+i][pos.GetY()+j] = c
			}
		}
	}
}

func createRenderables() []rendering.Renderable {
	renderables := make([]rendering.Renderable, 0)
	renderable := NewText("I am prajwal", *rendering.NewPosition(1, 81))
	renderable.AddAnimation(animation.NewSlideLeftAnimation(80, 2000))
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
		for i := 0; i < len(newframe.objects); i++ {
			renderable := newframe.objects[i]
			animations := renderable.GetAnimations()
			for _, anim := range animations {
				anim.Animate(renderable.GetPosition())
			}
		}
		rt.curFrame = newframe
		rt.RenderInSpace()
		str := ""
		for i := 0; i < len(rt.space); i++ {
			if i > 0 {
				str += "\n"
			}
			row := rt.space[i]
			for j := 0; j < len(row); j++ {
				str += string(row[j])
			}
		}
		textView.Clear()
		// fmt.Print(str)
		fmt.Fprint(textView, str)
		time.Sleep(time.Second / time.Duration(60))
	}
}

func main() {
	fmt.Println("This is the birdgame")
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
