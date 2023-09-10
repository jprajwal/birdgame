package main

type Space struct {
	buf     [][]rune
	rowSize int
	colSize int
}

func NewSpace(rowSize, colSize int) *Space {
	space := make([][]rune, rowSize)
	for i := 0; i < rowSize; i++ {
		space[i] = make([]rune, colSize)
	}
	toReturn := &Space{buf: space, rowSize: rowSize, colSize: colSize}
	toReturn.Clear()
	return toReturn
}

func (s *Space) Clear() {
	for i := 0; i < s.rowSize; i++ {
		for j := 0; j < s.colSize; j++ {
			s.buf[i][j] = ' '
		}
	}
}

func (s *Space) RowSize() int {
	return s.rowSize
}

func (s *Space) ColSize() int {
	return s.colSize
}

func (s *Space) SetDataAt(i, j int, data rune) {
	if i < 0 || i >= s.rowSize || j < 0 || j >= s.colSize {
		return
	}
	s.buf[i][j] = data
}

func (s *Space) GetDataAsString() string {
	data := ""
	for i := 0; i < s.rowSize; i++ {
		if i > 0 {
			data += "\n"
		}
		for j := 0; j < s.colSize; j++ {
			data += string(s.buf[i][j])
		}
	}
	return data

}
