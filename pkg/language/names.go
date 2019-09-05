package language

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

// GenerateNameList generates a list of names appropriate for the language
func (language Language) GenerateNameList(nameType string) ([]string, error) {
	var names []string
	var endings []string

	if nameType == "male" {
		endings = language.Category.MasculineEndings
	} else if nameType == "female" {
		endings = language.Category.FeminineEndings
	} else {
		for i := 0; i < 5; i++ {
			finisher, err := randomSyllable(language.Category, "finisher")
			if err != nil {
				err = fmt.Errorf("Could not generate name list: %w", err)
				return []string{}, err
			}
			endings = append(endings, finisher)
		}
	}

	for i := 0; i < 10; i++ {
		ending, err := random.String(endings)
		if err != nil {
			err = fmt.Errorf("Could not generate name list: %w", err)
			return []string{}, err
		}
		name, err := language.RandomName()
		if err != nil {
			err = fmt.Errorf("Could not generate name list: %w", err)
			return []string{}, err
		}
		name += ending
		if !slices.StringIn(name, names) {
			names = append(names, name)
		}
	}

	return names, nil
}

func mutateName(name string) string {
	mutation := randomMutation()

	name = strings.Replace(name, mutation.From, mutation.To, 1)

	return name
}

func randomLanguageName(category Category) (string, error) {
	var name string
	var syllables []string
	skewLonger := false

	if rand.Intn(10) > 3 {
		skewLonger = true
	}

	randomLength := rand.Intn(category.WordLength) + 1

	if skewLonger {
		randomLength++
	}

	role := "connector"

	for i := 0; i < randomLength; i++ {
		if randomLength-i == 1 {
			role = "finisher"
		}
		syllable, err := randomSyllable(category, role)
		if err != nil {
			err = fmt.Errorf("Could not generate language name: %w", err)
			return "", err
		}
		syllables = append(syllables, syllable)
	}

	for _, syllable := range syllables {
		name += syllable
	}

	chance := rand.Intn(10) + 1
	if chance > 3 {
		name = mutateName(name)
	}

	return name, nil
}

// RandomGenderedName generates a random gendered first name
func (language Language) RandomGenderedName(gender string) (string, error) {
	var endings []string

	name, err := language.RandomName()
	if err != nil {
		err = fmt.Errorf("Could not generate random gendered name: %w", err)
		return "", err
	}

	if gender == "male" {
		endings = language.Category.MasculineEndings
	} else {
		endings = language.Category.FeminineEndings
	}

	ending, err := random.String(endings)
	if err != nil {
		err = fmt.Errorf("Could not generate random gendered name: %w", err)
		return "", err
	}
	name = name + ending

	return name, nil
}

// RandomName generates a random name using the language
func (language Language) RandomName() (string, error) {
	var name string
	var syllables []string
	skewLonger := false

	if rand.Intn(10) > 7 {
		skewLonger = true
	}

	randomLength := rand.Intn(language.Category.WordLength) + 1

	if skewLonger {
		randomLength++
	}

	role := "connector"
	shouldIUseAnApostrophe := 0

	for i := 0; i < randomLength; i++ {
		if randomLength-i == 1 {
			role = "finisher"
		}
		syllable, err := randomSyllable(language.Category, role)
		if err != nil {
			err = fmt.Errorf("Could not generate conjugation rules: %w", err)
			return "", err
		}

		if language.Category.UsesApostrophes {
			shouldIUseAnApostrophe = rand.Intn(10)
			if shouldIUseAnApostrophe > 8 {
				syllable += "'"
			}
		}

		syllables = append(syllables, syllable)
	}

	for _, syllable := range syllables {
		name += syllable
	}

	chance := rand.Intn(10) + 1
	if chance > 8 {
		name = mutateName(name)
	}

	name = strings.Title(name)

	return name, nil
}
