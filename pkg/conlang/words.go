package conlang

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/language"
)

// GenerateWordList creates a word list for a language
func GenerateWordList(ctx context.Context, langCategory Category) (map[string]string, error) {
	wordList := make(map[string]string)

	wordTypes := getAllWordTypes()

	for _, t := range wordTypes {
		for _, w := range t.WordList {
			for {
				newWord, err := randomWord(ctx, langCategory, t.MaxSyllables)
				if err != nil {
					err = fmt.Errorf("Could not generate word list: %w", err)
					return wordList, err
				}
				if !IsInWordList(newWord, wordList) {
					wordList[w], err = randomWord(ctx, langCategory, t.MaxSyllables)
					if err != nil {
						err = fmt.Errorf("Could not generate word list: %w", err)
						return wordList, err
					}
					break
				}
			}
		}
	}

	return wordList, nil
}

// AddNounToWordList adds a new noun to an existing word list for a language
func AddNounToWordList(ctx context.Context, lang language.Language, langCategory Category, word string) (map[string]string, error) {
	wordList := lang.WordList

	for {
		newWord, err := randomWord(ctx, langCategory, 3)
		if err != nil {
			err = fmt.Errorf("Could not add noun to word list: %w", err)
			return wordList, err
		}
		if !IsInWordList(newWord, wordList) {
			wordList[word], err = randomWord(ctx, langCategory, 3)
			if err != nil {
				err = fmt.Errorf("Could not add noun to word list: %w", err)
				return wordList, err
			}
			break
		}
	}

	return wordList, nil
}
