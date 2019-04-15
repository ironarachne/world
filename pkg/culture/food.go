package culture

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/random"
)

// FoodStyle is a cultural food style
type FoodStyle struct {
	CommonBases       []string
	CommonSpices      []string
	CookingTechniques []string
	CommonDishes      []string
	EatingTraits      []string
	Breads            []string
}

func (culture Culture) generateFoodStyle() FoodStyle {
	foodStyle := FoodStyle{}

	for _, r := range culture.HomeClimate.Resources {
		if r.Type == "meat" {
			foodStyle.CommonBases = append(foodStyle.CommonBases, r.Name)
		} else if r.Type == "spice" {
			foodStyle.CommonSpices = append(foodStyle.CommonSpices, r.Name)
		} else if r.Type == "fruit" {
			foodStyle.CommonBases = append(foodStyle.CommonBases, r.Name)
		} else if r.Type == "vegetable" {
			foodStyle.CommonBases = append(foodStyle.CommonBases, r.Name)
		}
	}

	if !inSlice("salt", foodStyle.CommonSpices) {
		foodStyle.CommonSpices = append(foodStyle.CommonSpices, "salt")
	}

	foodStyle.CookingTechniques = randomCookingTechniques()

	foodStyle.CommonDishes = foodStyle.randomDishes()
	foodStyle.Breads = culture.randomBreads()
	foodStyle.EatingTraits = randomEatingTraits()

	return foodStyle
}

func randomCookingTechniques() []string {
	var techniques []string
	var technique string

	potentialTechniques := []string{
		"baked",
		"basted",
		"broiled",
		"curried",
		"dried",
		"fried",
		"raw",
		"roasted",
		"slow-roasted",
		"stewed",
	}

	for i := 0; i < 3; i++ {
		technique = random.String(potentialTechniques)
		if !inSlice(technique, techniques) {
			techniques = append(techniques, technique)
		}
	}

	return techniques
}

func (culture Culture) randomBread() string {
	bread := ""
	breadTypes := []string{
		"flat",
		"round",
		"thin",
		"brick-like",
	}

	flavors := []string{
		"bitter",
		"grainy",
		"hearty",
		"nutty",
		"savory",
		"sweet",
	}
	grains := climate.ListResourcesOfType("grain", culture.HomeClimate.Resources)

	if len(grains) > 0 {
		grain := grains[rand.Intn(len(grains))]

		bread = random.String(flavors) + " " + random.String(breadTypes) + " " + grain.Name + " bread"
	}

	return bread
}

func (culture Culture) randomBreads() []string {
	var bread string
	var breads []string

	grains := climate.ListResourcesOfType("grain", culture.HomeClimate.Resources)
	if len(grains) > 0 {
		numberOfBreads := rand.Intn(3) + 1
		for i := 0; i < numberOfBreads; i++ {
			bread = culture.randomBread()
			if !inSlice(bread, breads) {
				breads = append(breads, bread)
			}
		}
	}

	return breads
}

func (style FoodStyle) randomDishes() []string {
	var dishes []string
	var dish string

	flavors := []string{
		"aromatic",
		"bitter",
		"pungent",
		"salty",
		"saucy",
		"savory",
		"sour",
		"spicy",
		"sweet",
	}

	for i := 0; i < 5; i++ {
		dish = random.String(flavors) + " " + random.String(style.CookingTechniques) + " " + random.String(style.CommonBases) + " with " + random.String(style.CommonSpices)

		if !inSlice(dish, dishes) {
			dishes = append(dishes, dish)
		}
	}

	return dishes
}

func randomEatingTraits() []string {
	var traits []string
	var trait string
	var typesOfTraits []string
	var t string

	utensils := []string{
		"chopsticks",
		"fork",
		"hands",
		"knife and fork",
		"knife",
		"spoon",
	}

	spices := []string{
		"heavy",
		"moderate",
		"minimal",
	}

	heat := []string{
		"hot",
		"warm",
		"cold",
		"chilled",
	}

	customs := []string{
		"eat communal meals",
		"eat alone",
		"belch after eating",
		"thank the cook after eating",
	}

	potentialTraits := []string{
		"utensils",
		"spices",
		"heat",
		"customs",
	}

	for i := 0; i < 2; i++ {
		t = random.String(potentialTraits)
		if !inSlice(t, typesOfTraits) {
			typesOfTraits = append(typesOfTraits, t)
			if t == "utensils" {
				trait = "eat with " + random.String(utensils)
			} else if t == "spices" {
				trait = random.String(spices) + " spice"
			} else if t == "heat" {
				trait = "food is usually " + random.String(heat)
			} else if t == "customs" {
				trait = random.String(customs)
			}
			traits = append(traits, trait)
		}
	}

	return traits
}
