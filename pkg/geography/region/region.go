package region

import (
	"context"
	"math"

	"github.com/ironarachne/world/pkg/geometry"
	"github.com/ironarachne/world/pkg/random"
)

// Region is a geographic area.
type Region struct {
	Description               string `json:"description"`
	Altitude                  int    `json:"altitude"`    // -99-99, 0 is sea level
	Humidity                  int    `json:"humidity"`    // 0-99
	Temperature               int    `json:"temperature"` // 0-99
	NearestOceanDistance      int    `json:"nearest_ocean_distance"`
	NearestOceanDirection     int    `json:"nearest_ocean_direction"`
	NearestMountainsDistance  int    `json:"nearest_mountains_distance"`
	NearestMountainsDirection int    `json:"nearest_mountains_direction"`
	DistanceToEquator         int    `json:"distance_to_equator"` // 0 is on equator, -99 is south pole, 99 is north pole
}

// Generate procedurally generates a random region.
func Generate(ctx context.Context) Region {
	region := RandomTemperate(ctx)

	return region
}

// GenerateSpecific generates a region based on specific characteristics
func GenerateSpecific(ctx context.Context, temperature int, humidity int, altitude int, distance int) Region {
	region := Region{}

	region.DistanceToEquator = distance
	region.Temperature = temperature
	region.Humidity = humidity
	region.Altitude = altitude

	// TODO: Replace the following with real data gleaned from the world
	region.NearestOceanDistance = random.Intn(ctx, 100)
	region.NearestOceanDirection = geometry.RandomDirection(ctx)
	region.NearestMountainsDirection = geometry.OppositeDirection(region.NearestOceanDirection)
	region.NearestMountainsDistance = random.Intn(ctx, 100)

	region.Description = region.Describe()

	return region
}

// RandomTemperate returns a random region that is appropriate for life
func RandomTemperate(ctx context.Context) Region {
	region := Region{}

	region.DistanceToEquator = random.Intn(ctx, 100) - 50
	region.Altitude = random.Intn(ctx, 50) + 10
	region.NearestOceanDistance = random.Intn(ctx, 100)
	region.NearestOceanDirection = geometry.RandomDirection(ctx)
	region.NearestMountainsDirection = geometry.OppositeDirection(region.NearestOceanDirection)
	region.NearestMountainsDistance = random.Intn(ctx, 100)
	region.Temperature = GetTemperature(region.DistanceToEquator, region.Altitude)
	region.Humidity = GetHumidity(region.Altitude, region.NearestOceanDistance)

	region.Description = region.Describe()

	return region
}

// GetHumidity calculates a region's humidity based on its altitude and its distance from the nearest ocean
func GetHumidity(altitude int, oceanDistance int) int {
	if oceanDistance == 0 {
		return 100
	}
	humidity := 100 - (altitude / 2) - (oceanDistance / 2)

	if humidity > 100 {
		humidity = 100
	}

	if humidity < 0 {
		humidity = 0
	}

	return humidity
}

// GetTemperature calculates a temperature for a region given its distance from the equator and its altitude
func GetTemperature(distanceToEquator int, altitude int) int {
	temperature := 100 - int(math.Abs(float64(distanceToEquator))) - (altitude / 2)
	if temperature < 0 {
		temperature = 0
	}
	if temperature > 99 {
		temperature = 99
	}

	return temperature
}
