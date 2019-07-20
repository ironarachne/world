package climate

// Season is a season
type Season struct {
	Name              string
	TemperatureChange int
	HumidityChange    int
	WeatherProfile    WeatherProfile
}

func (climate Climate) getSeasons() []Season {
	var seasons []Season
	var newSeason Season

	fourSeasons := []Season{
		{
			Name:              "spring",
			HumidityChange:    0,
			TemperatureChange: 0,
		},
		{
			Name:              "summer",
			HumidityChange:    1,
			TemperatureChange: 3,
		},
		{
			Name:              "autumn",
			HumidityChange:    -1,
			TemperatureChange: 1,
		},
		{
			Name:              "winter",
			HumidityChange:    -2,
			TemperatureChange: -2,
		},
	}

	twoSeasons := []Season{
		{
			Name:              "dry",
			HumidityChange:    -1,
			TemperatureChange: +1,
		},
		{
			Name:              "wet",
			HumidityChange:    +1,
			TemperatureChange: 0,
		},
	}

	if climate.Humidity > 5 && climate.Temperature > 6 {
		for _, s := range twoSeasons {
			newSeason = s
			newSeason.WeatherProfile = climate.generateWeatherProfileForSeason(s)
			seasons = append(seasons, newSeason)
		}
		return seasons
	}

	for _, s := range fourSeasons {
		newSeason = s
		newSeason.WeatherProfile = climate.generateWeatherProfileForSeason(s)
		seasons = append(seasons, newSeason)
	}

	return seasons
}
