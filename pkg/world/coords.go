package world

type coord struct {
	x, y int
}

func (i coord) adjacent(width int, height int) []coord {
	var adjacent []coord

	if i.x > 0 {
		adjacent = append(adjacent, coord{x: i.x-1, y: i.y})
	}

	if i.x < width-1 {
		adjacent = append(adjacent, coord{x: i.x+1, y: i.y})
	}

	if i.y > 0 {
		adjacent = append(adjacent, coord{x: i.x, y: i.y-1})
	}

	if i.y < height-1 {
		adjacent = append(adjacent, coord{x: i.x, y: i.y+1})
	}

	return adjacent
}

func coordInSlice(c coord, target []coord) bool {
	for _, t := range target {
		if c.x == t.x && c.y == t.y {
			return true
		}
	}

	return false
}
