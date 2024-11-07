package spiral

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"math"
)

type ImageSpiralHandler struct {
	size    int
	quality int
	out     io.Writer
}

func (ish ImageSpiralHandler) HandleSpiralData(spiralPoints []Point) error {
	img := image.NewRGBA(image.Rect(0, 0, ish.size, ish.size))

	// fill background with while color
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// fill image with spiral points
	r := getStartingR(spiralPoints)
	scale := float64(ish.size-1) / (2 * r)
	offset := float64(ish.size-1) / 2
	for _, p := range spiralPoints {
		x := getImageCoordinate(p.X, scale, offset, false)
		y := getImageCoordinate(p.Y, scale, offset, true)
		img.Set(x, y, color.Black)
	}

	jpegOptions := &jpeg.Options{Quality: ish.quality}
	err := jpeg.Encode(ish.out, img, jpegOptions)
	return err
}

func getStartingR(spiralPoints []Point) float64 {
	var r float64
	for _, sp := range spiralPoints {
		x := math.Abs(sp.X)
		y := math.Abs(sp.Y)
		if x > r {
			r = x
		}
		if y > r {
			r = y
		}
	}
	return r
}

func getImageCoordinate(value, scale, offset float64, flipAxis bool) int {
	if flipAxis {
		value *= -1
	}
	return int(math.Floor(value*scale + offset))
}

func NewImageSpiralHandler(size, quality int, out io.Writer) ImageSpiralHandler {
	return ImageSpiralHandler{size, quality, out}
}
