package town

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/buildings"
	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/goods"
)

// Town is a town
type Town struct {
	Name             string
	Population       int
	BuildingStyle    buildings.BuildingStyle
	Category         Category
	Climate          climate.Climate
	Culture          culture.Culture
	Mayor            character.Character
	NotableProducers []goods.Producer
	Exports          []goods.TradeGood
	Imports          []goods.TradeGood
}

func (town Town) generateMayor() character.Character {
	mayor := character.Generate(town.Culture)
	mayor = mayor.ChangeAge(rand.Intn(30) + 30)

	return mayor
}

func (town Town) generateRandomExports() []goods.TradeGood {
	var exports []goods.TradeGood

	exports = goods.GenerateExportTradeGoods(town.Category.MinExports, town.Category.MaxExports, town.NotableProducers, town.Climate.Resources)

	return exports
}

func (town Town) generateRandomImports() []goods.TradeGood {
	var imports []goods.TradeGood

	foreignClimate := climate.GetForeignClimate(town.Climate)

	imports = goods.GenerateImportTradeGoods(town.Category.MinImports, town.Category.MaxImports, foreignClimate.Resources)

	return imports
}

func generateRandomPopulation(category Category) int {
	sizeIncrement := category.MaxSize - category.MinSize

	return rand.Intn(sizeIncrement) + category.MinSize
}

func (town Town) generateTownName() string {
	name := town.Culture.Language.RandomName()
	return name
}

// Generate generates a random town
func Generate(category string, biome string, originCulture culture.Culture) Town {
	town := Town{}

	if category == "random" {
		town.Category = getRandomWeightedCategory()
	} else {
		town.Category = getCategoryByName(category)
	}
	if biome == "random" {
		town.Climate = climate.Generate()
	} else {
		town.Climate = climate.GetClimate(biome)
	}

	town.Culture = originCulture
	town.Name = town.Culture.Language.RandomName()

	town.Population = generateRandomPopulation(town.Category)

	town.BuildingStyle = town.Culture.BuildingStyle

	mayor := town.generateMayor()
	mayor.FirstName = town.Culture.Language.RandomGenderedName(mayor.Gender.Name)
	mayor.LastName = town.Culture.Language.RandomName()
	town.Mayor = mayor

	possibleProducers := goods.GetPossibleProducers(town.Climate.Resources, town.Population)

	numberOfProducers := 0

	if town.Population < 20 {
		numberOfProducers = 1
	} else if town.Population < 50 {
		numberOfProducers = 3
	} else if town.Population < 100 {
		numberOfProducers = 4
	} else if town.Population < 500 {
		numberOfProducers = 6
	} else if town.Population < 1000 {
		numberOfProducers = 10
	} else if town.Population < 5000 {
		numberOfProducers = 20
	} else if town.Population < 10000 {
		numberOfProducers = 40
	} else {
		numberOfProducers = int(town.Population / 250)
	}

	town.NotableProducers = goods.GetRandomProducers(numberOfProducers, possibleProducers)

	town.Exports = town.generateRandomExports()
	town.Imports = town.generateRandomImports()

	return town
}

// Random generates a completely random town
func Random() Town {
	randomCulture := culture.Random()

	town := Generate("random", "random", randomCulture)

	return town
}
