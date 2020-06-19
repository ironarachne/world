/*
Package language provides structures and methods for dealing with languages, whether
fictional or otherwise, in fantasy worlds.
*/
package language

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
	"github.com/ironarachne/world/pkg/writing"
)

// Language is a spoken language
type Language struct {
	Name                    string            `json:"name"`
	Adjective               string            `json:"adjective"`
	Descriptors             []string          `json:"descriptors"`
	Description             string            `json:"description"`
	SamplePhrase            string            `json:"sample_phrase"`
	SamplePhraseTranslation string            `json:"sample_phrase_translation"`
	WritingSystem           writing.System    `json:"writing_system"`
	WordList                map[string]string `json:"word_list"`
	FemaleFirstNames        []string          `json:"female_first_names"`
	MaleFirstNames          []string          `json:"male_first_names"`
	FamilyNames             []string          `json:"family_names"`
	TownNames               []string          `json:"town_names"`
	VerbConjugationRules    ConjugationRules  `json:"conjugation_rules"`
	IsTonal                 bool              `json:"is_tonal"`
	NewWordPrefixes         []string          `json:"new_word_prefixes"`
	NewWordSuffixes         []string          `json:"new_word_suffixes"`
}

// RandomFemaleFirstName returns a random female first name
func (language Language) RandomFemaleFirstName(ctx context.Context) (string, error) {
	firstName, err := random.String(ctx, language.FemaleFirstNames)
	if err != nil {
		err = fmt.Errorf("could not select random female first name: %w", err)
		return "", err
	}

	return firstName, nil
}

// RandomMaleFirstName returns a random male first name
func (language Language) RandomMaleFirstName(ctx context.Context) (string, error) {
	firstName, err := random.String(ctx, language.MaleFirstNames)
	if err != nil {
		err = fmt.Errorf("could not select random male first name: %w", err)
		return "", err
	}

	return firstName, nil
}

// RandomFamilyName returns a random male first name
func (language Language) RandomFamilyName(ctx context.Context) (string, error) {
	familyName, err := random.String(ctx, language.FamilyNames)
	if err != nil {
		err = fmt.Errorf("could not select random family name: %w", err)
		return "", err
	}

	return familyName, nil
}

// RandomTownName returns a random male first name
func (language Language) RandomTownName(ctx context.Context) (string, error) {
	townName, err := random.String(ctx, language.TownNames)
	if err != nil {
		err = fmt.Errorf("could not select random town name: %w", err)
		return "", err
	}

	return townName, nil
}

// RandomNameList returns a list of N unique names of the given type
func (language Language) RandomNameList(ctx context.Context, numberOfNames int, nameType string) ([]string, error) {
	var name string
	var names []string
	var err error

	if nameType == "female" {
		for i := 0; i < numberOfNames; i++ {
			name, err = language.RandomFemaleFirstName(ctx)
			if err != nil {
				err = fmt.Errorf("could not generate name list: %w", err)
				return []string{}, err
			}
			if !slices.StringIn(name, names) {
				names = append(names, name)
			} else {
				i--
			}
		}
	} else if nameType == "male" {
		for i := 0; i < numberOfNames; i++ {
			name, err = language.RandomMaleFirstName(ctx)
			if err != nil {
				err = fmt.Errorf("could not generate name list: %w", err)
				return []string{}, err
			}
			if !slices.StringIn(name, names) {
				names = append(names, name)
			} else {
				i--
			}
		}
	} else if nameType == "family" {
		for i := 0; i < numberOfNames; i++ {
			name, err = language.RandomFamilyName(ctx)
			if err != nil {
				err = fmt.Errorf("could not generate name list: %w", err)
				return []string{}, err
			}
			if !slices.StringIn(name, names) {
				names = append(names, name)
			} else {
				i--
			}
		}
	} else if nameType == "town" {
		for i := 0; i < numberOfNames; i++ {
			name, err = language.RandomTownName(ctx)
			if err != nil {
				err = fmt.Errorf("could not generate name list: %w", err)
				return []string{}, err
			}
			if !slices.StringIn(name, names) {
				names = append(names, name)
			} else {
				i--
			}
		}
	}

	return names, nil
}

// NewWord returns a new word for the language using random components
func (language Language) NewWord(ctx context.Context) (string, error) {
	prefix, err := random.String(ctx, language.NewWordPrefixes)
	if err != nil {
		err = fmt.Errorf("could not generate new word: %w", err)
		return "", err
	}
	suffix, err := random.String(ctx, language.NewWordSuffixes)
	if err != nil {
		err = fmt.Errorf("could not generate new word: %w", err)
		return "", err
	}
	word := prefix + suffix

	return word, nil
}
