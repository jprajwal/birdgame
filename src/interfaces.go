package main

import "fmt"

type Cell interface {
	fmt.Stringer
	Clear()
	Set(rune)
	Get() rune
	Clone() Cell
}

type Field interface {
	Fill(Cell)
	Clear()
	CopyFieldAt(int, int, Field)
	GetContent(int, int) Cell
	Width() int
	Height() int
	Clone() Field
}

type Positionable interface {
	GetPositionX() int
	GetPositionY() int
	SetPositionX(int)
	SetPositionY(int)
}

type Renderable interface {
	GetData() [][]rune
}

type Drawable interface {
	Draw(Field)
}

type Object interface {
	Drawable
	Positionable
}

type SpaceKeyListenable interface {
	OnSpaceKeyPressed()
}

type EventHandler interface {
	Notify()
}

