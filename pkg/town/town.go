/*
Package town implements towns in fantasy worlds
*/
package town

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/geography"
	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/random"
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

func (town Town) generateMayor(ctx context.Context) (character.Character, error) {
	mayor, err := character.Generate(ctx, town.Culture)
	if err != nil {
		err = fmt.Errorf(mayorGenerationError, err)
		return character.Character{}, err
	}
	mayor = mayor.ChangeAge(ctx, random.Intn(ctx, 30)+30)

	firstName, err := town.Culture.Language.RandomMaleFirstName(ctx)
	if err != nil {
		err = fmt.Errorf(mayorGenerationError, err)
		return character.Character{}, err
	}

	if mayor.Gender.Name == "female" {
		firstName, err = town.Culture.Language.RandomFemaleFirstName(ctx)
		if err != nil {
			err = fmt.Errorf(mayorGenerationError, err)
			return character.Character{}, err
		}
	}

	mayor.FirstName = firstName
	lastName, err := town.Culture.Language.RandomFamilyName(ctx)
	if err != nil {
		err = fmt.Errorf(mayorGenerationError, err)
		return character.Character{}, err
	}
	mayor.LastName = lastName

	return mayor, nil
}

func (town Town) generateRandomExports(ctx context.Context) []goods.TradeGood {
	var exports []goods.TradeGood

	exports = goods.GenerateExportTradeGoods(ctx, town.Category.MinExports, town.Category.MaxExports, town.Resources)

	return exports
}

func (town Town) generateRandomImports(ctx context.Context) ([]goods.TradeGood, error) {
	var imports []goods.TradeGood

	foreignArea, err := geography.Generate(ctx)
	if err != nil {
		err = fmt.Errorf("failed to generate random imports: %w", err)
		return []goods.TradeGood{}, err
	}

	resources := foreignArea.GetResources()

	imports, err = goods.GenerateImportTradeGoods(ctx, town.Category.MinImports, town.Category.MaxImports, resources)
	if err != nil {
		err = fmt.Errorf("failed to generate random imports: %w", err)
		return []goods.TradeGood{}, err
	}

	return imports, nil
}

func generateRandomPopulation(ctx context.Context, category Category) int {
	sizeIncrement := category.MaxSize - category.MinSize

	population := random.Intn(ctx, sizeIncrement) + category.MinSize

	return population
}

// Generate generates a random town
func Generate(ctx context.Context, categoryName string, area geography.Area, originCulture culture.Culture) (Town, error) {
	var newProducers []profession.Profession
	var producers []profession.Profession
	var newResources []resource.Resource

	town := Town{}

	if categoryName == "random" {
		townCategory, err := getRandomWeightedCategory(ctx)
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

	name, err := town.Culture.Language.RandomTownName(ctx)
	if err != nil {
		err = fmt.Errorf(townGenerationError, err)
		return Town{}, err
	}
	town.Name = name

	town.Population = generateRandomPopulation(ctx, town.Category)

	town.BuildingStyle = town.Culture.BuildingStyle

	mayor, err := town.generateMayor(ctx)
	if err != nil {
		err = fmt.Errorf(townGenerationError, err)
		return Town{}, err
	}
	town.Mayor = mayor

	resources := area.GetResources()

	for i := 0; i < town.Category.ProductionIterations; i++ {
		newProducers, err = getProducers(ctx, town.Population, resources)
		if err != nil {
			err = fmt.Errorf(townGenerationError, err)
			return Town{}, err
		}
		newResources, err = goods.Produce(ctx, newProducers, resources)
		if err != nil {
			err = fmt.Errorf(townGenerationError, err)
			return Town{}, err
		}
		resources = append(resources, newResources...)
		producers = append(producers, newProducers...)
	}

	town.Resources = resources
	town.NotableProducers = producers

	town.Exports = town.generateRandomExports(ctx)
	imports, err := town.generateRandomImports(ctx)
	if err != nil {
		err = fmt.Errorf(townGenerationError, err)
		return Town{}, err
	}
	town.Imports = imports

	return town, nil
}

// Random generates a completely random town
func Random(ctx context.Context) (Town, error) {
	area, err := geography.Generate(ctx)
	if err != nil {
		err = fmt.Errorf(randomTownError, err)
		return Town{}, err
	}

	randomCulture, err := culture.Generate(ctx, area)
	if err != nil {
		err = fmt.Errorf(randomTownError, err)
		return Town{}, err
	}

	town, err := Generate(ctx, "random", area, randomCulture)
	if err != nil {
		err = fmt.Errorf(randomTownError, err)
		return Town{}, err
	}

	return town, nil
}
