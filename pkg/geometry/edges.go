package geometry

// Edge is an undirected line
type Edge struct {
	A Point
	B Point
}

// Length returns the length of an edge
func (edge Edge) Length() float64 {
	distance := Distance(edge.A, edge.B)

	return distance
}

// Slope returns the slope of an edge
func (edge Edge) Slope() float64 {
	slope := (edge.B.Y- edge.A.Y)/(edge.B.X- edge.A.X)

	return slope
}

// GetUniqueEdges returns all the unique edges in a slice
func GetUniqueEdges(edges []Edge) []Edge {
	var unique []Edge

	for _, e := range edges {
		if !e.InSlice(unique) {
			unique = append(unique, e)
		}
	}

	return unique
}

func (edge Edge) Equals(b Edge) bool {
	if (edge.A.Equals(b.A) && edge.B.Equals(b.B)) || (edge.A.Equals(b.B) && edge.B.Equals(b.A)) {
		return true
	}
	return false
}

func (edge Edge) InSlice(edges []Edge) bool {
	for _, e := range edges {
		if edge.Equals(e) {
			return true
		}
	}
	return false
}
