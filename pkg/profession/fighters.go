package profession

func fighters() []Profession {
	professions := []Profession{
		{
			Name:        "cleric",
			Description: "Clerics are divine messengers who carry out the will of the gods.",
			Tags: []string{
				"divine",
				"fighter",
			},
		},
		{
			Name: "paladin",
			Description: "Paladins are holy warriors who wield divine magic in addition to their fighing skill.",
			Tags: []string{
				"divine",
				"fighter",
			},
		},
		{
			Name: "soldier",
			Description: "Soldiers are formally trained fighters with allegiance to a fighting unit.",
			Tags: []string{
				"fighter",
			},
		},
		{
			Name: "warrior",
			Description: "Warriors are fighters who have no formal training, but have gained experience in battle.",
			Tags: []string{
				"fighter",
			},
		},
	}

	return professions
}
