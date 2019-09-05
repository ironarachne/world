package language

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

// Language is a fantasy language
type Language struct {
	Name                 string
	Adjective            string
	Descriptors          []string
	Category             Category
	IsTonal              bool
	WritingSystem        WritingSystem
	WordList             map[string]string
	VerbConjugationRules ConjugationRules
	PracticeSuffix       string
}

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
		err = fmt.Errorf("Could not generate language adjective: %w", err)
		return "", err
	}

	adjective += suffix

	return adjective, nil
}

// Generate creates a random language
func Generate() (Language, error) {
	var language Language

	combinedChance := rand.Intn(100)
	if combinedChance > 70 {
		language.Category = randomCombinedCategory()
	} else {
		language.Category = randomCategory()
	}

	name, err := randomLanguageName(language.Category)
	if err != nil {
		err = fmt.Errorf("Could not generate language: %w", err)
		return Language{}, err
	}
	language.Name = strings.Title(name)
	language.Descriptors = append(language.Descriptors, language.Category.Descriptors...)
	adjective, err := deriveLanguageAdjective(language.Name)
	if err != nil {
		err = fmt.Errorf("Could not generate language: %w", err)
		return Language{}, err
	}
	language.Adjective = adjective

	tonalChance := rand.Intn(10) + 1
	if tonalChance > 7 {
		language.IsTonal = true
	} else {
		language.IsTonal = false
	}
	if language.IsTonal {
		language.Descriptors = append(language.Descriptors, "tonal")
	}

	writingSystem, err := randomWritingSystem()
	if err != nil {
		err = fmt.Errorf("Could not generate language: %w", err)
		return Language{}, err
	}
	language.WritingSystem = writingSystem
	language.WritingSystem.Name = language.Adjective
	wordList, err := language.GenerateWordList()
	if err != nil {
		err = fmt.Errorf("Could not generate language: %w", err)
		return Language{}, err
	}
	language.WordList = wordList
	verbConjugationRules, err := language.deriveConjugationRules()
	if err != nil {
		err = fmt.Errorf("Could not generate language: %w", err)
		return Language{}, err
	}
	language.VerbConjugationRules = verbConjugationRules
	suffix, err := randomSyllable(language.Category, "finisher")
	if err != nil {
		err = fmt.Errorf("Could not generate language: %w", err)
		return Language{}, err
	}
	language.PracticeSuffix = suffix

	return language, nil
}

func (language Language) deriveConjugationRules() (ConjugationRules, error) {
	continuousSuffix, err := randomSyllable(language.Category, "finisher")
	if err != nil {
		err = fmt.Errorf("Could not generate conjugation rules: %w", err)
		return ConjugationRules{}, err
	}
	pastSuffix, err := randomSyllable(language.Category, "finisher")
	if err != nil {
		err = fmt.Errorf("Could not generate conjugation rules: %w", err)
		return ConjugationRules{}, err
	}

	rules := ConjugationRules{
		SimplePresent:            "{{.Root}}",
		SimplePast:               "{{.Root}}" + pastSuffix,
		SimpleFuture:             "{{.Root}}",
		PresentContinuous:        "{{.Root}}" + continuousSuffix,
		PastContinuous:           "{{.Root}}" + continuousSuffix,
		FutureContinuous:         "{{.Root}}" + continuousSuffix,
		PresentPerfect:           "{{.Root}}" + pastSuffix,
		PastPerfect:              "{{.Root}}" + pastSuffix,
		FuturePerfect:            "{{.Root}}" + pastSuffix,
		PresentPerfectContinuous: "{{.Root}}" + continuousSuffix,
		PastPerfectContinuous:    "{{.Root}}" + continuousSuffix,
		FuturePerfectContinuous:  "{{.Root}}" + continuousSuffix,
	}

	return rules, nil
}
