package language

import (
	"strings"

	"github.com/ironarachne/world/pkg/words"
)

// RosettaStone generates a version of a common phrase in the language
func (language Language) RosettaStone() string {
	phrase := "Hello! It is good to see you, friend."

	translation := words.CapitalizeFirst(language.TranslatePhrase(phrase))

	return translation
}

// TranslatePhrase parses a phrase and returns the translated version
func (language Language) TranslatePhrase(phrase string) string {
	punctuation := ""
	translation := ""
	translatedWord := ""
	wordOnly := ""
	skipNext := false
	titleNext := false

	phrase = strings.ToLower(phrase)

	words := strings.Split(phrase, " ")

	for i, w := range words {
		if !skipNext {
			if strings.ContainsAny(w, "!?.,") {
				wordOnly = strings.TrimRight(w, "!?.,")
				punctuation = string(w[strings.IndexAny(w, "!?.,")])
			} else {
				wordOnly = w
				punctuation = ""
			}

			if wordOnly == "is" {
				wordOnly = "to be"
			}

			if wordOnly == "to" && i < len(words)-2 && language.WordList["to "+words[i+1]] != "" {
				wordOnly = "to " + words[i+1]
				skipNext = true
			}

			wordOnly = getSingular(wordOnly)
			translatedWord = language.WordList[wordOnly]
			if titleNext {
				translatedWord = strings.Title(translatedWord)
			}

			translation += translatedWord + punctuation + " "

			if strings.ContainsAny(w, "!?.") {
				titleNext = true
			} else {
				titleNext = false
			}
		} else {
			skipNext = false
		}
	}

	// TODO: Remove trailing space

	return translation
}

// TranslateWord returns the translated version of the word
func (language Language) TranslateWord(word string) string {
	word = strings.ToLower(word)

	return language.WordList[word]
}

func getSingular(word string) string {
	if strings.HasSuffix(word, "es") {
		word = strings.TrimSuffix(word, "es")
	} else if strings.HasSuffix(word, "s") {
		word = strings.TrimSuffix(word, "s")
	}

	return word
}
