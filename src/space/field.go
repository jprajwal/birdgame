package space

type Field struct {
	x Dimension2D
	y Dimension2D
	buffer [][]rune
}

func make2D(x, y Dimension2D) [][]rune {
	made := make([][]rune, x)
	var i Dimension2D
	for i = 0; i < x; i++ {
		made[i] = make([]rune, y)
	}
	return made
}

func NewField(x, y Dimension2D) *Field {
	return &Field{
		x: x,
		y: y,
		buffer: make2D(x, y),
	}
}

func (f *Field) GetMatterAt(x Dimension2D, y Dimension2D) string {
	if x >= Dimension2D(f.x) || y >= Dimension2D(f.y) {
		return ""
	}
	return string(f.buffer[x][y])
}

func (f *Field) GetRowCount() Dimension2D {
	return Dimension2D(f.x)
}

func (f *Field) GetColumnCount() Dimension2D {
	return Dimension2D(f.y)
}
