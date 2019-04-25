package pantheon

import (
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

func getGeneralAppearances() []string {
	appearances := []string{
		"anguished",
		"attractive",
		"bald",
		"beautiful",
		"blind",
		"bulky",
		"chiselled",
		"clean",
		"cute",
		"elfin",
		"filthy",
		"gargantuan",
		"hairy",
		"handsome",
		"hefty",
		"homely",
		"lantern-jawed",
		"muscular",
		"one-armed",
		"plain",
		"plump",
		"pretty",
		"rotund",
		"short",
		"slender",
		"stocky",
		"stout",
		"sweaty",
		"tall",
		"thin",
		"tidy",
		"weak",
	}

	return appearances
}

func getRandomGeneralAppearances(max int) []string {
	var appearances []string
	var appearance string

	possibleAppearances := getGeneralAppearances()

	for i := 0; i < max; i++ {
		appearance = random.String(possibleAppearances)
		if !slices.StringIn(appearance, appearances) {
			appearances = append(appearances, appearance)
		}
	}

	return appearances
}
