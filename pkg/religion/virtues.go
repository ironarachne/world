package religion

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

func getAllVirtues() []string {
	all := []string{
		"acetism",
		"balance",
		"bravery",
		"carefulness",
		"civic duty",
		"cleanliness",
		"community",
		"compassion",
		"education",
		"faith",
		"generosity",
		"honesty",
		"honor",
		"hope",
		"justice",
		"love",
		"loyalty",
		"nobility",
		"order",
		"prudence",
		"respect",
		"self-reflection",
		"strength",
		"subtlety",
		"temperance",
		"wealth",
	}

	return all
}

func getRandomVirtues() ([]string, error) {
	all := getAllVirtues()
	virtues := []string{}

	numberOfVirtues := rand.Intn(4) + 2

	for i := 0; i < numberOfVirtues; i++ {
		virtue, err := random.String(all)
		if err != nil {
			err = fmt.Errorf("Could not generate random virtues: %w", err)
			return []string{}, err
		}
		if !slices.StringIn(virtue, virtues) {
			virtues = append(virtues, virtue)
		} else {
			i--
		}
	}

	return virtues, nil
}
