package race

func getAllRaces() []Race {
	races := []Race{}

	races = append(races, getDwarves()...)
	races = append(races, getElves()...)
	races = append(races, getHalflings()...)
	races = append(races, getHumans()...)

	return races
}
