package race_test

import (
	"ch15/race"
	"testing"
)

func TestRace(t *testing.T) {
	result := race.Race(false)
	if result != 5000 {
		t.Errorf("got: %d, want: 5000", result)
	}
}
