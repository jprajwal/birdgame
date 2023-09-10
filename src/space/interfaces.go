package space

import "fmt"

type Dimension2D int

type Cell2D interface {
	fmt.Stringer
}

type Location2D interface {
	getX() Dimension2D
	getY() Dimension2D
	setX(Dimension2D)
	setY(Dimension2D)
}

type Field2D interface {
	GetMatterAt(Dimension2D, Dimension2D) Cell2D
	GetRowCount() Dimension2D
	GetColumnCount() Dimension2D
	Clear()
}

type Object2D interface {
	Field2D
	Location2D
}

type Space2D interface {
	Field2D
	IncludeObject2D(Object2D)
	IterateObjects2D() Object2DIterator
}

type Object2DIterator interface {
	Next() Object2D
}
