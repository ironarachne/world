package profession

func finishers() []Profession {
	professions := []Profession{
		{
			Name:        "carpenter",
			Description: "Carpenters produce furniture and other works from finished lumber.",
			Tags: []string{
				"crafter",
			},
		},
		{
			Name:        "cobbler",
			Description: "Cobblers make shoes from leather and other materials.",
			Tags: []string{
				"crafter",
			},
		},
		{
			Name:        "mason",
			Description: "Masons architect and build structures from finished stone.",
			Tags: []string{
				"crafter",
			},
		},
		{
			Name:        "tailor",
			Description: "Tailors take fabrics and make them into clothing.",
			Tags: []string{
				"crafter",
			},
		},
	}

	return professions
}