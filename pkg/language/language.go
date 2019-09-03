package language

import (
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

func deriveLanguageAdjective(name string) string {
	var suffix string

	adjective := name
	lastCharacter := adjective[len(adjective)-1:]

	potentialSuffixes := []string{"n", "lese", "ish"}

	if slices.StringIn(lastCharacter, consonants) {
		potentialSuffixes = []string{"ish", "ian", "an", "i", "ese"}
	}

	suffix = random.String(potentialSuffixes)

	adjective += suffix

	return adjective
}

// Generate creates a random language
func Generate() Language {
	var language Language

	combinedChance := rand.Intn(100)
	if combinedChance > 70 {
		language.Category = randomCombinedCategory()
	} else {
		language.Category = randomCategory()
	}

	language.Name = strings.Title(randomLanguageName(language.Category))
	language.Descriptors = append(language.Descriptors, language.Category.Descriptors...)
	language.Adjective = deriveLanguageAdjective(language.Name)

	tonalChance := rand.Intn(10) + 1
	if tonalChance > 7 {
		language.IsTonal = true
	} else {
		language.IsTonal = false
	}
	if language.IsTonal {
		language.Descriptors = append(language.Descriptors, "tonal")
	}

	language.WritingSystem = randomWritingSystem()
	language.WritingSystem.Name = language.Adjective
	language.WordList = language.GenerateWordList()
	language.VerbConjugationRules = language.deriveConjugationRules()
	language.PracticeSuffix = randomSyllable(language.Category, "finisher")

	return language
}

func (language Language) deriveConjugationRules() ConjugationRules {
	continuousSuffix := randomSyllable(language.Category, "finisher")
	pastSuffix := randomSyllable(language.Category, "finisher")

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

	return rules
}
