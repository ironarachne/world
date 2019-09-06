package profession

func finishers() []Profession {
	professions := []Profession{
		{
			Name:        "apothecary",
			Description: "Apothecaries make healing salves and ointments.",
			Tags: []string{
				"crafter",
			},
		},
		{
			Name:        "bowyer",
			Description: "Bowyers craft bows for use in hunting and combat.",
			Tags: []string{
				"crafter",
			},
		},
		{
			Name:        "carpenter",
			Description: "Carpenters produce furniture and other works from finished lumber.",
			Tags: []string{
				"crafter",
			},
		},
		{
			Name:        "cartwright",
			Description: "Cartwrights build carts, wagons, and other land vehicles.",
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
