package spiral_test

import (
	"spiral"
	"testing"
)

type TestSpiralHandler struct {
	points []spiral.Point
}

func (tsh *TestSpiralHandler) HandleSpiralData(points []spiral.Point) {
	tsh.points = points
}

func newTestSpiralHandler() *TestSpiralHandler {
	return &TestSpiralHandler{}
}

func TestWriteSpiral(t *testing.T) {
	tsh := newTestSpiralHandler()
	spiral.WriteSpiral(tsh, 4, 4, 1)

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