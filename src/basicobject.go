package main

type BasicObject struct {
	data *BasicField
	posX int
	posY int
}

func NewBasicObject(data [][]rune, posX, posY int) *BasicObject {
	height := len(data)
	width := 0
	for i := 0; i < height; i++ {
		if len(data[i]) > width {
			width = len(data[i])
		}
	}
	field := NewBasicField(width, height)
	for i := 0; i < field.Height(); i++ {
		for j := 0; j < field.Width(); j++ {
			if c := field.GetContent(i, j); c != nil {
				c.Set(data[i][j])
			}
		}
	}
	return &BasicObject{data: field, posX: posX, posY: posY}
}

func (o *BasicObject) GetPositionX() int {
	return o.posX
}

func (o *BasicObject) GetPositionY() int {
	return o.posY
}

func (o *BasicObject) SetPositionX(x int) {
	o.posX = x
}

func (o *BasicObject) SetPositionY(y int) {
	o.posY = y
}

func (o *BasicObject) Draw(f Field) {
	f.CopyFieldAt(o.posX, o.posY, o.data)
}
