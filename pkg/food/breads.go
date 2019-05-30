package food

import (
	"math/rand"
	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

func generateBread(originClimate climate.Climate) string {
	bread := ""
	breadTypes := []string{
		"brick-like",
		"flat",
		"long",
		"round",
		"spongy",
		"thin",
	}

	flavors := []string{
		"bitter",
		"grainy",
		"hearty",
		"nutty",
		"savory",
		"sweet",
	}
	grains := climate.ListResourcesOfType("grain", originClimate.Resources)

	if len(grains) > 0 {
		grain := grains[rand.Intn(len(grains))]

		bread = random.String(flavors) + " " + random.String(breadTypes) + " " + grain.Name + " bread"
	}

	return bread
}

func randomBreads(originClimate climate.Climate) []string {
	var bread string
	var breads []string

	grains := climate.ListResourcesOfType("grain", originClimate.Resources)
	if len(grains) > 0 {
		numberOfBreads := rand.Intn(3) + 1
		for i := 0; i < numberOfBreads; i++ {
			bread = generateBread(originClimate)
			if !slices.StringIn(bread, breads) {
				breads = append(breads, bread)
			}
		}
	}

	return breads
}