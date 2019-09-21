package race

import (
	"github.com/ironarachne/world/pkg/age"
	"github.com/ironarachne/world/pkg/dice"
	"github.com/ironarachne/world/pkg/size"
	"github.com/ironarachne/world/pkg/species"
	"github.com/ironarachne/world/pkg/trait"
)

func getGnomes() []species.Species {
	common := getGnomeCommonTraitTemplates()
	possible := getGnomePossibleTraitTemplates()
	ageCategories := getGnomeAgeCategories()

	races := []species.Species{
		{
			Name:           "gnome",
			PluralName:     "gnomes",
			Adjective:      "gnome",
			CommonTraits:   common,
			PossibleTraits: possible,
			AgeCategories:  ageCategories,
			Commonality:    10,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Tags: []string{
				"gnome",
			},
		},
	}

	return races
}

func getGnomeAgeCategories() []age.Category {
	heightDice := dice.Dice{Number: 2, Sides: 4}
	weightDice := dice.Dice{Number: 2, Sides: 4}
	adultSizeCategory := size.GetCategoryByName("small")
	childSizeCategory := size.GetCategoryByName("small")
	infantSizeCategory := size.GetCategoryByName("tiny")

	categories := []age.Category{
		{
			Name:             "adult",
			MinAge:           36,
			MaxAge:           99,
			MaleHeightBase:   36,
			FemaleHeightBase: 34,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   35,
			FemaleWeightBase: 30,
			WeightModifier:   1,
			WeightRangeDice:  weightDice,
			SizeCategory:     adultSizeCategory,
			Commonality:      150,
		},
		{
			Name:             "elderly",
			MinAge:           100,
			MaxAge:           200,
			MaleHeightBase:   36,
			FemaleHeightBase: 34,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   30,
			FemaleWeightBase: 25,
			WeightModifier:   1,
			WeightRangeDice:  weightDice,
			SizeCategory:     adultSizeCategory,
			Commonality:      50,
		},
		{
			Name:             "young adult",
			MinAge:           20,
			MaxAge:           35,
			MaleHeightBase:   36,
			FemaleHeightBase: 34,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   35,
			FemaleWeightBase: 30,
			WeightModifier:   1,
			WeightRangeDice:  weightDice,
			SizeCategory:     adultSizeCategory,
			Commonality:      100,
		},
		{
			Name:             "teenager",
			MinAge:           13,
			MaxAge:           19,
			MaleHeightBase:   30,
			FemaleHeightBase: 28,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   30,
			FemaleWeightBase: 25,
			WeightModifier:   1,
			WeightRangeDice:  weightDice,
			SizeCategory:     adultSizeCategory,
			Commonality:      60,
		},
		{
			Name:             "child",
			MinAge:           2,
			MaxAge:           12,
			MaleHeightBase:   22,
			FemaleHeightBase: 20,
			HeightRangeDice:  heightDice,
			MaleWeightBase:   20,
			FemaleWeightBase: 15,
			WeightModifier:   1,
			WeightRangeDice:  weightDice,
			SizeCategory:     childSizeCategory,
			Commonality:      10,
		},
		{
			Name:             "infant",
			MinAge:           0,
			MaxAge:           1,
			MaleHeightBase:   15,
			FemaleHeightBase: 15,
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

func getGnomeCommonTraitTemplates() []trait.Template {
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

func getGnomePossibleTraitTemplates() []trait.Template {
	templates := []trait.Template{
		{
			Name: "nose shape",
			PossibleValues: []string{
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

