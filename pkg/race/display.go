package race

// SimplifiedRace is a simplified version of a race for display
type SimplifiedRace struct {
	Name        string `json:"name"`
	Adjective   string `json:"adjective"`
	Description string `json:"description"`
}

// Simplify returns a simplified version of a race
func (race Race) Simplify() SimplifiedRace {
	return SimplifiedRace{
		Name:        race.Name,
		Adjective:   race.Adjective,
		Description: race.Describe(),
	}
}

// Describe returns a description of a race
func (race Race) Describe() string {
	description := ""

	description += race.PluralName + " are " + race.Size.Name + ", averaging " + heightToString(race.getAverageHeight()) + " tall. "

	return description
}
