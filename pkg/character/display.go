package character

import (
	"bytes"
	"strings"
	"text/template"

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
	descriptionObject := character.compileDescription()
	descriptionTemplate := randomDescriptionTemplate()

	var tplOutput bytes.Buffer

	tmpl, err := template.New(descriptionObject.FullName).Funcs(template.FuncMap{
		"caseStart": func(word string) string {
			return strings.Title(word)
		},
		"pronoun": func(word string) string {
			phrase := words.Pronoun(word) + " " + word
			return phrase
		},
	}).Parse(descriptionTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&tplOutput, descriptionObject)
	if err != nil {
		panic(err)
	}
	description := tplOutput.String()

	return description
}
