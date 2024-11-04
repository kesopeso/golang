package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func secondsInRadiants(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func minutesInRadiants(t time.Time) float64 {
	return (secondsInRadiants(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
}

func hoursInRadiants(t time.Time) float64 {
	return (minutesInRadiants(t) / 12) + (math.Pi / (6 / float64(t.Hour()%12)))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadiants(t)
	return angleToPoint(angle)
}

func minuteHandPoint(t time.Time) Point {
	angle := minutesInRadiants(t)
	return angleToPoint(angle)
}

func hourHandPoint(t time.Time) Point {
	angle := hoursInRadiants(t)
	return angleToPoint(angle)
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
