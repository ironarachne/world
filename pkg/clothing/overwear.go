package clothing

func getOverwear() []Item {
	return []Item{
		Item{
			Name:           "short coat",
			Type:           "overwear",
			MaterialType:   "fabric",
			IdealHeatLevel: "cold",
			Layer:          2,
		},
		Item{
			Name:           "cloak",
			Type:           "overwear",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          3,
		},
		Item{
			Name:           "long coat",
			Type:           "overwear",
			MaterialType:   "fabric",
			IdealHeatLevel: "cold",
			Layer:          2,
		},
		Item{
			Name:           "hooded cloak",
			Type:           "overwear",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          3,
		},
	}
}
