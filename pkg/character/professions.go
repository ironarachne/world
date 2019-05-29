package character

import (
	"github.com/ironarachne/world/pkg/goods"
	"github.com/ironarachne/world/pkg/random"
)

func getAllProfessions() []string {
	professions := []string{
		"bartender",
		"cook",
		"farmer",
		"guard",
		"guildmaster",
		"healer",
		"mage",
		"merchant",
	}

	producers := goods.ProducerList()

	professions = append(professions, producers...)

	return professions
}

func getRandomProfession() string {
	professions := getAllProfessions()

	return random.String(professions)
}
