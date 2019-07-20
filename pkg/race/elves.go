package race

import "github.com/ironarachne/world/pkg/size"

func getElves() []Race {
	heightBase := 60
	weightBase := 100
	sizeCategory := size.GetCategoryByName("medium")

	femaleHeightModifier := -6
	femaleWeightModifier := -50
	heightRange := 6
	maleHeightModifier := 0
	maleWeightModifier := 0
	weightRange := 50

	return []Race{
		{
			Name:       "elf",
			PluralName: "elves",
			Adjective:  "elven",
			AgeCategories: []AgeCategory{
				{
					Name:        "adult",
					MinAge:      76,
					MaxAge:      269,
					Commonality: 12,
				},
				{
					Name:        "elderly",
					MinAge:      270,
					MaxAge:      500,
					Commonality: 1,
				},
				{
					Name:        "young adult",
					MinAge:      40,
					MaxAge:      75,
					Commonality: 2,
				},
				{
					Name:        "teenager",
					MinAge:      23,
					MaxAge:      39,
					Commonality: 1,
				},
				{
					Name:        "child",
					MinAge:      2,
					MaxAge:      22,
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
					"chiseled",
					"oval",
					"triangular",
				},
				EarShapes: []string{
					"slightly pointed",
					"steeply pointed",
					"long and pointed",
				},
				EyeShapes: []string{
					"almond",
					"narrow",
					"slanted",
				},
				NoseShapes: []string{
					"long",
					"pointed",
					"thin",
					"slender",
				},
				MouthShapes: []string{
					"narrow",
					"thin",
				},
				FacialHairStyles: []string{
					"short goatee",
					"long goatee",
					"clean shaven",
				},
				HairColors: []string{
					"auburn",
					"grey",
					"white",
					"silver",
					"light blonde",
				},
				HairConsistencies: []string{
					"thick",
					"thin",
				},
				FemaleHairStyles: []string{
					"shoulder-length",
					"ponytail",
					"topknot",
					"braided",
					"back-length",
				},
				MaleHairStyles: []string{
					"shoulder-length",
					"short-cropped",
					"ponytail",
					"topknot",
					"braided",
					"back-length",
				},
				SkinColors: []string{
					"light",
					"tan",
					"pale",
				},
				EyeColors: []string{
					"amber",
					"blue",
					"brown",
					"gold",
					"green",
					"hazel",
					"grey",
					"purple",
					"silver",
				},
				UniqueTraits: []string{
					"are thin, almost ethereal",
				},
			},
			Commonality:  2,
			SizeCategory: sizeCategory,
		},
	}
}
