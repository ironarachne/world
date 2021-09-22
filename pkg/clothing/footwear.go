package clothing

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

func getBootTemplates() []ItemTemplate {
	return []ItemTemplate{
		{
			Name:         "boots",
			Type:         "footwear",
			MaterialType: "hide",
			PrefixModifiers: []string{
				"knee-high",
				"short",
				"tall",
				"thigh-high",
				"heavy",
				"buckled",
				"strapped",
			},
			SuffixModifiers: []string{
				"with thick soles",
				"with thin soles",
			},
		},
	}
}

func getShoeTemplates() []ItemTemplate {
	return []ItemTemplate{
		{
			Name:         "sandals",
			Type:         "footwear",
			MaterialType: "hide",
			PrefixModifiers: []string{
				"thin",
				"thick",
				"strapped",
			},
			SuffixModifiers: []string{
				"with thin soles",
				"with thick soles",
				"with ankle strapping",
			},
		},
		{
			Name:         "turnshoes",
			Type:         "footwear",
			MaterialType: "hide",
			PrefixModifiers: []string{
				"light",
				"heavy",
				"thick",
				"pointed",
				"rectangular",
			},
			SuffixModifiers: []string{
				"with stiff heels",
				"with stiff soles",
				"with folded-over tops",
			},
		},
	}
}

func getRandomBoots(ctx context.Context) (Item, error) {
	potentialTemplates := getBootTemplates()

	template := potentialTemplates[random.Intn(ctx, len(potentialTemplates))]

	item, err := getItemFromTemplate(ctx, template)
	if err != nil {
		err = fmt.Errorf("Could not get random boots: %w", err)
		return Item{}, err
	}

	return item, nil
}

func getRandomShoes(ctx context.Context) (Item, error) {
	potentialTemplates := getShoeTemplates()

	template := potentialTemplates[random.Intn(ctx, len(potentialTemplates))]

	item, err := getItemFromTemplate(ctx, template)
	if err != nil {
		err = fmt.Errorf("Could not get random shoes: %w", err)
		return Item{}, err
	}

	return item, nil
}
