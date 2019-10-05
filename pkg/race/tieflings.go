package race

import (
	"github.com/ironarachne/world/pkg/age"
	"github.com/ironarachne/world/pkg/dice"
	"github.com/ironarachne/world/pkg/size"
	"github.com/ironarachne/world/pkg/species"
	"github.com/ironarachne/world/pkg/trait"
)

func getTiefling() []species.Species {
	common := getTieflingCommonTraitTemplates()
	possible := getTieflingPossibleTraitTemplates()
	ageCategories := getTieflingAgeCategories()

	return []species.Species {
		{
			Name:           "tiefling",
			PluralName:     "tieflings",
			Adjective:      "tiefling",
			CommonTraits:   common,
			PossibleTraits: possible,
			AgeCategories:  ageCategories,
			Commonality:    10,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Tags: []string{
				"tiefling",
			},
		},
	}
}


func getTieflingAgeCategories() []age.Category {
	heightDice := dice.Dice{Number: 2, Sides: 8}
	weightDice := dice.Dice{Number: 2, Sides: 10}
	adultSizeCategory := size.GetCategoryByName("medium")
	childSizeCategory := size.GetCategoryByName("small")
	infantSizeCategory := size.GetCategoryByName("tiny")

	categories := []age.Category{
		{
			Name:             "elderly",
			MinAge:           70,
			MaxAge:           110,
			MaleHeightBase:   62,
			FemaleHeightBase: 57,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   140,
			FemaleWeightBase: 105,
			WeightModifier:   5,
			WeightRangeDice:  weightDice,
			SizeCategory:     adultSizeCategory,
			Commonality:      10,
		},
		{
			Name:             "adult",
			MinAge:           26,
			MaxAge:           69,
			MaleHeightBase:   62,
			FemaleHeightBase: 57,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   140,
			FemaleWeightBase: 105,
			WeightModifier:   5,
			WeightRangeDice:  weightDice,
			SizeCategory:     adultSizeCategory,
			Commonality:      30,
		},
		{
			Name:             "young adult",
			MinAge:           20,
			MaxAge:           25,
			MaleHeightBase:   62,
			FemaleHeightBase: 57,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   140,
			FemaleWeightBase: 105,
			WeightModifier:   5,
			WeightRangeDice:  weightDice,
			SizeCategory:     adultSizeCategory,
			Commonality:      150,
		},
		{
			Name:             "teenager",
			MinAge:           13,
			MaxAge:           19,
			MaleHeightBase:   61,
			FemaleHeightBase: 59,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   92,
			FemaleWeightBase: 89,
			WeightModifier:   5,
			WeightRangeDice:  weightDice,
			SizeCategory:     adultSizeCategory,
			Commonality:      100,
		},
		{
			Name:             "child",
			MinAge:           2,
			MaxAge:           12,
			MaleHeightBase:   42,
			FemaleHeightBase: 42,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   36,
			FemaleWeightBase: 36,
			WeightModifier:   1,
			WeightRangeDice:  weightDice,
			SizeCategory:     childSizeCategory,
			Commonality:      10,
		},
		{
			Name:             "infant",
			MinAge:           0,
			MaxAge:           1,
			MaleHeightBase:   18,
			FemaleHeightBase: 18,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   5,
			FemaleWeightBase: 5,
			WeightModifier:   1,
			WeightRangeDice:  weightDice,
			SizeCategory:     infantSizeCategory,
			Commonality:      1,
		},
	}

	return categories
}

func getTieflingCommonTraitTemplates() []trait.Template {
	templates := []trait.Template{
		{
			Name: "eye color",
			PossibleValues: []string{
				"black",
				"gold",
				"red",
				"silver",
				"white",
			},
			PossibleDescriptors: []string{
				"{{.Value}} eyes",
			},
			Tags: []string{
				"appearance",
				"physical",
				"eyes",
			},
		},
		{
			Name: "hair color",
			PossibleValues: []string{
				"dark red",
				"brown",
				"dark brown",
				"black",
				"dark blue",
				"purple",
			},
			PossibleDescriptors: []string{
				"{{.Value}} hair",
			},
			Tags: []string{
				"appearance",
				"physical",
				"hair",
			},
		},
		{
			Name: "skin color",
			PossibleValues: []string{
				"light",
				"tan",
				"olive",
				"bronze",
				"brown",
				"dark",
				"maroon",
				"red",
			},
			PossibleDescriptors: []string{
				"{{.Value}} skin",
			},
			Tags: []string{
				"appearance",
				"physical",
				"skin",
			},
		},
		{
			Name: "horns",
			PossibleValues: []string{
				"curling",
				"spiral upward",
				"spiral downard",
				"long straight",
				"short straight",
			},
			PossibleDescriptors: []string{
				"{{.Value}} horns",
			},
			Tags: []string{
				"appearance",
				"physical",
				"horns",
			},
		},
	}

	return templates
}

func getTieflingPossibleTraitTemplates() []trait.Template {
	templates := []trait.Template{
		{
			Name: "nose shape",
			PossibleValues: []string{
				"aquiline",
				"broad",
				"flat",
				"long",
				"narrow",
				"pointed",
				"upturned",
			},
			PossibleDescriptors: []string{
				"{{.Value}} nose",
			},
			Tags: []string{
				"appearance",
				"physical",
				"nose",
			},
		},
	}

	return templates
}