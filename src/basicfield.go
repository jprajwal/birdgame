package main

type BasicField struct {
	buf [][]Cell
	width int
	height int
}

func NewBasicField(width, height int) *BasicField {
	field := &BasicField{buf: make2DSlice[Cell](width, height), width: width, height: height}
	field.Fill(NewDefaultBasicCell())
	return field
}

func (bf *BasicField) Fill(c Cell) {
	for i := 0; i < bf.width; i++ {
		for j := 0; j < bf.height; j++ {
			bf.buf[i][j] = c.Clone()
		}
	}
}

func (bf *BasicField) Clear() {
	for i := 0; i < bf.width; i++ {
		for j := 0; j < bf.height; j++ {
			bf.buf[i][j].Clear()
		}
	}
}

func (bf *BasicField) CopyFieldAt(x, y int, f Field) {
	for i := 0; i < f.Width(); i++ {
		for j := 0; j < f.Height(); j++ {
			if x+i >= bf.width || y+j >= bf.height || x+i < 0 || y+j < 0 {
				continue
			}
			bf.buf[x+i][y+j] = f.GetContent(i, j)
		}
	}
}

func (bf *BasicField) GetContent(x, y int) Cell {
	if x < 0 || x > bf.width || y < 0 || y > bf.height {
		return nil
	}
	return bf.buf[x][y]
}

func (bf *BasicField) Width() int {
	return bf.width
}

func (bf *BasicField) Height() int {
	return bf.height
}

func (bf *BasicField) Clone() Field {
	return &BasicField{ buf: copy2DSlice[Cell](bf.buf), width: bf.width, height: bf.height}
}
