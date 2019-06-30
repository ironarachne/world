package clothing

func getTops() []Item {
	return []Item{
		Item{
			Name:           "short tunic",
			Type:           "top",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "knee-length tunic",
			Type:           "top",
			MaterialType:   "fabric",
			IdealHeatLevel: "any",
			Layer:          1,
		},
		Item{
			Name:           "ankle-length tunic",
			Type:           "top",
			MaterialType:   "fabric",
			IdealHeatLevel: "cold",
			Layer:          1,
		},
		Item{
			Name:           "sleeveless shirt",
			Type:           "top",
			MaterialType:   "fabric",
			IdealHeatLevel: "warm",
			Layer:          1,
		},
		Item{
			Name:           "vest",
			Type:           "top",
			MaterialType:   "hide",
			IdealHeatLevel: "any",
			Layer:          2,
		},
	}
}
