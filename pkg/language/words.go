package language

import (
	"bytes"
	"fmt"
	"text/template"
)

// ConjugationRules is a set of templates to be applied to a word depending on its tense
type ConjugationRules struct {
	SimplePresent            string
	SimplePast               string
	SimpleFuture             string
	PresentContinuous        string
	PastContinuous           string
	FutureContinuous         string
	PresentPerfect           string
	PastPerfect              string
	FuturePerfect            string
	PresentPerfectContinuous string
	PastPerfectContinuous    string
	FuturePerfectContinuous  string
}

// Word is a simple wrapper for use in conjugations
type Word struct {
	Root string
}

// Conjugate returns the conjugated form of a root word given a tense
func (language Language) Conjugate(root string, tense string) string {
	template := ""

	if tense == "simple present" {
		template = language.VerbConjugationRules.SimplePresent
	} else if tense == "simple past" {
		template = language.VerbConjugationRules.SimplePast
	} else if tense == "simple future" {
		template = language.VerbConjugationRules.SimpleFuture
	} else if tense == "present continuous" {
		template = language.VerbConjugationRules.PresentContinuous
	} else if tense == "past continuous" {
		template = language.VerbConjugationRules.PastContinuous
	} else if tense == "future continuous" {
		template = language.VerbConjugationRules.FutureContinuous
	} else if tense == "present perfect" {
		template = language.VerbConjugationRules.PresentPerfect
	} else if tense == "past perfect" {
		template = language.VerbConjugationRules.PastPerfect
	} else if tense == "future perfect" {
		template = language.VerbConjugationRules.FuturePerfect
	} else if tense == "present perfect continuous" {
		template = language.VerbConjugationRules.PresentPerfectContinuous
	} else if tense == "past perfect continuous" {
		template = language.VerbConjugationRules.PastPerfectContinuous
	} else if tense == "future perfect continuous" {
		template = language.VerbConjugationRules.FuturePerfectContinuous
	}
	word := renderWord(root, template)

	return word
}

// GenerateWordList creates a word list for a language
func (language Language) GenerateWordList() (map[string]string, error) {
	wordList := make(map[string]string)

	wordTypes := getAllWordTypes()

	for _, t := range wordTypes {
		for _, w := range t.WordList {
			for {
				newWord, err := language.randomWord(t.MaxSyllables)
				if err != nil {
					err = fmt.Errorf("Could not generate word list: %w", err)
					return wordList, err
				}
				if !IsInWordList(newWord, wordList) {
					wordList[w], err = language.randomWord(t.MaxSyllables)
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
func (language Language) AddNounToWordList(word string) (map[string]string, error) {
	wordList := language.WordList

	for {
		newWord, err := language.randomWord(3)
		if err != nil {
			err = fmt.Errorf("Could not add noun to word list: %w", err)
			return wordList, err
		}
		if !IsInWordList(newWord, wordList) {
			wordList[word], err = language.randomWord(3)
			if err != nil {
				err = fmt.Errorf("Could not add noun to word list: %w", err)
				return wordList, err
			}
			break
		}
	}

	return wordList, nil
}

// MakePractice changes a word into an "-ism" form word
func (language Language) MakePractice(word string) string {
	practice := word + language.PracticeSuffix

	return practice
}

func renderWord(word string, wordTemplate string) string {
	var tplOutput bytes.Buffer

	data := Word{
		Root: word,
	}

	tmpl, _ := template.New(word).Parse(wordTemplate)
	tmpl.Execute(&tplOutput, data)

	templateOutput := tplOutput.String()

	return templateOutput
}
