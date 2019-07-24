package town

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/resource"

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
	NotableProducers []profession.Profession
	Exports          []goods.TradeGood
	Imports          []goods.TradeGood
	Resources        []resource.Resource
}

func (town Town) generateMayor() character.Character {
	mayor := character.Generate(town.Culture)
	mayor = mayor.ChangeAge(rand.Intn(30) + 30)

	return mayor
}

func (town Town) generateRandomExports() []goods.TradeGood {
	var exports []goods.TradeGood

	exports = goods.GenerateExportTradeGoods(town.Category.MinExports, town.Category.MaxExports, town.Resources)

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
	var newProducers []profession.Profession
	var producers []profession.Profession
	var newResources []resource.Resource

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

	resources := town.Climate.Resources

	for i := 0; i < town.Category.ProductionIterations; i++ {
		newProducers = getProducers(town.Population, resources)
		newResources = goods.Produce(newProducers, resources)
		resources = append(resources, newResources...)
		producers = append(producers, newProducers...)
	}

	town.Resources = resources
	town.NotableProducers = producers

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
