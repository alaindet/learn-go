package clockface

import (
	"testing"
	"time"
)

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1970, time.January, 1, 0, 0, 30, 0, time.UTC)

	result := SecondHand(tm)
	expected := Point{X: Center, Y: Center + SecondHandLength}

	if result != expected {
		t.Errorf("Got %v expected %v", result, expected)
	}
}
