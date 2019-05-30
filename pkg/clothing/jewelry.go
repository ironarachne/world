package clothing

import (
	"math/rand"
	"strings"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/random"
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

	materials := []string{}
	gems := []string{}

	for _, r := range originClimate.Resources {
		if r.Type == "metal ingot" {
			materials = append(materials, strings.TrimSuffix(r.Name, " ingot"))
		} else if r.Type == "metal bar" {
			materials = append(materials, strings.TrimSuffix(r.Name, " bar"))
		} else if r.Type == "gem" {
			gems = append(gems, r.Name)
		}
	}

	numberOfJewelryPieces := rand.Intn(4) + 1

	for i := 0; i < numberOfJewelryPieces; i++ {
		jewelryItem = random.String(descriptors) + " " + random.String(materials) + " " + random.String(foundations)
		if len(gems) > 0 {
			gemProbability = rand.Intn(10) + 1
			if gemProbability > 5 {
				jewelryItem += " " + random.String(settings) + " " + random.String(gems)
			}
		}

		jewelry = append(jewelry, jewelryItem)
	}

	return jewelry
}
