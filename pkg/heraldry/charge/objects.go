package charge

func getObjectCharges() []Raster {
	charges := []Raster{
		{
			Identifier: "anchor",
			Name:       "anchor",
			Noun:       "anchor",
			NounPlural: "anchors",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"anchor",
				"object",
				"water",
				"navy",
				"sailor",
			},
		},
		{
			Identifier: "annulet",
			Name:       "annulet",
			Noun:       "annulet",
			NounPlural: "annulets",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"annulet",
				"object",
			},
		},
		{
			Identifier: "axe",
			Name:       "axe",
			Noun:       "axe",
			NounPlural: "axes",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"axe",
				"object",
				"weapon",
				"aggressive",
			},
		},
		{
			Identifier: "galleon",
			Name:       "galleon",
			Noun:       "galleon",
			NounPlural: "galleons",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"galleon",
				"object",
				"water",
				"navy",
				"sailor",
			},
		},
	}

	return charges
}
