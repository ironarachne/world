/*
Package gender provides structures, tools, and methods for dealing with
gender, including description and pronoun generation.
*/
package gender

import (
	"context"

	"github.com/ironarachne/world/pkg/random"
)

// Gender is a gender and its accompanying descriptors
type Gender struct {
	Name                 string `json:"name"`
	Noun                 string `json:"noun"`
	PluralNoun           string `json:"plural_noun"`
	AdolescentNoun       string `json:"adolescent_noun"`
	PluralAdolescentNoun string `json:"plural_adolescent_noun"`
	ObjectPronoun        string `json:"object_pronoun"`
	PossessivePronoun    string `json:"possessive_pronoun"`
	ReflexivePronoun     string `json:"reflexive_pronoun"`
	SubjectPronoun       string `json:"subject_pronoun"`
}

// All returns all genders
func All() []Gender {
	genders := []Gender{}
	genders = append(genders, Female())
	genders = append(genders, Male())

	return genders
}

// Female returns the female gender
func Female() Gender {
	gender := Gender{
		Name:                 "female",
		Noun:                 "woman",
		PluralNoun:           "women",
		AdolescentNoun:       "girl",
		PluralAdolescentNoun: "girls",
		ObjectPronoun:        "her",
		PossessivePronoun:    "her",
		ReflexivePronoun:     "herself",
		SubjectPronoun:       "she",
	}

	return gender
}

// Male returns the male gender
func Male() Gender {
	gender := Gender{
		Name:                 "male",
		Noun:                 "man",
		PluralNoun:           "men",
		AdolescentNoun:       "boy",
		PluralAdolescentNoun: "boys",
		ObjectPronoun:        "her",
		PossessivePronoun:    "his",
		ReflexivePronoun:     "himself",
		SubjectPronoun:       "he",
	}

	return gender
}

// Opposite returns the opposite gender from a given gender
func (gender Gender) Opposite() Gender {
	if gender.Name == "male" {
		return Female()
	}

	return Male()
}

// Random returns a random gender
func Random(ctx context.Context) Gender {
	genders := All()

	return genders[random.Intn(ctx, len(genders))]
}
