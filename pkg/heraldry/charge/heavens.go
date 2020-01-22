package charge

func getHeavensCharges() []Raster {
	charges := []Raster{
		{
			Identifier: "estoille",
			Name:       "estoille",
			Noun:       "estoille",
			NounPlural: "estoilles",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"estoille",
				"heavens",
				"star",
			},
		},
		{
			Identifier: "sun-in-splendor",
			Name:       "sun in splendor",
			Noun:       "sun",
			NounPlural: "suns",
			Descriptor: "in splendor",
			SingleOnly: false,
			Tags: []string{
				"sun",
				"heavens",
			},
		},
	}

	return charges
}
