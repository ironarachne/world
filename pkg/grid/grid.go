package grid

// Grid is a grid
type Grid struct {
	Coordinates [][]Coordinate
	Height      int
	Width       int
}

// Generate procedurally generates a grid
func Generate(height int, width int) Grid {
	var coordinate Coordinate
	var coordinates [][]Coordinate
	var row []Coordinate

	grid := Grid{}

	for i := 0; i < height; i++ {
		row = []Coordinate{}
		for j := 0; j < width; j++ {
			coordinate = Coordinate{
				X: j,
				Y: i,
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
