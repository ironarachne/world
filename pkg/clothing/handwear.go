package clothing

import "math/rand"

func getHandwearTemplates() []ItemTemplate {
	return []ItemTemplate{
		ItemTemplate{
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
			IdealHeatLevel: "any",
		},
		ItemTemplate{
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
			IdealHeatLevel: "any",
		},
	}
}

func getRandomHandwear() Item {
	potentialTemplates := getHandwearTemplates()

	template := potentialTemplates[rand.Intn(len(potentialTemplates))]

	return getItemFromTemplate(template)
}
