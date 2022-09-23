package clothing

import (
	"fmt"
	"math/rand"
)

func getHatTemplates() []ItemTemplate {
	return []ItemTemplate{
		{
			Name:         "hat",
			Type:         "hat",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"wide-brimmed",
				"narrow-brimmed",
				"conical",
				"cylindrical",
				"tricorn",
			},
			SuffixModifiers: []string{
				"with a wide brim",
				"with a narrow brim",
				"with ear flaps",
			},
		},
		{
			Name:         "cap",
			Type:         "hat",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"conical",
				"skull",
				"tight",
				"light",
				"heavy",
				"tapered",
			},
			SuffixModifiers: []string{
				"with straps",
			},
		},
		{
			Name:         "turban",
			Type:         "hat",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"wide",
				"tight",
				"complexly woven",
				"tightly woven",
				"loosely woven",
			},
			SuffixModifiers: []string{
				"with a draped neck cover",
			},
		},
	}
}

func getRandomHat() (Item, error) {
	potentialTemplates := getHatTemplates()

	template := potentialTemplates[rand.Intn(len(potentialTemplates))]

	item, err := getItemFromTemplate(template)
	if err != nil {
		err = fmt.Errorf("Could not get random hat: %w", err)
		return Item{}, err
	}

	return item, nil
}
