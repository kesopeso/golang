package spiral

type SpiralHandler interface {
	HandleSpiralData([]Point)
}

func WriteSpiral(sh SpiralHandler, startingR float64, totalIterations, totalCircles int) {
	var spiralPoints []Point
	for i := 0; i <= totalIterations; i++ {
		spiralPoints = append(spiralPoints, newSpiralPoint(startingR, i, totalIterations, totalCircles))
	}
	sh.HandleSpiralData(spiralPoints)
}

func newSpiralPoint(startingR float64, currentIteration, totalIterations, totalCircles int) Point {
	r := startingR * float64(totalIterations-currentIteration) / float64(totalIterations)
	point := NewPoint(currentIteration, totalIterations, totalCircles)
	return Point{r * point.X, r * point.Y}
}
