package clothing

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

const jewelryError = "failed to generate jewelry: %w"

func generateJewelry(ctx context.Context) ([]string, error) {
	var chanceOfAdornment int
	var descriptor string
	var err error
	var jewelryItem string
	var itemType string
	var setting string

	jewelry := []string{}

	mainMaterials := []string{
		"metal",
		"hide",
		"bone",
	}

	secondaryComponents := []string{
		"gems",
		"beads",
	}

	descriptors := []string{
		"brilliant",
		"entwined",
		"gaudy",
		"large",
		"lustrous",
		"ornate",
		"simple",
		"straight",
		"thin",
		"twisting",
	}

	itemTypes := []string{
		"anklets",
		"bracelets",
		"chokers",
		"necklaces",
		"rings",
	}

	settings := []string{
		"adorned with",
		"decorated with",
		"set with",
	}

	numberOfJewelryPieces := random.Intn(ctx, 4) + 1

	primaryMaterial, err := random.String(ctx, mainMaterials)
	if err != nil {
		err = fmt.Errorf(jewelryError, err)
		return []string{}, err
	}
	primaryComponent, err := random.String(ctx, secondaryComponents)
	if err != nil {
		err = fmt.Errorf(jewelryError, err)
		return []string{}, err
	}

	for i := 0; i < numberOfJewelryPieces; i++ {
		descriptor, err = random.String(ctx, descriptors)
		if err != nil {
			err = fmt.Errorf(jewelryError, err)
			return []string{}, err
		}
		itemType, err = random.String(ctx, itemTypes)
		if err != nil {
			err = fmt.Errorf(jewelryError, err)
			return []string{}, err
		}
		setting, err = random.String(ctx, settings)
		if err != nil {
			err = fmt.Errorf(jewelryError, err)
			return []string{}, err
		}
		jewelryItem = descriptor + " " + primaryMaterial + " " + itemType

		chanceOfAdornment = random.Intn(ctx, 100)
		if chanceOfAdornment > 50 {
			jewelryItem += " " + setting + " " + primaryComponent
		}

		jewelry = append(jewelry, jewelryItem)
	}

	return jewelry, nil
}
