package character

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/ironarachne/world/pkg/words"
)

// SimplifiedCharacter is a simplified version of a character
type SimplifiedCharacter struct {
	Name        string `json:"name"`
	Titles      string `json:"titles"`
	Blazon      string `json:"blazon"`
	Device      string `json:"device"`
	Description string `json:"description"`
}

// Simplify returns a simplified version of a character
func (character Character) Simplify() (SimplifiedCharacter, error) {
	description, err := character.Describe()
	if err != nil {
		err = fmt.Errorf("Could not generate face shape: %w", err)
		return SimplifiedCharacter{}, err
	}

	titles := words.CombinePhrases(character.Titles)

	simplified := SimplifiedCharacter{
		Name:        character.FirstName + " " + character.LastName,
		Titles:      titles,
		Blazon:      character.Heraldry.Blazon,
		Description: description,
		Device:      character.Heraldry.ImageURL,
	}

	if character.Title != "" {
		simplified.Name = words.Title(character.Title) + " " + simplified.Name
	}

	return simplified, nil
}

// RandomSimplified returns a random simplified character
func RandomSimplified() (SimplifiedCharacter, error) {
	character, err := Random()
	if err != nil {
		err = fmt.Errorf("Could not generate simplified character: %w", err)
		return SimplifiedCharacter{}, err
	}

	simplified, err := character.Simplify()
	if err != nil {
		err = fmt.Errorf("Could not generate simplified character: %w", err)
		return SimplifiedCharacter{}, err
	}

	return simplified, nil
}

// Describe returns a prose description of a character based on his or her traits and attributes
func (character Character) Describe() (string, error) {
	descriptionObject, err := character.compileDescription()
	if err != nil {
		err = fmt.Errorf("Could not generate character description: %w", err)
		return "", err
	}
	descriptionTemplate, err := randomDescriptionTemplate()
	if err != nil {
		err = fmt.Errorf("Could not generate character description: %w", err)
		return "", err
	}

	var tplOutput bytes.Buffer

	tmpl, err := template.New(descriptionObject.FullName).Funcs(template.FuncMap{
		"caseStart": func(word string) string {
			return words.Title(word)
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

	return description, nil
}
