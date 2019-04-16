package language

import "math/rand"

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
