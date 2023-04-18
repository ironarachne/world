package clothing

import (
	"fmt"
	"math/rand"
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

func getRandomOverwear() (Item, error) {
	potentialTemplates := getOverwearTemplates()

	template := potentialTemplates[rand.Intn(len(potentialTemplates))]

	item, err := getItemFromTemplate(template)
	if err != nil {
		err = fmt.Errorf("Could not get random overwear: %w", err)
		return Item{}, err
	}

	return item, nil
}
