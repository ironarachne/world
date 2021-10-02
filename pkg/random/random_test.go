package random

import (
	"context"
	"testing"
)

func TestNew(t *testing.T) {
	seed := "testing"

	r := New(seed)

	result := r.Intn(5)

	if result != 2 {
		t.Error(result)
	}
}

func TestWithSeed(t *testing.T) {
	ctx := context.Background()
	seed := "testing"

	newCtx := WithSeed(ctx, seed)

	r := Intn(newCtx, 5)

	if r != 2 {
		t.Error(r)
	}
}
