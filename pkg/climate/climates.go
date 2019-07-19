package climate

import (
	"math"
	"math/rand"

	"github.com/ironarachne/world/pkg/animal"
	"github.com/ironarachne/world/pkg/mineral"
	"github.com/ironarachne/world/pkg/plant"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/slices"
	"github.com/ironarachne/world/pkg/soil"
	"github.com/ironarachne/world/pkg/words"
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
	Animals            []animal.Animal
	Fish               []animal.Fish
	Gems               []mineral.Mineral
	Metals             []mineral.Mineral
	OtherMinerals      []mineral.Mineral
	Plants             []plant.Plant
	Resources          []resource.Resource
	Seasons            []Season
	Soils              []soil.Soil
	Stones             []mineral.Mineral
	Trees              []plant.Tree
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

func getClimateByName(name string) Climate {
	climates := getAllClimates()

	for _, c := range climates {
		if c.Name == name {
			return c
		}
	}

	panic("No such climate")
}

func getRandomClimate() Climate {
	climates := getAllClimates()

	return climates[rand.Intn(len(climates)-1)]
}

func (climate Climate) getDescription() string {
	description := "This is "

	firstLetter := string(climate.Name[0])

	if slices.StringIn(firstLetter, []string{"a", "e", "i", "o", "u"}) {
		description += "an "
	} else {
		description += "a "
	}
	description += climate.Name + " region"

	waterComponents := []string{}
	if climate.HasLakes {
		lakes := []string{"many smaller lakes", "a few large lakes", "an enormous lake"}
		waterComponents = append(waterComponents, random.String(lakes))
	}
	if climate.HasRivers {
		rivers := []string{"a few rivers", "a major river", "many small rivers and streams"}
		waterComponents = append(waterComponents, random.String(rivers))
	}
	if climate.HasWetlands {
		wetlands := []string{"several bogs", "extensive wetlands", "scattered wetlands"}
		waterComponents = append(waterComponents, random.String(wetlands))
	}
	waterComponents = slices.ShuffleStringSlice(waterComponents)

	if len(waterComponents) > 0 {
		description += " with "
		description += words.CombinePhrases(waterComponents)
	}

	if climate.HasOcean {
		description += ". It lies on the coast of "
		oceanTypes := []string{"a large ocean", "a gulf", "a major sea", "a smaller sea", "an ocean"}
		description += random.String(oceanTypes)

		beachTypes := []string{"sandy beaches", "long sandy beaches", "rocky beaches", "sheer cliffs", "low cliffs", "sandy beaches", "beautiful sandy beaches"}
		description += " with " + random.String(beachTypes)
	}

	weather := ". The weather is "
	weatherItem := ""
	weatherItems := []string{}

	for _, s := range climate.Seasons {
		weatherItem = s.WeatherProfile.Description + " during the " + s.Name
		if len(climate.Seasons) == 2 {
			weatherItem += " season"
		}
		weatherItems = append(weatherItems, weatherItem)
	}
	weather += words.CombinePhrases(weatherItems)

	description += weather + "."

	predatorCount := 0
	birdCount := 0

	for _, a := range climate.Animals {
		if a.EatsAnimals {
			predatorCount++
		}
		if a.AnimalType == "bird" {
			birdCount++
		}
	}

	manyBirds := rand.Intn(10)
	manyPredators := rand.Intn(10)

	if predatorCount > 1 && manyPredators > 8 {
		description += " There are a number of predators in the region."
	}

	if birdCount > 1 && manyBirds > 8 {
		description += " The skies are full of birds."
	}

	return description
}

func (climate Climate) getCurrentHumidity(season Season) int {
	return climate.Humidity + season.HumidityChange
}

func (climate Climate) getCurrentTemperature(season Season) int {
	return climate.Temperature + season.TemperatureChange
}

func (climate Climate) populate() Climate {
	resources := []resource.Resource{}

	animals := climate.getFilteredAnimals()
	metals := mineral.Metals()
	gems := mineral.Gems()
	plants := climate.getFilteredPlants()
	stones := mineral.Stones()
	trees := climate.getFilteredTrees()

	climate.Seasons = climate.getSeasons()

	lakeChance := rand.Intn(10)
	riverChance := rand.Intn(10)
	oceanChance := rand.Intn(10)
	wetlandsChance := rand.Intn(10)

	if lakeChance > 3 {
		climate.HasLakes = true
	}
	if riverChance > 2 {
		climate.HasRivers = true
	}
	if oceanChance > 8 {
		climate.HasOcean = true
	}
	if wetlandsChance > 8 {
		climate.HasWetlands = true
	}

	soils := climate.getFilteredSoils()

	if climate.HasLakes || climate.HasRivers || climate.HasOcean {
		climate.Fish = climate.getFish()
		for _, i := range climate.Fish {
			resources = append(resources, i.Resources...)
		}
	} else {
		climate.Fish = []animal.Fish{}
	}

	climate.Animals = animal.Random(climate.MaxAnimals, animals)
	climate.Metals = mineral.Random(climate.MaxMetals, metals)
	climate.Gems = mineral.Random(climate.MaxGems, gems)
	climate.OtherMinerals = mineral.OtherMinerals()
	climate.Plants = plant.Random(climate.MaxPlants-1, plants)
	climate.Plants = append(climate.Plants, plant.RandomFabric())
	climate.Soils = soil.Random(climate.MaxSoils, soils)
	climate.Stones = mineral.Random(climate.MaxStones, stones)
	climate.Trees = plant.RandomTrees(climate.MaxTrees, trees)

	for _, i := range climate.Animals {
		resources = append(resources, i.Resources...)
	}
	for _, i := range climate.Metals {
		resources = append(resources, i.Resources...)
	}
	for _, i := range climate.Gems {
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

	climate.Resources = resources
	climate.Description = climate.getDescription()

	climate.Habitability = climate.calculateHabitability()

	return climate
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

func getAllClimates() []Climate {
	climates := []Climate{
		Climate{
			Name:               "coniferous forest",
			Adjective:          "coniferous forest",
			Temperature:        4,
			HasDeciduousTrees:  false,
			HasConiferousTrees: true,
			Humidity:           6,
			MaxAnimals:         15,
			MaxMetals:          2,
			MaxFish:            6,
			MaxGems:            2,
			MaxPlants:          15,
			MaxSoils:           4,
			MaxStones:          1,
			MaxTrees:           6,
		},
		Climate{
			Name:               "deciduous forest",
			Adjective:          "deciduous forest",
			Temperature:        5,
			HasDeciduousTrees:  true,
			HasConiferousTrees: false,
			Humidity:           5,
			MaxAnimals:         15,
			MaxMetals:          2,
			MaxFish:            6,
			MaxGems:            2,
			MaxPlants:          15,
			MaxSoils:           4,
			MaxStones:          1,
			MaxTrees:           7,
		},
		Climate{
			Name:               "desert",
			Adjective:          "desert",
			Temperature:        9,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           0,
			MaxAnimals:         5,
			MaxMetals:          3,
			MaxFish:            2,
			MaxGems:            2,
			MaxPlants:          3,
			MaxSoils:           2,
			MaxStones:          1,
			MaxTrees:           0,
		},
		Climate{
			Name:               "grassland",
			Adjective:          "grassland",
			Temperature:        5,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           3,
			MaxAnimals:         10,
			MaxMetals:          2,
			MaxFish:            6,
			MaxGems:            2,
			MaxPlants:          15,
			MaxSoils:           3,
			MaxStones:          1,
			MaxTrees:           3,
		},
		Climate{
			Name:               "marshland",
			Adjective:          "marshy",
			Temperature:        7,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           9,
			HasWetlands:        true,
			MaxAnimals:         15,
			MaxMetals:          1,
			MaxFish:            6,
			MaxGems:            1,
			MaxPlants:          10,
			MaxSoils:           4,
			MaxStones:          0,
			MaxTrees:           3,
		},
		Climate{
			Name:               "tropical",
			Adjective:          "tropical",
			Temperature:        9,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           7,
			MaxAnimals:         16,
			MaxMetals:          1,
			MaxFish:            6,
			MaxGems:            4,
			MaxPlants:          16,
			MaxSoils:           3,
			MaxStones:          1,
			MaxTrees:           3,
		},
		Climate{
			Name:               "mountain",
			Adjective:          "mountainous",
			Temperature:        4,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           4,
			MaxAnimals:         5,
			MaxMetals:          10,
			MaxFish:            3,
			MaxGems:            2,
			MaxPlants:          5,
			MaxSoils:           2,
			MaxStones:          5,
			MaxTrees:           3,
		},
		Climate{
			Name:               "rainforest",
			Adjective:          "rainforest",
			Temperature:        9,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           9,
			MaxAnimals:         16,
			MaxMetals:          1,
			MaxFish:            6,
			MaxGems:            2,
			MaxPlants:          16,
			MaxSoils:           3,
			MaxStones:          1,
			MaxTrees:           3,
		},
		Climate{
			Name:               "savanna",
			Adjective:          "savanna",
			Temperature:        9,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           5,
			MaxAnimals:         9,
			MaxMetals:          2,
			MaxFish:            6,
			MaxGems:            3,
			MaxPlants:          6,
			MaxSoils:           2,
			MaxStones:          1,
			MaxTrees:           3,
		},
		Climate{
			Name:               "steppe",
			Adjective:          "steppe",
			Temperature:        7,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           3,
			MaxAnimals:         9,
			MaxMetals:          3,
			MaxFish:            6,
			MaxGems:            3,
			MaxPlants:          8,
			MaxSoils:           2,
			MaxStones:          3,
			MaxTrees:           3,
		},
		Climate{
			Name:               "taiga",
			Adjective:          "taiga",
			Temperature:        3,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           3,
			MaxAnimals:         9,
			MaxMetals:          2,
			MaxFish:            6,
			MaxGems:            3,
			MaxPlants:          8,
			MaxSoils:           2,
			MaxStones:          1,
			MaxTrees:           5,
		},
		Climate{
			Name:               "tundra",
			Adjective:          "tundra",
			Temperature:        1,
			HasDeciduousTrees:  true,
			HasConiferousTrees: true,
			Humidity:           3,
			MaxAnimals:         6,
			MaxMetals:          3,
			MaxFish:            6,
			MaxGems:            1,
			MaxPlants:          7,
			MaxSoils:           1,
			MaxStones:          2,
			MaxTrees:           3,
		},
	}

	return climates
}
