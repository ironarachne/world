package race

import (
	"github.com/ironarachne/world/pkg/age"
	"github.com/ironarachne/world/pkg/dice"
	"github.com/ironarachne/world/pkg/size"
	"github.com/ironarachne/world/pkg/species"
	"github.com/ironarachne/world/pkg/trait"
)

func getHalflings() []species.Species {
	common := getHalflingCommonTraitTemplates()
	possible := getHalflingPossibleTraitTemplates()
	ageCategories := getHalflingAgeCategories()

	return []species.Species{
		{
			Name:           "halfling",
			PluralName:     "halflings",
			Adjective:      "halfling",
			CommonTraits:   common,
			PossibleTraits: possible,
			AgeCategories:  ageCategories,
			Commonality:    2,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Tags: []string{
				"halfling",
			},
		},
	}
}

func getHalflingAgeCategories() []age.Category {
	heightDice := dice.Dice{Number: 1, Sides: 4}
	weightDice := dice.Dice{Number: 1, Sides: 4}
	adultSizeCategory := size.GetCategoryByName("small")
	childSizeCategory := size.GetCategoryByName("small")
	infantSizeCategory := size.GetCategoryByName("tiny")

	categories := []age.Category{
		{
			Name:                 "adult",
			MinAge:               36,
			MaxAge:               99,
			MaleHeightModifier:   1,
			FemaleHeightModifier: 1,
			HeightRangeDice:      heightDice,
			MaleWeightModifier:   1,
			FemaleWeightModifier: 1,
			WeightRangeDice:      weightDice,
			SizeCategory:         adultSizeCategory,
			Commonality:          12,
		},
		{
			Name:                 "elderly",
			MinAge:               100,
			MaxAge:               200,
			MaleHeightModifier:   1,
			FemaleHeightModifier: 1,
			HeightRangeDice:      heightDice,
			MaleWeightModifier:   1,
			FemaleWeightModifier: 1,
			WeightRangeDice:      weightDice,
			SizeCategory:         adultSizeCategory,
			Commonality:          1,
		},
		{
			Name:                 "young adult",
			MinAge:               20,
			MaxAge:               35,
			MaleHeightModifier:   1,
			FemaleHeightModifier: 1,
			HeightRangeDice:      heightDice,
			MaleWeightModifier:   1,
			FemaleWeightModifier: 1,
			WeightRangeDice:      weightDice,
			SizeCategory:         adultSizeCategory,
			Commonality:          2,
		},
		{
			Name:                 "teenager",
			MinAge:               13,
			MaxAge:               19,
			MaleHeightModifier:   1,
			FemaleHeightModifier: 1,
			HeightRangeDice:      heightDice,
			MaleWeightModifier:   1,
			FemaleWeightModifier: 1,
			WeightRangeDice:      weightDice,
			SizeCategory:         adultSizeCategory,
			Commonality:          1,
		},
		{
			Name:                 "child",
			MinAge:               2,
			MaxAge:               12,
			MaleHeightModifier:   1,
			FemaleHeightModifier: 1,
			HeightRangeDice:      heightDice,
			MaleWeightModifier:   1,
			FemaleWeightModifier: 1,
			WeightRangeDice:      weightDice,
			SizeCategory:         childSizeCategory,
			Commonality:          1,
		},
		{
			Name:                 "infant",
			MinAge:               0,
			MaxAge:               1,
			MaleHeightModifier:   1,
			FemaleHeightModifier: 1,
			HeightRangeDice:      heightDice,
			MaleWeightModifier:   1,
			FemaleWeightModifier: 1,
			WeightRangeDice:      weightDice,
			SizeCategory:         infantSizeCategory,
			Commonality:          1,
		},
	}

	return categories
}

func getHalflingCommonTraitTemplates() []trait.Template {
	templates := []trait.Template{
		{
			Name: "eye color",
			PossibleValues: []string{
				"amber",
				"blue",
				"brown",
				"gold",
				"green",
				"hazel",
				"grey",
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
				"black",
				"auburn",
				"red",
				"grey",
				"brown",
				"dark brown",
				"light brown",
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
	}

	return templates
}

func getHalflingPossibleTraitTemplates() []trait.Template {
	templates := []trait.Template{
		{
			Name: "nose shape",
			PossibleValues: []string{
				"aquiline",
				"broad",
				"bulbous",
				"flat",
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
