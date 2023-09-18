package main

type BasicCell struct {
	r rune
}

func NewBasicCell(r rune) *BasicCell {
	return &BasicCell{r: r}
}

func NewDefaultBasicCell() *BasicCell {
	return &BasicCell{}
}

func (bc *BasicCell) String() string {
	return string(bc.r)
}

func (bc *BasicCell) Clear() {
	bc.r = rune(0)
}

func (bc *BasicCell) Set(r rune) {
	bc.r = r
}

func (bc *BasicCell) Get() rune {
	return bc.r
}

func (bc *BasicCell) Clone() Cell {
	return NewBasicCell(bc.r)
}
