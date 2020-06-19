/*
Package food provides methods and tools for generating fantasy
food styles.
*/
package food

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/geography"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/slices"
)

const foodStyleError = "failed to generate food style: %w"

// Style is a cultural food style
type Style struct {
	Breads             []string `json:"breads"`
	CommonBases        []string `json:"common_bases"`
	CommonDessertBases []string `json:"common_dessert_bases"`
	CommonDesserts     []string `json:"common_desserts"`
	CommonMainDishes   []string `json:"common_main_dishes"`
	CommonSpices       []string `json:"common_spices"`
	CookingTechniques  []string `json:"cooking_techniques"`
	EatingTraits       []string `json:"eating_traits"`
	DessertTypes       []string `json:"dessert_types"`
	Noodles            []string `json:"noodles"`
}

// GenerateStyle procedurally generates a style of food
func GenerateStyle(ctx context.Context, resources []resource.Resource) (Style, error) {
	chanceForGoldFlakes := 0
	style := Style{}
	style.CommonDessertBases = append(style.CommonDessertBases, "sweetened")

	for _, r := range resources {
		if r.HasTag("meat") {
			style.CommonBases = append(style.CommonBases, r.Name)
		} else if r.HasTag("spice") || r.HasTag("herb") {
			if r.Name == "gold" {
				chanceForGoldFlakes = random.Intn(ctx, 100)
				if chanceForGoldFlakes > 89 {
					style.CommonSpices = append(style.CommonSpices, "gold flakes")
				}
			} else {
				style.CommonSpices = append(style.CommonSpices, r.Name)
			}
		} else if r.HasTag("egg") {
			style.CommonBases = append(style.CommonBases, r.Name)
			style.CommonDessertBases = append(style.CommonDessertBases, r.Name)
			if !slices.StringIn("custard", style.DessertTypes) {
				style.DessertTypes = append(style.DessertTypes, "custard")
			}
		} else if r.HasTag("milk") {
			if !slices.StringIn("pudding", style.DessertTypes) {
				style.DessertTypes = append(style.DessertTypes, "pudding")
			}
			if !slices.StringIn("yogurt", style.DessertTypes) {
				style.DessertTypes = append(style.DessertTypes, "yogurt")
			}
		} else if r.HasTag("fruit") {
			style.CommonDessertBases = append(style.CommonDessertBases, r.Name)
			if !slices.StringIn("fruit", style.DessertTypes) {
				style.DessertTypes = append(style.DessertTypes, "fruit")
			}
		} else if r.HasTag("vegetable") || r.HasTag("squash") {
			style.CommonBases = append(style.CommonBases, r.Name)
		} else if r.HasTag("grain") {
			style.Noodles = append(style.Noodles, getNoodles(r.Name)...)
			if !slices.StringIn("pie", style.DessertTypes) {
				style.DessertTypes = append(style.DessertTypes, "pie")
			}
			if !slices.StringIn("crisp", style.DessertTypes) {
				style.DessertTypes = append(style.DessertTypes, "crisp")
			}
			if !slices.StringIn("tart", style.DessertTypes) {
				style.DessertTypes = append(style.DessertTypes, "tart")
			}
			if !slices.StringIn("cake", style.DessertTypes) {
				style.DessertTypes = append(style.DessertTypes, "cake")
			}
		}
	}

	if !slices.StringIn("salt", style.CommonSpices) {
		style.CommonSpices = append(style.CommonSpices, "salt")
	}

	techniques, err := randomTechniques(ctx, 3)
	if err != nil {
		err = fmt.Errorf(foodStyleError, err)
		return Style{}, err
	}
	style.CookingTechniques = techniques

	desserts, err := style.randomDesserts(ctx)
	if err != nil {
		err = fmt.Errorf(foodStyleError, err)
		return Style{}, err
	}
	style.CommonDesserts = desserts
	mainDishes, err := style.randomMainDishes(ctx)
	if err != nil {
		err = fmt.Errorf(foodStyleError, err)
		return Style{}, err
	}
	style.CommonMainDishes = mainDishes
	breads, err := randomBreads(ctx, resources)
	if err != nil {
		err = fmt.Errorf(foodStyleError, err)
		return Style{}, err
	}
	style.Breads = breads
	traits, err := randomEatingTraits(ctx)
	if err != nil {
		err = fmt.Errorf(foodStyleError, err)
		return Style{}, err
	}
	style.EatingTraits = traits

	return style, nil
}

// Random generates a completely random style of food
func Random(ctx context.Context) (Style, error) {
	area, err := geography.Generate(ctx)
	if err != nil {
		err = fmt.Errorf("failed to generate random food style: %w", err)
		return Style{}, err
	}

	resources := area.GetResources()

	style, err := GenerateStyle(ctx, resources)
	if err != nil {
		err = fmt.Errorf("failed to generate random food style: %w", err)
		return Style{}, err
	}

	return style, nil
}
