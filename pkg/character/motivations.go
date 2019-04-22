package character

import "github.com/ironarachne/world/pkg/random"

func getAllMotivations() []string {
	motivations := []string{
		"achieve destiny",
		"ambition",
		"conquest",
		"conspiracy",
		"control",
		"creativity",
		"debauchery",
		"desire to better oneself",
		"desperation",
		"duty",
		"escape destiny",
		"fame",
		"family",
		"fear",
		"freedom",
		"friends",
		"greatness",
		"greed",
		"happiness",
		"hate",
		"honor",
		"justice",
		"knowledge",
		"laziness",
		"love",
		"money",
		"pleasure",
		"power",
		"religious devotion",
		"revenge",
		"rivalry",
		"romance",
		"safety",
		"stability",
		"survival",
		"to distinguish oneself",
		"to gain acceptance",
		"truth",
	}

	return motivations
}

func getRandomMotivation() string {
	motivations := getAllMotivations()

	return random.String(motivations)
}
