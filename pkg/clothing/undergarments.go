package clothing

func getUndergarments() []Item {
	return []Item{
		Item{
			Name:           "loincloth",
			Type:           "undergarment",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          0,
		},
	}
}