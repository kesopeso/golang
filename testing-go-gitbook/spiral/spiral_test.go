package spiral_test

import (
	"spiral"
	"testing"
)

type TestSpiralHandler struct {
	points []spiral.Point
}

func (tsh *TestSpiralHandler) HandleSpiralData(points []spiral.Point) error {
	tsh.points = points
	return nil
}

func newTestSpiralHandler() *TestSpiralHandler {
	return &TestSpiralHandler{}
}

func TestWriteSpiral(t *testing.T) {
	tsh := newTestSpiralHandler()
	sd := spiral.NewSpiralData(4, 4, 1)
	spiral.WriteSpiral(tsh, sd)

	want := []spiral.Point{
		{0, 4},
		{3, 0},
		{0, -2},
		{-1, 0},
		{0, 0},
	}
	got := tsh.points

	if !pointsRouglyEqual(got, want) {
		t.Errorf("result mismatch, got %v, want %v", got, want)
	}
}

func pointsRouglyEqual(a, b []spiral.Point) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !pointRoughlyEquals(a[i], b[i]) {
			return false
		}
	}

	return true
}
