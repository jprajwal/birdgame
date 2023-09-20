package main

import "time"

func UpdateGravity(interval time.Duration, object HasVerticalVelocity) {
	object.SetVerticalVelocity(object.GetVerticalVelocity() - MetrePerSecond(G*interval.Seconds()))
}

func GetDistanceTravelled(interval time.Duration, velocity MetrePerSecond) Metre {
	return Metre(float64(velocity) * float64(interval.Seconds()))
}
