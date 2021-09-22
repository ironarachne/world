package climate

import (
	"context"

	"github.com/ironarachne/world/pkg/geography/region"
	"github.com/ironarachne/world/pkg/geometry"
)

// Climate is a geographic climate
type Climate struct {
	CloudCover             int    `json:"cloud_cover"`             // 0-99
	WindStrength           int    `json:"wind_strength"`           // 0-99
	WindDirection          int    `json:"wind_direction"`          // 0-7
	PrecipitationAmount    int    `json:"precipitation_amount"`    // 0-99
	PrecipitationFrequency int    `json:"precipitation_frequency"` // 0-99
	PrecipitationType      string `json:"precipitation_type"`      // 0-99
}

// DescribeClouds gives a textual description for the clouds
func (c Climate) DescribeClouds() string {
	var description string

	if c.CloudCover < 10 {
		description = "no clouds"
		return description
	} else if c.CloudCover < 30 {
		description = "few clouds"
	} else if c.CloudCover < 50 {
		description = "some clouds"
	} else if c.CloudCover < 70 {
		description = "many clouds"
	} else {
		description = "frequently overcast skies"
	}

	return description
}

// Generate procedurally generates a geographic climate based on a region.
func Generate(ctx context.Context, r region.Region) Climate {
	c := Climate{}

	// cloud cover increases further away from mountains
	// cloud cover increases if mountains are to the east
	// cloud cover increases with temperature
	// cloud cover increases with wind strength
	// wind strength increases with temperature
	// wind strength increases closer to mountains
	// wind moves away from mountains
	// wind slows down going uphill
	// wind speeds up going downhill
	c.WindDirection = geometry.OppositeDirection(r.NearestMountainsDirection)
	c.WindStrength = getWindStrength(r.NearestMountainsDistance, r.NearestOceanDistance)
	c.CloudCover = getCloudCover(r.Temperature, c.WindStrength, r.NearestMountainsDistance)
	c.PrecipitationAmount = getPrecipitationAmount(r.Temperature, r.Humidity)
	c.PrecipitationFrequency = getPrecipitationFrequency(c.CloudCover, c.PrecipitationAmount)
	c.PrecipitationType = getPrecipitationType(r.Temperature)

	return c
}

func getCloudCover(temperature int, windStrength int, mountainDistance int) int {
	cloudCover := (temperature / 3) + (windStrength / 3) + (mountainDistance / 2)
	if cloudCover > 99 {
		cloudCover = 99
	}

	return cloudCover
}

func getPrecipitationAmount(temperature int, humidity int) int {
	amount := (temperature / 2) + int(float64(humidity)*0.7)
	if amount > 99 {
		amount = 99
	}

	return amount
}

func getPrecipitationFrequency(cloudCover int, amount int) int {
	frequency := (cloudCover / 3) + (amount / 3)

	return frequency
}

func getPrecipitationType(temperature int) string {
	if temperature < 30 {
		return "snow"
	}

	return "rain"
}

func getWindStrength(mountainDistance int, oceanDistance int) int {
	windStrength := (mountainDistance / 2) + (oceanDistance / 4)
	if windStrength > 99 {
		windStrength = 99
	}

	return windStrength
}
