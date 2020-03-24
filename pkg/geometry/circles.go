package geometry

// Circle is a circle
type Circle struct {
	Center Point
	Radius float64
}

// InCircle checks to see if the given point is in the given circle
func (p Point) InCircle(c Circle) bool {
	if ((p.X - c.Center.X) * (p.X - c.Center.X)) + ((p.Y - c.Center.Y) * (p.Y - c.Center.Y)) < c.Radius * c.Radius {
		return true
	}

	return false
}
