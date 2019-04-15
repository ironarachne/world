package climate

import "math/rand"

// Generate generates a climate
func Generate() Climate {
	climate := getRandomClimate()
	climate = climate.populate()

	return climate
}

// GetClimate returns a specific climate
func GetClimate(name string) Climate {
	climate := getClimateByName(name)
	climate = climate.populate()

	return climate
}

// GetForeignClimate gets a random climate that's different from the given one
func GetForeignClimate(climate Climate) Climate {
	var possibleClimates []Climate

	climates := getAllClimates()

	for _, c := range climates {
		if c.Name != climate.Name {
			possibleClimates = append(possibleClimates, c)
		}
	}

	foreignClimate := possibleClimates[rand.Intn(len(possibleClimates)-1)]
	foreignClimate = foreignClimate.populate()

	return foreignClimate
}
