package food

import (
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

func getAllTechniques() []string {
	return []string{
		"baked",
		"basted",
		"broiled",
		"curried",
		"dried",
		"fried",
		"raw",
		"roasted",
		"slow-roasted",
		"stewed",
	}
}

func randomTechniques(maxTechniques int) []string {
	var techniques []string
	var technique string

	potentialTechniques := getAllTechniques()

	if maxTechniques < 1 {
		return []string{}
	}

	for i := 0; i < maxTechniques; i++ {
		technique = random.String(potentialTechniques)
		if !slices.StringIn(technique, techniques) {
			techniques = append(techniques, technique)
		}
	}

	return techniques
}