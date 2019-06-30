package clothing

func getFull() []Item {
	return []Item{
		Item{
			Name:           "robe",
			Type:           "full",
			MaterialType:   "fabric",
			IdealHeatLevel: "cold",
			Layer:          1,
		},
		Item{
			Name:           "dress",
			Type:           "full",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
	}
}