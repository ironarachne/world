package clothing

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

const itemError = "failed to get item from template: %w"
const outfitError = "failed to generate outfit: %w"

// ItemTemplate is a pattern for constructing an item
type ItemTemplate struct {
	Name            string   `json:"name"`
	Type            string   `json:"type"`
	MaterialType    string   `json:"material_type"`
	PrefixModifiers []string `json:"prefix_modifiers"`
	SuffixModifiers []string `json:"suffix_modifiers"`
}

// Item is a type of clothing item
type Item struct {
	Name           string `json:"name"`
	Type           string `json:"type"`
	MaterialType   string `json:"material_type"`
	PrefixModifier string `json:"prefix_modifier"`
	SuffixModifier string `json:"suffix_modifier"`
}

// GenerateOutfit generates a random outfit based on environment temperature and gender
func GenerateOutfit(ctx context.Context, temperature int, gender string) ([]Item, error) {
	var err error
	var item Item
	items := []Item{}

	chanceOfFull := random.Intn(ctx, 100)
	if gender == "female" {
		chanceOfFull += 30
	}

	if chanceOfFull > 50 {
		if gender == "female" {
			dressChance := random.Intn(ctx, 100)
			if dressChance > 30 {
				item, err = getRandomDress(ctx)
				if err != nil {
					err = fmt.Errorf(outfitError, err)
					return []Item{}, err
				}
				items = append(items, item)
			} else {
				item, err = getRandomRobe(ctx)
				if err != nil {
					err = fmt.Errorf(outfitError, err)
					return []Item{}, err
				}
				items = append(items, item)
			}
		} else {
			item, err = getRandomRobe(ctx)
			if err != nil {
				err = fmt.Errorf(outfitError, err)
				return []Item{}, err
			}
			items = append(items, item)
		}
	} else {
		item, err = getRandomTop(ctx)
		if err != nil {
			err = fmt.Errorf(outfitError, err)
			return []Item{}, err
		}
		items = append(items, item)
		item, err = getRandomBottom(ctx)
		if err != nil {
			err = fmt.Errorf(outfitError, err)
			return []Item{}, err
		}
		items = append(items, item)
	}

	if temperature < 5 {
		item, err = getRandomHandwear(ctx)
		if err != nil {
			err = fmt.Errorf(outfitError, err)
			return []Item{}, err
		}
		items = append(items, item)
		item, err = getRandomOverwear(ctx)
		if err != nil {
			err = fmt.Errorf(outfitError, err)
			return []Item{}, err
		}
		items = append(items, item)
		item, err = getRandomBoots(ctx)
		if err != nil {
			err = fmt.Errorf(outfitError, err)
			return []Item{}, err
		}
		items = append(items, item)
	} else {
		footwearChance := random.Intn(ctx, 100)
		if footwearChance > 50 {
			item, err = getRandomShoes(ctx)
			if err != nil {
				err = fmt.Errorf(outfitError, err)
				return []Item{}, err
			}
			items = append(items, item)
		} else {
			item, err = getRandomBoots(ctx)
			if err != nil {
				err = fmt.Errorf(outfitError, err)
				return []Item{}, err
			}
			items = append(items, item)
		}
	}

	hatChance := random.Intn(ctx, 100)
	if hatChance > 60 {
		item, err = getRandomHat(ctx)
		if err != nil {
			err = fmt.Errorf(outfitError, err)
			return []Item{}, err
		}
		items = append(items, item)
	}

	waistChance := random.Intn(ctx, 100)
	if waistChance > 20 {
		item, err = getRandomWaist(ctx)
		if err != nil {
			err = fmt.Errorf(outfitError, err)
			return []Item{}, err
		}
		items = append(items, item)
	}

	return items, nil
}

func getItemFromTemplate(ctx context.Context, template ItemTemplate) (Item, error) {
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

	modifier, err := random.StringFromThresholdMap(ctx, weights)
	if err != nil {
		err = fmt.Errorf(itemError, err)
		return Item{}, err
	}

	if modifier == "prefix" {
		prefix, err := random.String(ctx, template.PrefixModifiers)
		if err != nil {
			err = fmt.Errorf(itemError, err)
			return Item{}, err
		}
		item.PrefixModifier = prefix
	} else if modifier == "suffix" {
		suffix, err := random.String(ctx, template.SuffixModifiers)
		if err != nil {
			err = fmt.Errorf(itemError, err)
			return Item{}, err
		}
		item.SuffixModifier = suffix
	}

	return item, nil
}
