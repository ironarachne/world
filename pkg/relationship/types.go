package relationship

// Type is a type of relationship
type Type struct {
	Name        string   `json:"name"`
	Descriptors []string `json:"descriptors"`
	InverseName string   `json:"inverse_name"`
	Disallows   []string `json:"disallows"`
}

// AllTypes returns all relationship types
func AllTypes() []Type {
	types := []Type{
		{
			Name:        "parent",
			InverseName: "child",
			Disallows:   []string{},
			Descriptors: []string{
				"is a parent of",
				"is parent of",
			},
		},
		{
			Name:        "child",
			InverseName: "parent",
			Disallows:   []string{},
			Descriptors: []string{
				"is a child of",
				"is the child of",
				"is child of",
			},
		},
		{
			Name:        "friend",
			InverseName: "friend",
			Disallows: []string{
				"enemy",
				"lover",
			},
			Descriptors: []string{
				"is a confidant of",
				"is a friend of",
				"is an ally of",
			},
		},
		{
			Name:        "enemy",
			InverseName: "enemy",
			Disallows: []string{
				"friend",
				"lover",
			},
			Descriptors: []string{
				"is the enemy of",
				"is the hated foe of",
				"wars against",
			},
		},
		{
			Name:        "lover",
			InverseName: "lover",
			Disallows: []string{
				"enemy",
				"friend",
			},
			Descriptors: []string{
				"desires",
				"is a lover of",
				"loves",
				"trysts with",
			},
		},
		{
			Name:        "negative opinion",
			InverseName: "negative opinion",
			Disallows: []string{
				"friend",
				"lover",
				"positive opinion",
			},
			Descriptors: []string{
				"despises",
				"dislikes",
				"distrusts",
				"envies",
				"fears",
				"hates",
				"is bored by",
				"is chagrined by",
				"is suspicious of",
				"suspects",
				"tends to avoid",
				"worries about",
			},
		},
		{
			Name:        "positive opinion",
			InverseName: "positive opinion",
			Disallows: []string{
				"negative opinion",
				"enemy",
			},
			Descriptors: []string{
				"cherishes",
				"enjoys the company of",
				"is amused by",
				"is hopeful for",
				"owes a debt to",
				"respects",
				"supports",
				"trusts",
			},
		},
		{
			Name:        "rival",
			InverseName: "rival",
			Disallows:   []string{},
			Descriptors: []string{
				"fights with",
				"is the rival of",
			},
		},
	}

	return types
}
