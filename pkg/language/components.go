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
		Mutation{
			"s",
			"ss",
		},
		Mutation{
			"s",
			"sh",
		},
		Mutation{
			"f",
			"ff",
		},
		Mutation{
			"f",
			"fh",
		},
		Mutation{
			"g",
			"gh",
		},
		Mutation{
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
