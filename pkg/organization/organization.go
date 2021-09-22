/*
Package organization implements social organizations and tools for dealing with them
*/
package organization

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	"github.com/ironarachne/world/pkg/age"
	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/heraldry"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/species"
)

// NameData is name data
type NameData struct {
	FirstPart  string `json:"first_part"`
	SecondPart string `json:"second_part"`
}

// Member is a member of an organization
type Member struct {
	CharacterData character.Character `json:"character_data"`
	Rank          Rank                `json:"rank"`
}

// Organization is an organization
type Organization struct {
	Name           string                    `json:"name"`
	Type           Type                      `json:"type"`
	Superior       string                    `json:"superior"` // Superior is the name of an organization that this organization owes allegiance to
	Heraldry       heraldry.SimplifiedDevice `json:"heraldry"`
	Leader         Member                    `json:"leader"`
	LeaderType     string                    `json:"leader_type"`
	NotableMembers []Member                  `json:"notable_members"`
	SizeClass      SizeClass                 `json:"size_class"`
	Size           int                       `json:"size"`
	PrimaryTrait   string                    `json:"primary_trait"`
}

func (org Organization) setLeaderType(ctx context.Context) string {
	leaderType := "individual"

	if org.Type.CanBeLedByGroup {
		x := random.Intn(ctx, 10)

		if x >= 9 {
			leaderType = "group"
		}
	}

	return leaderType
}

func (org Organization) setName(ctx context.Context) (string, error) {
	var tplOutput bytes.Buffer

	firstPart, err := random.String(ctx, org.Type.NameFirstParts)
	if err != nil {
		err = fmt.Errorf("Could not set organization name: %w", err)
		return "", err
	}
	secondPart, err := random.String(ctx, org.Type.NameSecondParts)
	if err != nil {
		err = fmt.Errorf("Could not set organization name: %w", err)
		return "", err
	}
	nameData := NameData{firstPart, secondPart}

	tmpl, err := template.New("orgname").Parse(org.Type.NameTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&tplOutput, nameData)
	if err != nil {
		panic(err)
	}
	name := tplOutput.String()

	return name, nil
}

func (org Organization) setTrait(ctx context.Context) (string, error) {
	trait, err := random.String(ctx, org.Type.PossibleTraits)
	if err != nil {
		err = fmt.Errorf("Could not set organization trait: %w", err)
		return "", err
	}
	return trait, nil
}

func (org Organization) getLeader(ctx context.Context, originCulture culture.Culture) (Member, error) {
	leaderData, err := character.Generate(ctx, originCulture)
	if err != nil {
		err = fmt.Errorf("Could not set organization leader: %w", err)
		return Member{}, err
	}
	leaderRank := org.Type.Ranks[0]
	leaderData = leaderData.ChangeAge(ctx, GetModifiedMemberAge(ctx, leaderRank, leaderData.Race))
	prof, err := org.Type.GetRandomMemberProfession(ctx)
	if err != nil {
		err = fmt.Errorf("Failed to get notable organization members: %w", err)
		return Member{}, err
	}
	leaderData.Profession = prof
	leaderData.Title = org.Type.LeaderTitle

	leader := Member{
		CharacterData: leaderData,
		Rank:          leaderRank,
	}

	return leader, nil
}

func (org Organization) getNotableMembers(ctx context.Context, originCulture culture.Culture) ([]Member, error) {
	var memberData Member
	var memberRank Rank
	members := []Member{}

	maxMemberCount := org.Size - 1
	if maxMemberCount > 10 {
		maxMemberCount = random.Intn(ctx, 9) + 1
	}

	for i := 0; i < maxMemberCount; i++ {
		member, err := character.Generate(ctx, originCulture)
		if err != nil {
			err = fmt.Errorf("Failed to get notable organization members: %w", err)
			return []Member{}, err
		}
		memberRank = org.Type.GetRandomMemberRank(ctx, members)
		member = member.ChangeAge(ctx, GetModifiedMemberAge(ctx, memberRank, member.Race))
		prof, err := org.Type.GetRandomMemberProfession(ctx)
		if err != nil {
			err = fmt.Errorf("Failed to get notable organization members: %w", err)
			return []Member{}, err
		}
		member.Profession = prof
		memberRank = org.Type.GetRandomMemberRank(ctx, members)
		memberData = Member{
			CharacterData: member,
			Rank:          memberRank,
		}
		members = append(members, memberData)
	}

	return members, nil
}

// Generate generates a org
func Generate(ctx context.Context, originCulture culture.Culture) (Organization, error) {
	org := Organization{}

	orgType, err := getRandomType(ctx)
	if err != nil {
		err = fmt.Errorf("Failed to generate organization: %w", err)
		return Organization{}, err
	}
	org.Type = orgType
	org.SizeClass = getRandomSizeClass(ctx, org.Type.MinSize, org.Type.MaxSize)
	org.Size = random.Intn(ctx, org.SizeClass.MaxSize-org.SizeClass.MinSize) + org.SizeClass.MinSize
	name, err := org.setName(ctx)
	if err != nil {
		err = fmt.Errorf("Failed to generate organization: %w", err)
		return Organization{}, err
	}
	org.Name = name
	org.LeaderType = org.setLeaderType(ctx)
	leader, err := org.getLeader(ctx, originCulture)
	if err != nil {
		err = fmt.Errorf("Failed to generate organization: %w", err)
		return Organization{}, err
	}
	org.Leader = leader
	trait, err := org.setTrait(ctx)
	if err != nil {
		err = fmt.Errorf("Failed to generate organization: %w", err)
		return Organization{}, err
	}
	org.PrimaryTrait = trait
	notableMembers, err := org.getNotableMembers(ctx, originCulture)
	if err != nil {
		err = fmt.Errorf("Failed to generate organization: %w", err)
		return Organization{}, err
	}
	org.NotableMembers = notableMembers

	return org, nil
}

// Random generates a completely random organization
func Random(ctx context.Context) (Organization, error) {
	originCulture, err := culture.Random(ctx)
	if err != nil {
		err = fmt.Errorf("Failed to generate random organization: %w", err)
		return Organization{}, err
	}

	org, err := Generate(ctx, originCulture)
	if err != nil {
		err = fmt.Errorf("Failed to generate random organization: %w", err)
		return Organization{}, err
	}

	return org, nil
}

// GetModifiedMemberAge returns an age appropriate for the given rank and race
func GetModifiedMemberAge(ctx context.Context, rank Rank, memberRace species.Species) int {
	var randomAgeCategory string

	if len(rank.AgeCategories) == 1 {
		randomAgeCategory = rank.AgeCategories[0]
	} else {
		randomAgeCategory = rank.AgeCategories[random.Intn(ctx, len(rank.AgeCategories))]
	}

	ageCategory := age.GetCategoryByName(randomAgeCategory, memberRace.AgeCategories)
	age := age.GetRandomAge(ctx, ageCategory)
	modifiedAge := rank.AgeModifier * float64(age)

	return int(modifiedAge)
}
