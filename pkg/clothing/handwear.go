package clothing

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func getHandwearTemplates() []ItemTemplate {
	return []ItemTemplate{
		{
			Name:         "gloves",
			Type:         "handwear",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"long",
				"short",
				"elbow-length",
				"thick",
				"thin",
				"fingerless",
			},
			SuffixModifiers: []string{
				"that reach up to the elbow",
			},
		},
		{
			Name:         "mittens",
			Type:         "handwear",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"heavy",
				"light",
			},
			SuffixModifiers: []string{
				"with thick lining",
			},
		},
	}
}

func getRandomHandwear(ctx context.Context) (Item, error) {
	potentialTemplates := getHandwearTemplates()

	template := potentialTemplates[random.Intn(ctx, len(potentialTemplates))]

	item, err := getItemFromTemplate(ctx, template)
	if err != nil {
		err = fmt.Errorf("Could not get random handwear: %w", err)
		return Item{}, err
	}

	return item, nil
}
