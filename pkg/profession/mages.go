package profession

func mages() []Profession {
	professions := []Profession{
		{
			Name:        "alchemist",
			Description: "Alchemists take raw materials and use magic to turn them into new substances.",
			Tags: []string{
				"crafter",
				"mage",
			},
		},
		{
			Name:        "necromancer",
			Description: "Necromancers use magic to speak with and sometimes raise the dead.",
			Tags: []string{
				"death",
				"mage",
			},
		},
		{
			Name:        "scrollmaker",
			Description: "Scrollmakers enscribe sorcerous incantations onto paper for later use by skilled practitioners.",
			Tags: []string{
				"crafter",
				"mage",
			},
		},
		{
			Name:        "sorcerer",
			Description: "Sorcerers wield intuitive magic.",
			Tags: []string{
				"mage",
			},
		},
		{
			Name:        "warlock",
			Description: "Warlocks wield magic borne of strange and terrible origins.",
			Tags: []string{
				"mage",
			},
		},
		{
			Name:        "wizard",
			Description: "Wizards study magic as a science and wield it in complex and powerful ways.",
			Tags: []string{
				"mage",
			},
		},
	}

	return professions
}
