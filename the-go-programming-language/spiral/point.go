package spiral

import "math"

type Point struct {
	X float64
	Y float64
}

func NewPoint(currentIteration, totalIterations, totalCircles int) Point {
	angle := AngleInRadians(currentIteration, totalIterations, totalCircles)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func AngleInRadians(currentIteration, totalIterations, totalCircles int) float64 {
	if totalCircles <= 0 || totalIterations <= 0 {
		return 0
	}

	if currentIteration > totalIterations {
		currentIteration = totalIterations
	}

	progress := float64(currentIteration) / float64(totalIterations)
	totalAngle := float64(totalCircles) * 2 * math.Pi
	result := progress * totalAngle
	result = math.Mod(result, 2*math.Pi)

	return result
}
