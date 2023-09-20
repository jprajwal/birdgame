package main

var BIRD = [][]rune{{'_', '_', '\\', '_', ','}}

type Bird struct {
	obj BasicAnimatableObject
	vv  MetrePerSecond
}

func NewBird(x, y int) *Bird {
	return &Bird{obj: *NewBasicAnimatableObject(NewBasicObject(copy2DSlice(BIRD), x, y))}
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
	b.vv = MetrePerSecond(BIRD_JUMP_VELOCITY)
}

func (b *Bird) GetVerticalVelocity() MetrePerSecond {
	return b.vv
}

func (b *Bird) SetVerticalVelocity(v MetrePerSecond) {
	b.vv = v
}

func (b *Bird) Animate(f Field) {
	b.obj.Animate(f)
}

func (b *Bird) AddAnimation(animation Animatable) {
	b.obj.AddAnimation(animation)
}
