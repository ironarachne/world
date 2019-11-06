package language

import (
	"fmt"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
	"github.com/ironarachne/world/pkg/writing"
)

// Language is a spoken language
type Language struct {
	Name                 string
	Adjective            string
	Descriptors          []string
	Description          string
	WritingSystem        writing.System
	WordList             map[string]string
	FemaleFirstNames     []string
	MaleFirstNames       []string
	FamilyNames          []string
	TownNames            []string
	VerbConjugationRules ConjugationRules
	IsTonal              bool
}

// RandomFemaleFirstName returns a random female first name
func (language Language) RandomFemaleFirstName() (string, error) {
	firstName, err := random.String(language.FemaleFirstNames)
	if err != nil {
		err = fmt.Errorf("could not select random female first name: %w", err)
		return "", err
	}

	return firstName, nil
}

// RandomMaleFirstName returns a random male first name
func (language Language) RandomMaleFirstName() (string, error) {
	firstName, err := random.String(language.MaleFirstNames)
	if err != nil {
		err = fmt.Errorf("could not select random male first name: %w", err)
		return "", err
	}

	return firstName, nil
}

// RandomFamilyName returns a random male first name
func (language Language) RandomFamilyName() (string, error) {
	familyName, err := random.String(language.FamilyNames)
	if err != nil {
		err = fmt.Errorf("could not select random family name: %w", err)
		return "", err
	}

	return familyName, nil
}

// RandomTownName returns a random male first name
func (language Language) RandomTownName() (string, error) {
	townName, err := random.String(language.TownNames)
	if err != nil {
		err = fmt.Errorf("could not select random town name: %w", err)
		return "", err
	}

	return townName, nil
}

// RandomNameList returns a list of N unique names of the given type
func (language Language) RandomNameList(numberOfNames int, nameType string) ([]string, error) {
	var name string
	var names []string
	var err error

	if nameType == "female" {
		for i := 0; i < numberOfNames; i++ {
			name, err = language.RandomFemaleFirstName()
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
			name, err = language.RandomMaleFirstName()
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
			name, err = language.RandomFamilyName()
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
			name, err = language.RandomTownName()
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
