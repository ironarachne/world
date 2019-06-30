package clothing

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
)

// ItemTemplate is a pattern for constructing an item
type ItemTemplate struct {
	Name            string
	Type            string
	MaterialType    string
	PrefixModifiers []string
	SuffixModifiers []string
}

// Item is a type of clothing item
type Item struct {
	Name           string
	Type           string
	Material       string
	MaterialType   string
	PrefixModifier string
	SuffixModifier string
}

func addMaterials(items []Item, hides []string, fabrics []string) []Item {
	var newItem Item
	var result []Item

	for _, i := range items {
		newItem = i
		if i.MaterialType == "fabric" {
			newItem.Material = random.String(fabrics)
			result = append(result, newItem)
		} else if i.MaterialType == "hide" {
			newItem.Material = random.String(hides)
			result = append(result, newItem)
		}
	}

	return result
}

// GenerateOutfit generates a random outfit based on materials, temperature, and gender
func GenerateOutfit(temperature int, hides []string, fabrics []string, gender string) []Item {
	items := []Item{}

	chanceOfFull := rand.Intn(100)
	if gender == "female" {
		chanceOfFull += 30
	}

	if chanceOfFull > 50 {
		if gender == "female" {
			dressChance := rand.Intn(100)
			if dressChance > 30 {
				items = append(items, getRandomDress())
			} else {
				items = append(items, getRandomRobe())
			}
		} else {
			items = append(items, getRandomRobe())
		}
	} else {
		items = append(items, getRandomTop())
		items = append(items, getRandomBottom())
	}

	if temperature < 5 {
		items = append(items, getRandomHandwear())
		items = append(items, getRandomOverwear())
		items = append(items, getRandomBoots())
	} else {
		footwearChance := rand.Intn(100)
		if footwearChance > 50 {
			items = append(items, getRandomShoes())
		} else {
			items = append(items, getRandomBoots())
		}
	}

	hatChance := rand.Intn(100)
	if hatChance > 60 {
		items = append(items, getRandomHat())
	}

	waistChance := rand.Intn(100)
	if waistChance > 20 {
		items = append(items, getRandomWaist())
	}

	items = addMaterials(items, hides, fabrics)

	return items
}

func getItemFromTemplate(template ItemTemplate) Item {
	item := Item{
		Name:         template.Name,
		Type:         template.Type,
		MaterialType: template.MaterialType,
	}

	weights := map[string]int{
		"prefix": len(template.PrefixModifiers),
		"suffix": len(template.SuffixModifiers),
		"none":   3,
	}

	modifier := random.StringFromThresholdMap(weights)

	if modifier == "prefix" {
		item.PrefixModifier = random.String(template.PrefixModifiers)
	} else if modifier == "suffix" {
		item.SuffixModifier = random.String(template.SuffixModifiers)
	}

	return item
}
