package main

import "time"

type GravityAnimation struct {
	lastTime time.Time
	object   ObjectWithVerticalVelocity
	distance Metre
}

func NewGravityAnimation(object ObjectWithVerticalVelocity) *GravityAnimation {
	return &GravityAnimation{lastTime: time.Now(), object: object}
}

func (ga *GravityAnimation) Animate(_ Field) {
	elapsed := time.Since(ga.lastTime)
	ga.lastTime = ga.lastTime.Add(elapsed)
	UpdateGravity(elapsed, ga.object)
	ga.distance += GetDistanceTravelled(elapsed, ga.object.GetVerticalVelocity())
	scaledDistance := ScaleDistance(ga.distance)
	if scaledDistance != 0 {
		ga.distance = ga.distance - Metre(scaledDistance)
		ga.object.SetPositionX(ga.object.GetPositionX() - scaledDistance)
	}
}
