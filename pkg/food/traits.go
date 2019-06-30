package food

import (
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

func randomEatingTraits() []string {
	var traits []string
	var trait string
	var typesOfTraits []string
	var t string

	utensils := []string{
		"chopsticks",
		"fork",
		"hands",
		"knife and fork",
		"knife",
		"spoon",
	}

	spices := []string{
		"heavy",
		"moderate",
		"minimal",
	}

	heat := []string{
		"hot",
		"warm",
		"cold",
		"chilled",
	}

	customs := []string{
		"eat communal meals",
		"eat alone",
		"belch after eating",
		"thank the cook after eating",
	}

	potentialTraits := []string{
		"utensils",
		"spices",
		"heat",
		"customs",
	}

	for i := 0; i < 2; i++ {
		t = random.String(potentialTraits)
		if !slices.StringIn(t, typesOfTraits) {
			typesOfTraits = append(typesOfTraits, t)
			if t == "utensils" {
				trait = "eat with " + random.String(utensils)
			} else if t == "spices" {
				trait = "use " + random.String(spices) + " spice"
			} else if t == "heat" {
				trait = "serve food " + random.String(heat)
			} else if t == "customs" {
				trait = random.String(customs)
			}
			traits = append(traits, trait)
		}
	}

	return traits
}
