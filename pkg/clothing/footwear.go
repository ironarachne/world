package clothing

import "math/rand"

func getBootTemplates() []ItemTemplate {
	return []ItemTemplate{
		ItemTemplate{
			Name:         "boots",
			Type:         "footwear",
			MaterialType: "hide",
			PrefixModifiers: []string{
				"knee-high",
				"short",
				"tall",
				"thigh-high",
				"heavy",
				"buckled",
				"strapped",
			},
			SuffixModifiers: []string{
				"with thick soles",
				"with thin soles",
			},
		},
	}
}

func getShoeTemplates() []ItemTemplate {
	return []ItemTemplate{
		ItemTemplate{
			Name:         "sandals",
			Type:         "footwear",
			MaterialType: "hide",
			PrefixModifiers: []string{
				"thin",
				"thick",
				"strapped",
			},
			SuffixModifiers: []string{
				"with thin soles",
				"with thick soles",
				"with ankle strapping",
			},
		},
		ItemTemplate{
			Name:         "turnshoes",
			Type:         "footwear",
			MaterialType: "hide",
			PrefixModifiers: []string{
				"light",
				"heavy",
				"thick",
				"pointed",
				"rectangular",
			},
			SuffixModifiers: []string{
				"with stiff heels",
				"with stiff soles",
				"with folded-over tops",
			},
		},
	}
}

func getRandomBoots() Item {
	potentialTemplates := getBootTemplates()

	template := potentialTemplates[rand.Intn(len(potentialTemplates))]

	return getItemFromTemplate(template)
}

func getRandomShoes() Item {
	potentialTemplates := getShoeTemplates()

	template := potentialTemplates[rand.Intn(len(potentialTemplates))]

	return getItemFromTemplate(template)
}
