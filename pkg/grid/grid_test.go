package grid

import "testing"

func TestGenerate(t *testing.T) {
	grid := Generate(10, 10)

	if len(grid.Coordinates) != 10 {
		t.Error("failed to generate correct number of rows of coordinates")
	}

	if len(grid.Coordinates[0]) != 10 {
		t.Error("failed to generate correct number of columns of coordinates")
	}

	if grid.Height != 10 {
		t.Error("failed to set grid height to 10")
	}

	if grid.Width != 10 {
		t.Error("failed to set grid width to 10")
	}
}
