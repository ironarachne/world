package random

import (
	"context"
	"testing"
)

func TestIntFromThresholdMap(t *testing.T) {
	items := map[int]int{
		1: 5,
		2: 5,
		3: 1000,
	}

	randomItem, err := IntFromThresholdMap(context.Background(), items)
	if err != nil {
		t.Error("Failed to get random item from weighted int slice")
	}
	if randomItem != 3 {
		t.Errorf("Failed to get expected 'random' result from weighted int slice: got %d instead of 3", randomItem)
	}
}
