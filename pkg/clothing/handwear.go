package clothing

import (
	"fmt"
	"math/rand"
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

func getRandomHandwear() (Item, error) {
	potentialTemplates := getHandwearTemplates()

	template := potentialTemplates[rand.Intn(len(potentialTemplates))]

	item, err := getItemFromTemplate(template)
	if err != nil {
		err = fmt.Errorf("Could not get random handwear: %w", err)
		return Item{}, err
	}

	return item, nil
}
