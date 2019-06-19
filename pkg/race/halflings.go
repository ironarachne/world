package race

func getHalflings() []Race {
	baseSize := getSizeByName("small")

	femaleHeightModifier := 0
	femaleWeightModifier := 20
	heightRange := 6
	maleHeightModifier := 6
	maleWeightModifier := 50
	weightRange := 100

	return []Race{
		Race{
			Name:       "halfling",
			PluralName: "halflings",
			Adjective:  "halfling",
			AgeCategories: []AgeCategory{
				AgeCategory{
					Name:        "adult",
					MinAge:      36,
					MaxAge:      99,
					Commonality: 12,
				},
				AgeCategory{
					Name:        "elderly",
					MinAge:      100,
					MaxAge:      200,
					Commonality: 1,
				},
				AgeCategory{
					Name:        "young adult",
					MinAge:      20,
					MaxAge:      35,
					Commonality: 2,
				},
				AgeCategory{
					Name:        "teenager",
					MinAge:      13,
					MaxAge:      19,
					Commonality: 1,
				},
				AgeCategory{
					Name:        "child",
					MinAge:      2,
					MaxAge:      12,
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
				UniqueTraits: []string{},
			},
			Commonality: 2,
			Size:        baseSize,
		},
	}
}
