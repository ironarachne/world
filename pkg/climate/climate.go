package climate

import (
	"fmt"
	"math/rand"
)

// Generate generates a climate
func Generate() (Climate, error) {
	rawClimate := getRandomClimate()
	climate, err := rawClimate.populate()
	if err != nil {
		err = fmt.Errorf("Could not generate climate: %w", err)
		return Climate{}, err
	}

	return climate, nil
}

// GetClimate returns a specific climate
func GetClimate(name string) (Climate, error) {
	rawClimate := getClimateByName(name)
	climate, err := rawClimate.populate()
	if err != nil {
		err = fmt.Errorf("Could not generate climate by name: %w", err)
		return Climate{}, err
	}

	return climate, nil
}

// GetForeignClimate gets a random climate that's different from the given one
func GetForeignClimate(climate Climate) (Climate, error) {
	var possibleClimates []Climate

	climates := getAllClimates()

	for _, c := range climates {
		if c.Name != climate.Name {
			possibleClimates = append(possibleClimates, c)
		}
	}

	foreignClimate := possibleClimates[rand.Intn(len(possibleClimates))]
	newClimate, err := foreignClimate.populate()
	if err != nil {
		err = fmt.Errorf("Could not generate foreign climate: %w", err)
		return Climate{}, err
	}

	return newClimate, nil
}
