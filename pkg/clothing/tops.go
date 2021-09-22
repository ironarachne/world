package clothing

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func getTopTemplates() []ItemTemplate {
	return []ItemTemplate{
		{
			Name:         "tunic",
			Type:         "top",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"short",
				"long",
				"knee-length",
				"ankle-length",
				"waist-length",
				"sleeveless",
			},
			SuffixModifiers: []string{
				"with broad sleeves",
				"with short sleeves",
				"with no sleeves",
			},
		},
		{
			Name:         "shirt",
			Type:         "top",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"short",
				"long",
				"sleeveless",
			},
			SuffixModifiers: []string{
				"with broad sleeves",
				"with short sleeves",
				"with no sleeves",
			},
		},
	}
}

func getRandomTop(ctx context.Context) (Item, error) {
	potentialTemplates := getTopTemplates()

	template := potentialTemplates[random.Intn(ctx, len(potentialTemplates))]

	item, err := getItemFromTemplate(ctx, template)
	if err != nil {
		err = fmt.Errorf("Could not get random top: %w", err)
		return Item{}, err
	}

	return item, nil
}
