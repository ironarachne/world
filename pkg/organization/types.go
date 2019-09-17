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
	AgeModifier float64
	AgeCategory string
}

// Type is a type of organization
type Type struct {
	Name              string
	PossibleTraits    []string
	CanBeLedByGroup   bool
	LeaderTitle       string
	MemberProfessions []profession.Profession
	NameFirstParts    []string
	NameSecondParts   []string
	NameTemplate      string
	Ranks             []Rank
	MaxSize           int
	MinSize           int
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
			Name:    "adventuring company",
			MaxSize: 50,
			MinSize: 2,
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
			LeaderTitle:       "captain",
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
					Title:       "Captain",
					Precedence:  0,
					MaxNumber:   1,
					AgeModifier: 1.25,
					AgeCategory: "young adult",
				},
				{
					Title:       "Adventurer",
					Precedence:  1,
					MaxNumber:   0,
					AgeModifier: 1.0,
					AgeCategory: "young adult",
				},
			},
		},
		{
			Name:    "church",
			MaxSize: 1000,
			MinSize: 10,
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
			LeaderTitle:       "high priest",
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
					Title:       "High Priest",
					Precedence:  0,
					MaxNumber:   1,
					AgeModifier: 1.5,
					AgeCategory: "adult",
				},
				{
					Title:       "Priest",
					Precedence:  1,
					MaxNumber:   0,
					AgeModifier: 1.25,
					AgeCategory: "adult",
				},
				{
					Title:       "Acolyte",
					Precedence:  2,
					MaxNumber:   0,
					AgeModifier: 0.9,
					AgeCategory: "young adult",
				},
			},
		},
		{
			Name:    "mercenary company",
			MaxSize: 500,
			MinSize: 20,
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
			LeaderTitle:       "captain",
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
					Title:       "Captain",
					Precedence:  0,
					MaxNumber:   1,
					AgeModifier: 1.0,
					AgeCategory: "adult",
				},
				{
					Title:       "Mercenary",
					Precedence:  1,
					MaxNumber:   0,
					AgeModifier: 1.0,
					AgeCategory: "young adult",
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
	wizardSociety, err := getWizardSociety()
	if err != nil {
		err = fmt.Errorf("Failed to generate organization types: %w", err)
		return []Type{}, err
	}

	types = append(types, guild)
	types = append(types, school)
	types = append(types, wizardSociety)

	return types, nil
}

func getCraftingGuild() (Type, error) {
	guild := Type{
		Name:    "guild",
		MaxSize: 150,
		MinSize: 10,
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
		NameTemplate: "{{.FirstPart}} {{.SecondPart}}'s Guild",
		Ranks: []Rank{
			{
				Title:       "Guild Leader",
				Precedence:  0,
				MaxNumber:   1,
				AgeModifier: 1.5,
				AgeCategory: "adult",
			},
			{
				Title:       "Artisan",
				Precedence:  1,
				MaxNumber:   0,
				AgeModifier: 1.1,
				AgeCategory: "adult",
			},
			{
				Title:       "Apprentice Artisan",
				Precedence:  2,
				MaxNumber:   0,
				AgeModifier: 0.7,
				AgeCategory: "young adult",
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

func getWizardSociety() (Type, error) {
	org := Type{
		Name:    "wizard society",
		MaxSize: 100,
		MinSize: 10,
		PossibleTraits: []string{
			"aggressive",
			"aloof",
			"ambitious",
			"closeted",
			"dangerous",
			"eldritch",
			"guarded",
			"knowledgeable",
			"mysterious",
			"powerful",
			"reckless",
			"secretive",
		},
		CanBeLedByGroup: true,
		LeaderTitle:     "archmage",
		NameFirstParts: []string{
			"Order",
			"Brotherhood",
			"Society",
		},
		NameSecondParts: []string{
			"Arcane Arts",
			"Arcane Endeavors",
			"Eldritch Arts",
			"Mysteries of the Arcane",
			"Mysteries",
			"Mystical Arts",
			"Sorcerous Endeavors",
			"the Arcane Circle",
			"the Eldritch Mysteries",
			"the Mystical",
		},
		NameTemplate: "{{.FirstPart}} of {{.SecondPart}}",
		Ranks: []Rank{
			{
				Title:       "Archmage",
				Precedence:  0,
				MaxNumber:   1,
				AgeModifier: 4.0,
				AgeCategory: "elderly",
			},
			{
				Title:       "Adept",
				Precedence:  1,
				MaxNumber:   0,
				AgeModifier: 2.5,
				AgeCategory: "adult",
			},
			{
				Title:       "Master",
				Precedence:  2,
				MaxNumber:   0,
				AgeModifier: 2.0,
				AgeCategory: "adult",
			},
			{
				Title:       "Journeyman",
				Precedence:  3,
				MaxNumber:   0,
				AgeModifier: 1.25,
				AgeCategory: "young adult",
			},
			{
				Title:       "Apprentice",
				Precedence:  4,
				MaxNumber:   0,
				AgeModifier: 1.0,
				AgeCategory: "teenager",
			},
		},
	}

	wizards := profession.ByTag("wizard")
	memberProfessions, err := profession.RandomSet(3, wizards)
	if err != nil {
		err = fmt.Errorf("Failed to generate wizard society: %w", err)
		return Type{}, err
	}
	org.MemberProfessions = memberProfessions

	return org, nil
}

func getWizardSchool() (Type, error) {
	org := Type{
		Name:    "school of wizardry",
		MaxSize: 500,
		MinSize: 50,
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
		LeaderTitle:     "headmaster",
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
				Title:       "Headmaster",
				Precedence:  0,
				MaxNumber:   1,
				AgeModifier: 3.0,
				AgeCategory: "elderly",
			},
			{
				Title:       "Instructor",
				Precedence:  1,
				MaxNumber:   0,
				AgeModifier: 2.0,
				AgeCategory: "adult",
			},
			{
				Title:       "Student",
				Precedence:  2,
				MaxNumber:   0,
				AgeModifier: 1.0,
				AgeCategory: "teenager",
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

func getRandomType() (Type, error) {
	orgTypes, err := getAllTypes()
	if err != nil {
		err = fmt.Errorf("Failed to generate organization type: %w", err)
		return Type{}, err
	}

	randomOrgType := orgTypes[rand.Intn(len(orgTypes))]

	return randomOrgType, nil
}
