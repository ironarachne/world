package organization

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/ironarachne/world/pkg/profession"
)

// Rank is a rank held within an organization
type Rank struct {
	Title       string
	Precedence  int
	MaxNumber   int
	MemberNames []string
}

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
	Ranks             []Rank
}

func getAllTypes() ([]Type, error) {
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
			Ranks: []Rank{
				{
					Title:      "Captain",
					Precedence: 0,
					MaxNumber:  1,
				},
				{
					Title:      "Adventurer",
					Precedence: 1,
					MaxNumber:  0,
				},
			},
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
			Ranks: []Rank{
				{
					Title:      "High Priest",
					Precedence: 0,
					MaxNumber:  1,
				},
				{
					Title:      "Priest",
					Precedence: 1,
					MaxNumber:  0,
				},
				{
					Title:      "Acolyte",
					Precedence: 2,
					MaxNumber:  0,
				},
			},
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
			Ranks: []Rank{
				{
					Title:      "Captain",
					Precedence: 0,
					MaxNumber:  1,
				},
				{
					Title:      "Mercenary",
					Precedence: 1,
					MaxNumber:  0,
				},
			},
		},
	}

	guild, err := getCraftingGuild()
	if err != nil {
		err = fmt.Errorf("Failed to generate organization types: %w", err)
		return []Type{}, err
	}
	school, err := getWizardSchool()
	if err != nil {
		err = fmt.Errorf("Failed to generate organization types: %w", err)
		return []Type{}, err
	}

	types = append(types, guild)
	types = append(types, school)

	return types, nil
}

func getCraftingGuild() (Type, error) {
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
		Ranks: []Rank{
			{
				Title:      "Guild Leader",
				Precedence: 0,
				MaxNumber:  1,
			},
			{
				Title:      "Artisan",
				Precedence: 1,
				MaxNumber:  0,
			},
		},
	}

	crafters := profession.ByTag("crafter")
	memberProfessions, err := profession.RandomSet(1, crafters)
	if err != nil {
		err = fmt.Errorf("Failed to generate guild: %w", err)
		return Type{}, err
	}
	guild.MemberProfessions = memberProfessions
	guild.NameSecondParts = []string{
		strings.Title(memberProfessions[0].Name),
	}

	return guild, nil
}

func getWizardSchool() (Type, error) {
	org := Type{
		Name: "school of wizardry",
		PossibleTraits: []string{
			"aloof",
			"ambitious",
			"august",
			"closeted",
			"eldritch",
			"esoteric",
			"guarded",
			"knowledgeable",
			"powerful",
			"reckless",
			"regimented",
			"wise",
		},
		CanBeLedByGroup: true,
		LeaderMaxAge:    100,
		LeaderMinAge:    50,
		LeaderTitle:     "headmaster",
		MemberMaxAge:    120,
		MemberMinAge:    10,
		NameFirstParts: []string{
			"Academy",
			"Institute",
			"School",
		},
		NameSecondParts: []string{
			"Arcane Arts",
			"Arcane Endeavors",
			"Arcane Studies",
			"Eldritch Investigations",
			"Magic Studies",
			"Mysteries of the Arcane",
			"Mystical Arts",
			"Mystical Studies",
			"Sorcerous Endeavors",
			"Wizarding Arts",
			"Wizardry and Magic",
		},
		NameTemplate: "{{.FirstPart}} of {{.SecondPart}}",
		Ranks: []Rank{
			{
				Title:      "Headmaster",
				Precedence: 0,
				MaxNumber:  1,
			},
			{
				Title:      "Instructor",
				Precedence: 1,
				MaxNumber:  0,
			},
			{
				Title:      "Student",
				Precedence: 2,
				MaxNumber:  0,
			},
		},
	}

	wizards := profession.ByTag("wizard")
	memberProfessions, err := profession.RandomSet(3, wizards)
	if err != nil {
		err = fmt.Errorf("Failed to generate wizard school: %w", err)
		return Type{}, err
	}
	org.MemberProfessions = memberProfessions

	return org, nil
}

// GetRandomMemberRank returns a random appropriate rank that is NOT a leader
func (t Type) GetRandomMemberRank() Rank {
	var possibleRanks []Rank

	for _, r := range t.Ranks {
		if r.Precedence != 0 {
			possibleRanks = append(possibleRanks, r)
		}
	}

	if len(possibleRanks) == 1 {
		return possibleRanks[0]
	}

	rank := possibleRanks[rand.Intn(len(possibleRanks))]

	return rank
}

// GetRandomMemberProfession returns a random profession from those available to the organization
func (t Type) GetRandomMemberProfession() (profession.Profession, error) {
	prof, err := profession.RandomSet(1, t.MemberProfessions)
	if err != nil {
		err = fmt.Errorf("Failed to generate organization type: %w", err)
		return profession.Profession{}, err
	}

	return prof[0], nil
}

// GetRandomLeaderAge returns an appropriate age for a leader of this group
func (t Type) GetRandomLeaderAge() int {
	age := rand.Intn(t.LeaderMaxAge-t.LeaderMinAge) + t.LeaderMinAge
	return age
}

// GetRandomMemberAge returns an appropriate age for a member of this group
func (t Type) GetRandomMemberAge() int {
	age := rand.Intn(t.MemberMaxAge-t.MemberMinAge) + t.MemberMinAge
	return age
}

func getRandomType() (Type, error) {
	orgTypes, err := getAllTypes()
	if err != nil {
		err = fmt.Errorf("Failed to generate organization type: %w", err)
		return Type{}, err
	}

	randomOrgType := orgTypes[rand.Intn(len(orgTypes))]

	return randomOrgType, nil
}
