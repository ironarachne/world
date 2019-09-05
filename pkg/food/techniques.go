package food

import (
	"fmt"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

func getAllTechniques() []string {
	return []string{
		"baked",
		"basted",
		"broiled",
		"curried",
		"deep fried",
		"dried",
		"lightly fried",
		"raw",
		"roasted",
		"slow-roasted",
		"stewed",
	}
}

func randomTechniques(maxTechniques int) ([]string, error) {
	var err error
	var techniques []string
	var technique string

	potentialTechniques := getAllTechniques()

	if maxTechniques < 1 {
		err = fmt.Errorf("Requested less than one food technique")
		return []string{}, err
	}

	for i := 0; i < maxTechniques; i++ {
		technique, err = random.String(potentialTechniques)
		if err != nil {
			err = fmt.Errorf("Could not generate food techniques: %w", err)
			return []string{}, err
		}
		if !slices.StringIn(technique, techniques) {
			techniques = append(techniques, technique)
		}
	}

	return techniques, nil
}
