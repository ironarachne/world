/*
Package organization implements social organizations and tools for dealing with them
*/
package organization

import (
	"bytes"
	"fmt"
	"html/template"
	"math/rand"

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

func (org Organization) setLeaderType() string {
	leaderType := "individual"

	if org.Type.CanBeLedByGroup {
		x := rand.Intn(10)

		if x >= 9 {
			leaderType = "group"
		}
	}

	return leaderType
}

func (org Organization) setName() (string, error) {
	var tplOutput bytes.Buffer

	firstPart, err := random.String(org.Type.NameFirstParts)
	if err != nil {
		err = fmt.Errorf("Could not set organization name: %w", err)
		return "", err
	}
	secondPart, err := random.String(org.Type.NameSecondParts)
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

func (org Organization) setTrait() (string, error) {
	trait, err := random.String(org.Type.PossibleTraits)
	if err != nil {
		err = fmt.Errorf("Could not set organization trait: %w", err)
		return "", err
	}
	return trait, nil
}

func (org Organization) getLeader(originCulture culture.Culture) (Member, error) {
	leaderData, err := character.Generate(originCulture)
	if err != nil {
		err = fmt.Errorf("Could not set organization leader: %w", err)
		return Member{}, err
	}
	leaderRank := org.Type.Ranks[0]
	leaderData = leaderData.ChangeAge(GetModifiedMemberAge(leaderRank, leaderData.Race))
	prof, err := org.Type.GetRandomMemberProfession()
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

func (org Organization) getNotableMembers(originCulture culture.Culture) ([]Member, error) {
	var memberData Member
	var memberRank Rank
	members := []Member{}

	maxMemberCount := org.Size - 1
	if maxMemberCount > 10 {
		maxMemberCount = rand.Intn(9) + 1
	}

	for i := 0; i < maxMemberCount; i++ {
		member, err := character.Generate(originCulture)
		if err != nil {
			err = fmt.Errorf("Failed to get notable organization members: %w", err)
			return []Member{}, err
		}
		memberRank = org.Type.GetRandomMemberRank(members)
		member = member.ChangeAge(GetModifiedMemberAge(memberRank, member.Race))
		prof, err := org.Type.GetRandomMemberProfession()
		if err != nil {
			err = fmt.Errorf("Failed to get notable organization members: %w", err)
			return []Member{}, err
		}
		member.Profession = prof
		memberRank = org.Type.GetRandomMemberRank(members)
		memberData = Member{
			CharacterData: member,
			Rank:          memberRank,
		}
		members = append(members, memberData)
	}

	return members, nil
}

// Generate generates a org
func Generate(originCulture culture.Culture) (Organization, error) {
	org := Organization{}

	orgType, err := getRandomType()
	if err != nil {
		err = fmt.Errorf("Failed to generate organization: %w", err)
		return Organization{}, err
	}
	org.Type = orgType
	org.SizeClass = getRandomSizeClass(org.Type.MinSize, org.Type.MaxSize)
	org.Size = rand.Intn(org.SizeClass.MaxSize-org.SizeClass.MinSize) + org.SizeClass.MinSize
	name, err := org.setName()
	if err != nil {
		err = fmt.Errorf("Failed to generate organization: %w", err)
		return Organization{}, err
	}
	org.Name = name
	org.LeaderType = org.setLeaderType()
	leader, err := org.getLeader(originCulture)
	if err != nil {
		err = fmt.Errorf("Failed to generate organization: %w", err)
		return Organization{}, err
	}
	org.Leader = leader
	trait, err := org.setTrait()
	if err != nil {
		err = fmt.Errorf("Failed to generate organization: %w", err)
		return Organization{}, err
	}
	org.PrimaryTrait = trait
	notableMembers, err := org.getNotableMembers(originCulture)
	if err != nil {
		err = fmt.Errorf("Failed to generate organization: %w", err)
		return Organization{}, err
	}
	org.NotableMembers = notableMembers

	return org, nil
}

// Random generates a completely random organization
func Random() (Organization, error) {
	originCulture, err := culture.Random()
	if err != nil {
		err = fmt.Errorf("Failed to generate random organization: %w", err)
		return Organization{}, err
	}

	org, err := Generate(originCulture)
	if err != nil {
		err = fmt.Errorf("Failed to generate random organization: %w", err)
		return Organization{}, err
	}

	return org, nil
}

// GetModifiedMemberAge returns an age appropriate for the given rank and race
func GetModifiedMemberAge(rank Rank, memberRace species.Species) int {
	var randomAgeCategory string

	if len(rank.AgeCategories) == 1 {
		randomAgeCategory = rank.AgeCategories[0]
	} else {
		randomAgeCategory = rank.AgeCategories[rand.Intn(len(rank.AgeCategories))]
	}

	ageCategory := age.GetCategoryByName(randomAgeCategory, memberRace.AgeCategories)
	age := age.GetRandomAge(ageCategory)
	modifiedAge := rank.AgeModifier * float64(age)

	return int(modifiedAge)
}
