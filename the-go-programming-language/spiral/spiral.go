package spiral

import (
	"fmt"
	"io"
	"math"
)

func WriteSpiral(w io.Writer, startingR float64, totalIterations, totalCircles int) (n int, err error) {
	var resultText string
	for i := 0; i <= totalIterations; i++ {
		spiralPoint := newSpiralPoint(startingR, i, totalIterations, totalCircles)
		resultText += fmt.Sprintf("%.10f#%.10f", spiralPoint.X, spiralPoint.Y)
		if i != totalIterations {
			resultText += "\n"
		}
	}
	n, err = fmt.Fprint(w, resultText)
	return
}

func newSpiralPoint(startingR float64, currentIteration, totalIterations, totalCircles int) Point {
	r := startingR * float64(totalIterations-currentIteration) / float64(totalIterations)
	point := NewPoint(currentIteration, totalIterations, totalCircles)
	return Point{roundRoughlyZero(r * point.X), roundRoughlyZero(r * point.Y)}
}

func roundRoughlyZero(num float64) float64 {
	const zeroThreshold = 1e-10
	if math.Abs(num) < zeroThreshold {
		return 0
	}
	return num
}
