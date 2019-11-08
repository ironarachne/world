package climate

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/words"
)

func (climate Climate) getDescription() (string, error) {
	description := "This is "

	description += words.Pronoun(climate.Name) + " " + climate.Name + " region"

	weather := ". The weather is "
	weatherItem := ""
	weatherItems := []string{}

	for _, s := range climate.Seasons {
		weatherItem = s.WeatherProfile.Description + " during the " + s.Name
		if len(climate.Seasons) == 2 {
			weatherItem += " season"
		}
		weatherItems = append(weatherItems, weatherItem)
	}
	weather += words.CombinePhrases(weatherItems)

	description += weather + "."

	predatorCount := 0
	birdCount := 0

	for _, a := range climate.Animals {
		if a.HasTag("carnivore") {
			predatorCount++
		}
		if a.HasTag("bird") {
			birdCount++
		}
	}

	manyBirds := rand.Intn(10)
	manyPredators := rand.Intn(10)

	if predatorCount > 1 && manyPredators > 8 {
		description += " There are a number of predators in the region."
	}

	if birdCount > 1 && manyBirds > 8 {
		description += " The skies are full of birds."
	}

	return description, nil
}
