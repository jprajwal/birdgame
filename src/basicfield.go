package main

type BasicField struct {
	buf    [][]Cell
	width  int
	height int
}

func NewBasicField(width, height int) *BasicField {
	field := &BasicField{buf: make2DSlice[Cell](height, width), width: width, height: height}
	field.Fill(NewDefaultBasicCell())
	return field
}

func (bf *BasicField) Fill(c Cell) {
	for i := 0; i < bf.height; i++ {
		for j := 0; j < bf.width; j++ {
			bf.buf[i][j] = c.Clone()
		}
	}
}

func (bf *BasicField) Clear() {
	for i := 0; i < bf.height; i++ {
		for j := 0; j < bf.width; j++ {
			bf.buf[i][j].Clear()
		}
	}
}

func (bf *BasicField) CopyFieldAt(x, y int, f Field) {
	for i := 0; i < f.Height(); i++ {
		for j := 0; j < f.Width(); j++ {
			if x+i >= bf.height || y+j >= bf.width || x+i < 0 || y+j < 0 {
				continue
			}
			bf.buf[x+i][y+j] = f.GetContent(i, j)
		}
	}
}

func (bf *BasicField) GetContent(x, y int) Cell {
	if x < 0 || x >= bf.height || y < 0 || y >= bf.width {
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
	return &BasicField{buf: copy2DSlice[Cell](bf.buf), width: bf.width, height: bf.height}
}
