package language

import (
	"bytes"
	"html/template"
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
func Conjugate(lang Language, root string, tense string) string {
	template := ""

	if tense == "simple present" {
		template = lang.VerbConjugationRules.SimplePresent
	} else if tense == "simple past" {
		template = lang.VerbConjugationRules.SimplePast
	} else if tense == "simple future" {
		template = lang.VerbConjugationRules.SimpleFuture
	} else if tense == "present continuous" {
		template = lang.VerbConjugationRules.PresentContinuous
	} else if tense == "past continuous" {
		template = lang.VerbConjugationRules.PastContinuous
	} else if tense == "future continuous" {
		template = lang.VerbConjugationRules.FutureContinuous
	} else if tense == "present perfect" {
		template = lang.VerbConjugationRules.PresentPerfect
	} else if tense == "past perfect" {
		template = lang.VerbConjugationRules.PastPerfect
	} else if tense == "future perfect" {
		template = lang.VerbConjugationRules.FuturePerfect
	} else if tense == "present perfect continuous" {
		template = lang.VerbConjugationRules.PresentPerfectContinuous
	} else if tense == "past perfect continuous" {
		template = lang.VerbConjugationRules.PastPerfectContinuous
	} else if tense == "future perfect continuous" {
		template = lang.VerbConjugationRules.FuturePerfectContinuous
	}
	word := renderWord(root, template)

	return word
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
