package profession

func refiners() []Profession {
	professions := []Profession{
		{
			Name:        "animal trainer",
			Description: "Animal trainers tame animals and turn them into disciplined followers.",
			Tags: []string{
				"crafter",
			},
		},
		{
			Name:        "baker",
			Description: "Bakers bake bread and other oven-derived goods from grain and related ingredients.",
			Tags: []string{
				"crafter",
			},
		},
		{
			Name:        "brewer",
			Description: "Brewers turn grain into ale, beer, and related beverages.",
			Tags: []string{
				"crafter",
			},
		},
		{
			Name:        "distiller",
			Description: "Distillers make hard alcoholic beverages.",
			Tags: []string{
				"crafter",
			},
		},
		{
			Name:        "farmer",
			Description: "Farmers till the earth and practice animal husbandry.",
			Tags: []string{
				"harvester",
			},
		},
		{
			Name: "glassmaker",
			Description: "Glassmakers turn sand and other materials into glass.",
			Tags: []string{
				"crafter",
			},
		},
		{
			Name:        "logger",
			Description: "Loggers cut down trees to make lumber.",
			Tags: []string{
				"harvester",
			},
		},
		{
			Name:        "miller",
			Description: "Millers take raw grain and turn it into flour.",
			Tags: []string{
				"crafter",
			},
		},
		{
			Name:        "miner",
			Description: "Miners extract raw ore from the earth.",
			Tags: []string{
				"harvester",
			},
		},
		{
			Name:        "potter",
			Description: "Potters turn raw clay into finished pots and other ceramics.",
			Tags: []string{
				"crafter",
			},
		},
		{
			Name:        "spinner",
			Description: "Spinners create thread and cloth from raw fibers.",
			Tags: []string{
				"crafter",
			},
		},
		{
			Name:        "tanner",
			Description: "Tanners turn raw animal skin into cured hides.",
			Tags: []string{
				"crafter",
			},
		},
		{
			Name:        "vintner",
			Description: "Vintners make wine.",
			Tags: []string{
				"crafter",
			},
		},
	}

	return professions
}