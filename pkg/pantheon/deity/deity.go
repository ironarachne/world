package deity

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/gender"
	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/pantheon/domain"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/relationship"
)

// Deity is a fictional god or goddess
type Deity struct {
	Name              string                      `json:"name"`
	Description       string                      `json:"description"`
	Domains           []domain.Domain             `json:"domains"`
	Appearance        string                      `json:"appearance"`
	Gender            gender.Gender               `json:"gender"`
	PersonalityTraits []string                    `json:"personality_traits"`
	Relationships     []relationship.Relationship `json:"relationships"`
	HolyItem          string                      `json:"holy_item"`
	HolySymbol        string                      `json:"holy_symbol"`
}

const deityError = "failed to generate deity: %w"

func (deity Deity) getRandomHolyItem(ctx context.Context) (string, error) {
	options := []string{
		"amulet",
		"necklace",
		"orb",
		"ring",
		"staff",
	}

	for _, d := range deity.Domains {
		options = append(options, d.HolyItems...)
	}

	holyItem, err := random.String(ctx, options)
	if err != nil {
		err = fmt.Errorf("Could not generate random holy item: %w", err)
		return "", err
	}

	return holyItem, nil
}

func getGenericHolySymbols() []string {
	options := []string{
		"circle divided in three",
		"closed eye",
		"open eye",
		"pair of circles",
		"pair of triangles",
		"square",
		"triangle mirrored",
		"triangle",
		"trio of slanted lines",
	}

	return options
}

func (deity Deity) getRandomHolySymbol(ctx context.Context) (string, error) {
	options := []string{}
	genericOptions := getGenericHolySymbols()

	if len(deity.Domains) == 0 {
		options = genericOptions
	} else {
		for _, d := range deity.Domains {
			options = append(options, d.HolySymbols...)
		}
		if len(options) == 0 {
			options = genericOptions
		}
	}

	holySymbol, err := random.String(ctx, options)
	if err != nil {
		err = fmt.Errorf("Could not generate random holy symbol: %w", err)
		return "", err
	}

	return holySymbol, nil
}

// ByName returns a deity with the given name from the given set
func ByName(name string, set []Deity) (Deity, error) {
	for _, d := range set {
		if d.Name == name {
			return d, nil
		}
	}

	err := fmt.Errorf("Deity with name " + name + " is not present in set")

	return Deity{}, err
}

// Exclude returns a set of deities with the given one removed
func Exclude(deity Deity, set []Deity) []Deity {
	var result []Deity

	for _, d := range set {
		if d.Name != deity.Name {
			result = append(result, d)
		}
	}

	return result
}

// Generate generates a random deity
func Generate(ctx context.Context, lang language.Language, possibleDomains []domain.Domain) (Deity, error) {
	var deity Deity
	var d domain.Domain
	var err error
	var name string

	numberOfDomains := random.Intn(ctx, 3) + 1

	for i := 0; i < numberOfDomains; i++ {
		d, err = domain.Random(ctx, possibleDomains)

		// Only add domain if it isn't already in Domains slice
		if !domain.InSlice(d, deity.Domains) {
			deity.Domains = append(deity.Domains, d)
		}
	}

	appearances, err := getRandomGeneralAppearances(ctx, 3)
	if err != nil {
		err = fmt.Errorf(deityError, err)
		return Deity{}, err
	}
	appearances = append(appearances, domain.AllAppearancesForDomains(deity.Domains)...)

	appearance, err := random.String(ctx, appearances)
	if err != nil {
		err = fmt.Errorf(deityError, err)
		return Deity{}, err
	}
	deity.Appearance = appearance
	deity.Gender = gender.Random(ctx)

	traits, err := deity.getRandomTraits(ctx)
	if err != nil {
		err = fmt.Errorf(deityError, err)
		return Deity{}, err
	}
	deity.PersonalityTraits = traits

	holyItem, err := deity.getRandomHolyItem(ctx)
	if err != nil {
		err = fmt.Errorf(deityError, err)
		return Deity{}, err
	}
	deity.HolyItem = holyItem
	holySymbol, err := deity.getRandomHolySymbol(ctx)
	if err != nil {
		err = fmt.Errorf(deityError, err)
		return Deity{}, err
	}
	deity.HolySymbol = holySymbol

	if deity.Gender.Name == "female" {
		name, err = lang.RandomFemaleFirstName(ctx)
		if err != nil {
			err = fmt.Errorf(deityError, err)
			return Deity{}, err
		}
	} else if deity.Gender.Name == "male" {
		name, err = lang.RandomMaleFirstName(ctx)
		if err != nil {
			err = fmt.Errorf(deityError, err)
			return Deity{}, err
		}
	} else {
		name, err = lang.RandomFemaleFirstName(ctx)
		if err != nil {
			err = fmt.Errorf(deityError, err)
			return Deity{}, err
		}
	}

	deity.Name = name
	deity.Description = deity.Describe()

	return deity, nil
}

// Random returns a random deity from a set
func Random(ctx context.Context, set []Deity) (Deity, error) {
	if len(set) == 0 {
		err := fmt.Errorf("tried to get a random deity from an empty set")
		return Deity{}, err
	}

	if len(set) == 1 {
		return set[0], nil
	}

	deity := set[random.Intn(ctx, len(set))]

	return deity, nil
}

// Replace replaces an existing deity in a given set
func Replace(deity Deity, set []Deity) []Deity {
	var result []Deity

	for _, d := range set {
		if deity.Name != d.Name {
			result = append(result, d)
		}
	}

	result = append(result, deity)

	return result
}
