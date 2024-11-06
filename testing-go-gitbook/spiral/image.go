package spiral

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"math"
)

const (
	whiteIndex = 0
	blackIndex = 1
)

var pallete = []color.Color{color.White, color.Black}

type ImageSpiralHandler struct {
	size    int
	quality int
	out     io.Writer
}

func (ish ImageSpiralHandler) HandleSpiralData(r float64, spiralPoints []Point) error {
	img := image.NewRGBA(image.Rect(0, 0, ish.size, ish.size))

	// fill background with while color
	backgroundColor := pallete[whiteIndex]
	draw.Draw(img, img.Bounds(), &image.Uniform{backgroundColor}, image.Point{}, draw.Src)

	// fill image with spiral points
	scale := float64(ish.size-1) / (2 * r)
	offset := float64(ish.size-1) / 2
	lineColor := pallete[blackIndex]
	for _, p := range spiralPoints {
		x := getImageCoordinate(p.X, scale, offset, false)
		y := getImageCoordinate(p.Y, scale, offset, true)
		img.Set(x, y, lineColor)
	}

	jpegOptions := &jpeg.Options{Quality: ish.quality}
	err := jpeg.Encode(ish.out, img, jpegOptions)
	return err
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
