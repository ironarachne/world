package geometry

import "testing"

func TestGetUniqueEdges(t *testing.T) {
	edges := []Edge{
		{
			A: Point{
				X: 1,
				Y: 1,
			},
			B: Point{
				X: 3,
				Y: 1,
			},
		},
		{
			A: Point{
				X: 1,
				Y: 1,
			},
			B: Point{
				X: 3,
				Y: 1,
			},
		},
		{
			A: Point{
				X: 3,
				Y: 1,
			},
			B: Point{
				X: 1,
				Y: 1,
			},
		},
		{
			A: Point{
				X: 2,
				Y: 2,
			},
			B: Point{
				X: 4,
				Y: 2,
			},
		},
	}

	expectedCount := 2

	uniqueEdges := GetUniqueEdges(edges)

	if len(uniqueEdges) != expectedCount {
		t.Error("failed to find accurate number of unique edges")
	}
}

func TestIsEdgeEqual(t *testing.T) {
	edgeA := Edge{
		A: Point{
			X: 3,
			Y: 1,
		},
		B: Point{
			X: 2,
			Y: 4,
		},
	}

	edgeB := Edge{
		A: Point{
			X: 3,
			Y: 1,
		},
		B: Point{
			X: 2,
			Y: 4,
		},
	}

	edgeC := Edge{
		A: Point{
			X: 2,
			Y: 4,
		},
		B: Point{
			X: 3,
			Y: 1,
		},
	}

	edgeD := Edge{
		A: Point{
			X: 5,
			Y: 4,
		},
		B: Point{
			X: 6,
			Y: 7,
		},
	}

	if edgeA.Equals(edgeB) != true {
		t.Error("failed to detect equivalent edges")
	}

	if edgeA.Equals(edgeC) != true {
		t.Error("failed to detect equivalent edges")
	}

	if edgeA.Equals(edgeD) == true {
		t.Error("failed to recognize different edges")
	}
}

func TestIsEdgeInSlice(t *testing.T) {
	edges := []Edge{
		{
			A: Point{
				X: 1,
				Y: 1,
			},
			B: Point{
				X: 3,
				Y: 1,
			},
		},
		{
			A: Point{
				X: 1,
				Y: 1,
			},
			B: Point{
				X: 3,
				Y: 1,
			},
		},
		{
			A: Point{
				X: 3,
				Y: 1,
			},
			B: Point{
				X: 1,
				Y: 1,
			},
		},
		{
			A: Point{
				X: 2,
				Y: 2,
			},
			B: Point{
				X: 4,
				Y: 2,
			},
		},
	}

	edgeA := Edge{
		A: Point{
			X: 1,
			Y: 1,
		},
		B: Point{
			X: 3,
			Y: 1,
		},
	}

	edgeB := Edge{
		A: Point{
			X: 10,
			Y: 5,
		},
		B: Point{
			X: 3,
			Y: 1,
		},
	}

	if edgeA.InSlice(edges) != true {
		t.Error("failed to detect edge present in slice")
	}

	if edgeB.InSlice(edges) == true {
		t.Error("failed to detect edge was not present in slice")
	}
}
