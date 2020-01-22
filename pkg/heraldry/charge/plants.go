package charge

func getPlantCharges() []Raster {
	charges := []Raster{
		{
			Identifier: "acorn-slipped-and-leaved",
			Name:       "acorn slipped and leaved",
			Noun:       "acorn",
			NounPlural: "acorns",
			Descriptor: "slipped and leaved",
			SingleOnly: false,
			Tags: []string{
				"acorn",
				"plant",
				"harvest",
				"farmer",
			},
		},
	}

	return charges
}
