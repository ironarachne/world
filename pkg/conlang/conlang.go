/*
Package conlang provides methods and tools for procedurally generating
constructed languages. It makes use of the language package for underlying
structure.
*/
package conlang

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
	"github.com/ironarachne/world/pkg/writing"
)

const languageError = "failed to generate language: %w"

var (
	consonants = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}
	breaths    = []string{"h", "th", "f", "ch", "sh"}
	fricatives = []string{"f", "v", "th", "ð", "s", "z", "ch", "zh"}
	glides     = []string{"j", "w"}
	glottals   = []string{"g", "k", "ch"}
	growls     = []string{"br", "tr", "gr", "dr", "kr"}
	liquids    = []string{"l", "r"}
	nasals     = []string{"m", "n", "ng"}
	sibilants  = []string{"s", "f"}
	stops      = []string{"p", "b", "t", "d", "k", "g"}
	velars     = []string{"k", "g", "ng", "w"}
)

func deriveLanguageAdjective(name string) (string, error) {
	var suffix string

	adjective := name
	lastCharacter := adjective[len(adjective)-1:]

	potentialSuffixes := []string{"n", "lese", "ish"}

	if slices.StringIn(lastCharacter, consonants) {
		potentialSuffixes = []string{"ish", "ian", "an", "i", "ese"}
	}

	suffix, err := random.String(potentialSuffixes)
	if err != nil {
		err = fmt.Errorf("failed to generate language adjective: %w", err)
		return "", err
	}

	adjective += suffix

	return adjective, nil
}

// Generate creates a random language
func Generate() (language.Language, Category, error) {
	var lang language.Language
	var langCategory Category
	var prefix string
	var suffix string

	// Does this language combine multiple language categories?
	combinedChance := rand.Intn(100)
	if combinedChance > 70 {
		langCategory = randomCombinedCategory()
	} else {
		langCategory = randomCategory()
	}

	// Language name, adjective, and descriptors
	name, err := randomLanguageName(langCategory)
	if err != nil {
		err = fmt.Errorf(languageError, err)
		return language.Language{}, Category{}, err
	}
	lang.Name = strings.Title(name)
	lang.Descriptors = append(lang.Descriptors, langCategory.Descriptors...)
	adjective, err := deriveLanguageAdjective(lang.Name)
	if err != nil {
		err = fmt.Errorf(languageError, err)
		return language.Language{}, Category{}, err
	}
	lang.Adjective = adjective

	// Is this a tonal language?
	tonalChance := rand.Intn(10) + 1
	if tonalChance > 7 {
		lang.IsTonal = true
	} else {
		lang.IsTonal = false
	}
	if lang.IsTonal {
		lang.Descriptors = append(lang.Descriptors, "tonal")
	}

	// Writing system generation
	writingSystem, err := writing.Generate()
	if err != nil {
		err = fmt.Errorf(languageError, err)
		return language.Language{}, Category{}, err
	}
	lang.WritingSystem = writingSystem
	lang.WritingSystem.Name = lang.Adjective

	// Word list generation
	wordList, err := GenerateWordList(langCategory)
	if err != nil {
		err = fmt.Errorf(languageError, err)
		return language.Language{}, Category{}, err
	}
	lang.WordList = wordList

	// Verb conjugation rules generation
	verbConjugationRules, err := deriveConjugationRules(langCategory)
	if err != nil {
		err = fmt.Errorf(languageError, err)
		return language.Language{}, Category{}, err
	}
	lang.VerbConjugationRules = verbConjugationRules

	// Affix generation for new words
	newWordPrefixes := []string{}

	for i := 0; i < 50; i++ {
		prefix, err = randomSyllable(langCategory, "connector")
		if !slices.StringIn(prefix, newWordPrefixes) {
			newWordPrefixes = append(newWordPrefixes, prefix)
		}
	}

	newWordSuffixes := []string{}

	for i := 0; i < 50; i++ {
		suffix, err = randomSyllable(langCategory, "finisher")
		if !slices.StringIn(suffix, newWordSuffixes) {
			newWordSuffixes = append(newWordSuffixes, suffix)
		}
	}

	// Name generation
	femaleNames, err := GenerateNameList(100, langCategory, "female")
	if err != nil {
		err = fmt.Errorf(languageError, err)
		return language.Language{}, Category{}, err
	}
	maleNames, err := GenerateNameList(100, langCategory, "male")
	if err != nil {
		err = fmt.Errorf(languageError, err)
		return language.Language{}, Category{}, err
	}
	familyNames, err := GenerateNameList(100, langCategory, "family")
	if err != nil {
		err = fmt.Errorf(languageError, err)
		return language.Language{}, Category{}, err
	}
	townNames, err := GenerateNameList(100, langCategory, "town")
	if err != nil {
		err = fmt.Errorf(languageError, err)
		return language.Language{}, Category{}, err
	}
	lang.FemaleFirstNames = femaleNames
	lang.MaleFirstNames = maleNames
	lang.FamilyNames = familyNames
	lang.TownNames = townNames
	lang.NewWordPrefixes = newWordPrefixes
	lang.NewWordSuffixes = newWordSuffixes
	lang.Description = lang.Describe()
	lang.SamplePhrase = "Hello! It is good to see you, friend."
	lang.SamplePhraseTranslation = lang.TranslatePhrase(lang.SamplePhrase)

	return lang, langCategory, nil
}

func deriveConjugationRules(langCategory Category) (language.ConjugationRules, error) {
	rootTemplate := "{{.Root}}"

	continuousSuffix, err := randomSyllable(langCategory, "finisher")
	if err != nil {
		err = fmt.Errorf("Could not generate conjugation rules: %w", err)
		return language.ConjugationRules{}, err
	}
	pastSuffix, err := randomSyllable(langCategory, "finisher")
	if err != nil {
		err = fmt.Errorf("Could not generate conjugation rules: %w", err)
		return language.ConjugationRules{}, err
	}

	rules := language.ConjugationRules{
		SimplePresent:            rootTemplate,
		SimplePast:               rootTemplate + pastSuffix,
		SimpleFuture:             rootTemplate,
		PresentContinuous:        rootTemplate + continuousSuffix,
		PastContinuous:           rootTemplate + continuousSuffix,
		FutureContinuous:         rootTemplate + continuousSuffix,
		PresentPerfect:           rootTemplate + pastSuffix,
		PastPerfect:              rootTemplate + pastSuffix,
		FuturePerfect:            rootTemplate + pastSuffix,
		PresentPerfectContinuous: rootTemplate + continuousSuffix,
		PastPerfectContinuous:    rootTemplate + continuousSuffix,
		FuturePerfectContinuous:  rootTemplate + continuousSuffix,
	}

	return rules, nil
}
