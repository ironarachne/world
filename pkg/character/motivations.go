package character

import "github.com/ironarachne/world/pkg/random"

func getAllMotivations() []string {
	motivations := []string{
		"debauchery",
		"fame",
		"family",
		"friends",
		"greatness",
		"happiness",
		"justice",
		"knowledge",
		"laziness",
		"money",
		"power",
		"revenge",
		"truth",
	}

	return motivations
}

func getRandomMotivation() string {
	motivations := getAllMotivations()

	return random.String(motivations)
}
