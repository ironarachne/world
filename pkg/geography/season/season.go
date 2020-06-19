package season

import (
	"context"
	"fmt"
	"github.com/ironarachne/world/pkg/geography/climate"
	"github.com/ironarachne/world/pkg/geography/region"
	"github.com/ironarachne/world/pkg/words"
)

// Season is a meteorological season
type Season struct {
	Name                         string `json:"name"`
	Description                  string `json:"description"`
	TemperatureChange            int    `json:"temperature_change"`
	HumidityChange               int    `json:"humidity_change"`
	DaylightHours                int    `json:"daylight_hours"`
	PrecipitationType            string `json:"precipitation_type"`
	PrecipitationFrequencyChange int    `json:"precipitation_frequency"`
	PrecipitationAmountChange    int    `json:"precipitation_amount"`
	MonthBegin                   int    `json:"month_begin"`
	MonthEnd                     int    `json:"month_end"`
}

func (s Season) describe(c climate.Climate, r region.Region) (string, error) {
	var name string

	if s.Name == "dry" || s.Name == "wet" {
		name = "The " + s.Name + " season"
	} else {
		name = words.CapitalizeFirst(s.Name)
	}
	description := name + " is "

	temperatureWord := getWordForTemperature(r.Temperature + s.TemperatureChange)

	description += temperatureWord + ", with "

	cloudDescription := c.DescribeClouds()

	description += cloudDescription + ". The weather is "

	pAmount := s.PrecipitationAmountChange + c.PrecipitationAmount
	pFreq := s.PrecipitationFrequencyChange + c.PrecipitationFrequency

	if pFreq < 10 {
		description += "usually dry"
	} else if pFreq < 50 {
		description += "sometimes " + s.PrecipitationType + "y"
	} else if pFreq < 70 {
		description += "often " + s.PrecipitationType + "y, with "
		if pAmount < 30 {
			description += "light storms"
		} else if pAmount < 50 {
			description += "moderate storms"
		} else {
			description += "heavy storms"
		}
	} else {
		description += "persistently " + s.PrecipitationType + "y, with "
		if pAmount < 30 {
			description += "moderate storms"
		} else if pAmount < 50 {
			description += "heavy storms"
		} else {
			description += "fierce storms"
		}
	}

	daylightHours := words.NumberToWord(s.DaylightHours)

	description += ". There are " + daylightHours + " hours of daylight each day."

	return description, nil
}

func getDaylightHours(distance int, month int) int {
	modifiers := []float64{
		0.75,
		1,
		0.75,
		0.5,
		0,
		0.5,
		-0.75,
		-1,
		-0.75,
		0.5,
		0,
		0.5,
	}
	variance := (float64(distance) / 99) * 12
	monthModifier := modifiers[month-1]

	hours := int(12 + (variance * monthModifier))

	return hours
}

func getWordForTemperature(temperature int) string {
	if temperature < 10 {
		return "frigid"
	} else if temperature < 30 {
		return "cold"
	} else if temperature < 50 {
		return "mild"
	} else if temperature < 70 {
		return "warm"
	} else if temperature < 90 {
		return "hot"
	}

	return "really hot"
}

func getPrecipitationTypeForTemperature(temperature int) string {
	if temperature < 30 {
		return "snow"
	}

	return "rain"
}

// Generate generates a set of seasons based on a climate and region
func Generate(ctx context.Context, c climate.Climate, r region.Region) ([]Season, error) {
	var seasons []Season

	base := []Season{
		{
			Name:                         "spring",
			TemperatureChange:            0,
			HumidityChange:               0,
			PrecipitationFrequencyChange: +10,
			PrecipitationAmountChange:    +10,
			MonthBegin:                   3,
			MonthEnd:                     5,
		},
		{
			Name:                         "summer",
			TemperatureChange:            +10,
			HumidityChange:               0,
			PrecipitationFrequencyChange: 0,
			PrecipitationAmountChange:    0,
			MonthBegin:                   6,
			MonthEnd:                     8,
		},
		{
			Name:                         "autumn",
			TemperatureChange:            0,
			HumidityChange:               0,
			PrecipitationFrequencyChange: 0,
			PrecipitationAmountChange:    0,
			MonthBegin:                   9,
			MonthEnd:                     11,
		},
		{
			Name:                         "winter",
			TemperatureChange:            -10,
			HumidityChange:               -10,
			PrecipitationFrequencyChange: -10,
			PrecipitationAmountChange:    -10,
			MonthBegin:                   12,
			MonthEnd:                     2,
		},
	}

	if r.DistanceToEquator < 20 && r.DistanceToEquator > -20 {
		base = []Season{
			{
				Name:                         "wet",
				TemperatureChange:            0,
				HumidityChange:               +10,
				PrecipitationFrequencyChange: +10,
				PrecipitationAmountChange:    +10,
				MonthBegin:                   5,
				MonthEnd:                     10,
			},
			{
				Name:                         "dry",
				TemperatureChange:            0,
				HumidityChange:               0,
				PrecipitationFrequencyChange: 0,
				PrecipitationAmountChange:    0,
				MonthBegin:                   11,
				MonthEnd:                     4,
			},
		}
	}

	for _, s := range base {
		s.PrecipitationType = getPrecipitationTypeForTemperature(r.Temperature + s.TemperatureChange)
		s.DaylightHours = getDaylightHours(r.DistanceToEquator, getMiddleMonth(s.MonthBegin, s.MonthEnd))
		description, err := s.describe(c, r)
		if err != nil {
			err = fmt.Errorf("failed to generate season: %w", err)
			return []Season{}, err
		}
		s.Description = description

		seasons = append(seasons, s)
	}

	return seasons, nil
}

func getMiddleMonth(begin int, end int) int {
	var mid int

	if begin <= end {
		mid = (begin + end) / 2
	} else {
		mid = (begin + end + 12) / 2
		if mid > 12 {
			mid = mid - 12
		}
	}

	return mid
}
