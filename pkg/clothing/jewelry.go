package clothing

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
)

func generateJewelry(originClimate climate.Climate) []string {
	var jewelryItem string
	var gemProbability int

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

	metals := resource.ListOfType("metal", originClimate.Resources)
	gems := resource.ListOfType("gem", originClimate.Resources)

	numberOfJewelryPieces := rand.Intn(4) + 1

	primaryMaterial := resource.Random(metals)
	primaryGem := resource.Random(gems)

	for i := 0; i < numberOfJewelryPieces; i++ {
		jewelryItem = random.String(descriptors) + " " + primaryMaterial.Name + " " + random.String(foundations)
		if len(gems) > 0 {
			gemProbability = rand.Intn(10) + 1
			if gemProbability > 5 {
				jewelryItem += " " + random.String(settings) + " " + primaryGem.Name
			}
		}

		jewelry = append(jewelry, jewelryItem)
	}

	return jewelry
}
