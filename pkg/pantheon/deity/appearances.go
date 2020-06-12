package deity

import (
	"context"
	"fmt"

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

func getRandomGeneralAppearances(ctx context.Context, max int) ([]string, error) {
	var appearances []string

	possibleAppearances := getGeneralAppearances()

	for i := 0; i < max; i++ {
		appearance, err := random.String(ctx, possibleAppearances)
		if err != nil {
			err = fmt.Errorf("Could not generate random deity appearance: %w", err)
			return []string{}, err
		}
		if !slices.StringIn(appearance, appearances) {
			appearances = append(appearances, appearance)
		}
	}

	return appearances, nil
}
