package town

import (
	"math/rand"

	"github.com/ironarachne/chargen"
	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/random"
)

// Town is a town
type Town struct {
	Name       string
	Population int
	Category   TownCategory
	Climate    climate.Climate
	Culture    culture.Culture
	Mayor      character.Character
	Exports    []TradeGood
	Imports    []TradeGood
}

// TownCategory is a type of town
type TownCategory struct {
	Name       string
	MinSize    int
	MaxSize    int
	MinExports int
	MaxExports int
	MinImports int
	MaxImports int
}

func (town Town) generateMayor() character.Character {
	mayor := character.GenerateCharacterOfCulture(town.Culture)
	mayor.AgeCategory = "adult"
	mayor.Age = 30 + rand.Intn(30)

	return mayor
}

func generateRandomCategory() TownCategory {
	categoryName := random.ItemFromThresholdMap(townCategoryOptions)

	category := townCategories[categoryName]

	return category
}

func (town Town) generateRandomExports() []TradeGood {
	var exports []TradeGood

	exports = generateTradeGoods(town.Category.MinExports, town.Category.MaxExports, town.Climate.Resources)

	return exports
}

func (town Town) generateRandomImports() []TradeGood {
	var imports []TradeGood

	foreignClimate := climate.GetForeignClimate(town.Climate)

	imports = generateTradeGoods(town.Category.MinImports, town.Category.MaxImports, foreignClimate.Resources)

	return imports
}

func generateRandomPopulation(category TownCategory) int {
	sizeIncrement := category.MaxSize - category.MinSize

	return rand.Intn(sizeIncrement) + category.MinSize
}

func (town Town) generateTownName() string {
	name := town.Culture.Language.RandomName()
	return name
}

// GenerateTown generates a random town
func GenerateTown(category string, climate string) Town {
	town := Town{}

	if category == "random" {
		town.Category = generateRandomCategory()
	} else {
		town.Category = townCategories[category]
	}
	if climate == "random" {
		town.Climate = climate.Generate()
	} else {
		town.Climate = climate.GetClimate(climate)
	}

	culture := culture.GenerateCulture()
	culture = culture.SetClimate(town.Climate.Name)
	town = SetCulture(culture, town)

	town.Mayor = town.generateMayor()
	town.Name = town.generateTownName()

	town.Exports = town.generateRandomExports()
	town.Imports = town.generateRandomImports()
	town.Population = generateRandomPopulation(town.Category)

	return town
}

// SetCulture sets the culture of a town and recalculates some things
func SetCulture(culture culture.Culture, town Town) Town {
	newTown := town

	newTown.Culture = culture
	newTown.Name = newTown.Culture.Language.RandomName()
	newTown.Mayor.FirstName = newTown.Culture.Language.RandomGenderedName(town.Mayor.Gender)
	newTown.Mayor.LastName = newTown.Culture.Language.RandomName()

	return newTown
}
