package conlang

import (
	"context"
	"fmt"
	"strings"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

const nameListError = "failed to generate name list: %w"

// GenerateNameList generates a list of names appropriate for the language
func GenerateNameList(ctx context.Context, numberOfNames int, langCategory Category, nameType string) ([]string, error) {
	var names []string
	var endings []string

	if nameType == "male" {
		endings = langCategory.MasculineEndings
	} else if nameType == "female" {
		endings = langCategory.FeminineEndings
	} else {
		for i := 0; i < 5; i++ {
			finisher, err := randomSyllable(ctx, langCategory, "finisher")
			if err != nil {
				err = fmt.Errorf(nameListError, err)
				return []string{}, err
			}
			endings = append(endings, finisher)
		}
	}

	for i := 0; i < numberOfNames; i++ {
		ending, err := random.String(ctx, endings)
		if err != nil {
			err = fmt.Errorf(nameListError, err)
			return []string{}, err
		}
		name, err := RandomName(ctx, langCategory)
		if err != nil {
			err = fmt.Errorf(nameListError, err)
			return []string{}, err
		}
		name += ending
		if !slices.StringIn(name, names) {
			names = append(names, name)
		}
	}

	return names, nil
}

func mutateName(ctx context.Context, name string) string {
	mutation := randomMutation(ctx)

	name = strings.Replace(name, mutation.From, mutation.To, 1)

	return name
}

func randomLanguageName(ctx context.Context, category Category) (string, error) {
	var name string
	var syllables []string
	skewLonger := false

	if random.Intn(ctx, 10) > 3 {
		skewLonger = true
	}

	randomLength := random.Intn(ctx, category.WordLength) + 1

	if skewLonger {
		randomLength++
	}

	role := "connector"

	for i := 0; i < randomLength; i++ {
		if randomLength-i == 1 {
			role = "finisher"
		}
		syllable, err := randomSyllable(ctx, category, role)
		if err != nil {
			err = fmt.Errorf("failed to generate language name: %w", err)
			return "", err
		}
		syllables = append(syllables, syllable)
	}

	for _, syllable := range syllables {
		name += syllable
	}

	chance := random.Intn(ctx, 10) + 1
	if chance > 3 {
		name = mutateName(ctx, name)
	}

	return name, nil
}

// RandomGenderedName generates a random gendered first name
func RandomGenderedName(ctx context.Context, langCategory Category, gender string) (string, error) {
	var endings []string

	name, err := RandomName(ctx, langCategory)
	if err != nil {
		err = fmt.Errorf("failed to generate random gendered name: %w", err)
		return "", err
	}

	if gender == "male" {
		endings = langCategory.MasculineEndings
	} else {
		endings = langCategory.FeminineEndings
	}

	ending, err := random.String(ctx, endings)
	if err != nil {
		err = fmt.Errorf("failed to generate random gendered name: %w", err)
		return "", err
	}
	name = name + ending

	return name, nil
}

// RandomName generates a random name using the language
func RandomName(ctx context.Context, langCategory Category) (string, error) {
	var name string
	var syllables []string
	skewLonger := false

	if random.Intn(ctx, 10) > 7 {
		skewLonger = true
	}

	randomLength := random.Intn(ctx, langCategory.WordLength) + 1

	if skewLonger {
		randomLength++
	}

	role := "connector"
	shouldIUseAnApostrophe := 0

	for i := 0; i < randomLength; i++ {
		if randomLength-i == 1 {
			role = "finisher"
		}
		syllable, err := randomSyllable(ctx, langCategory, role)
		if err != nil {
			err = fmt.Errorf("failed to generate conjugation rules: %w", err)
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
		name += syllable
	}

	chance := random.Intn(ctx, 10) + 1
	if chance > 8 {
		name = mutateName(ctx, name)
	}

	name = strings.Title(name)

	return name, nil
}
