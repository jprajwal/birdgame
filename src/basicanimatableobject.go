package main

type BasicAnimatableObject struct {
	object     Object
	animations []Animatable
}

func NewBasicAnimatableObject(object Object) *BasicAnimatableObject {
	return &BasicAnimatableObject{object: object, animations: make([]Animatable, 0)}
}

func (o *BasicAnimatableObject) GetPositionX() int {
	return o.object.GetPositionX()
}

func (o *BasicAnimatableObject) GetPositionY() int {
	return o.object.GetPositionY()
}

func (o *BasicAnimatableObject) SetPositionX(x int) {
	o.object.SetPositionX(x)
}

func (o *BasicAnimatableObject) SetPositionY(y int) {
	o.object.SetPositionY(y)
}

func (o *BasicAnimatableObject) Draw(f Field) {
	o.object.Draw(f)
}

func (o *BasicAnimatableObject) Animate(f Field) {
	for i := 0; i < len(o.animations); i++ {
		o.animations[i].Animate(f)
	}
}

func (o *BasicAnimatableObject) AddAnimation(animation Animatable) {
	o.animations = append(o.animations, animation)
}
