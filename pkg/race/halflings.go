package race

import "github.com/ironarachne/world/pkg/size"

func getHalflings() []Race {
	heightBase := 48
	weightBase := 50
	sizeCategory := size.GetCategoryByName("small")

	femaleHeightModifier := 0
	femaleWeightModifier := 20
	heightRange := 6
	maleHeightModifier := 6
	maleWeightModifier := 50
	weightRange := 100

	return []Race{
		{
			Name:       "halfling",
			PluralName: "halflings",
			Adjective:  "halfling",
			AgeCategories: []AgeCategory{
				{
					Name:        "adult",
					MinAge:      36,
					MaxAge:      99,
					Commonality: 12,
				},
				{
					Name:        "elderly",
					MinAge:      100,
					MaxAge:      200,
					Commonality: 1,
				},
				{
					Name:        "young adult",
					MinAge:      20,
					MaxAge:      35,
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
					"mutton chops",
					"clean shaven",
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
				UniqueTraits: []string{
					"have large, hairy feet",
				},
			},
			Commonality:  2,
			SizeCategory: sizeCategory,
		},
	}
}
