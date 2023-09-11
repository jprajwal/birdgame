package space

type Location struct {
	x, y Dimension2D
}

func NewLocation(x, y Dimension2D) *Location {
	return &Location{ x: x, y: y }
}

func (l *Location) getX() Dimension2D {
	return l.x
}

func (l *Location) getY() Dimension2D {
	return l.y
}

func (l *Location) setX(x Dimension2D) {
	l.x = x
}

func (l *Location) setY(y Dimension2D) {
	l.y = y
}

