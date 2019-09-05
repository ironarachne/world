package language

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
)

// Mutation is a word mutation
type Mutation struct {
	From string
	To   string
}

func getRandomWeightedVowel() string {
	vowels := map[string]int{
		"a": 18,
		"e": 20,
		"i": 13,
		"o": 6,
		"u": 2,
	}

	return random.StringFromThresholdMap(vowels)
}

func randomMutation() Mutation {
	rules := []Mutation{
		{
			"s",
			"ss",
		},
		{
			"s",
			"sh",
		},
		{
			"f",
			"ff",
		},
		{
			"f",
			"fh",
		},
		{
			"g",
			"gh",
		},
		{
			"l",
			"l'",
		},
	}

	return rules[rand.Intn(len(rules)-1)]
}

func randomSyllable(category Category, role string) (string, error) {
	syllable, err := random.String(category.Initiators)
	if err != nil {
		err = fmt.Errorf("Could not generate syllable: %w", err)
		return "", err
	}
	syllable += getRandomWeightedVowel()
	expand := rand.Intn(10) + 1
	if expand > 2 {
		if role == "connector" {
			connector, err := random.String(category.Connectors)
			if err != nil {
				err = fmt.Errorf("Could not generate syllable: %w", err)
				return "", err
			}
			syllable += connector
		} else {
			finisher, err := random.String(category.Finishers)
			if err != nil {
				err = fmt.Errorf("Could not generate syllable: %w", err)
				return "", err
			}
			syllable += finisher
		}
	}

	return syllable, nil
}

func (language Language) randomWord(maxSyllables int) (string, error) {
	var word string
	var syllables []string
	numSyllables := 1

	role := "connector"
	shouldIUseAnApostrophe := 0

	if maxSyllables > 1 {
		numSyllables = rand.Intn(maxSyllables) + 1
	}

	for i := 0; i < numSyllables; i++ {
		if numSyllables-i == 1 {
			role = "finisher"
		}
		syllable, err := randomSyllable(language.Category, role)
		if err != nil {
			err = fmt.Errorf("Could not generate word: %w", err)
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
		word += syllable
	}

	return word, nil
}
