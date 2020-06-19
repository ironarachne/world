package food

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func (style Style) getRandomTreatment(ctx context.Context) (string, error) {
	treatments := []string{
		" soup",
		" stew",
	}

	if len(style.Noodles) > 0 {
		for _, n := range style.Noodles {
			treatments = append(treatments, " on "+n+"s")
			treatments = append(treatments, " in "+n+" soup")
		}
	}

	treatment, err := random.String(ctx, treatments)
	if err != nil {
		err = fmt.Errorf("Could not generate food treatment: %w", err)
		return "", err
	}

	chanceOfPlain := random.Intn(ctx, 10)

	if chanceOfPlain > 4 {
		treatment = ""
	}

	return treatment, nil
}
