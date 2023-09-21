package main

import (
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type GameView struct {
	width           int
	height          int
	data            *BasicField
	changedCallback func()
	spaceKeyEh      EventHandler
	objects         []AnimatableObject
	rwlock          sync.RWMutex
}

func NewGameView(x, y int) *GameView {
	return &GameView{
		data:    NewBasicField(x, y),
		objects: make([]AnimatableObject, 0),
	}
}

func (b *GameView) AddObject(obj AnimatableObject) {
	b.objects = append(b.objects, obj)
}

func (b *GameView) RegisterSpaceKeyEventHandler(eh EventHandler) {
	b.spaceKeyEh = eh
}

func (b *GameView) Render() {
	newField := NewBasicField(b.height, b.width)

	b.rwlock.RLock()
	for i := 0; i < len(b.objects); i++ {
		object := b.objects[i]
		object.Animate(newField)
		object.Draw(newField)
	}
	b.rwlock.RUnlock()

	b.rwlock.Lock()
	b.data = newField
	b.rwlock.Unlock()

	b.changedCallback()
}

func (b *GameView) RunRenderLoop(fps int) {
	for {
		interval := 1 / fps
		time.Sleep(time.Duration(interval) * time.Second)
		b.Render()
	}
}

func (b *GameView) Draw(screen tcell.Screen) {
	b.rwlock.RLock()
	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			screen.SetContent(j, i, b.data.GetContent(i, j).Get(), nil, tcell.StyleDefault)
		}
	}
	b.rwlock.RUnlock()
}

func (b *GameView) GetRect() (int, int, int, int) {
	b.rwlock.RLock()
	defer b.rwlock.RUnlock()
	return 0, 0, b.width, b.height
}

func (b *GameView) SetRect(x, y, width, height int) {
	b.rwlock.Lock()
	old := b.data
	b.data = NewBasicField(height, width)
	b.data.CopyFieldAt(0, 0, old)
	b.width = width
	b.height = height
	b.rwlock.Unlock()
}

func (b *GameView) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	handler := func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		if event.Key() != tcell.KeyRune {
			return
		}
		switch event.Rune() {
		case ' ':
			b.spaceKeyEh.Notify()
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
