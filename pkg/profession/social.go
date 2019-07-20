package profession

func social() []Profession {
	professions := []Profession{
		{
			Name: "bard",
			Description: "Bards are minstrels who travel and entertain.",
			Tags: []string{
				"musician",
				"social",
			},
		},
		{
			Name:        "mayor",
			Description: "Mayors coordinate a town or city.",
			Tags: []string{
				"leader",
				"social",
			},
		},
		{
			Name:        "noble",
			Description: "Nobles act as military coordinators, protectors, and leaders.",
			Tags: []string{
				"leader",
				"social",
			},
		},
		{
			Name:        "priest",
			Description: "Priests interpret the will of the gods.",
			Tags: []string{
				"divine",
				"social",
			},
		},
		{
			Name:        "shaman",
			Description: "Shamans speak for the spirits.",
			Tags: []string{
				"spiritual",
				"social",
			},
		},
	}

	return professions
}