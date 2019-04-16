package language

import (
	"math/rand"
	"strings"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

// Language is a fantasy language
type Language struct {
	Name          string
	Adjective     string
	Descriptors   []string
	Category      Category
	IsTonal       bool
	WritingSystem WritingSystem
}

// Category is a style of language
type Category struct {
	Name             string
	WordLength       int
	UsesApostrophes  bool
	Initiators       []string
	Connectors       []string
	Finishers        []string
	MasculineEndings []string
	FeminineEndings  []string
}

// Mutation is a word mutation
type Mutation struct {
	From string
	To   string
}

// WritingSystem is a system of writing
type WritingSystem struct {
	Name           string
	Classification string
	StrokeType     string
}

var (
	consonants = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}
	vowels     = []string{"a", "e", "i", "o", "u"}
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

func deriveAdjective(name string) string {
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

	language.Category = randomCategory()
	language.Name = strings.Title(randomLanguageName(language.Category))
	language.Descriptors = append(language.Descriptors, language.Category.Name)
	language.Adjective = deriveAdjective(language.Name)

	tonalChance := rand.Intn(10) + 1
	if tonalChance > 7 {
		language.IsTonal = true
	} else {
		language.IsTonal = false
	}

	language.WritingSystem = randomWritingSystem()
	language.WritingSystem.Name = language.Adjective

	return language
}

// GenerateNameList generates a list of names appropriate for the language
func (language Language) GenerateNameList(nameType string) []string {
	var names []string
	var name string
	var endings []string

	if nameType == "male" {
		endings = language.Category.MasculineEndings
	} else if nameType == "female" {
		endings = language.Category.FeminineEndings
	} else {
		for i := 0; i < 5; i++ {
			endings = append(endings, randomSyllable(language.Category, "finisher"))
		}
	}

	for i := 0; i < 10; i++ {
		name = language.RandomName() + random.String(endings)
		if !slices.StringIn(name, names) {
			names = append(names, name)
		}
	}

	return names
}

func mutateName(name string) string {
	mutation := randomMutation()

	name = strings.Replace(name, mutation.From, mutation.To, 1)

	return name
}

func randomCategory() Category {
	categories := []Category{
		Category{
			Name:             "musical",
			WordLength:       2,
			UsesApostrophes:  false,
			Initiators:       fricatives,
			Connectors:       liquids,
			Finishers:        sibilants,
			MasculineEndings: []string{"ion", "on", "en", "o"},
			FeminineEndings:  []string{"i", "a", "ia", "ila"},
		},
		Category{
			Name:             "guttural",
			WordLength:       1,
			UsesApostrophes:  false,
			Initiators:       append(glottals, growls...),
			Connectors:       velars,
			Finishers:        velars,
			MasculineEndings: []string{"ur", "ar", "ach", "ag"},
			FeminineEndings:  []string{"a", "agi"},
		},
		Category{
			Name:             "abrupt",
			WordLength:       2,
			UsesApostrophes:  true,
			Initiators:       stops,
			Connectors:       fricatives,
			Finishers:        liquids,
			MasculineEndings: []string{"on", "en", "an"},
			FeminineEndings:  []string{"a", "e", "et"},
		},
		Category{
			Name:             "nasal",
			WordLength:       2,
			UsesApostrophes:  false,
			Initiators:       glottals,
			Connectors:       stops,
			Finishers:        nasals,
			MasculineEndings: []string{"een", "oon", "in", "en"},
			FeminineEndings:  []string{"ini", "nia", "mia", "mi"},
		},
		Category{
			Name:             "rhythmic",
			WordLength:       2,
			UsesApostrophes:  false,
			Initiators:       append(glottals, fricatives...),
			Connectors:       append(liquids, fricatives...),
			Finishers:        liquids,
			MasculineEndings: []string{"ior", "iun", "ayan", "ar"},
			FeminineEndings:  []string{"oa", "ua", "lia", "li"},
		},
		Category{
			Name:             "graceful",
			WordLength:       2,
			UsesApostrophes:  false,
			Initiators:       consonants,
			Connectors:       append(liquids, fricatives...),
			Finishers:        append(stops, glottals...),
			MasculineEndings: []string{"em", "amn", "astor", "est", "and"},
			FeminineEndings:  []string{"eela", "aela", "ali", "eli", "oli", "oa", "ea"},
		},
		Category{
			Name:             "breathy",
			WordLength:       1,
			UsesApostrophes:  false,
			Initiators:       append(breaths, fricatives...),
			Connectors:       append(breaths, glides...),
			Finishers:        append(stops, glottals...),
			MasculineEndings: []string{"esh", "ashem", "eh", "ih", "ah"},
			FeminineEndings:  []string{"eshi", "eha", "ala", "asha", "iha"},
		},
	}
	return categories[rand.Intn(len(categories))]
}

func randomLanguageName(category Category) string {
	var name string
	var syllables []string
	skewLonger := false

	if rand.Intn(10) > 3 {
		skewLonger = true
	}

	randomLength := rand.Intn(category.WordLength) + 1

	if skewLonger {
		randomLength++
	}

	role := "connector"

	for i := 0; i < randomLength; i++ {
		if randomLength-i == 1 {
			role = "finisher"
		}
		syllables = append(syllables, randomSyllable(category, role))
	}

	for _, syllable := range syllables {
		name += syllable
	}

	chance := rand.Intn(10) + 1
	if chance > 3 {
		name = mutateName(name)
	}

	return name
}

func randomMutation() Mutation {
	rules := []Mutation{
		Mutation{
			"s",
			"ss",
		},
		Mutation{
			"s",
			"sh",
		},
		Mutation{
			"f",
			"ff",
		},
		Mutation{
			"f",
			"fh",
		},
		Mutation{
			"g",
			"gh",
		},
		Mutation{
			"l",
			"l'",
		},
	}

	return rules[rand.Intn(len(rules)-1)]
}

// RandomGenderedName generates a random gendered first name
func (language Language) RandomGenderedName(gender string) string {
	var endings []string

	name := language.RandomName()

	if gender == "male" {
		endings = language.Category.MasculineEndings
	} else {
		endings = language.Category.FeminineEndings
	}

	name = name + random.String(endings)

	return name
}

// RandomName generates a random name using the language
func (language Language) RandomName() string {
	var name string
	var syllables []string
	skewLonger := false

	if rand.Intn(10) > 7 {
		skewLonger = true
	}

	randomLength := rand.Intn(language.Category.WordLength) + 1

	if skewLonger {
		randomLength++
	}

	role := "connector"
	syllable := ""
	shouldIUseAnApostrophe := 0

	for i := 0; i < randomLength; i++ {
		if randomLength-i == 1 {
			role = "finisher"
		}
		syllable = randomSyllable(language.Category, role)

		if language.Category.UsesApostrophes {
			shouldIUseAnApostrophe = rand.Intn(10)
			if shouldIUseAnApostrophe > 8 {
				syllable += "'"
			}
		}

		syllables = append(syllables, syllable)
	}

	for _, syllable := range syllables {
		name += syllable
	}

	chance := rand.Intn(10) + 1
	if chance > 8 {
		name = mutateName(name)
	}

	name = strings.Title(name)

	return name
}

func randomSyllable(category Category, role string) string {
	syllable := random.String(category.Initiators) + random.String(vowels)
	expand := rand.Intn(10) + 1
	if expand > 2 {
		if role == "connector" {
			syllable += random.String(category.Connectors)
		} else {
			syllable += random.String(category.Finishers)
		}
	}

	return syllable
}

func randomStrokeType() string {
	strokeTypes := []string{
		"arcs and loops",
		"dots and slashes",
		"flowing",
		"geometric shapes",
		"right angles",
		"runic",
	}

	return random.String(strokeTypes)
}

func randomWritingSystem() WritingSystem {
	var writingSystem WritingSystem

	classifications := []string{
		"abjad",
		"abugida",
		"alphabet",
		"ideograms",
		"pictograms",
		"semanto-phonetic",
		"syllabary",
	}

	writingSystem.Classification = random.String(classifications)
	writingSystem.StrokeType = randomStrokeType()

	return writingSystem
}
