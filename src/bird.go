package main

var BIRD = [][]rune{{'_', '_', '\\', '_', ','}}

type Bird struct {
	obj BasicObject
}

func NewBird(x, y int) *Bird {
	return &Bird{obj: *NewBasicObject(copy2DSlice(BIRD), x, y)}
}

func (b *Bird) GetPositionX() int {
	return b.obj.GetPositionX()
}

func (b *Bird) GetPositionY() int {
	return b.obj.GetPositionY()
}

func (b *Bird) SetPositionX(x int) {
	b.obj.SetPositionX(x)
}

func (b *Bird) SetPositionY(y int) {
	b.obj.SetPositionY(y)
}

func (b *Bird) Draw(f Field) {
	b.obj.Draw(f)
}

func (b *Bird) OnSpaceKeyPressed() {
	b.obj.SetPositionX(b.obj.GetPositionX() - 1)
}
