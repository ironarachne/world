package clothing

func getWaist() []Item {
	return []Item{
		Item{
			Name:           "belt",
			Type:           "waist",
			MaterialType:   "hide",
			IdealHeatLevel: "any",
			Layer:          3,
		},
		Item{
			Name:           "sash",
			Type:           "waist",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          3,
		},
	}
}
