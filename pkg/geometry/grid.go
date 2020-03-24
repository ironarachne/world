/*
Package geometry provides a geometry coordinate system and attendant functions for
dealing with grids.
*/
package geometry

// Grid is a geometry
type Grid struct {
	Coordinates [][]Point
	Height      int
	Width       int
}

// GenerateGrid procedurally generates a grid
func GenerateGrid(height int, width int) Grid {
	var coordinate Point
	var coordinates [][]Point
	var row []Point

	grid := Grid{}

	for i := 0; i < height; i++ {
		row = []Point{}
		for j := 0; j < width; j++ {
			coordinate = Point{
				X: float64(j),
				Y: float64(i),
			}
			row = append(row, coordinate)
		}
		coordinates = append(coordinates, row)
	}

	grid.Coordinates = coordinates
	grid.Height = height
	grid.Width = width

	return grid
}
