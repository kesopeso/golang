package spiral_test

import (
	"bytes"
	"spiral"
	"testing"
)

func TestWriteSpiral(t *testing.T) {
	var b bytes.Buffer
	spiral.WriteSpiral(&b, 4, 4, 1)

	want := `0.0000000000#4.0000000000
3.0000000000#0.0000000000
0.0000000000#-2.0000000000
-1.0000000000#0.0000000000
0.0000000000#0.0000000000`
	got := b.String()

	if got != want {
		t.Errorf("draw result mismatch, got %v, want %v", got, want)
	}
}
