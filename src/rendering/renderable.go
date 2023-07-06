package rendering

import "github.com/jprajwal/birdgame/animation"

type Position struct {
	x, y int
}

func NewPosition(x, y int) *Position {
	return &Position{
		x: x,
		y: y,
	}
}

func (pos *Position) MoveHorizontal(delta int) {
	pos.y += delta
}

func (pos *Position) MoveVertical(delta int) {
	pos.x += delta
}

func (pos *Position) GetX() int {
	return pos.x
}

func (pos *Position) GetY() int {
	return pos.y
}

type Renderable interface {
	GetPosition() *Position
	GetAnimations() []animation.Animation
	GetRenderData() [][]rune
}
