package organization

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
)

// Type is a type of organization
type Type struct {
	Name              string
	PossibleTraits    []string
	CanBeLedByGroup   bool
	LeaderMaxAge      int
	LeaderMinAge      int
	LeaderTitle       string
	MemberMaxAge      int
	MemberMinAge      int
	MemberProfessions []string
	NameFirstParts    []string
	NameSecondParts   []string
	NameTemplate      string
}

func getAllTypes() []Type {
	types := []Type{
		Type{
			Name: "adventuring company",
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
			LeaderMaxAge:    50,
			LeaderMinAge:    25,
			LeaderTitle:     "captain",
			MemberMaxAge:    50,
			MemberMinAge:    15,
			MemberProfessions: []string{
				"barbarian",
				"bard",
				"cleric",
				"mage",
				"paladin",
				"rogue",
				"warrior",
			},
			NameFirstParts: []string{
				"Black",
				"Burning",
				"Crimson",
				"Free",
				"Golden",
				"Iron",
				"Righteous",
				"Silver",
				"Wandering",
				"White",
			},
			NameSecondParts: []string{
				"Axes",
				"Blades",
				"Coins",
				"Company",
				"Dragons",
				"Giants",
				"Lords",
				"Pikes",
				"Swords",
				"Wolves",
				"Wyverns",
			},
			NameTemplate: "{{.FirstPart}} {{.SecondPart}}",
		},
		Type{
			Name: "church",
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
			LeaderMaxAge:    80,
			LeaderMinAge:    35,
			LeaderTitle:     "high priest",
			MemberMaxAge:    90,
			MemberMinAge:    15,
			MemberProfessions: []string{
				"acolyte",
				"monk",
				"priest",
			},
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
		Type{
			Name: "guild",
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
			LeaderMaxAge:    70,
			LeaderMinAge:    30,
			LeaderTitle:     "guild leader",
			MemberMaxAge:    90,
			MemberMinAge:    20,
			MemberProfessions: []string{
				"artisan",
				"craftsman",
				"financier",
				"merchant",
			},
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
		Type{
			Name: "mercenary company",
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
			LeaderMaxAge:    50,
			LeaderMinAge:    25,
			LeaderTitle:     "captain",
			MemberMaxAge:    50,
			MemberMinAge:    15,
			MemberProfessions: []string{
				"archer",
				"cavalryman",
				"heavy footman",
				"knife fighter",
				"pikeman",
				"swordsman",
			},
			NameFirstParts: []string{
				"Black",
				"Burning",
				"Crimson",
				"Free",
				"Gilded",
				"Iron",
				"Ivory",
				"Red",
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

// GetRandomMemberProfession returns a random profession as a string
func (t Type) GetRandomMemberProfession() string {
	return random.String(t.MemberProfessions)
}

// GetRandomLeaderAge returns an appropriate age for a leader of this group
func (t Type) GetRandomLeaderAge() int {
	return rand.Intn(t.LeaderMaxAge-t.LeaderMinAge) + t.LeaderMinAge
}

// GetRandomMemberAge returns an appropriate age for a member of this group
func (t Type) GetRandomMemberAge() int {
	return rand.Intn(t.MemberMaxAge-t.MemberMinAge) + t.MemberMinAge
}

func getRandomType() Type {
	orgTypes := getAllTypes()

	return orgTypes[rand.Intn(len(orgTypes)-1)]
}
