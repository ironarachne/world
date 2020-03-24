package region

import (
	"github.com/ironarachne/world/pkg/geometry"
	"math"
	"math/rand"
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
func Generate() Region {
	region := RandomTemperate()

	return region
}

// GenerateSpecific generates a region based on specific characteristics
func GenerateSpecific(temperature int, humidity int, altitude int, distance int) Region {
	region := Region{}

	region.DistanceToEquator = distance
	region.Temperature = temperature
	region.Humidity = humidity
	region.Altitude = altitude

	// TODO: Replace the following with real data gleaned from the world
	region.NearestOceanDistance = rand.Intn(100)
	region.NearestOceanDirection = geometry.RandomDirection()
	region.NearestMountainsDirection = geometry.OppositeDirection(region.NearestOceanDirection)
	region.NearestMountainsDistance = rand.Intn(100)

	region.Description = region.Describe()

	return region
}

// RandomTemperate returns a random region that is appropriate for life
func RandomTemperate() Region {
	region := Region{}

	region.DistanceToEquator = rand.Intn(100) - 50
	region.Altitude = rand.Intn(50) + 10
	region.NearestOceanDistance = rand.Intn(100)
	region.NearestOceanDirection = geometry.RandomDirection()
	region.NearestMountainsDirection = geometry.OppositeDirection(region.NearestOceanDirection)
	region.NearestMountainsDistance = rand.Intn(100)
	region.Temperature = getTemperature(region.DistanceToEquator, region.Altitude)
	region.Humidity = getHumidity(region.Altitude, region.NearestOceanDistance)

	region.Description = region.Describe()

	return region
}

func getHumidity(altitude int, oceanDistance int) int {
	humidity := 100 - (altitude / 2) - (oceanDistance / 2)

	return humidity
}

func getTemperature(distanceToEquator int, altitude int) int {
	temperature := 100 - int(math.Abs(float64(distanceToEquator))) - (altitude / 2)
	if temperature < 0 {
		temperature = 0
	}
	if temperature > 99 {
		temperature = 99
	}

	return temperature
}
