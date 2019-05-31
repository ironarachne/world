package character

import (
	"strconv"
	"strings"

	"github.com/ironarachne/world/pkg/words"
)

// SimplifiedCharacter is a simplified version of a character
type SimplifiedCharacter struct {
	Name        string `json:"name"`
	Blazon      string `json:"blazon"`
	Device      string `json:"device"`
	Description string `json:"description"`
}

// Simplify returns a simplified version of a character
func (character Character) Simplify() SimplifiedCharacter {
	simplified := SimplifiedCharacter{
		Name:        character.FirstName + " " + character.LastName,
		Blazon:      character.Heraldry.Blazon,
		Device:      character.Heraldry.Device,
		Description: character.Describe(),
	}

	if character.Title != "" {
		simplified.Name = strings.Title(character.Title) + " " + simplified.Name
	}

	return simplified
}

// RandomSimplified returns a random simplified character
func RandomSimplified() SimplifiedCharacter {
	character := Random()

	return character.Simplify()
}

// Describe returns a prose description of a character based on his or her traits and attributes
func (character Character) Describe() string {
	description := ""
	noun := ""

	if character.Title != "" {
		description += strings.Title(character.Title) + " "
	}

	if character.AgeCategory.Name == "child" || character.AgeCategory.Name == "infant" {
		noun = character.Gender.AdolescentNoun
	} else {
		noun = character.Gender.Noun
	}

	description += character.FirstName + " " + character.LastName + " is a " + strconv.Itoa(character.Age) + "-year-old " + character.Culture.Adjective + " " + character.Race.Adjective + " " + noun + " with "
	description += character.HairColor + ", " + character.HairStyle + " hair and " + character.SkinColor + " skin. "
	if character.FacialHair != "none" {
		description += strings.Title(character.Gender.SubjectPronoun) + " has a " + character.FacialHair + ". "
	}
	description += strings.Title(character.Gender.SubjectPronoun) + " has " + character.EyeColor + " " + character.EyeShape + " eyes, a " + character.NoseShape + " nose, and a "
	description += character.MouthShape + " mouth. " + character.FirstName + " is " + character.HeightSimplified() + " tall and weighs " + strconv.Itoa(character.Weight) + " lbs. "
	description += strings.Title(character.Gender.SubjectPronoun) + " is motivated by " + character.Motivation + ". "
	description += "While " + character.Gender.SubjectPronoun + " is " + words.CombinePhrases(character.PositiveTraits) + ", "
	description += character.Gender.SubjectPronoun + " has also been described as " + words.CombinePhrases(character.NegativeTraits) + ". "
	description += character.FirstName + "'s hobby is " + character.Hobby.Name + ", and " + character.Gender.SubjectPronoun + " is " + words.Pronoun(character.Profession) + " " + character.Profession + ". "

	if character.Heraldry.Blazon != "" {
		description += strings.Title(character.Gender.PossessivePronoun) + " coat of arms is described as \"" + character.Heraldry.Blazon + ".\""
	}

	return description
}
