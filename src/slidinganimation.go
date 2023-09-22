package main

import "time"

type SlidingAnimation struct {
	left     bool
	v        MetrePerSecond
	obj      Slidable
	lastTime time.Time
	distance Metre
}

func NewLeftSlidingAnimation(obj Slidable, velocity MetrePerSecond) *SlidingAnimation {
	return newSlidingAnimation(obj, velocity, true)
}

func NewRightSlidingAnimation(obj Slidable, velocity MetrePerSecond) *SlidingAnimation {
	return newSlidingAnimation(obj, velocity, false)
}

func newSlidingAnimation(obj Slidable, velocity MetrePerSecond, left bool) *SlidingAnimation {
	return &SlidingAnimation{left: left, v: velocity, obj: obj, lastTime: time.Now(), distance: 0}
}

func (sa *SlidingAnimation) Animate(f Field) {
	elapsed := time.Since(sa.lastTime)
	sa.lastTime = sa.lastTime.Add(elapsed)
	sa.distance += GetDistanceTravelled(elapsed, sa.v)
	scaledDistance := ScaleDistance(sa.distance)
	if scaledDistance != 0 {
		sa.distance -= Metre(scaledDistance)
		if sa.left {
			scaledDistance *= -1
		}
		sa.obj.Slide(scaledDistance)
	}
}
