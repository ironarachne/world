package conlang

import (
	"context"

	"github.com/ironarachne/world/pkg/random"
)

// Category is a style of constructed language
type Category struct {
	Descriptors      []string
	WordLength       int
	UsesApostrophes  bool
	Initiators       []string
	Connectors       []string
	Finishers        []string
	MasculineEndings []string
	FeminineEndings  []string
}

func isCategoryInSlice(cat Category, categories []Category) bool {
	for _, c := range categories {
		if cat.Descriptors[0] == c.Descriptors[0] {
			return true
		}
	}
	return false
}

func randomCombinedCategory(ctx context.Context) Category {
	var cat Category

	categories := []Category{}

	for i := 0; i < 2; i++ {
		cat = randomCategory(ctx)
		if !isCategoryInSlice(cat, categories) {
			categories = append(categories, cat)
		} else {
			i--
		}
	}

	newCategory := Category{}

	for _, c := range categories {
		newCategory.Descriptors = append(newCategory.Descriptors, c.Descriptors...)
		newCategory.WordLength = c.WordLength
		newCategory.UsesApostrophes = c.UsesApostrophes
		newCategory.Initiators = append(newCategory.Initiators, c.Initiators...)
		newCategory.Connectors = append(newCategory.Connectors, c.Connectors...)
		newCategory.Finishers = append(newCategory.Finishers, c.Finishers...)
		newCategory.MasculineEndings = append(newCategory.MasculineEndings, c.MasculineEndings...)
		newCategory.FeminineEndings = append(newCategory.FeminineEndings, c.FeminineEndings...)
	}

	return newCategory
}

func getAllCategories() []Category {
	return []Category{
		{
			Descriptors:      []string{"musical"},
			WordLength:       2,
			UsesApostrophes:  false,
			Initiators:       fricatives,
			Connectors:       liquids,
			Finishers:        sibilants,
			MasculineEndings: []string{"ion", "on", "en", "o"},
			FeminineEndings:  []string{"i", "a", "ia", "ila"},
		},
		{
			Descriptors:      []string{"guttural"},
			WordLength:       1,
			UsesApostrophes:  false,
			Initiators:       append(glottals, growls...),
			Connectors:       velars,
			Finishers:        velars,
			MasculineEndings: []string{"ur", "ar", "ach", "ag"},
			FeminineEndings:  []string{"a", "agi"},
		},
		{
			Descriptors:      []string{"abrupt"},
			WordLength:       2,
			UsesApostrophes:  true,
			Initiators:       stops,
			Connectors:       fricatives,
			Finishers:        liquids,
			MasculineEndings: []string{"on", "en", "an"},
			FeminineEndings:  []string{"a", "e", "et"},
		},
		{
			Descriptors:      []string{"nasal"},
			WordLength:       2,
			UsesApostrophes:  false,
			Initiators:       glottals,
			Connectors:       stops,
			Finishers:        nasals,
			MasculineEndings: []string{"een", "oon", "in", "en"},
			FeminineEndings:  []string{"ini", "nia", "mia", "mi"},
		},
		{
			Descriptors:      []string{"rhythmic"},
			WordLength:       2,
			UsesApostrophes:  false,
			Initiators:       append(glottals, fricatives...),
			Connectors:       append(liquids, fricatives...),
			Finishers:        liquids,
			MasculineEndings: []string{"ior", "iun", "ayan", "ar"},
			FeminineEndings:  []string{"oa", "ua", "lia", "li"},
		},
		{
			Descriptors:      []string{"graceful"},
			WordLength:       2,
			UsesApostrophes:  false,
			Initiators:       consonants,
			Connectors:       append(liquids, fricatives...),
			Finishers:        append(stops, glottals...),
			MasculineEndings: []string{"em", "amn", "astor", "est", "and"},
			FeminineEndings:  []string{"eela", "aela", "ali", "eli", "oli", "oa", "ea"},
		},
		{
			Descriptors:      []string{"breathy"},
			WordLength:       1,
			UsesApostrophes:  false,
			Initiators:       append(breaths, fricatives...),
			Connectors:       append(breaths, glides...),
			Finishers:        append(stops, glottals...),
			MasculineEndings: []string{"esh", "ashem", "eh", "ih", "ah"},
			FeminineEndings:  []string{"eshi", "eha", "ala", "asha", "iha"},
		},
	}
}

func randomCategory(ctx context.Context) Category {
	categories := getAllCategories()

	return categories[random.Intn(ctx, len(categories))]
}
