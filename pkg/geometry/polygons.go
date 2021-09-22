package geometry

// Polygon is a polygon
type Polygon struct {
	Edges []Edge
	Points []Point
}

// PolygonFromPoints returns a polygon built from the given slice of points
func PolygonFromPoints(points []Point) Polygon {
	var edge Edge
	var polygon Polygon

	sorted := SortPointsByDistance(points)

	for i:=0;i<len(sorted);i++ {
		if i<len(sorted)-1 {
			edge = Edge{
				A: Point{
					X: sorted[i].X,
					Y: sorted[i].Y,
				},
				B: Point{
					X: sorted[i+1].X,
					Y: sorted[i+1].Y,
				},
			}
		} else {
			edge = Edge{
				A: Point{
					X: sorted[i].X,
					Y: sorted[i].Y,
				},
				B: Point{
					X: sorted[0].X,
					Y: sorted[0].Y,
				},
			}
		}
		polygon.Edges = append(polygon.Edges, edge)
	}

	polygon.Points = sorted

	return polygon
}
