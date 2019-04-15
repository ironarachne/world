package organization

import "math/rand"

// OrganizationType is a type of organization
type OrganizationType struct {
	Name            string
	PossibleTraits  []string
	CanBeLedByGroup bool
	LeaderTitle     string
	NameFirstParts  []string
	NameSecondParts []string
	NameTemplate    string
}

func getAllTypes() []OrganizationType {
	types := []OrganizationType{
		OrganizationType{
			Name: "an adventuring company",
			PossibleTraits: []string{
				"aggressive",
				"avaricious",
				"charismatic",
				"curious",
				"forceful",
				"heroic",
				"honorable",
				"reckless",
				"reserved",
				"selfless",
			},
			CanBeLedByGroup: false,
			LeaderTitle:     "captain",
			NameFirstParts: []string{
				"Black",
				"Burning",
				"Crimson",
				"Free",
				"Righteous",
				"Silver",
				"Wandering",
				"White",
			},
			NameSecondParts: []string{
				"Axes",
				"Blade",
				"Coin",
				"Company",
				"Dragons",
				"Giants",
				"Lords",
				"Pike",
				"Swords",
				"Wolves",
				"Wyverns",
			},
			NameTemplate: "{{.FirstPart}} {{.SecondPart}}",
		},
		OrganizationType{
			Name: "a church",
			PossibleTraits: []string{
				"penitent",
				"helpful",
				"selfless",
				"controlling",
				"merciful",
				"merciless",
				"penniless",
				"proselytizing",
			},
			CanBeLedByGroup: true,
			LeaderTitle:     "high priest",
			NameFirstParts: []string{
				"Holy",
				"Glorious",
				"Exalted",
				"Humble",
				"Penitent",
				"Righteous",
			},
			NameSecondParts: []string{
				"Divine Hand",
				"Dove",
				"Eye",
				"Flame",
				"Forest",
				"Four Truths",
				"Iron Path",
				"Light",
				"Meek",
				"Path",
				"River",
				"Sword",
				"Three Truths",
				"Truth",
			},
			NameTemplate: "{{.FirstPart}} Church of the {{.SecondPart}}",
		},
		OrganizationType{
			Name: "a guild",
			PossibleTraits: []string{
				"ambitious",
				"avaricious",
				"charitable",
				"fair",
				"frugal",
				"lazy",
				"manipulative",
				"observant",
				"productive",
			},
			CanBeLedByGroup: true,
			LeaderTitle:     "guild leader",
			NameFirstParts: []string{
				"August",
				"East Wind",
				"Global",
				"Imperial",
				"Incorporated",
				"North Wind",
				"Royal",
				"South Wind",
				"West Wind",
			},
			NameSecondParts: []string{
				"Alchemist",
				"Apothecary",
				"Artist",
				"Blacksmith",
				"Bowyer",
				"Carpenter",
				"Fletcher",
				"Furrier",
				"Mason",
				"Merchant",
				"Moneylender",
				"Painter",
				"Sculptor",
				"Tanner",
			},
			NameTemplate: "{{.FirstPart}} {{.SecondPart}}'s Guild",
		},
		OrganizationType{
			Name: "a mercenary company",
			PossibleTraits: []string{
				"aggressive",
				"belligerent",
				"avaricious",
				"neutral",
				"honorable",
				"dishonorable",
				"trustworthy",
			},
			CanBeLedByGroup: true,
			LeaderTitle:     "captain",
			NameFirstParts: []string{
				"Black",
				"Burning",
				"Crimson",
				"Free",
				"Gilded",
				"Ivory",
				"Silver",
				"White",
			},
			NameSecondParts: []string{
				"Army",
				"Axes",
				"Company",
				"Dragons",
				"Giants",
				"Legion",
				"Lords",
				"Sentinels",
				"Swords",
				"Wolves",
				"Wyverns",
			},
			NameTemplate: "{{.FirstPart}} {{.SecondPart}}",
		},
	}

	return types
}

func getRandomType() OrganizationType {
	orgTypes := getAllTypes()

	return orgTypes[rand.Intn(len(orgTypes)-1)]
}
