package organization

import (
	"bytes"
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

func (org Organization) setName() string {
	var tplOutput bytes.Buffer

	firstPart := random.String(org.Type.NameFirstParts)
	secondPart := random.String(org.Type.NameSecondParts)
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

	return name
}

func (org Organization) setTrait() string {
	return random.String(org.Type.PossibleTraits)
}

func (org Organization) getLeader(originCulture culture.Culture) character.Character {
	leader := character.Generate(originCulture)
	leader = leader.ChangeAge(org.Type.GetRandomLeaderAge())
	leader.Profession = org.Type.GetRandomMemberProfession()
	leader.Title = org.Type.LeaderTitle

	return leader
}

func (org Organization) getNotableMembers(originCulture culture.Culture) []character.Character {
	var member character.Character
	members := []character.Character{}

	maxMemberCount := org.Size - 1
	if maxMemberCount > 10 {
		maxMemberCount = rand.Intn(9) + 1
	}

	for i := 0; i < maxMemberCount; i++ {
		member = character.Generate(originCulture)
		member = member.ChangeAge(org.Type.GetRandomMemberAge())
		member.Profession = org.Type.GetRandomMemberProfession()
		members = append(members, member)
	}

	return members
}

// Generate generates a org
func Generate(originCulture culture.Culture) Organization {
	org := Organization{}

	org.Type = getRandomType()
	org.SizeClass = getRandomSizeClass()
	org.Size = rand.Intn(org.SizeClass.MaxSize-org.SizeClass.MinSize) + org.SizeClass.MinSize
	org.Name = org.setName()
	org.LeaderType = org.setLeaderType()
	org.Leader = org.getLeader(originCulture)
	org.PrimaryTrait = org.setTrait()
	org.NotableMembers = org.getNotableMembers(originCulture)

	return org
}

// Random generates a completely random organization
func Random() Organization {
	originCulture := culture.Generate()

	org := Generate(originCulture)

	return org
}
