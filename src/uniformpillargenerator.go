package main

type UniformPillarGenerator struct {
	distance int
	gapSize  int
}

func NewUniformPillarGenerator(distance, gapSize int) *UniformPillarGenerator {
	return &UniformPillarGenerator{distance: distance, gapSize: gapSize}
}

func (pg *UniformPillarGenerator) Generate(f Field, pc PillarContainer) Object {
	lastPos := pc.GetLastPillarPosition()
	newPillar := NewPillar(int(f.Height()/2), pg.gapSize, 0, lastPos+pg.distance)
	return newPillar
}
