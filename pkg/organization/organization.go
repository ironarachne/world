package organization

import (
	"bytes"
	"html/template"
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
)

// NameData is name data
type NameData struct {
	FirstPart  string
	SecondPart string
}

// Organization is an organization
type Organization struct {
	Name         string
	Type         OrganizationType
	LeaderType   string
	SizeClass    SizeClass
	Size         int
	PrimaryTrait string
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

// Generate generates a org
func Generate() Organization {
	org := Organization{}

	org.Type = getRandomType()
	org.SizeClass = getRandomSizeClass()
	org.Size = rand.Intn(org.SizeClass.MaxSize-org.SizeClass.MinSize) + org.SizeClass.MinSize
	org.Name = org.setName()
	org.LeaderType = org.setLeaderType()
	org.PrimaryTrait = org.setTrait()

	return org
}
