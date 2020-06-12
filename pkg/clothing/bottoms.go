package clothing

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func getBottomTemplates() []ItemTemplate {
	return []ItemTemplate{
		{
			Name:         "breeches",
			Type:         "bottom",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"short",
				"long",
				"loose",
				"baggy",
			},
			SuffixModifiers: []string{
				"tied at the ankles",
				"bunched at the hips",
				"with flared ends",
				"tied at the waist",
			},
		},
		{
			Name:         "pants",
			Type:         "bottom",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"ankle-length",
				"baggy",
				"flared",
				"long",
				"narrow",
				"short",
				"straight",
				"tight-fitting",
			},
			SuffixModifiers: []string{
				"tied at the waist",
				"tied at the ankles",
				"gathered at the hips",
			},
		},
		{
			Name:         "pantaloons",
			Type:         "bottom",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"baggy",
				"wide",
			},
			SuffixModifiers: []string{
				"bunched at the hips",
				"flared at the ends",
			},
		},
		{
			Name:         "skirt",
			Type:         "bottom",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"knee-length",
				"ankle-length",
				"short",
				"pleated",
			},
			SuffixModifiers: []string{
				"down to the knees",
				"flared at the hem",
				"tied at the waist",
			},
		},
		{
			Name:         "trews",
			Type:         "bottom",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"tight",
				"ribbed",
			},
			SuffixModifiers: []string{
				"with trim down the sides",
				"with reinforced stitching",
			},
		},
		{
			Name:         "kilt",
			Type:         "bottom",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"belted",
				"calf-length",
				"knee-length",
				"pleated",
			},
			SuffixModifiers: []string{
				"belted around the waist",
				"down to the calves",
				"down to the knees",
			},
		},
	}
}

func getRandomBottom(ctx context.Context) (Item, error) {
	potentialTemplates := getBottomTemplates()

	template := potentialTemplates[random.Intn(ctx, len(potentialTemplates))]

	item, err := getItemFromTemplate(ctx, template)
	if err != nil {
		err = fmt.Errorf("Could not get random bottom: %w", err)
		return Item{}, err
	}

	return item, nil
}
