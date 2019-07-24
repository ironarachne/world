package profession

func blacksmiths() []Profession {
	professions := []Profession{
		{
			Name:        "armorsmith",
			Description: "Armorsmiths make armor from raw materials.",
			Tags: []string{
				"crafter",
				"smith",
			},
		},
		{
			Name:        "blacksmith",
			Description: "Blacksmiths turn raw metal into useful materials.",
			Tags: []string{
				"crafter",
				"refiner",
				"smith",
			},
		},
		{
			Name:        "farrier",
			Description: "Farriers make horseshoes and perform veterinarian surgery.",
			Tags: []string{
				"crafter",
				"smith",
			},
		},
		{
			Name:        "weaponsmith",
			Description: "Weaponsmiths specialize in making weapons.",
			Tags: []string{
				"crafter",
				"smith",
			},
		},
	}

	return professions
}
