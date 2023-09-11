package space

type Field struct {
	x Dimension2D
	y Dimension2D
	buffer [][]Cell2D
}

func make2D(x, y Dimension2D) [][]Cell2D {
	made := make([][]Cell2D, x)
	var i Dimension2D
	for i = 0; i < x; i++ {
		made[i] = make([]Cell2D, y)
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
	if x >= f.x || y >= f.y {
		return ""
	}
	return f.buffer[x][y].String()
}

func (f *Field) SetMatterAt(x Dimension2D, y Dimension2D, c Cell2D) {
	if x < 0 || y < 0 {
		return
	}
	f.buffer[x][y] = c
}

func (f *Field) GetRowCount() Dimension2D {
	return Dimension2D(f.x)
}

func (f *Field) GetColumnCount() Dimension2D {
	return Dimension2D(f.y)
}

func CopyFieldToField(loc Location2D, src Field2D, dst Field2D) Field2D {
	var i, j Dimension2D
	x, y := loc.getX(), loc.getY()

	for i = 0; i < src.GetRowCount(); i++ {
		for j = 0; j < src.GetColumnCount(); j++ {
			if x+i < 0 || y+j < 0 {
				continue
			}
			dst.SetMatterAt(x+i, y+j, src.GetMatterAt(i, j))
		}
	}
	return dst
}
