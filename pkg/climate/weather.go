package climate

import (
	"math/rand"
)

// WeatherProfile is a summary of a type of weather
type WeatherProfile struct {
	Description         string
	CloudCover          int
	CloudFrequency      int
	PrecipitationAmount int
	PrecipitationTypes  []string
	WindLevel           int
}

func (climate Climate) generateWeatherProfileForSeason(season Season) WeatherProfile {
	var weatherProfile WeatherProfile

	description := ""

	temp := climate.getCurrentTemperature(season)
	humidity := climate.getCurrentHumidity(season)

	if temp < 4 {
		weatherProfile.PrecipitationTypes = []string{"snow"}
	} else if temp < 5 {
		weatherProfile.PrecipitationTypes = []string{"snow", "sleet"}
	} else if temp < 6 {
		weatherProfile.PrecipitationTypes = []string{"hail", "sleet", "rain"}
	} else {
		weatherProfile.PrecipitationTypes = []string{"rain"}
	}

	if season.TemperatureChange > 1 {
		weatherProfile.CloudCover = 0
	} else {
		weatherProfile.CloudCover = rand.Intn(10)
	}

	weatherProfile.PrecipitationAmount = humidity

	weatherProfile.WindLevel = rand.Intn(10) + 1
	weatherProfile.CloudCover = weatherProfile.CloudCover + weatherProfile.PrecipitationAmount
	weatherProfile.CloudFrequency = weatherProfile.CloudCover + weatherProfile.PrecipitationAmount

	if weatherProfile.PrecipitationAmount > 5 && inSlice("rain", weatherProfile.PrecipitationTypes) {
		description = "rainy"
	} else if weatherProfile.PrecipitationAmount > 5 && inSlice("snow", weatherProfile.PrecipitationTypes) {
		description = "snowy"
	} else if temp > 3 && temp < 8 {
		description = "pleasant"
	} else if temp < 2 {
		description = "freezing"
	} else if temp < 4 {
		description = "cold"
	} else if temp > 8 && humidity > 7 {
		description = "sweltering"
	} else if temp > 8 {
		description = "very hot"
	} else if temp > 7 {
		description = "hot"
	} else {
		description = "mild"
	}

	description += " with "

	if weatherProfile.CloudCover < 5 {
		description += "lots of sun"
	} else if weatherProfile.CloudCover < 8 && weatherProfile.CloudFrequency < 8 {
		description += "some sun but frequent clouds"
	} else if weatherProfile.CloudCover < 10 && weatherProfile.CloudFrequency < 10 {
		description += "lots of cloudy days"
	} else {
		description += "many overcast days"
	}

	if weatherProfile.WindLevel > 8 {
		description += " and fierce winds"
	} else if weatherProfile.WindLevel > 6 {
		description += " and strong winds"
	} else if weatherProfile.WindLevel > 4 {
		description += " and light winds"
	}

	weatherProfile.Description = description

	return weatherProfile
}
