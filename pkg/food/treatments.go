package food

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
)

func (style Style) getRandomTreatment() string {
	treatments := []string{
		" soup",
		" stew",
	}

	if len(style.Noodles) > 0 {
		for _, n := range style.Noodles {
			treatments = append(treatments, " on "+n)
			treatments = append(treatments, " in "+n+" soup")
		}
	}

	treatment := random.String(treatments)

	chanceOfPlain := rand.Intn(10)

	if chanceOfPlain > 4 {
		treatment = ""
	}

	return treatment
}
