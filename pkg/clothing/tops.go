package clothing

import (
	"fmt"
	"math/rand"
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

func getRandomTop() (Item, error) {
	potentialTemplates := getTopTemplates()

	template := potentialTemplates[rand.Intn(len(potentialTemplates))]

	item, err := getItemFromTemplate(template)
	if err != nil {
		err = fmt.Errorf("Could not get random top: %w", err)
		return Item{}, err
	}

	return item, nil
}
