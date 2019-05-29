package character

import "github.com/ironarachne/world/pkg/random"

func getAllMotivations() []string {
	motivations := []string{
		"a desire to better oneself",
		"achieving destiny",
		"ambition",
		"balance in all things",
		"conquest",
		"conspiracy",
		"control",
		"creativity",
		"debauchery",
		"desperation",
		"distinguishing oneself",
		"duty",
		"escaping destiny",
		"fairness",
		"faith",
		"fame",
		"family",
		"fear",
		"freedom",
		"friends",
		"gaining acceptance",
		"greatness",
		"greed",
		"happiness",
		"hate",
		"honor",
		"justice",
		"knowledge",
		"laziness",
		"love",
		"loyalty",
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
		"the pursuit of perfection",
		"truth",
	}

	return motivations
}

func getRandomMotivation() string {
	motivations := getAllMotivations()

	return random.String(motivations)
}
