package race

import "github.com/ironarachne/world/pkg/size"

func getDwarves() []Race {
	heightBase := 48
	weightBase := 50
	sizeCategory := size.GetCategoryByName("medium")

	femaleHeightModifier := 3
	femaleWeightModifier := 50
	heightRange := 6
	maleHeightModifier := 6
	maleWeightModifier := 150
	weightRange := 100

	return []Race{
		Race{
			Name:       "dwarf",
			PluralName: "dwarves",
			Adjective:  "dwarven",
			AgeCategories: []AgeCategory{
				AgeCategory{
					Name:        "adult",
					MinAge:      76,
					MaxAge:      269,
					Commonality: 12,
				},
				AgeCategory{
					Name:        "elderly",
					MinAge:      270,
					MaxAge:      500,
					Commonality: 1,
				},
				AgeCategory{
					Name:        "young adult",
					MinAge:      40,
					MaxAge:      75,
					Commonality: 2,
				},
				AgeCategory{
					Name:        "teenager",
					MinAge:      23,
					MaxAge:      39,
					Commonality: 1,
				},
				AgeCategory{
					Name:        "child",
					MinAge:      2,
					MaxAge:      22,
					Commonality: 1,
				},
				AgeCategory{
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
					"square",
					"rectangular",
				},
				EarShapes: []string{
					"slightly pointed",
					"rounded",
					"slightly tapered",
				},
				EyeShapes: []string{
					"broad",
					"oval",
					"narrow",
				},
				NoseShapes: []string{
					"wide",
					"large",
					"hooked",
					"strong",
				},
				MouthShapes: []string{
					"wide",
					"thin",
				},
				FacialHairStyles: []string{
					"braided beard",
					"bristly beard",
					"enormous beard",
					"long beard",
					"robust beard",
					"wide beard",
				},
				HairColors: []string{
					"auburn",
					"black",
					"dark brown",
					"dusty brown",
					"grey",
					"orange",
					"red",
					"white",
				},
				HairConsistencies: []string{
					"thick",
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
					"bronzed",
					"tan",
					"ruddy",
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
					"usually have beards",
					"are stocky but well-muscled",
				},
			},
			Commonality:  2,
			SizeCategory: sizeCategory,
		},
	}
}
