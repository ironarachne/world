package race

func getElves() []Race {
	baseSize := getSizeByName("medium")

	femaleHeightModifier := -6
	femaleWeightModifier := -50
	heightRange := 6
	maleHeightModifier := 0
	maleWeightModifier := 0
	weightRange := 50

	return []Race{
		Race{
			Name:       "elf",
			PluralName: "elves",
			Adjective:  "elven",
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
				MaxFemaleHeight: baseSize.HeightBase + femaleHeightModifier + heightRange,
				MinFemaleHeight: baseSize.HeightBase + femaleHeightModifier,
				MaxMaleHeight:   baseSize.HeightBase + maleHeightModifier + heightRange,
				MinMaleHeight:   baseSize.HeightBase + maleHeightModifier,
				MaxFemaleWeight: baseSize.WeightBase + femaleWeightModifier + weightRange,
				MinFemaleWeight: baseSize.WeightBase + femaleWeightModifier,
				MaxMaleWeight:   baseSize.WeightBase + maleWeightModifier + weightRange,
				MinMaleWeight:   baseSize.WeightBase + maleWeightModifier,
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
				UniqueTraits: []string{},
			},
			Commonality: 2,
			Size:        baseSize,
		},
	}
}
