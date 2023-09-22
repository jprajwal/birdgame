package main

type BasicPillarContainer struct {
	pillars []Object
	obj     BasicAnimatableObject
	pg      PillarGenerator
}

func NewBasicPillarContainer() *BasicPillarContainer {
	return &BasicPillarContainer{
		pillars: make([]Object, 0),
		obj:     *NewBasicAnimatableObject(NewBasicObject(nil, 0, 0)),
	}
}

func (pc *BasicPillarContainer) GetPositionX() int {
	return pc.obj.GetPositionX()
}

func (pc *BasicPillarContainer) GetPositionY() int {
	return pc.obj.GetPositionY()
}

func (pc *BasicPillarContainer) SetPositionX(x int) {
	pc.obj.SetPositionX(x)
}

func (pc *BasicPillarContainer) SetPositionY(y int) {
	pc.obj.SetPositionY(y)
}

func (pc *BasicPillarContainer) Draw(f Field) {
	for lastPos := pc.GetLastPillarPosition(); lastPos < f.Width(); lastPos = pc.GetLastPillarPosition() {
		newPillar := pc.pg.Generate(f, pc)
		pc.pillars = append(pc.pillars, newPillar)
	}
	for i := 0; i < len(pc.pillars); i++ {
		pc.pillars[i].Draw(f)
	}
}

func (pc *BasicPillarContainer) Animate(f Field) {
	pc.obj.Animate(f)
}

func (pc *BasicPillarContainer) AddAnimation(a Animatable) {
	pc.obj.AddAnimation(a)
}

func (pc *BasicPillarContainer) Slide(by int) {
	toRemove := make([]int, 0)
	for i := 0; i < len(pc.pillars); i++ {
		object := pc.pillars[i]
		object.SetPositionY(object.GetPositionY() + by)
		if object.GetPositionX() < 0 {
			toRemove = append(toRemove, i)
		}
	}
	for i := len(toRemove) - 1; i >= 0; i-- {
		index := toRemove[i]
		pc.pillars = RemoveFromSlice(pc.pillars, index)
	}
}

// func (pc *BasicPillarContainer) GetPillars() []Object {
// 	toReturn := make([]Object, len(pc.pillars))
// 	for i := 0; i < len(pc.pillars); i++ {
// 		toReturn[i] = pc.pillars[i]
// 	}
// 	return toReturn
// }

func (pc *BasicPillarContainer) GetLastPillarPosition() int {
	if len(pc.pillars) == 0 {
		return 0
	}
	return pc.pillars[len(pc.pillars)-1].GetPositionY()
}

func (pc *BasicPillarContainer) RegisterPillarGenerator(pg PillarGenerator) {
	pc.pg = pg
}
