package race

import "github.com/ironarachne/world/pkg/size"

func getHumans() []Race {
	heightBase := 60
	weightBase := 100
	sizeCategory := size.GetCategoryByName("medium")

	femaleHeightModifier := 0
	femaleWeightModifier := 0
	heightRange := 12
	maleHeightModifier := 6
	maleWeightModifier := 50
	weightRange := 100

	races := []Race{
		{
			Name:       "human",
			PluralName: "humans",
			Adjective:  "human",
			AgeCategories: []AgeCategory{
				{
					Name:        "adult",
					MinAge:      26,
					MaxAge:      69,
					Commonality: 12,
				},
				{
					Name:        "elderly",
					MinAge:      70,
					MaxAge:      110,
					Commonality: 1,
				},
				{
					Name:        "young adult",
					MinAge:      20,
					MaxAge:      25,
					Commonality: 2,
				},
				{
					Name:        "teenager",
					MinAge:      13,
					MaxAge:      19,
					Commonality: 1,
				},
				{
					Name:        "child",
					MinAge:      2,
					MaxAge:      12,
					Commonality: 1,
				},
				{
					Name:        "infant",
					MinAge:      0,
					MaxAge:      1,
					Commonality: 1,
				},
			},
			Appearance: Appearance{
				MaxFemaleHeight: heightBase + femaleHeightModifier + heightRange,
				MinFemaleHeight: heightBase + femaleHeightModifier,
				MaxMaleHeight:   heightBase + maleHeightModifier + heightRange,
				MinMaleHeight:   heightBase + maleHeightModifier,
				MaxFemaleWeight: weightBase + femaleWeightModifier + weightRange,
				MinFemaleWeight: weightBase + femaleWeightModifier,
				MaxMaleWeight:   weightBase + maleWeightModifier + weightRange,
				MinMaleWeight:   weightBase + maleWeightModifier,
				FaceShapes: []string{
					"broad",
					"chiseled",
					"oval",
					"round",
					"square",
				},
				EarShapes: []string{
					"rounded",
					"slightly pointed",
				},
				EyeShapes: []string{
					"almond",
					"narrow",
					"oval",
					"round",
					"slanted",
				},
				NoseShapes: []string{
					"acquiline",
					"long",
					"narrow",
					"pointed",
					"thin",
					"wide",
				},
				MouthShapes: []string{
					"full",
					"narrow",
					"thin",
					"wide",
				},
				FacialHairStyles: []string{
					"full beard",
					"short goatee",
					"long goatee",
					"handlebar mustache",
					"long beard",
					"short beard",
					"mutton chops",
					"clean shaven",
					"trim mustache",
					"narrow mustache",
				},
				HairColors: []string{
					"black",
					"auburn",
					"red",
					"grey",
					"brown",
					"dark brown",
					"light brown",
				},
				HairConsistencies: []string{
					"thick",
					"thin",
				},
				FemaleHairStyles: []string{
					"shoulder-length",
					"short-cropped",
					"bald",
					"bowl-cut",
					"ponytail",
					"topknot",
					"braided",
					"back-length",
					"mohawk",
				},
				MaleHairStyles: []string{
					"shoulder-length",
					"short-cropped",
					"bald",
					"bowl-cut",
					"ponytail",
					"topknot",
					"braided",
					"back-length",
					"mohawk",
				},
				SkinColors: []string{
					"light",
					"tan",
					"olive",
					"bronze",
					"brown",
					"dark",
				},
				EyeColors: []string{
					"amber",
					"blue",
					"brown",
					"gold",
					"green",
					"hazel",
					"grey",
				},
				UniqueTraits: []string{},
			},
			Commonality:  10,
			SizeCategory: sizeCategory,
		},
	}

	return races
}
