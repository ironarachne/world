package species

// Simplified is a simplified version of a species meant for JSON display
type Simplified struct {
	Name        string `json:"name"`
	PluralName  string `json:"plural_name"`
	Adjective   string `json:"adjective"`
	Description string `json:"description"`
}

// Describe returns a description of a species
func (s Species) Describe() string {
	// TODO: Find a way to describe a species
	return ""
}

// Simplify returns a simplified version of a given species
func (s Species) Simplify() Simplified {
	description := s.Describe()

	sr := Simplified{
		Name:        s.Name,
		PluralName:  s.PluralName,
		Adjective:   s.Adjective,
		Description: description,
	}

	return sr
}
