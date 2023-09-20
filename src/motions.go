package main

import "time"

func UpdateGravity(interval time.Duration, velocity MetrePerSecond) MetrePerSecond {
	return velocity - MetrePerSecond(G*interval.Seconds())
}

func GetDistanceTravelled(interval time.Duration, velocity MetrePerSecond) Metre {
	return Metre(float64(velocity) * float64(interval.Seconds()))
}
