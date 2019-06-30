package clothing

import "math/rand"

func getWaistTemplates() []ItemTemplate {
	return []ItemTemplate{
		ItemTemplate{
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
			IdealHeatLevel: "any",
		},
		ItemTemplate{
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
			IdealHeatLevel: "any",
		},
	}
}

func getRandomWaist() Item {
	potentialTemplates := getWaistTemplates()

	template := potentialTemplates[rand.Intn(len(potentialTemplates))]

	return getItemFromTemplate(template)
}
