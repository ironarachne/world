package language

import (
	"bytes"
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
		template = language.ConjugationRules.SimplePresent
	} else if tense == "simple past" {
		template = language.ConjugationRules.SimplePast
	} else if tense == "simple future" {
		template = language.ConjugationRules.SimpleFuture
	} else if tense == "present continuous" {
		template = language.ConjugationRules.PresentContinuous
	} else if tense == "past continuous" {
		template = language.ConjugationRules.PastContinuous
	} else if tense == "future continuous" {
		template = language.ConjugationRules.FutureContinuous
	} else if tense == "present perfect" {
		template = language.ConjugationRules.PresentPerfect
	} else if tense == "past perfect" {
		template = language.ConjugationRules.PastPerfect
	} else if tense == "future perfect" {
		template = language.ConjugationRules.FuturePerfect
	} else if tense == "present perfect continuous" {
		template = language.ConjugationRules.PresentPerfectContinuous
	} else if tense == "past perfect continuous" {
		template = language.ConjugationRules.PastPerfectContinuous
	} else if tense == "future perfect continuous" {
		template = language.ConjugationRules.FuturePerfectContinuous
	}
	word := renderWord(root, template)

	return word
}

// GenerateWordList creates a word list for a language
func (language Language) GenerateWordList() map[string]string {
	newWord := ""
	wordList := make(map[string]string)

	wordTypes := getAllWordTypes()

	for _, t := range wordTypes {
		for _, w := range t.WordList {
			for {
				newWord = language.randomWord(t.NumSyllables)
				if !isInWordList(newWord, wordList) {
					wordList[w] = language.randomWord(t.NumSyllables)
					break
				}
			}
		}
	}

	return wordList
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
