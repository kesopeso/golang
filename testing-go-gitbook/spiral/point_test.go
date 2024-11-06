package spiral_test

import (
	// "math"
	"fmt"
	"math"
	"spiral"
	"testing"
)

func TestAngleInRadians(t *testing.T) {
	cases := []struct {
		currentIteration int
		totalIterations  int
		totalCircles     int
		want             float64
	}{
		{0, -2, 4, 0},
		{1, 4, 1, math.Pi / 2},
		{2, 4, 1, math.Pi},
		{3, 4, 1, 3 * math.Pi / 2},
		{4, 4, 1, 0},
		{1, 4, 2, math.Pi},
		{2, 4, 2, 0},
		{3, 4, 2, math.Pi},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("test #%d", i), func(t *testing.T) {
			got := spiral.AngleInRadians(c.currentIteration, c.totalIterations, c.totalCircles)
			if got != c.want {
				t.Errorf("angle mismatch, got %v, want %v", got, c.want)
			}
		})
	}
}

func TestNewPoint(t *testing.T) {
	cases := []struct {
		currentIteration int
		totalIterations  int
		totalCircles     int
		want             spiral.Point
	}{
		{0, -2, 4, spiral.Point{0, 1}},
		{0, 4, 1, spiral.Point{0, 1}},
		{1, 4, 1, spiral.Point{1, 0}},
		{2, 4, 1, spiral.Point{0, -1}},
		{3, 4, 1, spiral.Point{-1, 0}},
		{4, 4, 1, spiral.Point{0, 1}},
		{1, 4, 2, spiral.Point{0, -1}},
		{2, 4, 2, spiral.Point{0, 1}},
		{3, 8, 2, spiral.Point{-1, 0}},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("test #%d", i), func(t *testing.T) {
			got := spiral.NewPoint(c.currentIteration, c.totalIterations, c.totalCircles)
			if !pointRoughlyEquals(got, c.want) {
				t.Errorf("point mismatch, got %v, want %v", got, c.want)
			}
		})
	}
}

func pointRoughlyEquals(a, b spiral.Point) bool {
	return roughlyEquals(a.X, b.X) && roughlyEquals(a.Y, b.Y)
}

func roughlyEquals(a, b float64) bool {
	const eqThreshold float64 = 1e-10
	return math.Abs(a-b) < eqThreshold
}
