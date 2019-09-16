package organization

import (
	"bytes"
	"fmt"
	"html/template"
	"math/rand"

	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/random"
)

// NameData is name data
type NameData struct {
	FirstPart  string
	SecondPart string
}

// Member is a member of an organization
type Member struct {
	CharacterData character.Character
	Rank          Rank
}

// Organization is an organization
type Organization struct {
	Name           string
	Type           Type
	Leader         Member
	LeaderType     string
	NotableMembers []Member
	SizeClass      SizeClass
	Size           int
	PrimaryTrait   string
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
	leaderData = leaderData.ChangeAge(org.Type.GetRandomLeaderAge())
	prof, err := org.Type.GetRandomMemberProfession()
	if err != nil {
		err = fmt.Errorf("Failed to get notable organization members: %w", err)
		return Member{}, err
	}
	leaderData.Profession = prof
	leaderData.Title = org.Type.LeaderTitle
	leaderRank := org.Type.Ranks[0]

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
		member = member.ChangeAge(org.Type.GetRandomMemberAge())
		prof, err := org.Type.GetRandomMemberProfession()
		if err != nil {
			err = fmt.Errorf("Failed to get notable organization members: %w", err)
			return []Member{}, err
		}
		member.Profession = prof
		memberRank = org.Type.GetRandomMemberRank()
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
	org.SizeClass = getRandomSizeClass()
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
