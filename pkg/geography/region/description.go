package region

// Describe provides a prose description of a region
func (region Region) Describe() string {
	description := describeTemperature(region.Temperature) + " with " + describeHumidity(region.Humidity) + " humidity"

	return description
}

func describeHumidity(humidity int) string {
	if humidity < 10 {
		return "very dry"
	} else if humidity < 30 {
		return "dry"
	} else if humidity < 75 {
		return "moderate"
	} else if humidity < 90 {
		return "humid"
	}

	return "very humid"
}

func describeTemperature(temperature int) string {
	if temperature < 10 {
		return "frigid"
	} else if temperature < 30 {
		return "cold"
	} else if temperature < 60 {
		return "temperate"
	} else if temperature < 75 {
		return "warm"
	} else if temperature < 90 {
		return "hot"
	}

	return "very hot"
}
