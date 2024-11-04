package clockface_test

import (
	"bytes"
	"clockface"
	"encoding/xml"
	"testing"
	"time"
)

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

func TestSvgWriterHand(t *testing.T) {
	cases := []struct {
		hand string
		time time.Time
		line Line
	}{
		{
			"second",
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{
			"second",
			simpleTime(0, 0, 30),
			Line{150, 150, 150, 240},
		},
		{
			"minute",
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 70},
		},
		{
			"minute",
			simpleTime(0, 15, 0),
			Line{150, 150, 230, 150},
		},
		{
			"minute",
			simpleTime(0, 30, 0),
			Line{150, 150, 150, 230},
		},
		{
			"minute",
			simpleTime(0, 45, 0),
			Line{150, 150, 70, 150},
		},
		{
			"hour",
			simpleTime(6, 0, 0),
			Line{150, 150, 150, 220},
		},
		{
			"hour",
			simpleTime(12, 0, 0),
			Line{150, 150, 150, 80},
		},
		{
			"hour",
			simpleTime(9, 0, 0),
			Line{150, 150, 80, 150},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.hand, c.time), func(t *testing.T) {
			var b bytes.Buffer
			clockface.SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Error("line was not found", c.line)
			}
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(hand string, t time.Time) string {
	return hand + " hand, " + t.Format("15:04:05")
}

func containsLine(line Line, lines []Line) bool {
	for _, l := range lines {
		if line == l {
			return true
		}
	}
	return false
}
