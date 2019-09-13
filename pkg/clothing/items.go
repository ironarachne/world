package clothing

import (
	"fmt"
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

func addMaterials(items []Item, hides []string, fabrics []string) ([]Item, error) {
	var err error
	var material string
	var newItem Item
	var result []Item

	for _, i := range items {
		newItem = i
		if i.MaterialType == "fabric" {
			material, err = random.String(fabrics)
			if err != nil {
				err = fmt.Errorf("Could not set fabric material: %w", err)
				return []Item{}, err
			}
			newItem.Material = material
			result = append(result, newItem)
		} else if i.MaterialType == "hide" {
			material, err = random.String(hides)
			if err != nil {
				err = fmt.Errorf("Could not set hide material: %w", err)
				return []Item{}, err
			}
			newItem.Material = material
			result = append(result, newItem)
		}
	}

	return result, nil
}

// GenerateOutfit generates a random outfit based on materials, temperature, and gender
func GenerateOutfit(temperature int, hides []string, fabrics []string, gender string) ([]Item, error) {
	var err error
	var item Item
	items := []Item{}

	chanceOfFull := rand.Intn(100)
	if gender == "female" {
		chanceOfFull += 30
	}

	if chanceOfFull > 50 {
		if gender == "female" {
			dressChance := rand.Intn(100)
			if dressChance > 30 {
				item, err = getRandomDress()
				if err != nil {
					err = fmt.Errorf("Could not generate outfit: %w", err)
					return []Item{}, err
				}
				items = append(items, item)
			} else {
				item, err = getRandomRobe()
				if err != nil {
					err = fmt.Errorf("Could not generate outfit: %w", err)
					return []Item{}, err
				}
				items = append(items, item)
			}
		} else {
			item, err = getRandomRobe()
			if err != nil {
				err = fmt.Errorf("Could not generate outfit: %w", err)
				return []Item{}, err
			}
			items = append(items, item)
		}
	} else {
		item, err = getRandomTop()
		if err != nil {
			err = fmt.Errorf("Could not generate outfit: %w", err)
			return []Item{}, err
		}
		items = append(items, item)
		item, err = getRandomBottom()
		if err != nil {
			err = fmt.Errorf("Could not generate outfit: %w", err)
			return []Item{}, err
		}
		items = append(items, item)
	}

	if temperature < 5 {
		item, err = getRandomHandwear()
		if err != nil {
			err = fmt.Errorf("Could not generate outfit: %w", err)
			return []Item{}, err
		}
		items = append(items, item)
		item, err = getRandomOverwear()
		if err != nil {
			err = fmt.Errorf("Could not generate outfit: %w", err)
			return []Item{}, err
		}
		items = append(items, item)
		item, err = getRandomBoots()
		if err != nil {
			err = fmt.Errorf("Could not generate outfit: %w", err)
			return []Item{}, err
		}
		items = append(items, item)
	} else {
		footwearChance := rand.Intn(100)
		if footwearChance > 50 {
			item, err = getRandomShoes()
			if err != nil {
				err = fmt.Errorf("Could not generate outfit: %w", err)
				return []Item{}, err
			}
			items = append(items, item)
		} else {
			item, err = getRandomBoots()
			if err != nil {
				err = fmt.Errorf("Could not generate outfit: %w", err)
				return []Item{}, err
			}
			items = append(items, item)
		}
	}

	hatChance := rand.Intn(100)
	if hatChance > 60 {
		item, err = getRandomHat()
		if err != nil {
			err = fmt.Errorf("Could not generate outfit: %w", err)
			return []Item{}, err
		}
		items = append(items, item)
	}

	waistChance := rand.Intn(100)
	if waistChance > 20 {
		item, err = getRandomWaist()
		if err != nil {
			err = fmt.Errorf("Could not generate outfit: %w", err)
			return []Item{}, err
		}
		items = append(items, item)
	}

	finished, err := addMaterials(items, hides, fabrics)
	if err != nil {
		err = fmt.Errorf("Could not add materials: %w", err)
		return []Item{}, err
	}

	return finished, nil
}

func getItemFromTemplate(template ItemTemplate) (Item, error) {
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

	modifier, err := random.StringFromThresholdMap(weights)
	if err != nil {
		err = fmt.Errorf("Could not get item from template: %w", err)
		return Item{}, err
	}

	if modifier == "prefix" {
		prefix, err := random.String(template.PrefixModifiers)
		if err != nil {
			err = fmt.Errorf("Could not get item from template: %w", err)
			return Item{}, err
		}
		item.PrefixModifier = prefix
	} else if modifier == "suffix" {
		suffix, err := random.String(template.SuffixModifiers)
		if err != nil {
			err = fmt.Errorf("Could not get item from template: %w", err)
			return Item{}, err
		}
		item.SuffixModifier = suffix
	}

	return item, nil
}
