package language

import (
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

func randomSyllable(category Category, role string) string {
	syllable := random.String(category.Initiators) + getRandomWeightedVowel()
	expand := rand.Intn(10) + 1
	if expand > 2 {
		if role == "connector" {
			syllable += random.String(category.Connectors)
		} else {
			syllable += random.String(category.Finishers)
		}
	}

	return syllable
}

func (language Language) randomWord(maxSyllables int) string {
	var word string
	var syllables []string
	numSyllables := 1

	role := "connector"
	syllable := ""
	shouldIUseAnApostrophe := 0

	if maxSyllables > 1 {
		numSyllables = rand.Intn(maxSyllables) + 1
	}

	for i := 0; i < numSyllables; i++ {
		if numSyllables-i == 1 {
			role = "finisher"
		}
		syllable = randomSyllable(language.Category, role)

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

	return word
}
