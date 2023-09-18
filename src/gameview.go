package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type GameView struct {
	width int
	height int
	data *BasicField
	changedCallback func()
	spaceKeyEh      EventHandler
	objects         []Object
}

func NewGameView(x, y int) *GameView {
	return &GameView{
		data: NewBasicField(x, y),
		objects: make([]Object, 0),
	}
}

func (b *GameView) AddObject(obj Object) {
	b.objects = append(b.objects, obj)
}

func (b *GameView) RegisterSpaceKeyEventHandler(eh EventHandler) {
	b.spaceKeyEh = eh
}

func (b *GameView) Render() {
	newField := NewBasicField(b.width, b.height)
	
	for i := 0; i < len(b.objects); i++ {
		object := b.objects[i]
		object.Draw(newField)
	}

	b.data = newField
}

func (b *GameView) Draw(screen tcell.Screen) {
	b.Render()
	for i := 0; i < b.width; i++ {
		for j := 0; j < b.height; j++ {
			screen.SetContent(j, i, b.data.GetContent(i, j).Get(), nil, tcell.StyleDefault)
		}
	}
}

func (b *GameView) GetRect() (int, int, int, int) {
	return 0, 0, b.width, b.height
}

func (b *GameView) SetRect(x, y, width, height int) {
	old := b.data
	b.data = NewBasicField(width, height)
	b.data.CopyFieldAt(0, 0, old)
	b.width = width
	b.height = height
}

func (b *GameView) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	handler := func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		if event.Key() != tcell.KeyRune {
			return
		}
		switch event.Rune() {
		case ' ':
			b.spaceKeyEh.Notify()
			b.changedCallback()
		default:
		}
	}
	return handler
}

func (b *GameView) Focus(delegate func(p tview.Primitive)) {}

func (b *GameView) HasFocus() bool { return true }

func (b *GameView) Blur() {}

func (b *GameView) MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
	return nil
}

func (b *GameView) SetChangedFunc(callback func()) {
	b.changedCallback = callback
}
