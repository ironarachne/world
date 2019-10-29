package climate

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
	"github.com/ironarachne/world/pkg/words"
)

func getRandomBeachDescription() (string, error) {
	beaches := []string{
		"beautiful sandy beaches",
		"long sandy beaches",
		"low cliffs",
		"rocky beaches",
		"sandy beaches",
		"sandy beaches",
		"sheer cliffs",
	}

	beach, err := random.String(beaches)
	if err != nil {
		err = fmt.Errorf("Could not generate beach description: %w", err)
		return "", err
	}

	return beach, nil
}

func getRandomLakeDescription() (string, error) {
	lakes := []string{
		"a few large lakes",
		"an enormous lake",
		"many smaller lakes",
	}

	lake, err := random.String(lakes)
	if err != nil {
		err = fmt.Errorf("Could not generate lake description: %w", err)
		return "", err
	}

	return lake, nil
}

func getRandomOceanDescription() (string, error) {
	oceans := []string{
		"a large ocean",
		"a major sea",
		"a smaller sea",
		"an ocean",
	}

	ocean, err := random.String(oceans)
	if err != nil {
		err = fmt.Errorf("Could not generate ocean description: %w", err)
		return "", err
	}

	return ocean, nil
}

func getRandomRiverDescription() (string, error) {
	rivers := []string{
		"a few rivers",
		"a major river",
		"many small rivers and streams",
	}

	river, err := random.String(rivers)
	if err != nil {
		err = fmt.Errorf("Could not generate river description: %w", err)
		return "", err
	}

	return river, nil
}

func getRandomWetlandsDescription() (string, error) {
	wetlands := []string{
		"extensive wetlands",
		"scattered wetlands",
		"several bogs",
	}

	wetland, err := random.String(wetlands)
	if err != nil {
		err = fmt.Errorf("Could not generate wetlands description: %w", err)
		return "", err
	}

	return wetland, nil
}

func (climate Climate) getDescription() (string, error) {
	description := "This is "

	description += words.Pronoun(climate.Name) + " " + climate.Name + " region"

	waterComponents := []string{}
	if climate.HasLakes {
		lakes, err := getRandomLakeDescription()
		if err != nil {
			err = fmt.Errorf("Could not generate climate description: %w", err)
			return "", err
		}
		waterComponents = append(waterComponents, lakes)
	}
	if climate.HasRivers {
		rivers, err := getRandomRiverDescription()
		if err != nil {
			err = fmt.Errorf("Could not generate climate description: %w", err)
			return "", err
		}
		waterComponents = append(waterComponents, rivers)
	}
	if climate.HasWetlands {
		wetlands, err := getRandomWetlandsDescription()
		if err != nil {
			err = fmt.Errorf("Could not generate climate description: %w", err)
			return "", err
		}
		waterComponents = append(waterComponents, wetlands)
	}
	waterComponents = slices.ShuffleStringSlice(waterComponents)

	if len(waterComponents) > 0 {
		description += " with "
		description += words.CombinePhrases(waterComponents)
	}

	if climate.HasOcean {
		description += ". It lies on the coast of "
		ocean, err := getRandomOceanDescription()
		if err != nil {
			err = fmt.Errorf("Could not generate climate description: %w", err)
			return "", err
		}
		description += ocean

		beach, err := getRandomBeachDescription()
		if err != nil {
			err = fmt.Errorf("Could not generate climate description: %w", err)
			return "", err
		}
		description += " with " + beach
	}

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

	manyBirds := rand.Intn(10)
	manyPredators := rand.Intn(10)

	if manyPredators > 8 {
		description += " There are a number of predators in the region."
	}

	if manyBirds > 8 {
		description += " The skies are full of birds."
	}

	return description, nil
}
