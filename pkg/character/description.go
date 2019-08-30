package character

import (
	"strconv"
	"strings"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/words"
)

// Description is a description object of a character
type Description struct {
	Age               string
	Culture           string
	Eyes              string
	FacialHair        string
	FirstName         string
	FullName          string
	GenderNoun        string
	Hair              string
	Height            string
	Heraldry          string
	Hobby             string
	LastName          string
	Motivation        string
	Mouth             string
	NegativeTraits    string
	Nose              string
	PositiveTraits    string
	PossessivePronoun string
	Profession        string
	Race              string
	Skin              string
	SubjectPronoun    string
	Weight            string
}

func allDescriptionTemplates() []string {
	templates := []string{
		`{{.FullName}} is a {{.Age}}-year-old {{.Culture}} {{.Race}} {{.GenderNoun}} with {{.Hair}} hair and {{.Skin}} skin.
		{{caseStart .SubjectPronoun}} has {{.Eyes}} eyes, a {{.Nose}} nose, and a {{.Mouth}} mouth. {{if .FacialHair}}{{caseStart .SubjectPronoun}} has {{pronoun .FacialHair}}.{{end}}
		{{.FirstName}} is {{.Height}} tall and weighs {{.Weight}} lbs. {{caseStart .SubjectPronoun}} is motivated by {{.Motivation}}.
		While {{.SubjectPronoun}} is {{.PositiveTraits}}, {{.SubjectPronoun}} has also been described as {{.NegativeTraits}}.
		{{.FirstName}}'s hobby is {{.Hobby}} and {{.SubjectPronoun}} is {{pronoun .Profession}}.
		{{if .Heraldry}}{{.PossessivePronoun}} coat of arms is described "{{.Heraldry}}."{{end}}
		`,
		`{{.FullName}} is {{pronoun .Race}} {{.GenderNoun}} of {{.Age}} years. {{caseStart .SubjectPronoun}} is {{.Height}} and weighs {{.Weight}} lbs., with
		{{.Hair}} hair, {{.Eyes}} eyes, and {{.Skin}} skin. Motivated by {{.Motivation}}, {{.FirstName}} is {{.PositiveTraits}}, as well as {{.NegativeTraits}}.
		{{caseStart .SubjectPronoun}} is {{pronoun .Profession}}{{if .Heraldry}} and has a coat of arms of "{{.Heraldry}}." {{else}}.{{end}}`,
		`The {{.PositiveTraits}} {{.FullName}} is {{.Age}} years old and {{pronoun .Race}} {{.GenderNoun}}. {{caseStart .SubjectPronoun}} is {{pronoun .Profession}},
		and seeks {{.Motivation}}. {{caseStart .SubjectPronoun}} is {{.Height}} tall and weighs {{.Weight}} pounds. Despite a generally positive perception, some
		describe {{.FirstName}} as {{.NegativeTraits}}. {{caseStart .SubjectPronoun}} has a {{.Nose}} nose, a {{.Mouth}} mouth, and {{.Skin}} skin.
		{{caseStart .PossessivePronoun}} eyes are {{.Eyes}}.{{if .FacialHair}} {{caseStart .SubjectPronoun}} has {{pronoun .FacialHair}}.{{end}}`,
	}

	return templates
}

func randomDescriptionTemplate() string {
	all := allDescriptionTemplates()

	return random.String(all)
}

func (character Character) compileDescription() Description {
	description := Description{}

	description.Age = character.describeAge()
	description.Culture = character.describeCulture()
	description.Eyes = character.describeEyes()
	description.FacialHair = character.describeFacialHair()
	description.FirstName = character.FirstName
	description.FullName = character.describeFullName()
	description.GenderNoun = character.describeGenderNoun()
	description.Hair = character.describeHair()
	description.Height = character.describeHeight()
	description.Heraldry = character.describeHeraldry()
	description.Hobby = character.describeHobby()
	description.LastName = character.LastName
	description.Motivation = character.describeMotivation()
	description.Mouth = character.describeMouth()
	description.NegativeTraits = character.describeNegativeTraits()
	description.Nose = character.describeNose()
	description.PositiveTraits = character.describePositiveTraits()
	description.PossessivePronoun = character.Gender.PossessivePronoun
	description.Profession = character.describeProfession()
	description.Race = character.describeRace()
	description.Skin = character.describeSkin()
	description.SubjectPronoun = character.Gender.SubjectPronoun
	description.Weight = character.describeWeight()

	return description
}

func (character Character) describeAge() string {
	description := strconv.Itoa(character.Age)

	return description
}

func (character Character) describeCulture() string {
	description := character.Culture.Adjective

	return description
}

func (character Character) describeEyes() string {
	eyes := character.EyeColor + " and " + character.EyeShape

	return eyes
}

func (character Character) describeFacialHair() string {
	description := ""

	if character.FacialHair != "none" {
		description = character.FacialHair
	}

	return description
}

func (character Character) describeFullName() string {
	description := ""

	if character.Title != "" {
		description += strings.Title(character.Title) + " "
	}

	description += character.FirstName + " " + character.LastName

	return description
}

func (character Character) describeGenderNoun() string {
	description := ""

	if character.AgeCategory.Name == "child" || character.AgeCategory.Name == "infant" {
		description = character.Gender.AdolescentNoun
	} else {
		description = character.Gender.Noun
	}

	return description
}

func (character Character) describeHair() string {
	description := character.HairColor + " " + character.HairStyle

	return description
}

func (character Character) describeHeight() string {
	description := character.HeightSimplified()

	return description
}

func (character Character) describeHeraldry() string {
	description := character.Heraldry.Blazon

	return description
}

func (character Character) describeHobby() string {
	description := character.Hobby.Name

	return description
}

func (character Character) describeMotivation() string {
	description := character.Motivation

	return description
}

func (character Character) describeMouth() string {
	description := character.MouthShape

	return description
}

func (character Character) describeNegativeTraits() string {
	description := words.CombinePhrases(character.NegativeTraits)

	return description
}

func (character Character) describeNose() string {
	description := character.NoseShape

	return description
}

func (character Character) describePositiveTraits() string {
	description := words.CombinePhrases(character.PositiveTraits)

	return description
}

func (character Character) describeProfession() string {
	description := character.Profession.Name

	return description
}

func (character Character) describeRace() string {
	description := character.Race.Adjective

	return description
}

func (character Character) describeSkin() string {
	description := character.SkinColor

	return description
}

func (character Character) describeWeight() string {
	description := strconv.Itoa(character.Weight)

	return description
}
