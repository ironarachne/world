/*
Package town implements towns in fantasy worlds
*/
package town

import (
	"fmt"
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
	Name             string                  `json:"name"`
	Population       int                     `json:"population"`
	BuildingStyle    buildings.BuildingStyle `json:"building_style"`
	Category         Category                `json:"category"`
	Climate          climate.Climate         `json:"climate"`
	Culture          culture.Culture         `json:"culture"`
	Mayor            character.Character     `json:"mayor"`
	NotableProducers []profession.Profession `json:"notable_producers"`
	Exports          []goods.TradeGood       `json:"exports"`
	Imports          []goods.TradeGood       `json:"imports"`
	Resources        []resource.Resource     `json:"resources"`
}

func (town Town) generateMayor() (character.Character, error) {
	mayor, err := character.Generate(town.Culture)
	if err != nil {
		err = fmt.Errorf("Could not generate mayor: %w", err)
		return character.Character{}, err
	}
	mayor = mayor.ChangeAge(rand.Intn(30) + 30)

	firstName, err := town.Culture.Language.RandomMaleFirstName()
	if err != nil {
		err = fmt.Errorf("Could not generate mayor: %w", err)
		return character.Character{}, err
	}

	if mayor.Gender.Name == "female" {
		firstName, err = town.Culture.Language.RandomFemaleFirstName()
		if err != nil {
			err = fmt.Errorf("Could not generate mayor: %w", err)
			return character.Character{}, err
		}
	}

	mayor.FirstName = firstName
	lastName, err := town.Culture.Language.RandomFamilyName()
	if err != nil {
		err = fmt.Errorf("Could not generate mayor: %w", err)
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

	foreignClimate, err := climate.GenerateForeign(town.Climate)
	if err != nil {
		err = fmt.Errorf("Could not generate imports: %w", err)
		return []goods.TradeGood{}, err
	}

	imports, err = goods.GenerateImportTradeGoods(town.Category.MinImports, town.Category.MaxImports, foreignClimate.Resources)
	if err != nil {
		err = fmt.Errorf("Could not generate imports: %w", err)
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
func Generate(categoryName string, originClimate climate.Climate, originCulture culture.Culture) (Town, error) {
	var newProducers []profession.Profession
	var producers []profession.Profession
	var newResources []resource.Resource

	town := Town{}

	if categoryName == "random" {
		townCategory, err := getRandomWeightedCategory()
		if err != nil {
			err = fmt.Errorf("Could not generate town: %w", err)
			return Town{}, err
		}
		town.Category = townCategory
	} else {
		town.Category = getCategoryByName(categoryName)
	}

	town.Climate = originClimate
	town.Culture = originCulture

	name, err := town.Culture.Language.RandomTownName()
	if err != nil {
		err = fmt.Errorf("Could not generate town: %w", err)
		return Town{}, err
	}
	town.Name = name

	town.Population = generateRandomPopulation(town.Category)

	town.BuildingStyle = town.Culture.BuildingStyle

	mayor, err := town.generateMayor()
	if err != nil {
		err = fmt.Errorf("Could not generate town: %w", err)
		return Town{}, err
	}
	town.Mayor = mayor

	resources := town.Climate.Resources

	for i := 0; i < town.Category.ProductionIterations; i++ {
		newProducers, err = getProducers(town.Population, resources)
		if err != nil {
			err = fmt.Errorf("Could not generate town: %w", err)
			return Town{}, err
		}
		newResources, err = goods.Produce(newProducers, resources)
		if err != nil {
			err = fmt.Errorf("Could not generate town: %w", err)
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
		err = fmt.Errorf("Could not generate town: %w", err)
		return Town{}, err
	}
	town.Imports = imports

	return town, nil
}

// Random generates a completely random town
func Random() (Town, error) {
	randomCulture, err := culture.Random()
	if err != nil {
		err = fmt.Errorf("Could not generate random town: %w", err)
		return Town{}, err
	}

	town, err := Generate("random", randomCulture.HomeClimate, randomCulture)
	if err != nil {
		err = fmt.Errorf("Could not generate random town: %w", err)
		return Town{}, err
	}

	return town, nil
}
