package clothing

import (
	"fmt"
	"math/rand"
)

func getBootTemplates() []ItemTemplate {
	return []ItemTemplate{
		{
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
		{
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
		{
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

func getRandomBoots() (Item, error) {
	potentialTemplates := getBootTemplates()

	template := potentialTemplates[rand.Intn(len(potentialTemplates))]

	item, err := getItemFromTemplate(template)
	if err != nil {
		err = fmt.Errorf("Could not get random boots: %w", err)
		return Item{}, err
	}

	return item, nil
}

func getRandomShoes() (Item, error) {
	potentialTemplates := getShoeTemplates()

	template := potentialTemplates[rand.Intn(len(potentialTemplates))]

	item, err := getItemFromTemplate(template)
	if err != nil {
		err = fmt.Errorf("Could not get random shoes: %w", err)
		return Item{}, err
	}

	return item, nil
}
