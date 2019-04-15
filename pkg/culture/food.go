package culture

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/random"
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
		technique = random.Item(potentialTechniques)
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

		bread = random.Item(flavors) + " " + random.Item(breadTypes) + " " + grain.Name + " bread"
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
		dish = random.Item(flavors) + " " + random.Item(style.CookingTechniques) + " " + random.Item(style.CommonBases) + " with " + random.Item(style.CommonSpices)

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
		t = random.Item(potentialTraits)
		if !inSlice(t, typesOfTraits) {
			typesOfTraits = append(typesOfTraits, t)
			if t == "utensils" {
				trait = "eat with " + random.Item(utensils)
			} else if t == "spices" {
				trait = random.Item(spices) + " spice"
			} else if t == "heat" {
				trait = "food is usually " + random.Item(heat)
			} else if t == "customs" {
				trait = random.Item(customs)
			}
			traits = append(traits, trait)
		}
	}

	return traits
}
