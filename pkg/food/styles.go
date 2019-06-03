package food

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/slices"
)

// Style is a cultural food style
type Style struct {
	Breads            []string
	CommonBases       []string
	CommonDishes      []string
	CommonSpices      []string
	CookingTechniques []string
	EatingTraits      []string
	Noodles           []string
}

// GenerateStyle procedurally generates a style of food
func GenerateStyle(originClimate climate.Climate) Style {
	chanceForGoldFlakes := 0
	style := Style{}

	for _, r := range originClimate.Resources {
		if r.Type == "meat" {
			style.CommonBases = append(style.CommonBases, r.Name)
		} else if r.Type == "spice" {
			if r.Name == "gold" {
				chanceForGoldFlakes = rand.Intn(100)
				if chanceForGoldFlakes > 89 {
					style.CommonSpices = append(style.CommonSpices, "gold flakes")
				}
			} else {
				style.CommonSpices = append(style.CommonSpices, r.Name)
			}
		} else if r.Type == "fruit" {
			style.CommonBases = append(style.CommonBases, r.Name)
		} else if r.Type == "vegetable" {
			style.CommonBases = append(style.CommonBases, r.Name)
		} else if r.Type == "grain" {
			style.Noodles = append(style.Noodles, getNoodles(r.Name)...)
		}
	}

	if !slices.StringIn("salt", style.CommonSpices) {
		style.CommonSpices = append(style.CommonSpices, "salt")
	}

	style.CookingTechniques = randomTechniques(3)

	style.CommonDishes = style.randomDishes()
	style.Breads = randomBreads(originClimate)
	style.EatingTraits = randomEatingTraits()

	return style
}

// Random generates a completely random style of food
func Random() Style {
	originClimate := climate.Generate()
	return GenerateStyle(originClimate)
}
