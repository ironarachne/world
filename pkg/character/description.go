package character

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/ironarachne/world/pkg/measurement"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/words"
)

// Description is a description object of a character
type Description struct {
	Age               string
	Culture           string
	FirstName         string
	FullName          string
	GenderNoun        string
	Height            string
	Heraldry          string
	Hobby             string
	LastName          string
	Motivation        string
	NegativeTraits    string
	PositiveTraits    string
	PossessivePronoun string
	Profession        string
	Race              string
	SubjectPronoun    string
	Traits            string
	Weight            string
}

func allDescriptionTemplates() []string {
	templates := []string{
		`{{.FullName}} is a {{.Age}}-year-old {{.Culture}} {{.Race}} {{.GenderNoun}}.
		{{caseStart .SubjectPronoun}} has {{.Traits}}.
		{{.FirstName}} is {{.Height}} tall and weighs {{.Weight}} lbs. {{caseStart .SubjectPronoun}} is motivated by {{.Motivation}}.
		While {{.SubjectPronoun}} is {{.PositiveTraits}}, {{.SubjectPronoun}} has also been described as {{.NegativeTraits}}.
		{{.FirstName}}'s hobby is {{.Hobby}} and {{.SubjectPronoun}} is {{pronoun .Profession}}.
		{{if .Heraldry}}{{caseStart .PossessivePronoun}} coat of arms is described "{{.Heraldry}}."{{end}}
		`,
		`{{.FullName}} is {{pronoun .Race}} {{.GenderNoun}} of {{.Age}} years. {{caseStart .SubjectPronoun}} is {{.Height}} and weighs {{.Weight}} pounds, with
		{{.Traits}}. Motivated by {{.Motivation}}, {{.FirstName}} is {{.PositiveTraits}}, as well as {{.NegativeTraits}}.
		{{caseStart .SubjectPronoun}} is {{pronoun .Profession}}{{if .Heraldry}} and has a coat of arms of "{{.Heraldry}}." {{else}}.{{end}}`,
		`The {{.PositiveTraits}} {{.FullName}} is {{.Age}} years old and {{pronoun .Race}} {{.GenderNoun}}. {{caseStart .SubjectPronoun}} is {{pronoun .Profession}},
		and is driven by {{.Motivation}}. {{caseStart .SubjectPronoun}} is {{.Height}} tall and weighs {{.Weight}} lbs. Despite a generally positive perception, some
		describe {{.FirstName}} as {{.NegativeTraits}}. {{caseStart .SubjectPronoun}} has {{.Traits}}.`,
	}

	return templates
}

func randomDescriptionTemplate(ctx context.Context) (string, error) {
	all := allDescriptionTemplates()

	template, err := random.String(ctx, all)
	if err != nil {
		err = fmt.Errorf("Could not generate description template: %w", err)
		return "", err
	}
	return template, nil
}

func (character Character) compileDescription() (Description, error) {
	description := Description{}

	description.Age = character.describeAge()
	description.Culture = character.describeCulture()
	description.FirstName = character.FirstName
	description.FullName = character.describeFullName()
	description.GenderNoun = character.describeGenderNoun()
	description.Height = character.describeHeight()
	description.Heraldry = character.describeHeraldry()
	description.Hobby = character.describeHobby()
	description.LastName = character.LastName
	description.Motivation = character.describeMotivation()
	description.NegativeTraits = character.describeNegativeTraits()
	description.PositiveTraits = character.describePositiveTraits()
	description.PossessivePronoun = character.Gender.PossessivePronoun
	description.Profession = character.describeProfession()
	description.Race = character.describeRace()
	description.SubjectPronoun = character.Gender.SubjectPronoun
	traits, err := character.describeTraits()
	if err != nil {
		err = fmt.Errorf("Could not generate description template: %w", err)
		return Description{}, err
	}
	description.Traits = traits
	description.Weight = character.describeWeight()

	return description, nil
}

func (character Character) describeAge() string {
	description := strconv.Itoa(character.Age)

	return description
}

func (character Character) describeCulture() string {
	description := character.Culture.Adjective

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

func (character Character) describeHeight() string {
	description := measurement.ToString(character.Height)

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

func (character Character) describeNegativeTraits() string {
	description := words.CombinePhrases(character.NegativeTraits)

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

func (character Character) describeTraits() (string, error) {
	traits := []string{}

	for _, i := range character.PhysicalTraits {
		t, err := i.ToString()
		if err != nil {
			err = fmt.Errorf("Failed to describe traits: %w", err)
			return "", err
		}
		traits = append(traits, t)
	}

	description := words.CombinePhrases(traits)

	return description, nil
}

func (character Character) describeWeight() string {
	description := strconv.Itoa(character.Weight)

	return description
}
