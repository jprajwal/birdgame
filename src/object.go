package main


type BasicObject struct {
	data *BasicField
	posX int
	posY int
}

func NewBasicObject(data [][]rune, posX, posY int) *BasicObject {
	width := len(data)
	height := 0
	for i := 0; i < width; i++ {
		if len(data[i]) > height {
			height = len(data[i])
		}
	}
	field := NewBasicField(width, height)
	for i := 0; i < field.Width(); i++ {
		for j := 0; j < field.Height(); j++ {
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
