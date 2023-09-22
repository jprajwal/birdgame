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

type Metre float64
type MetrePerSecond float64

type HasVerticalVelocity interface {
	GetVerticalVelocity() MetrePerSecond
	SetVerticalVelocity(MetrePerSecond)
}

type Animatable interface {
	Animate(Field)
}

type AnimatableObject interface {
	Object
	Animatable
	AddAnimation(Animatable)
}

type ObjectWithVerticalVelocity interface {
	Object
	HasVerticalVelocity
}

type Slidable interface {
	Slide(int)
}

type PillarContainer interface {
	GetLastPillarPosition() int
	RegisterPillarGenerator(PillarGenerator)
}

type PillarGenerator interface {
	Generate(Field, PillarContainer) Object
}
