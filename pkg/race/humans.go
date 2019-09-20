package race

import (
	"github.com/ironarachne/world/pkg/age"
	"github.com/ironarachne/world/pkg/dice"
	"github.com/ironarachne/world/pkg/size"
	"github.com/ironarachne/world/pkg/species"
	"github.com/ironarachne/world/pkg/trait"
)

func getHumans() []species.Species {
	common := getHumanCommonTraitTemplates()
	possible := getHumanPossibleTraitTemplates()
	ageCategories := getHumanAgeCategories()

	races := []species.Species{
		{
			Name:           "human",
			PluralName:     "humans",
			Adjective:      "human",
			CommonTraits:   common,
			PossibleTraits: possible,
			AgeCategories:  ageCategories,
			Commonality:    10,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Tags: []string{
				"human",
			},
		},
	}

	return races
}

func getHumanAgeCategories() []age.Category {
	heightDice := dice.Dice{Number: 1, Sides: 4}
	weightDice := dice.Dice{Number: 1, Sides: 4}
	adultSizeCategory := size.GetCategoryByName("medium")
	childSizeCategory := size.GetCategoryByName("small")
	infantSizeCategory := size.GetCategoryByName("tiny")

	categories := []age.Category{
		{
			Name:                 "adult",
			MinAge:               26,
			MaxAge:               69,
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
			MinAge:               70,
			MaxAge:               110,
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
			MaxAge:               25,
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

func getHumanCommonTraitTemplates() []trait.Template {
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

func getHumanPossibleTraitTemplates() []trait.Template {
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
