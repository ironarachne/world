package clothing

func getFootwear() []Item {
	return []Item{
		Item{
			Name:           "sandals",
			Type:           "footwear",
			MaterialType:   "hide",
			IdealHeatLevel: "warm",
			Layer:          1,
		},
		Item{
			Name:           "knee-high boots",
			Type:           "footwear",
			MaterialType:   "hide",
			IdealHeatLevel: "cold",
			Layer:          1,
		},
		Item{
			Name:           "short boots",
			Type:           "footwear",
			MaterialType:   "hide",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "strapped boots",
			Type:           "footwear",
			MaterialType:   "hide",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "turnshoes",
			Type:           "footwear",
			MaterialType:   "hide",
			IdealHeatLevel: "any",
			Layer:          1,
		},
	}
}
