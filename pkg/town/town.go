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

func (town Town) generateMayor() (character.Character, error) {
	mayor, err := character.Generate(town.Culture)
	if err != nil {
		err = fmt.Errorf("Could not generate mayor: %w", err)
		return character.Character{}, err
	}
	mayor = mayor.ChangeAge(rand.Intn(30) + 30)

	return mayor, nil
}

func (town Town) generateRandomExports() []goods.TradeGood {
	var exports []goods.TradeGood

	exports = goods.GenerateExportTradeGoods(town.Category.MinExports, town.Category.MaxExports, town.Resources)

	return exports
}

func (town Town) generateRandomImports() ([]goods.TradeGood, error) {
	var imports []goods.TradeGood

	foreignClimate, err := climate.GetForeignClimate(town.Climate)
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

func (town Town) generateTownName() (string, error) {
	name, err := town.Culture.Language.RandomName()
	if err != nil {
		err = fmt.Errorf("Could not generate town name: %w", err)
		return "", err
	}
	return name, nil
}

// Generate generates a random town
func Generate(category string, originClimate climate.Climate, originCulture culture.Culture) (Town, error) {
	var newProducers []profession.Profession
	var producers []profession.Profession
	var newResources []resource.Resource

	town := Town{}

	if category == "random" {
		town.Category = getRandomWeightedCategory()
	} else {
		town.Category = getCategoryByName(category)
	}

	town.Climate = originClimate
	town.Culture = originCulture

	name, err := town.Culture.Language.RandomName()
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
	firstName, err := town.Culture.Language.RandomGenderedName(mayor.Gender.Name)
	if err != nil {
		err = fmt.Errorf("Could not generate town: %w", err)
		return Town{}, err
	}
	mayor.FirstName = firstName
	lastName, err := town.Culture.Language.RandomName()
	if err != nil {
		err = fmt.Errorf("Could not generate town: %w", err)
		return Town{}, err
	}
	mayor.LastName = lastName
	town.Mayor = mayor

	resources := town.Climate.Resources

	for i := 0; i < town.Category.ProductionIterations; i++ {
		newProducers = getProducers(town.Population, resources)
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
