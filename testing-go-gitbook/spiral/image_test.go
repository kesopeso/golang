package spiral_test

import (
	"bytes"
	// "image/color"
	"image/jpeg"
	"spiral"
	"testing"
)

func TestHandleSpiralData(t *testing.T) {
	imgSize := 5
	var buf bytes.Buffer
	ish := spiral.NewImageSpiralHandler(imgSize, 100, &buf)
	points := []spiral.Point{
		{0, 2.5},
		{0, 0},
	}
	err := ish.HandleSpiralData(points)
	if err != nil {
		t.Fatal("error should not occur", err)
	}

	img, err := jpeg.Decode(&buf)
	if err != nil {
		t.Fatal("error should not occur", err)
	}

	// check dimensions
	imgRowCount := img.Bounds().Dy()
	imgColumnCount := img.Bounds().Dx()
	if imgColumnCount != imgSize || imgRowCount != imgSize {
		t.Fatalf("expected image size %d x %d, got %d x %d", imgSize, imgSize, imgColumnCount, imgRowCount)
	}

	// shouldBeBlack := func(x, y int) bool {
	// 	blackPointsMap := map[[2]int]bool{
	// 		{0, 2}: true,
	// 		{2, 2}: true,
	// 	}
	// 	return blackPointsMap[[2]int{x, y}]
	// }

	// for x := 0; x < imgColumnCount; x++ {
	// 	for y := 0; y < imgRowCount; y++ {
	// 		gotR, gotG, gotB, gotA := img.At(x, y).RGBA()
	// 		wantR, wantG, wantB, wantA := color.White.RGBA()
	// 		if shouldBeBlack(x, y) {
	// 			wantR, wantG, wantB, wantA = color.Black.RGBA()
	// 		}
	// 		if gotR != wantR || gotG != wantG || gotB != wantB || gotA != wantA {
	// 			t.Errorf("color mismatch at %d:%d, got %v, want %v", x, y, [4]uint32{gotR, gotG, gotB, gotA}, [4]uint32{wantR, wantG, wantB, wantA})
	// 		}
	// 	}
	// }

}
