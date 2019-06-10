package climate

// SimplifiedClimate is a simplified version of Climate for display purposes
type SimplifiedClimate struct {
	Name        string
	Adjective   string
	Description string
}

// Simplify returns a simplified version of a Climate
func (climate Climate) Simplify() SimplifiedClimate {
	return SimplifiedClimate{
		Name:        climate.Name,
		Adjective:   climate.Adjective,
		Description: climate.Description,
	}
}
