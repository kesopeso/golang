package clockface

import (
	"fmt"
	"io"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 70
	clockCentreX     = 150
	clockCentreY     = 150

	svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
	bezel  = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
	svgEnd = `</svg>`
)

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

func secondHand(w io.Writer, t time.Time) {
	p := secondHandPoint(t)
	clockHand(w, "#f00", p, secondHandLength)
}

func minuteHand(w io.Writer, t time.Time) {
	p := minuteHandPoint(t)
	clockHand(w, "#000", p, minuteHandLength)
}

func hourHand(w io.Writer, t time.Time) {
	p := hourHandPoint(t)
	clockHand(w, "#0f0", p, hourHandLength)
}

func clockHand(w io.Writer, stroke string, p Point, handLength float64) {
	p = Point{p.X * handLength, p.Y * handLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:%s;stroke-width:3px;"/>`, p.X, p.Y, stroke)
}
