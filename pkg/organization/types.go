package organization

import (
	"github.com/ironarachne/world/pkg/profession"
	"math/rand"
	"strings"
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
	MemberProfessions []profession.Profession
	NameFirstParts    []string
	NameSecondParts   []string
	NameTemplate      string
}

func getAllTypes() []Type {
	var adventurers []profession.Profession

	divine := profession.ByTag("divine")
	fighters := profession.ByTag("fighter")
	mages := profession.ByTag("mage")

	adventurers = append(adventurers, fighters...)
	adventurers = append(adventurers, mages...)
	adventurers = append(adventurers, divine...)

	types := []Type{
		{
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
			CanBeLedByGroup:   false,
			LeaderMaxAge:      50,
			LeaderMinAge:      25,
			LeaderTitle:       "captain",
			MemberMaxAge:      50,
			MemberMinAge:      15,
			MemberProfessions: adventurers,
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
		{
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
			CanBeLedByGroup:   true,
			LeaderMaxAge:      80,
			LeaderMinAge:      35,
			LeaderTitle:       "high priest",
			MemberMaxAge:      90,
			MemberMinAge:      15,
			MemberProfessions: divine,
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
		{
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
			CanBeLedByGroup:   true,
			LeaderMaxAge:      50,
			LeaderMinAge:      25,
			LeaderTitle:       "captain",
			MemberMaxAge:      50,
			MemberMinAge:      15,
			MemberProfessions: fighters,
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

	guild := getCraftingGuild()

	types = append(types, guild)

	return types
}

func getCraftingGuild() Type {
	guild := Type{
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
		NameTemplate: "{{.FirstPart}} {{.SecondPart}}'s Guild",
	}

	crafters := profession.ByTag("crafter")
	memberProfessions := profession.RandomSet(1, crafters)
	guild.MemberProfessions = memberProfessions
	guild.NameSecondParts = []string{
		strings.Title(memberProfessions[0].Name),
	}

	return guild
}

// GetRandomMemberProfession returns a random profession from those available to the organization
func (t Type) GetRandomMemberProfession() profession.Profession {
	prof := profession.RandomSet(1, t.MemberProfessions)

	return prof[0]
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

	return orgTypes[rand.Intn(len(orgTypes))]
}
