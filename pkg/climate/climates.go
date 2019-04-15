package climate

import (
	"math/rand"

	"github.com/ironarachne/random"
)

// Climate is a climate
type Climate struct {
	Name              string
	Description       string
	Habitability      int
	Temperature       int
	Humidity          int
	HasWetlands       bool
	HasRivers         bool
	HasLakes          bool
	HasOcean          bool
	MaxAnimals        int
	MaxCommonMetals   int
	MaxGems           int
	MaxPlants         int
	MaxPreciousMetals int
	MaxStones         int
	Resources         []Resource
	Seasons           []Season
	Animals           []Animal
	Fish              []Fish
	CommonMetals      []Mineral
	Gems              []Mineral
	OtherMinerals     []Mineral
	Plants            []Plant
	PreciousMetals    []Mineral
	Stones            []Mineral
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

	if inSlice(firstLetter, []string{"a", "e", "i", "o", "u"}) {
		description += "an "
	} else {
		description += "a "
	}
	description += climate.Name + " region"

	waterComponents := []string{}
	if climate.HasLakes {
		lakes := []string{"many smaller lakes", "a few large lakes", "an enormous lake"}
		waterComponents = append(waterComponents, random.Item(lakes))
	}
	if climate.HasRivers {
		rivers := []string{"a few rivers", "a major river", "many small rivers and streams"}
		waterComponents = append(waterComponents, random.Item(rivers))
	}
	if climate.HasWetlands {
		wetlands := []string{"several bogs", "extensive wetlands", "scattered wetlands"}
		waterComponents = append(waterComponents, random.Item(wetlands))
	}
	waterComponents = shuffle(waterComponents)

	if len(waterComponents) > 0 {
		description += " with "
		description += combinePhrases(waterComponents)
	}

	if climate.HasOcean {
		description += ". It lies on the coast of "
		oceanTypes := []string{"a large ocean", "a gulf", "a major sea", "a smaller sea", "an ocean"}
		description += random.Item(oceanTypes)

		beachTypes := []string{"white sandy beaches", "long sandy beaches", "rocky beaches", "sheer cliffs", "low cliffs", "sandy beaches", "beautiful sandy beaches"}
		description += " with " + random.Item(beachTypes)
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
	weather += combinePhrases(weatherItems)

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
	resources := []Resource{}

	animals := climate.getFilteredAnimals()
	commonMetals := getCommonMetals()
	gems := getGems()
	plants := climate.getFilteredPlants()
	preciousMetals := getPreciousMetals()
	stones := getStones()

	climate.Seasons = climate.getSeasons()

	lakeChance := rand.Intn(10)
	riverChance := rand.Intn(10)
	oceanChance := rand.Intn(10)
	wetlandsChance := rand.Intn(10)

	if lakeChance > 6 {
		climate.HasLakes = true
	}
	if riverChance > 3 {
		climate.HasRivers = true
	}
	if oceanChance > 8 {
		climate.HasOcean = true
	}
	if wetlandsChance > 8 {
		climate.HasWetlands = true
	}

	if climate.HasLakes || climate.HasRivers || climate.HasOcean {
		climate.Fish = climate.getRandomFish()
		for _, i := range climate.Fish {
			resources = append(resources, resourcesFromFish(i)...)
		}
	} else {
		climate.Fish = []Fish{}
	}

	climate.Animals = getRandomAnimals(climate.MaxAnimals, animals)
	climate.CommonMetals = getRandomMinerals(climate.MaxCommonMetals, commonMetals)
	climate.Gems = getRandomMinerals(climate.MaxGems, gems)
	climate.OtherMinerals = getOtherMinerals()
	climate.Plants = getRandomPlants(climate.MaxPlants, plants)
	climate.PreciousMetals = getRandomMinerals(climate.MaxPreciousMetals, preciousMetals)
	climate.Stones = getRandomMinerals(climate.MaxStones, stones)

	for _, i := range climate.Animals {
		resources = append(resources, resourcesFromAnimal(i)...)
	}
	for _, i := range climate.CommonMetals {
		resources = append(resources, resourcesFromMineral(i)...)
	}
	for _, i := range climate.PreciousMetals {
		resources = append(resources, resourcesFromMineral(i)...)
	}
	for _, i := range climate.Gems {
		resources = append(resources, resourcesFromMineral(i)...)
	}
	for _, i := range climate.Stones {
		resources = append(resources, resourcesFromMineral(i)...)
	}
	for _, i := range climate.OtherMinerals {
		resources = append(resources, resourcesFromMineral(i)...)
	}
	for _, i := range climate.Plants {
		resources = append(resources, resourcesFromPlant(i)...)
	}

	climate.Resources = resources
	climate.Description = climate.getDescription()

	climate.Habitability = climate.calculateHabitability()

	return climate
}

func getAllClimates() []Climate {
	climates := []Climate{
		Climate{
			Name:              "coniferous forest",
			Temperature:       4,
			Humidity:          6,
			MaxAnimals:        15,
			MaxCommonMetals:   2,
			MaxGems:           2,
			MaxPlants:         15,
			MaxPreciousMetals: 1,
			MaxStones:         1,
		},
		Climate{
			Name:              "deciduous forest",
			Temperature:       5,
			Humidity:          5,
			MaxAnimals:        15,
			MaxCommonMetals:   2,
			MaxGems:           2,
			MaxPlants:         15,
			MaxPreciousMetals: 1,
			MaxStones:         1,
		},
		Climate{
			Name:              "desert",
			Temperature:       9,
			Humidity:          0,
			MaxAnimals:        5,
			MaxCommonMetals:   3,
			MaxGems:           2,
			MaxPlants:         3,
			MaxPreciousMetals: 1,
			MaxStones:         1,
		},
		Climate{
			Name:              "grassland",
			Temperature:       5,
			Humidity:          3,
			MaxAnimals:        10,
			MaxCommonMetals:   2,
			MaxGems:           2,
			MaxPlants:         15,
			MaxPreciousMetals: 1,
			MaxStones:         1,
		},
		Climate{
			Name:              "marshland",
			Temperature:       7,
			Humidity:          9,
			HasWetlands:       true,
			MaxAnimals:        15,
			MaxCommonMetals:   1,
			MaxGems:           1,
			MaxPlants:         10,
			MaxPreciousMetals: 1,
			MaxStones:         0,
		},
		Climate{
			Name:              "tropical",
			Temperature:       9,
			Humidity:          7,
			MaxAnimals:        16,
			MaxCommonMetals:   1,
			MaxGems:           4,
			MaxPlants:         16,
			MaxPreciousMetals: 1,
			MaxStones:         1,
		},
		Climate{
			Name:              "mountain",
			Temperature:       4,
			Humidity:          4,
			MaxAnimals:        5,
			MaxCommonMetals:   10,
			MaxGems:           2,
			MaxPlants:         5,
			MaxPreciousMetals: 5,
			MaxStones:         5,
		},
		Climate{
			Name:              "rainforest",
			Temperature:       9,
			Humidity:          9,
			MaxAnimals:        16,
			MaxCommonMetals:   1,
			MaxGems:           2,
			MaxPlants:         16,
			MaxPreciousMetals: 1,
			MaxStones:         1,
		},
		Climate{
			Name:              "savanna",
			Temperature:       9,
			Humidity:          5,
			MaxAnimals:        9,
			MaxCommonMetals:   2,
			MaxGems:           3,
			MaxPlants:         6,
			MaxPreciousMetals: 1,
			MaxStones:         1,
		},
		Climate{
			Name:              "steppe",
			Temperature:       7,
			Humidity:          3,
			MaxAnimals:        9,
			MaxCommonMetals:   3,
			MaxGems:           3,
			MaxPlants:         8,
			MaxPreciousMetals: 3,
			MaxStones:         3,
		},
		Climate{
			Name:              "taiga",
			Temperature:       3,
			Humidity:          3,
			MaxAnimals:        9,
			MaxCommonMetals:   2,
			MaxGems:           3,
			MaxPlants:         8,
			MaxPreciousMetals: 1,
			MaxStones:         1,
		},
		Climate{
			Name:              "tundra",
			Temperature:       1,
			Humidity:          3,
			MaxAnimals:        6,
			MaxCommonMetals:   3,
			MaxGems:           1,
			MaxPlants:         7,
			MaxPreciousMetals: 1,
			MaxStones:         2,
		},
	}

	return climates
}
