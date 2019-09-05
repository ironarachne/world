package clothing

import (
	"fmt"
	"math/rand"
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

func getRandomWaist() (Item, error) {
	potentialTemplates := getWaistTemplates()

	template := potentialTemplates[rand.Intn(len(potentialTemplates))]

	item, err := getItemFromTemplate(template)
	if err != nil {
		err = fmt.Errorf("Could not get random waist: %w", err)
		return Item{}, err
	}

	return item, nil
}
