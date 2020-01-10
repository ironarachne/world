package clothing

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
)

const jewelryError = "failed to generate jewelry: %w"

func generateJewelry(originClimate climate.Climate) ([]string, error) {
	var descriptor string
	var err error
	var foundation string
	var gemProbability int
	var jewelryItem string
	var setting string

	jewelry := []string{}

	descriptors := []string{
		"brilliant",
		"gaudy",
		"lustrous",
		"ornate",
		"simple",
	}

	foundations := []string{
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

	metals := resource.ByTag("metal ore", originClimate.Resources)
	gems := resource.ByTag("gem ore", originClimate.Resources)

	numberOfJewelryPieces := rand.Intn(4) + 1

	primaryMaterial := resource.Random(metals)
	primaryGem := resource.Random(gems)

	for i := 0; i < numberOfJewelryPieces; i++ {
		descriptor, err = random.String(descriptors)
		if err != nil {
			err = fmt.Errorf(jewelryError, err)
			return []string{}, err
		}
		foundation, err = random.String(foundations)
		if err != nil {
			err = fmt.Errorf(jewelryError, err)
			return []string{}, err
		}
		setting, err = random.String(settings)
		if err != nil {
			err = fmt.Errorf(jewelryError, err)
			return []string{}, err
		}
		jewelryItem = descriptor + " " + primaryMaterial.Name + " " + foundation
		if len(gems) > 0 {
			gemProbability = rand.Intn(10) + 1
			if gemProbability > 5 {
				jewelryItem += " " + setting + " " + primaryGem.Name
			}
		}

		jewelry = append(jewelry, jewelryItem)
	}

	return jewelry, nil
}
