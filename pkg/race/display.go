package race

// SimplifiedRace is a simplified version of a race for display
type SimplifiedRace struct {
	Name      string
	Adjective string
}

// Simplify returns a simplified version of a race
func (race Race) Simplify() SimplifiedRace {
	return SimplifiedRace{
		Name:      race.Name,
		Adjective: race.Adjective,
	}
}
