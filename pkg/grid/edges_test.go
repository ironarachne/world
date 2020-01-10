package grid

import "testing"

func TestGetUniqueEdges(t *testing.T) {
	edges := []Edge{
		{
			A: Coordinate{
				X: 1,
				Y: 1,
			},
			B: Coordinate{
				X: 3,
				Y: 1,
			},
		},
		{
			A: Coordinate{
				X: 1,
				Y: 1,
			},
			B: Coordinate{
				X: 3,
				Y: 1,
			},
		},
		{
			A: Coordinate{
				X: 3,
				Y: 1,
			},
			B: Coordinate{
				X: 1,
				Y: 1,
			},
		},
		{
			A: Coordinate{
				X: 2,
				Y: 2,
			},
			B: Coordinate{
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
		A: Coordinate{
			X: 3,
			Y: 1,
		},
		B: Coordinate{
			X: 2,
			Y: 4,
		},
	}

	edgeB := Edge{
		A: Coordinate{
			X: 3,
			Y: 1,
		},
		B: Coordinate{
			X: 2,
			Y: 4,
		},
	}

	edgeC := Edge{
		A: Coordinate{
			X: 2,
			Y: 4,
		},
		B: Coordinate{
			X: 3,
			Y: 1,
		},
	}

	edgeD := Edge{
		A: Coordinate{
			X: 5,
			Y: 4,
		},
		B: Coordinate{
			X: 6,
			Y: 7,
		},
	}

	if isEdgeEqual(edgeA, edgeB) != true {
		t.Error("failed to detect equivalent edges")
	}

	if isEdgeEqual(edgeA, edgeC) != true {
		t.Error("failed to detect equivalent edges")
	}

	if isEdgeEqual(edgeA, edgeD) == true {
		t.Error("failed to recognize different edges")
	}
}

func TestIsEdgeInSlice(t *testing.T) {
	edges := []Edge{
		{
			A: Coordinate{
				X: 1,
				Y: 1,
			},
			B: Coordinate{
				X: 3,
				Y: 1,
			},
		},
		{
			A: Coordinate{
				X: 1,
				Y: 1,
			},
			B: Coordinate{
				X: 3,
				Y: 1,
			},
		},
		{
			A: Coordinate{
				X: 3,
				Y: 1,
			},
			B: Coordinate{
				X: 1,
				Y: 1,
			},
		},
		{
			A: Coordinate{
				X: 2,
				Y: 2,
			},
			B: Coordinate{
				X: 4,
				Y: 2,
			},
		},
	}

	edgeA := Edge{
		A: Coordinate{
			X: 1,
			Y: 1,
		},
		B: Coordinate{
			X: 3,
			Y: 1,
		},
	}

	edgeB := Edge{
		A: Coordinate{
			X: 10,
			Y: 5,
		},
		B: Coordinate{
			X: 3,
			Y: 1,
		},
	}

	if isEdgeInSlice(edgeA, edges) != true {
		t.Error("failed to detect edge present in slice")
	}

	if isEdgeInSlice(edgeB, edges) == true {
		t.Error("failed to detect edge was not present in slice")
	}
}
