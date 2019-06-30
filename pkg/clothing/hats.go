package clothing

import "math/rand"

func getHatTemplates() []ItemTemplate {
	return []ItemTemplate{
		ItemTemplate{
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
		ItemTemplate{
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
		ItemTemplate{
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

func getRandomHat() Item {
	potentialTemplates := getHatTemplates()

	template := potentialTemplates[rand.Intn(len(potentialTemplates))]

	return getItemFromTemplate(template)
}
