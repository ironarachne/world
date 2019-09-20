package race

import (
	"github.com/ironarachne/world/pkg/age"
	"github.com/ironarachne/world/pkg/dice"
	"github.com/ironarachne/world/pkg/size"
	"github.com/ironarachne/world/pkg/species"
	"github.com/ironarachne/world/pkg/trait"
)

func getDwarves() []species.Species {

	common := getDwarfCommonTraitTemplates()
	possible := getDwarfPossibleTraitTemplates()
	ageCategories := getDwarfAgeCategories()

	races := []species.Species{
		{
			Name:           "hill dwarf",
			PluralName:     "hill dwarves",
			Adjective:      "hill dwarven",
			CommonTraits:   common,
			PossibleTraits: possible,
			AgeCategories:  ageCategories,
			Commonality:    2,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Tags: []string{
				"dwarf",
				"hill dwarf",
			},
		},
		{
			Name:           "mountain dwarf",
			PluralName:     "mountain dwarves",
			Adjective:      "mountain dwarven",
			CommonTraits:   common,
			PossibleTraits: possible,
			AgeCategories:  ageCategories,
			Commonality:    2,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Tags: []string{
				"dwarf",
				"mountain dwarf",
			},
		},
	}

	return races
}

func getDwarfAgeCategories() []age.Category {
	heightDice := dice.Dice{Number: 1, Sides: 4}
	weightDice := dice.Dice{Number: 1, Sides: 4}
	adultSizeCategory := size.GetCategoryByName("medium")
	childSizeCategory := size.GetCategoryByName("small")
	infantSizeCategory := size.GetCategoryByName("tiny")

	categories := []age.Category{
		{
			Name:                 "adult",
			MinAge:               76,
			MaxAge:               269,
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
			MinAge:               270,
			MaxAge:               500,
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
			MinAge:               40,
			MaxAge:               75,
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
			MinAge:               23,
			MaxAge:               39,
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
			MaxAge:               22,
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

func getDwarfCommonTraitTemplates() []trait.Template {
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

func getDwarfPossibleTraitTemplates() []trait.Template {
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
