package race

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/size"
)

// Race is a race. It can have variations within.
type Race struct {
	Name          string
	PluralName    string
	Adjective     string
	AgeCategories []AgeCategory
	Appearance    Appearance
	Commonality   int
	SizeCategory  size.Category
}

// GenerateSubrace generates a random subrace based on a parent race
func GenerateSubrace(parent Race) (Race, error) {
	race := parent

	appearance, err := parent.generateRandomSubraceAppearance()
	if err != nil {
		err = fmt.Errorf("Could not generate random subrace: %w", err)
		return Race{}, err
	}
	race.Appearance = appearance

	return race, nil
}

// Get returns a specific race
func Get(name string) Race {
	race := Race{}
	races := getAllRaces()

	for _, r := range races {
		if r.Name == name {
			return r
		}
	}

	return race
}

// GetRandom returns a random race from the list
func GetRandom() Race {
	races := getAllRaces()

	return races[rand.Intn(len(races))]
}

// GetRandomWeighted returns a random race, taking commonality into account
func GetRandomWeighted() (Race, error) {
	races := getAllRaces()

	weights := map[string]int{}

	for _, c := range races {
		weights[c.Name] = c.Commonality
	}

	name, err := random.StringFromThresholdMap(weights)
	if err != nil {
		err = fmt.Errorf("Failed to get random weighted race: %w", err)
		return Race{}, err
	}

	for _, c := range races {
		if c.Name == name {
			return c, nil
		}
	}

	err = fmt.Errorf("Failed to get random weighted race!")
	return Race{}, err
}

// RandomSimplified returns a random simplified race
func RandomSimplified() SimplifiedRace {
	races := getAllRaces()

	return races[rand.Intn(len(races))].Simplify()
}
