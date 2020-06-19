package conlang

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

const syllableError = "failed to generate syllable: %w"

// Mutation is a word mutation
type Mutation struct {
	From string
	To   string
}

func getRandomWeightedVowel(ctx context.Context) (string, error) {
	vowels := map[string]int{
		"a": 18,
		"e": 20,
		"i": 13,
		"o": 6,
		"u": 2,
	}

	vowel, err := random.StringFromThresholdMap(ctx, vowels)
	if err != nil {
		err = fmt.Errorf("failed to get random vowel: %w", err)
		return "", err
	}
	return vowel, nil
}

func randomMutation(ctx context.Context) Mutation {
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

	return rules[random.Intn(ctx, len(rules))]
}

func randomSyllable(ctx context.Context, category Category, role string) (string, error) {
	syllable, err := random.String(ctx, category.Initiators)
	if err != nil {
		err = fmt.Errorf(syllableError, err)
		return "", err
	}
	vowel, err := getRandomWeightedVowel(ctx)
	if err != nil {
		err = fmt.Errorf(syllableError, err)
		return "", err
	}
	syllable += vowel
	expand := random.Intn(ctx, 10) + 1
	if expand > 2 {
		if role == "connector" {
			connector, err := random.String(ctx, category.Connectors)
			if err != nil {
				err = fmt.Errorf(syllableError, err)
				return "", err
			}
			syllable += connector
		} else {
			finisher, err := random.String(ctx, category.Finishers)
			if err != nil {
				err = fmt.Errorf(syllableError, err)
				return "", err
			}
			syllable += finisher
		}
	}

	return syllable, nil
}

func randomWord(ctx context.Context, langCategory Category, maxSyllables int) (string, error) {
	var word string
	var syllables []string
	numSyllables := 1

	role := "connector"
	shouldIUseAnApostrophe := 0

	if maxSyllables > 1 {
		numSyllables = random.Intn(ctx, maxSyllables) + 1
	}

	for i := 0; i < numSyllables; i++ {
		if numSyllables-i == 1 {
			role = "finisher"
		}
		syllable, err := randomSyllable(ctx, langCategory, role)
		if err != nil {
			err = fmt.Errorf("Could not generate word: %w", err)
			return "", err
		}

		if langCategory.UsesApostrophes {
			shouldIUseAnApostrophe = random.Intn(ctx, 10)
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
