package food

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/slices"
)

// Style is a cultural food style
type Style struct {
	Breads             []string
	CommonBases        []string
	CommonDessertBases []string
	CommonDesserts     []string
	CommonMainDishes   []string
	CommonSpices       []string
	CookingTechniques  []string
	EatingTraits       []string
	DessertTypes       []string
	Noodles            []string
}

// GenerateStyle procedurally generates a style of food
func GenerateStyle(originClimate climate.Climate) Style {
	chanceForGoldFlakes := 0
	style := Style{}

	for _, r := range originClimate.Resources {
		if r.Type == "meat" {
			style.CommonBases = append(style.CommonBases, r.Name)
		} else if r.Type == "spice" || r.Type == "herb" {
			if r.Name == "gold" {
				chanceForGoldFlakes = rand.Intn(100)
				if chanceForGoldFlakes > 89 {
					style.CommonSpices = append(style.CommonSpices, "gold flakes")
				}
			} else {
				style.CommonSpices = append(style.CommonSpices, r.Name)
			}
		} else if r.Type == "egg" {
			style.CommonBases = append(style.CommonBases, r.Name)
		} else if r.Type == "milk" {
			if !slices.StringIn("pudding", style.DessertTypes) {
				style.DessertTypes = append(style.DessertTypes, "pudding")
			}
		} else if r.Type == "fruit" {
			style.CommonDessertBases = append(style.CommonDessertBases, r.Name)
			if !slices.StringIn("fruit", style.DessertTypes) {
				style.DessertTypes = append(style.DessertTypes, "fruit")
			}
		} else if r.Type == "vegetable" {
			style.CommonBases = append(style.CommonBases, r.Name)
		} else if r.Type == "flour" {
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

	style.CookingTechniques = randomTechniques(3)

	style.CommonDesserts = style.randomDesserts()
	style.CommonMainDishes = style.randomMainDishes()
	style.Breads = randomBreads(originClimate)
	style.EatingTraits = randomEatingTraits()

	return style
}

// Random generates a completely random style of food
func Random() Style {
	originClimate := climate.Generate()
	return GenerateStyle(originClimate)
}
