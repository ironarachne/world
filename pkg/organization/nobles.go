package organization

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/heraldry"
	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/random"
)

// GenerateNobleHouse generates a random noble house
func GenerateNobleHouse(ctx context.Context, originCulture culture.Culture) (Organization, error) {
	org := Organization{}

	orgType := getNobleType()
	org.Type = orgType
	org.SizeClass = getRandomSizeClass(ctx, org.Type.MinSize, org.Type.MaxSize)
	org.Size = random.Intn(ctx, org.SizeClass.MaxSize-org.SizeClass.MinSize) + org.SizeClass.MinSize

	org.LeaderType = org.setLeaderType(ctx)
	leader, err := org.getLeader(ctx, originCulture)
	if err != nil {
		err = fmt.Errorf("Failed to generate noble house: %w", err)
		return Organization{}, err
	}
	orgType.Ranks[0].Title = leader.CharacterData.Title
	org.Type = orgType
	leader.Rank.Title = leader.CharacterData.Title

	device, err := heraldry.Generate(ctx)
	if err != nil {
		err = fmt.Errorf("Could not generate noble house: %w", err)
		return Organization{}, err
	}
	leader.CharacterData.Heraldry = device.Simplify()
	houseName := "House of " + leader.CharacterData.LastName
	if leader.CharacterData.Gender.Name == "female" {
		leader.CharacterData.Titles = append(leader.CharacterData.Titles, "Matriarch of the "+houseName)
	} else {
		leader.CharacterData.Titles = append(leader.CharacterData.Titles, "Patriarch of the "+houseName)
	}

	org.Leader = leader
	org.Name = houseName
	org.Heraldry = org.Leader.CharacterData.Heraldry
	trait, err := org.setTrait(ctx)
	if err != nil {
		err = fmt.Errorf("Failed to generate noble house: %w", err)
		return Organization{}, err
	}
	org.PrimaryTrait = trait

	notableMembers, err := org.getNobleMembers(ctx, originCulture, leader.CharacterData)
	if err != nil {
		err = fmt.Errorf("Failed to generate noble house: %w", err)
		return Organization{}, err
	}
	org.NotableMembers = notableMembers

	return org, nil
}

func getNobleType() Type {
	professions, _ := profession.ByTag("noble")

	house := Type{
		Name:    "noble house",
		MaxSize: 14,
		MinSize: 6,
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
			"selfish",
			"selfless",
		},
		CanBeLedByGroup:   false,
		LeaderTitle:       "master",
		MemberProfessions: professions,
		NameFirstParts: []string{
			"House",
		},
		NameSecondParts: []string{
			"Nobles",
		},
		NameTemplate: "{{.FirstPart}} of {{.SecondPart}}",
		Ranks: []Rank{
			{
				Title:       "Master",
				Precedence:  0,
				MaxNumber:   1,
				AgeModifier: 1.25,
				AgeCategories: []string{
					"adult",
				},
			},
			{
				Title:       "Spouse",
				Precedence:  0,
				MaxNumber:   1,
				AgeModifier: 1.15,
				AgeCategories: []string{
					"adult",
				},
			},
			{
				Title:       "Heir",
				Precedence:  1,
				MaxNumber:   1,
				AgeModifier: 1.25,
				AgeCategories: []string{
					"young adult",
				},
			},
			{
				Title:       "Lord",
				Precedence:  2,
				MaxNumber:   0,
				AgeModifier: 1.0,
				AgeCategories: []string{
					"teenager",
				},
			},
			{
				Title:       "Lord",
				Precedence:  3,
				MaxNumber:   0,
				AgeModifier: 1.0,
				AgeCategories: []string{
					"child",
				},
			},
		},
	}

	return house
}

func (org Organization) getNobleMembers(ctx context.Context, originCulture culture.Culture, master character.Character) ([]Member, error) {
	var memberData Member
	var memberRank Rank
	var nobleTitle string

	members := []Member{}

	maxMemberCount := org.Size - 2

	spouseData, err := character.GenerateCompatibleMate(ctx, master)
	if err != nil {
		err = fmt.Errorf("Failed to generate noble house: %w", err)
		return []Member{}, err
	}
	spouseData.Heraldry = master.Heraldry
	spouseData.LastName = master.LastName
	spouseData.Profession = master.Profession

	spouse := Member{
		CharacterData: spouseData,
		Rank: Rank{
			Title:       "Spouse",
			Precedence:  0,
			MaxNumber:   1,
			AgeModifier: 1.0,
			AgeCategories: []string{
				"adult",
			},
		},
	}

	members = append(members, spouse)

	for i := 0; i < maxMemberCount; i++ {
		member, err := character.Generate(ctx, originCulture)
		if err != nil {
			err = fmt.Errorf("Failed to get noble house members: %w", err)
			return []Member{}, err
		}
		memberRank = org.Type.GetRandomMemberRank(ctx, members)
		member.Race = master.Race
		member.LastName = master.LastName
		member = member.ChangeAge(ctx, GetModifiedMemberAge(ctx, memberRank, member.Race))

		prof, err := org.Type.GetRandomMemberProfession(ctx)
		if err != nil {
			err = fmt.Errorf("Failed to get noble house members: %w", err)
			return []Member{}, err
		}
		member.Profession = prof
		member.Heraldry = master.Heraldry

		if member.Gender.Name == "female" {
			nobleTitle = "Lady"
		} else {
			nobleTitle = "Lord"
		}
		member.Title = nobleTitle
		memberData = Member{
			CharacterData: member,
			Rank:          memberRank,
		}
		members = append(members, memberData)
	}

	return members, nil
}
