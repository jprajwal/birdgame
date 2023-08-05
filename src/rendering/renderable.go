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

type Text struct {
	text       string
	pos        Position
	animations []animation.Animation
}

func NewText(text string, pos Position) *Text {
	return &Text{
		text:       text,
		pos:        pos,
		animations: make([]animation.Animation, 0),
	}
}

func (t *Text) AddAnimation(a animation.Animation) {
	t.animations = append(t.animations, a)
}

func (t *Text) GetPosition() *Position {
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
