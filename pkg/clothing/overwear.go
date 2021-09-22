package clothing

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func getOverwearTemplates() []ItemTemplate {
	return []ItemTemplate{
		{
			Name:         "coat",
			Type:         "overwear",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"long",
				"short",
				"thick",
				"ankle-length",
				"waist-length",
				"hip-length",
			},
			SuffixModifiers: []string{
				"with large lapels",
			},
		},
		{
			Name:         "cloak",
			Type:         "overwear",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"hooded",
				"long",
				"thick",
				"triangular",
				"half-circle",
				"full-circle",
				"hoodless",
			},
			SuffixModifiers: []string{
				"with a deep hood",
				"with a shallow hood",
				"with no hood",
			},
		},
	}
}

func getRandomOverwear(ctx context.Context) (Item, error) {
	potentialTemplates := getOverwearTemplates()

	template := potentialTemplates[random.Intn(ctx, len(potentialTemplates))]

	item, err := getItemFromTemplate(ctx, template)
	if err != nil {
		err = fmt.Errorf("Could not get random overwear: %w", err)
		return Item{}, err
	}

	return item, nil
}
