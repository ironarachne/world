package clothing

import "math/rand"

func getOverwearTemplates() []ItemTemplate {
	return []ItemTemplate{
		ItemTemplate{
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
			IdealHeatLevel: "cold",
		},
		ItemTemplate{
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
			IdealHeatLevel: "any",
		},
	}
}

func getRandomOverwear() Item {
	potentialTemplates := getOverwearTemplates()

	template := potentialTemplates[rand.Intn(len(potentialTemplates))]

	return getItemFromTemplate(template)
}
