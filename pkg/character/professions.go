package character

import "github.com/ironarachne/world/pkg/random"

func getAllProfessions() []string {
	professions := []string{
		"alchemist",
		"baker",
		"bartender",
		"blacksmith",
		"bowyer",
		"carpenter",
		"cook",
		"farmer",
		"ferrier",
		"fletcher",
		"guard",
		"guildmaster",
		"healer",
		"herbalist",
		"mage",
		"merchant",
	}

	return professions
}

func getRandomProfession() string {
	professions := getAllProfessions()

	return random.String(professions)
}
