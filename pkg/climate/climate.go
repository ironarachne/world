/*
Package climate provides tools for procedural generation of climates.
*/
package climate

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/ironarachne/world/pkg/mineral"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/soil"
	"github.com/ironarachne/world/pkg/species"
)

const climateError = "could not generate climate: %w"

// Generator is a terrestrial climate generator
type Generator struct {
	NameRoot        string
	NameVariants    []string
	Adjective       string
	TemperatureMin  int
	TemperatureMax  int
	HumidityMin     int
	HumidityMax     int
	ElevationMin    int
	ElevationMax    int
	AnimalMin       int
	AnimalMax       int
	PlantMin        int
	PlantMax        int
	TreeMin         int
	TreeMax         int
	RiverPrevalence int
	LakePrevalence  int
	SoilTags        []string
	AnimalTags      []string
	PlantTags       []string
	TreeTags        []string
}

// Climate is a terrestrial climate
type Climate struct {
	Name         string              `json:"name"`
	Adjective    string              `json:"adjective"`
	Description  string              `json:"description"`
	Habitability int                 `json:"habitability"`
	Temperature  int                 `json:"temperature"`
	Humidity     int                 `json:"humidity"`
	Elevation    int                 `json:"elevation"`
	Animals      []species.Species   `json:"animals"`
	Plants       []species.Species   `json:"plants"`
	Minerals     []mineral.Mineral   `json:"minerals"`
	Soils        []soil.Soil         `json:"soils"`
	Trees        []species.Species   `json:"trees"`
	Resources    []resource.Resource `json:"resources"`
	Seasons      []Season            `json:"seasons"`
}

// Generate procedurally generates a random climate
func Generate() (Climate, error) {
	generators := AllGenerators()
	gen := generators[rand.Intn(len(generators))]

	humidity := rand.Intn(gen.HumidityMax-gen.HumidityMin) + gen.HumidityMin
	temperature := rand.Intn(gen.TemperatureMax-gen.TemperatureMin) + gen.TemperatureMin
	elevation := rand.Intn(gen.ElevationMax-gen.ElevationMin) + gen.ElevationMin

	animals, err := gen.getAnimals(humidity, temperature)
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}
	plants, err := gen.getPlants(humidity, temperature)
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}
	soils, err := gen.getSoils()
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}
	trees, err := gen.getTrees(humidity, temperature)
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}
	minerals, err := gen.getMinerals()
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}

	nameVariant := gen.NameVariants[rand.Intn(len(gen.NameVariants))]

	name := nameVariant + " " + gen.NameRoot

	climate := Climate{
		Name:        name,
		Adjective:   gen.Adjective,
		Temperature: temperature,
		Humidity:    humidity,
		Elevation:   elevation,
		Animals:     animals,
		Plants:      plants,
		Minerals:    minerals,
		Soils:       soils,
		Trees:       trees,
	}

	climate.Seasons = climate.getSeasons()
	climate.Resources = getResources(climate)
	climate.Habitability = climate.calculateHabitability()
	climate.Description, err = climate.getDescription()
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}

	return climate, nil
}

// GenerateByName procedurally generates a random climate given a previously produced climate generator name
func GenerateByName(name string) (Climate, error) {
	gen := Generator{}
	generators := AllGenerators()

	for _, g := range generators {
		for _, v := range g.NameVariants {
			if name == v+" "+g.NameRoot {
				gen = g
			}
		}
	}

	humidity := rand.Intn(gen.HumidityMax-gen.HumidityMin) + gen.HumidityMin
	temperature := rand.Intn(gen.TemperatureMax-gen.TemperatureMin) + gen.TemperatureMin
	elevation := rand.Intn(gen.ElevationMax-gen.ElevationMin) + gen.ElevationMin

	animals, err := gen.getAnimals(humidity, temperature)
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}
	plants, err := gen.getPlants(humidity, temperature)
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}
	soils, err := gen.getSoils()
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}
	trees, err := gen.getTrees(humidity, temperature)
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}
	minerals, err := gen.getMinerals()
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}

	nameVariant := gen.NameVariants[rand.Intn(len(gen.NameVariants))]

	climateName := nameVariant + " " + gen.NameRoot

	climate := Climate{
		Name:        climateName,
		Adjective:   gen.Adjective,
		Temperature: temperature,
		Humidity:    humidity,
		Elevation:   elevation,
		Animals:     animals,
		Plants:      plants,
		Minerals:    minerals,
		Soils:       soils,
		Trees:       trees,
	}

	climate.Seasons = climate.getSeasons()
	climate.Resources = getResources(climate)
	climate.Habitability = climate.calculateHabitability()
	climate.Description, err = climate.getDescription()
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}

	return climate, nil
}

// GenerateForCharacteristics procedurally generates a random climate given three characteristics
func GenerateForCharacteristics(humidity int, temperature int, elevation int) (Climate, error) {
	var gen Generator
	var score int
	generators := AllGenerators()

	bestScore := 30

	for _, g := range generators {
		score = g.CharacteristicScore(humidity, temperature, elevation)
		if score < bestScore {
			bestScore = score
			gen = g
		}
	}

	animals, err := gen.getAnimals(humidity, temperature)
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}
	plants, err := gen.getPlants(humidity, temperature)
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}
	soils, err := gen.getSoils()
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}
	trees, err := gen.getTrees(humidity, temperature)
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}
	minerals, err := gen.getMinerals()
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}

	nameVariant := gen.NameVariants[rand.Intn(len(gen.NameVariants))]

	name := nameVariant + " " + gen.NameRoot

	climate := Climate{
		Name:        name,
		Adjective:   gen.Adjective,
		Temperature: temperature,
		Humidity:    humidity,
		Elevation:   elevation,
		Animals:     animals,
		Plants:      plants,
		Minerals:    minerals,
		Soils:       soils,
		Trees:       trees,
	}

	climate.Seasons = climate.getSeasons()
	climate.Resources = getResources(climate)
	climate.Habitability = climate.calculateHabitability()
	climate.Description, err = climate.getDescription()
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}

	return climate, nil
}

// GenerateForeign procedurally generates a random climate different from the given one
func GenerateForeign(currentClimate Climate) (Climate, error) {
	var generators []Generator
	var score int

	allGenerators := AllGenerators()

	for _, g := range allGenerators {
		score = g.CharacteristicScore(currentClimate.Humidity, currentClimate.Temperature, currentClimate.Elevation)
		if score > 5 {
			generators = append(generators, g)
		}
	}

	gen := generators[rand.Intn(len(generators))]

	humidity := rand.Intn(gen.HumidityMax-gen.HumidityMin) + gen.HumidityMin
	temperature := rand.Intn(gen.TemperatureMax-gen.TemperatureMin) + gen.TemperatureMin
	elevation := rand.Intn(gen.ElevationMax-gen.ElevationMin) + gen.ElevationMin

	animals, err := gen.getAnimals(humidity, temperature)
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}
	plants, err := gen.getPlants(humidity, temperature)
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}
	soils, err := gen.getSoils()
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}
	trees, err := gen.getTrees(humidity, temperature)
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}
	minerals, err := gen.getMinerals()
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}

	nameVariant := gen.NameVariants[rand.Intn(len(gen.NameVariants))]

	name := nameVariant + " " + gen.NameRoot

	climate := Climate{
		Name:        name,
		Adjective:   gen.Adjective,
		Temperature: temperature,
		Humidity:    humidity,
		Elevation:   elevation,
		Animals:     animals,
		Plants:      plants,
		Minerals:    minerals,
		Soils:       soils,
		Trees:       trees,
	}

	climate.Seasons = climate.getSeasons()
	climate.Resources = getResources(climate)
	climate.Habitability = climate.calculateHabitability()
	climate.Description, err = climate.getDescription()
	if err != nil {
		err = fmt.Errorf(climateError, err)
		return Climate{}, err
	}

	return climate, nil
}

func (climate Climate) calculateHabitability() int {
	habitability := 100

	if climate.Temperature > 9 || climate.Temperature < 4 {
		habitability -= 40
	}

	if climate.Humidity < 2 || climate.Humidity > 9 {
		habitability -= 10
	}

	if len(climate.Animals) < 4 {
		habitability -= 10
	}

	if len(climate.Plants) < 4 {
		habitability -= 10
	}

	totalWeatherScore := 0

	for _, s := range climate.Seasons {
		if s.WeatherProfile.WindLevel > 5 {
			totalWeatherScore += 5
		}
		if s.WeatherProfile.PrecipitationAmount < 3 || s.WeatherProfile.PrecipitationAmount > 9 {
			totalWeatherScore += 5
		}
	}

	habitability -= totalWeatherScore

	habitability += len(climate.Seasons) * 2

	return habitability
}

func getResources(climate Climate) []resource.Resource {
	resources := []resource.Resource{}

	for _, i := range climate.Animals {
		resources = append(resources, i.Resources...)
	}
	for _, i := range climate.Minerals {
		resources = append(resources, i.Resources...)
	}
	for _, i := range climate.Plants {
		resources = append(resources, i.Resources...)
	}
	for _, i := range climate.Soils {
		resources = append(resources, i.Resources...)
	}
	for _, i := range climate.Trees {
		resources = append(resources, i.Resources...)
	}

	return resources
}

func (climate Climate) getCurrentHumidity(season Season) int {
	return climate.Humidity + season.HumidityChange
}

func (climate Climate) getCurrentTemperature(season Season) int {
	return climate.Temperature + season.TemperatureChange
}

// CharacteristicScore grades a set of characteristics according to its match to a given climate generator.
// Lower is better. A perfect score (complete match) is 0.
func (gen Generator) CharacteristicScore(humidity int, temperature int, elevation int) int {
	score := 0

	middleHumidity := gen.HumidityMin + int((gen.HumidityMax-gen.HumidityMin)/2)
	middleTemperature := gen.TemperatureMin + int((gen.TemperatureMax-gen.TemperatureMin)/2)
	middleElevation := gen.ElevationMin + int((gen.ElevationMax-gen.ElevationMin)/2)

	score += int(math.Abs(float64(humidity - middleHumidity)))
	score += int(math.Abs(float64(temperature - middleTemperature)))
	score += int(math.Abs(float64(elevation - middleElevation)))

	return score
}
