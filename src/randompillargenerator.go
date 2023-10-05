package main

import "math/rand"

type RandomPillarGenerator struct {
	start, end, gapPosition, gapSize int
}

func NewRandomPillarGenerator(start, end, gapPosition, gapSize int) *RandomPillarGenerator {
	return &RandomPillarGenerator{start: start, end: end, gapPosition: gapPosition, gapSize: gapSize}
}

func (rpg *RandomPillarGenerator) Generate(f Field, pc PillarContainer) Object {
	pos := pc.GetLastPillarPosition() + rpg.start + rand.Intn(rpg.end-rpg.start)
	pillar := NewPillar(rpg.gapPosition, rpg.gapSize, 0, pos)
	return pillar
}
