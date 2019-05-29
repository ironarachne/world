package character

import (
	"strconv"
	"strings"
)

// SimplifiedCharacter is a simplified version of a character
type SimplifiedCharacter struct {
	Name string `json:"name"`
	Blazon string `json:"blazon"`
	Device string `json:"device"`
	Description string `json:"description"`
}

// Simplify returns a simplified version of a character
func (character Character) Simplify() SimplifiedCharacter {
	simplified := SimplifiedCharacter{
		Name: character.FirstName + " " + character.LastName,
		Blazon: character.Heraldry.Blazon,
		Device: character.Heraldry.Device,
		Description: strconv.Itoa(character.Age) + "-year-old " + character.Gender.Noun + " who is " + character.PositiveTraits[0] + " but " + character.NegativeTraits[0],
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