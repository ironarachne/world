package clothing

import (
	"math/rand"
)

func getTopTemplates() []ItemTemplate {
	return []ItemTemplate{
		ItemTemplate{
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
			IdealHeatLevel: "any",
		},
		ItemTemplate{
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
			IdealHeatLevel: "warm",
		},
	}
}

func getRandomTop() Item {
	potentialTemplates := getTopTemplates()

	template := potentialTemplates[rand.Intn(len(potentialTemplates))]

	return getItemFromTemplate(template)
}
