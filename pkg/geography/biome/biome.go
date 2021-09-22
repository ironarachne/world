package biome

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"

	"github.com/ironarachne/world/pkg/geography/climate"
	"github.com/ironarachne/world/pkg/geography/region"
	"github.com/ironarachne/world/pkg/random"
)

// Data is a structural element to hold biomes when loaded or displayed
type Data struct {
	Biomes []Biome `json:"biomes"`
}

// Biome is a biome
type Biome struct {
	Name                 string   `json:"name"`
	MatchScore           int      `json:"match_score"`
	FaunaPrevalence      int      `json:"fauna_prevalence"`
	FaunaTags            []string `json:"fauna_tags"`
	AltitudeMin          int      `json:"altitude_min"`
	AltitudeMax          int      `json:"altitude_max"`
	TemperatureMin       int      `json:"temperature_min"`
	TemperatureMax       int      `json:"temperature_max"`
	PrecipitationMin     int      `json:"precipitation_min"`
	PrecipitationMax     int      `json:"precipitation_max"`
	Type                 string   `json:"type"`                  // either terrestrial, freshwater, or marine
	VegetationPrevalence int      `json:"vegetation_prevalence"` // 0-99
	VegetationTags       []string `json:"vegetation_tags"`
}

// BiomeCriteria is a set of information used for evaluating biomes
type BiomeCriteria struct {
	Altitude      int
	Precipitation int
	Temperature   int
}

// Generate procedurally generates a biome for a given climate and region
func Generate(ctx context.Context, c climate.Climate, r region.Region) (Biome, error) {
	biomeType := getRandomBiomeType(ctx, r.Altitude)

	b, err := Match(biomeType, c.PrecipitationAmount, r.Temperature, r.Altitude)
	if err != nil {
		err = fmt.Errorf("failed to find appropriate biome: %w", err)
		return Biome{}, err
	}

	return b, nil
}

// Match generates a biome for a given biome type, precipitation amount, temperature, and altitude
func Match(biomeType string, precipitation int, temperature int, altitude int) (Biome, error) {
	var biomeScore int
	var appropriate []Biome

	biomes, err := Load()
	if err != nil {
		err = fmt.Errorf("failed to generate biome: %w", err)
		return Biome{}, err
	}

	for _, b := range biomes {
		biomeScore = b.Score(biomeType, precipitation, temperature, altitude)
		if biomeScore > 0 {
			b.MatchScore = biomeScore
			appropriate = append(appropriate, b)
		}
	}

	if len(appropriate) == 0 {
		err = fmt.Errorf("no appropriate biomes found, input was: T %d, A %d, P %d, Type: %s", temperature, altitude, precipitation, biomeType)
		return Biome{}, err
	}

	b, err := findBestScore(appropriate)
	if err != nil {
		err = fmt.Errorf("failed to find appropriate biome: %w", err)
		return Biome{}, err
	}

	return b, nil
}

// Load returns all predefined species of a given type from a JSON file on disk
func Load() ([]Biome, error) {
	var d Data

	jsonFile, err := os.Open(os.Getenv("WORLDAPI_DATA_PATH") + "/data/biomes.json")
	if err != nil {
		err = fmt.Errorf("could not open data file: %w", err)
		return []Biome{}, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &d)
	if err != nil {
		err = fmt.Errorf("could not process biome data: %w", err)
		return []Biome{}, err
	}

	all := d.Biomes

	if len(all) == 0 {
		err = fmt.Errorf("no biomes returned from database: biomes.json")
		return []Biome{}, err
	}

	return all, nil
}

// Score returns a numeric score for a biome based on selection criteria
func (b Biome) Score(biomeType string, precipitation int, temperature int, altitude int) int {
	score := 0

	if biomeType != b.Type {
		return 0
	}

	if precipitation > b.PrecipitationMax || precipitation < b.PrecipitationMin {
		return 0
	}

	if temperature > b.TemperatureMax || temperature < b.TemperatureMin {
		return 0
	}

	if altitude > b.AltitudeMax || altitude < b.AltitudeMin {
		return 0
	}

	precipitationOffset := int(math.Abs(float64((b.PrecipitationMin+b.PrecipitationMax)/2) - float64(precipitation)))
	score += 50 - precipitationOffset

	temperatureOffset := int(math.Abs(float64((b.TemperatureMin+b.TemperatureMax)/2) - float64(temperature)))
	score += 50 - temperatureOffset

	altitudeOffset := int(math.Abs(float64((b.AltitudeMin+b.AltitudeMax)/2) - float64(altitude)))
	score += 50 - altitudeOffset

	return score
}

func findBestScore(biomes []Biome) (Biome, error) {
	bestScore := 0
	best := Biome{}

	if len(biomes) == 0 {
		err := fmt.Errorf("no biomes given")
		return Biome{}, err
	}

	for _, b := range biomes {
		if b.MatchScore > bestScore {
			best = b
		}
	}

	return best, nil
}

func getRandomBiomeType(ctx context.Context, altitude int) string {
	if altitude < 0 {
		return "marine"
	}

	types := []string{
		"terrestrial",
		"freshwater",
	}

	biomeType := types[random.Intn(ctx, len(types))]

	return biomeType
}
