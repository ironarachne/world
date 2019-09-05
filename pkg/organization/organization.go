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

// Organization is an organization
type Organization struct {
	Name           string
	Type           Type
	Leader         character.Character
	LeaderType     string
	NotableMembers []character.Character
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

func (org Organization) getLeader(originCulture culture.Culture) (character.Character, error) {
	leader, err := character.Generate(originCulture)
	if err != nil {
		err = fmt.Errorf("Could not set organization leader: %w", err)
		return character.Character{}, err
	}
	leader = leader.ChangeAge(org.Type.GetRandomLeaderAge())
	leader.Profession = org.Type.GetRandomMemberProfession()
	leader.Title = org.Type.LeaderTitle

	return leader, nil
}

func (org Organization) getNotableMembers(originCulture culture.Culture) ([]character.Character, error) {
	members := []character.Character{}

	maxMemberCount := org.Size - 1
	if maxMemberCount > 10 {
		maxMemberCount = rand.Intn(9) + 1
	}

	for i := 0; i < maxMemberCount; i++ {
		member, err := character.Generate(originCulture)
		if err != nil {
			err = fmt.Errorf("Could not set organization name: %w", err)
			return []character.Character{}, err
		}
		member = member.ChangeAge(org.Type.GetRandomMemberAge())
		member.Profession = org.Type.GetRandomMemberProfession()
		members = append(members, member)
	}

	return members, nil
}

// Generate generates a org
func Generate(originCulture culture.Culture) (Organization, error) {
	org := Organization{}

	org.Type = getRandomType()
	org.SizeClass = getRandomSizeClass()
	org.Size = rand.Intn(org.SizeClass.MaxSize-org.SizeClass.MinSize) + org.SizeClass.MinSize
	name, err := org.setName()
	if err != nil {
		err = fmt.Errorf("Could not generate organization: %w", err)
		return Organization{}, err
	}
	org.Name = name
	org.LeaderType = org.setLeaderType()
	leader, err := org.getLeader(originCulture)
	if err != nil {
		err = fmt.Errorf("Could not generate organization: %w", err)
		return Organization{}, err
	}
	org.Leader = leader
	trait, err := org.setTrait()
	if err != nil {
		err = fmt.Errorf("Could not generate organization: %w", err)
		return Organization{}, err
	}
	org.PrimaryTrait = trait
	notableMembers, err := org.getNotableMembers(originCulture)
	if err != nil {
		err = fmt.Errorf("Could not generate organization: %w", err)
		return Organization{}, err
	}
	org.NotableMembers = notableMembers

	return org, nil
}

// Random generates a completely random organization
func Random() (Organization, error) {
	originCulture, err := culture.Random()
	if err != nil {
		err = fmt.Errorf("Could not generate random organization: %w", err)
		return Organization{}, err
	}

	org, err := Generate(originCulture)
	if err != nil {
		err = fmt.Errorf("Could not generate random organization: %w", err)
		return Organization{}, err
	}

	return org, nil
}
