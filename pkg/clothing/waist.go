package clothing

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func getWaistTemplates() []ItemTemplate {
	return []ItemTemplate{
		{
			Name:         "belt",
			Type:         "waist",
			MaterialType: "hide",
			PrefixModifiers: []string{
				"wide",
				"narrow",
				"banded",
			},
			SuffixModifiers: []string{
				"with a thick buckle",
				"with a narrow clasp",
			},
		},
		{
			Name:         "sash",
			Type:         "waist",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"broad",
				"narrow",
				"layered",
				"shoulder-to-waist",
			},
			SuffixModifiers: []string{
				"over the shoulder",
				"around the waist",
			},
		},
	}
}

func getRandomWaist(ctx context.Context) (Item, error) {
	potentialTemplates := getWaistTemplates()

	template := potentialTemplates[random.Intn(ctx, len(potentialTemplates))]

	item, err := getItemFromTemplate(ctx, template)
	if err != nil {
		err = fmt.Errorf("Could not get random waist: %w", err)
		return Item{}, err
	}

	return item, nil
}
