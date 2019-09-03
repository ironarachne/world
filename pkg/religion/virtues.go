package religion

import (
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

func getRandomVirtues() []string {
	var virtue string
	all := getAllVirtues()
	virtues := []string{}

	numberOfVirtues := rand.Intn(4) + 2

	for i := 0; i < numberOfVirtues; i++ {
		virtue = random.String(all)
		if !slices.StringIn(virtue, virtues) {
			virtues = append(virtues, virtue)
		} else {
			i--
		}
	}

	return virtues
}
