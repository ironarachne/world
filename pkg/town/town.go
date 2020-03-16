/*
Package town implements towns in fantasy worlds
*/
package town

import (
	"fmt"
	"github.com/ironarachne/world/pkg/geography"
	"math/rand"

	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/resource"

	"github.com/ironarachne/world/pkg/buildings"
	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/goods"
)

// Town is a town
type Town struct {
	Name             string                  `json:"name"`
	Population       int                     `json:"population"`
	BuildingStyle    buildings.BuildingStyle `json:"building_style"`
	Category         Category                `json:"category"`
	Geography        geography.Area          `json:"geography"`
	Culture          culture.Culture         `json:"culture"`
	Mayor            character.Character     `json:"mayor"`
	NotableProducers []profession.Profession `json:"notable_producers"`
	Exports          []goods.TradeGood       `json:"exports"`
	Imports          []goods.TradeGood       `json:"imports"`
	Resources        []resource.Resource     `json:"resources"`
}

const mayorGenerationError = "failed to generate mayor: %w"
const randomTownError = "failed to generate random town: %w"
const townGenerationError = "failed to generate town: %w"

func (town Town) generateMayor() (character.Character, error) {
	mayor, err := character.Generate(town.Culture)
	if err != nil {
		err = fmt.Errorf(mayorGenerationError, err)
		return character.Character{}, err
	}
	mayor = mayor.ChangeAge(rand.Intn(30) + 30)

	firstName, err := town.Culture.Language.RandomMaleFirstName()
	if err != nil {
		err = fmt.Errorf(mayorGenerationError, err)
		return character.Character{}, err
	}

	if mayor.Gender.Name == "female" {
		firstName, err = town.Culture.Language.RandomFemaleFirstName()
		if err != nil {
			err = fmt.Errorf(mayorGenerationError, err)
			return character.Character{}, err
		}
	}

	mayor.FirstName = firstName
	lastName, err := town.Culture.Language.RandomFamilyName()
	if err != nil {
		err = fmt.Errorf(mayorGenerationError, err)
		return character.Character{}, err
	}
	mayor.LastName = lastName

	return mayor, nil
}

func (town Town) generateRandomExports() []goods.TradeGood {
	var exports []goods.TradeGood

	exports = goods.GenerateExportTradeGoods(town.Category.MinExports, town.Category.MaxExports, town.Resources)

	return exports
}

func (town Town) generateRandomImports() ([]goods.TradeGood, error) {
	var imports []goods.TradeGood

	foreignArea, err := geography.Generate()
	if err != nil {
		err = fmt.Errorf("failed to generate random imports: %w", err)
		return []goods.TradeGood{}, err
	}

	resources := foreignArea.GetResources()

	imports, err = goods.GenerateImportTradeGoods(town.Category.MinImports, town.Category.MaxImports, resources)
	if err != nil {
		err = fmt.Errorf("failed to generate random imports: %w", err)
		return []goods.TradeGood{}, err
	}

	return imports, nil
}

func generateRandomPopulation(category Category) int {
	sizeIncrement := category.MaxSize - category.MinSize

	population := rand.Intn(sizeIncrement) + category.MinSize

	return population
}

// Generate generates a random town
func Generate(categoryName string, area geography.Area, originCulture culture.Culture) (Town, error) {
	var newProducers []profession.Profession
	var producers []profession.Profession
	var newResources []resource.Resource

	town := Town{}

	if categoryName == "random" {
		townCategory, err := getRandomWeightedCategory()
		if err != nil {
			err = fmt.Errorf(townGenerationError, err)
			return Town{}, err
		}
		town.Category = townCategory
	} else {
		town.Category = getCategoryByName(categoryName)
	}

	town.Geography = area
	town.Culture = originCulture

	name, err := town.Culture.Language.RandomTownName()
	if err != nil {
		err = fmt.Errorf(townGenerationError, err)
		return Town{}, err
	}
	town.Name = name

	town.Population = generateRandomPopulation(town.Category)

	town.BuildingStyle = town.Culture.BuildingStyle

	mayor, err := town.generateMayor()
	if err != nil {
		err = fmt.Errorf(townGenerationError, err)
		return Town{}, err
	}
	town.Mayor = mayor

	resources := area.GetResources()

	for i := 0; i < town.Category.ProductionIterations; i++ {
		newProducers, err = getProducers(town.Population, resources)
		if err != nil {
			err = fmt.Errorf(townGenerationError, err)
			return Town{}, err
		}
		newResources, err = goods.Produce(newProducers, resources)
		if err != nil {
			err = fmt.Errorf(townGenerationError, err)
			return Town{}, err
		}
		resources = append(resources, newResources...)
		producers = append(producers, newProducers...)
	}

	town.Resources = resources
	town.NotableProducers = producers

	town.Exports = town.generateRandomExports()
	imports, err := town.generateRandomImports()
	if err != nil {
		err = fmt.Errorf(townGenerationError, err)
		return Town{}, err
	}
	town.Imports = imports

	return town, nil
}

// Random generates a completely random town
func Random() (Town, error) {
	area, err := geography.Generate()
	if err != nil {
		err = fmt.Errorf(randomTownError, err)
		return Town{}, err
	}

	randomCulture, err := culture.Generate(area)
	if err != nil {
		err = fmt.Errorf(randomTownError, err)
		return Town{}, err
	}

	town, err := Generate("random", area, randomCulture)
	if err != nil {
		err = fmt.Errorf(randomTownError, err)
		return Town{}, err
	}

	return town, nil
}
