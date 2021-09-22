package buildings

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func getAllDecorations() []string {
	return []string{
		"statues stand outside doorways",
		"signs bearing the occupants' names stand out front",
		"geometric shapes are carved into the trim",
		"ceilings are often painted in ornate styles",
		"ceiling are often painted in geometric styles",
		"decorative pillars are placed inside and out, with wealthier homes having more",
		"wealthier houses and communal buildings often have small pools just outside",
		"people often plant bushes around their homes",
		"people often plant small trees around their homes",
		"people often plant large trees around their homes",
	}
}

func getRandomDecorations(ctx context.Context) (string, error) {
	decorations := getAllDecorations()

	decoration, err := random.String(ctx, decorations)

	if err != nil {
		err = fmt.Errorf("Could not find decoration: %w", err)
		return "", err
	}

	return decoration, nil
}
