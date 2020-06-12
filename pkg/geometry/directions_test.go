package geometry

import (
	"context"
	"testing"
)

func TestRandomDirection(t *testing.T) {
	d := RandomDirection(context.Background())

	if d > 7 || d < 0 {
		t.Error("generated out of bounds direction")
	}
}

func TestOppositeDirection(t *testing.T) {
	o := OppositeDirection(WEST)
	if o != EAST {
		t.Error("got incorrect direction for opposite direction")
	}

	o = OppositeDirection(NORTH)
	if o != SOUTH {
		t.Error("failed to get south as opposite of north")
	}

	o = OppositeDirection(SOUTH)
	if o != NORTH {
		t.Error("failed to get north as opposite of south")
	}
}
