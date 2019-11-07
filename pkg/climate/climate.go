package climate

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/ironarachne/world/pkg/fish"
	"github.com/ironarachne/world/pkg/insect"
	"github.com/ironarachne/world/pkg/mineral"
	"github.com/ironarachne/world/pkg/plant"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/soil"
	"github.com/ironarachne/world/pkg/species"
	"github.com/ironarachne/world/pkg/tree"
)

// Climate is a climate
type Climate struct {
	Name               string
	Adjective          string
	Description        string
	Habitability       int
	Temperature        int
	Humidity           int
	HasDeciduousTrees  bool
	HasConiferousTrees bool
	HasWetlands        bool
	HasRivers          bool
	HasLakes           bool
	HasOcean           bool
	MaxAnimals         int
	MaxFish            int
	MaxGems            int
	MaxMetals          int
	MaxPlants          int
	MaxSoils           int
	MaxStones          int
	MaxTrees           int
	Animals            []species.Species
	Fish               []fish.Fish
	Gems               []mineral.Mineral
	Insects            []insect.Insect
	Metals             []mineral.Mineral
	OtherMinerals      []mineral.Mineral
	Plants             []plant.Plant
	Resources          []resource.Resource
	Seasons            []Season
	Soils              []soil.Soil
	Stones             []mineral.Mineral
	Trees              []tree.Tree
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

	if climate.HasLakes {
		habitability += 5
	}
	if climate.HasRivers {
		habitability += 5
	}
	if climate.HasOcean {
		habitability += 5
	}
	if climate.HasWetlands {
		habitability -= 10
	}

	return habitability
}

// ByName returns a climate by name
func ByName(name string) (Climate, error) {
	climates := getAllClimates()

	for _, c := range climates {
		if c.Name == name {
			return c, nil
		}
	}

	err := fmt.Errorf("Failed to find climate with name " + name)
	return Climate{}, err
}

// Random returns a completely random climate
func Random() (Climate, error) {
	climates := getAllClimates()

	climate := climates[rand.Intn(len(climates))]

	populatedClimate, err := climate.populate()
	if err != nil {
		err = fmt.Errorf("Could not generate random climate: %w", err)
		return Climate{}, err
	}

	return populatedClimate, nil
}

// Generate generates a climate with a given name
func Generate(name string) (Climate, error) {
	rawClimate, err := ByName(name)
	if err != nil {
		err = fmt.Errorf("Could not generate climate by name: %w", err)
		return Climate{}, err
	}
	climate, err := rawClimate.populate()
	if err != nil {
		err = fmt.Errorf("Could not generate climate by name: %w", err)
		return Climate{}, err
	}

	return climate, nil
}

// GetForeignClimate gets a random climate that's different from the given one
func GetForeignClimate(climate Climate) (Climate, error) {
	var possibleClimates []Climate

	climates := getAllClimates()

	for _, c := range climates {
		if c.Name != climate.Name {
			possibleClimates = append(possibleClimates, c)
		}
	}

	foreignClimate := possibleClimates[rand.Intn(len(possibleClimates))]

	newClimate, err := foreignClimate.populate()
	if err != nil {
		err = fmt.Errorf("Could not generate foreign climate: %w", err)
		return Climate{}, err
	}

	return newClimate, nil
}

func (climate Climate) getResources() []resource.Resource {
	resources := []resource.Resource{}

	for _, i := range climate.Fish {
		resources = append(resources, i.Resources...)
	}
	for _, i := range climate.Animals {
		resources = append(resources, i.Resources...)
	}
	for _, i := range climate.Metals {
		resources = append(resources, i.Resources...)
	}
	for _, i := range climate.Gems {
		resources = append(resources, i.Resources...)
	}
	for _, i := range climate.Insects {
		resources = append(resources, i.Resources...)
	}
	for _, i := range climate.Stones {
		resources = append(resources, i.Resources...)
	}
	for _, i := range climate.OtherMinerals {
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

func (climate Climate) populate() (Climate, error) {
	gems := mineral.Gems()
	insects := climate.getFilteredInsects()
	metals := mineral.Metals()

	stones := mineral.Stones()
	trees := climate.getFilteredTrees()

	climate.Seasons = climate.getSeasons()

	lakeChance := rand.Intn(100)
	riverChance := rand.Intn(100)
	oceanChance := rand.Intn(100)
	wetlandsChance := rand.Intn(100)

	if lakeChance > 30 {
		climate.HasLakes = true
	}
	if riverChance > 20 {
		climate.HasRivers = true
	}
	if oceanChance > 80 {
		climate.HasOcean = true
	}
	if wetlandsChance > 80 {
		climate.HasWetlands = true
	}

	soils := climate.getFilteredSoils()

	if climate.HasLakes || climate.HasRivers || climate.HasOcean {
		climate.Fish = climate.getFish()
	} else {
		climate.Fish = []fish.Fish{}
	}

	climate.Insects = insect.RandomSubset(7, insects)
	filteredMetals, err := mineral.RandomWeightedSet(climate.MaxMetals, metals)
	if err != nil {
		err = fmt.Errorf("Could not populate climate: %w", err)
		return Climate{}, err
	}
	climate.Metals = filteredMetals
	climate.Gems = mineral.Random(climate.MaxGems, gems)
	climate.OtherMinerals = mineral.OtherMinerals()

	climate.Animals, err = climate.getAnimals()
	if err != nil {
		err = fmt.Errorf("Could not populate climate: %w", err)
		return Climate{}, err
	}

	climate.Plants, err = climate.getPlants()
	if err != nil {
		err = fmt.Errorf("Could not populate climate: %w", err)
		return Climate{}, err
	}

	climate.Soils = soil.Random(climate.MaxSoils, soils)
	climate.Stones = mineral.Random(climate.MaxStones, stones)
	climate.Trees = tree.RandomSubset(climate.MaxTrees, trees)

	resources := climate.getResources()
	climate.Resources = resources

	description, err := climate.getDescription()
	if err != nil {
		err = fmt.Errorf("Could not populate climate: %w", err)
		return Climate{}, err
	}
	climate.Description = description

	climate.Habitability = climate.calculateHabitability()

	return climate, nil
}

// GetClimateNameForConditions finds the closest matching climate for a given humidity and temperature and returns its name
func GetClimateNameForConditions(humidity int, temperature int) string {
	var name string

	scores := make(map[string]int)

	lowestScore := 20
	climates := getAllClimates()

	for _, c := range climates {
		scores[c.Name] = int(math.Abs(float64(c.Temperature-temperature)) + math.Abs(float64(c.Humidity-humidity)))
	}

	for n, s := range scores {
		if s < lowestScore {
			name = n
			lowestScore = s
		}
	}

	return name
}
