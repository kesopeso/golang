package clockface

import (
	"math"
	"testing"
	"time"
)

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(3, 0, 0), Point{1, 0}},
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(9, 0, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hourHandPoint(c.time)
			want := c.point
			if !roughlyEqualPoint(got, want) {
				t.Fatalf("got %+v, want %+v", got, want)
			}
		})
	}
}

func TestHoursInRadiants(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(18, 0, 0), math.Pi},
		{simpleTime(21, 0, 0), (math.Pi / 2) * 3},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hoursInRadiants(c.time)
			want := c.angle
			if got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestMinutesInRadiants(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 45, 0), (math.Pi / 2) * 3},
		{simpleTime(0, 7, 0), (math.Pi / 30) * 7},
		{simpleTime(0, 7, 7), (7 * (math.Pi / 30)) + (7 * (math.Pi / (30 * 60)))},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minutesInRadiants(c.time)
			want := c.angle
			if got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minuteHandPoint(c.time)
			want := c.point
			if !roughlyEqualPoint(got, want) {
				t.Fatalf("got %+v, want %+v", got, want)
			}
		})
	}
}

func TestSecondsInRadiants(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondsInRadiants(c.time)
			want := c.angle
			if got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondHandPoint(c.time)
			want := c.point
			if !roughlyEqualPoint(got, want) {
				t.Fatalf("got %+v, want %+v", got, want)
			}
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func roughlyEqualFloat64(a, b float64) bool {
	const eqThreshold = 1e-7
	return math.Abs(a-b) < eqThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}
