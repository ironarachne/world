package charge

func getPeopleCharges() []Raster {
	charges := []Raster{
		{
			Identifier: "chevalier-on-horseback",
			Name:       "chevalier on horseback",
			Noun:       "chevalier",
			NounPlural: "chevaliers",
			Descriptor: "on horseback",
			SingleOnly: false,
			Tags: []string{
				"chevalier",
				"knight",
				"people",
				"aggressive",
			},
		},
	}

	return charges
}
