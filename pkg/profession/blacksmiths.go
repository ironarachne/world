package profession

func blacksmiths() []Profession {
	professions := []Profession{
		{
			Name:        "armorsmith",
			Description: "Armorsmiths make armor from raw materials.",
			Tags: []string{
				"smith",
				"crafter",
			},
		},
		{
			Name:        "blacksmith",
			Description: "Blacksmiths turn raw metal into useful materials.",
			Tags: []string{
				"smith",
				"crafter",
			},
		},
		{
			Name:        "farrier",
			Description: "Farriers make horseshoes and perform veterinarian surgery.",
			Tags: []string{
				"smith",
				"crafter",
			},
		},
		{
			Name:        "weaponsmith",
			Description: "Weaponsmiths specialize in making weapons.",
			Tags: []string{
				"smith",
				"crafter",
			},
		},
	}

	return professions
}
