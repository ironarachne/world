package grid

// Edge is an undirected line
type Edge struct {
	A Coordinate
	B Coordinate
}

// GetUniqueEdges returns all the unique edges in a slice
func GetUniqueEdges(edges []Edge) []Edge {
	unique := []Edge{}

	for _, e := range edges {
		if !isEdgeInSlice(e, unique) {
			unique = append(unique, e)
		}
	}

	return unique
}

func isEdgeEqual(a Edge, b Edge) bool {
	if (Equals(a.A, b.A) && Equals(a.B, b.B)) || (Equals(a.A, b.B) && Equals(a.B, b.A)) {
		return true
	}
	return false
}

func isEdgeInSlice(edge Edge, edges []Edge) bool {
	for _, e := range edges {
		if isEdgeEqual(edge, e) {
			return true
		}
	}
	return false
}
