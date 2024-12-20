package spiral

type SpiralHandler interface {
	HandleSpiralData(spiralPoints []Point) error
}

type SpiralData struct {
	R               float64
	TotalIterations int
	TotalCircles    int
}

func WriteSpiral(sh SpiralHandler, sd SpiralData) error {
	var spiralPoints []Point
	for i := 0; i <= sd.TotalIterations; i++ {
		spiralPoints = append(spiralPoints, newSpiralPoint(i, sd))
	}
	err := sh.HandleSpiralData(spiralPoints)
	return err
}

func NewSpiralData(r float64, totalIterations, totalCircles int) SpiralData {
	return SpiralData{r, totalIterations, totalCircles}
}

func newSpiralPoint(currentIteration int, sd SpiralData) Point {
	r := sd.R * float64(sd.TotalIterations-currentIteration) / float64(sd.TotalIterations)
	point := NewPoint(currentIteration, sd.TotalIterations, sd.TotalCircles)
	return Point{r * point.X, r * point.Y}
}
