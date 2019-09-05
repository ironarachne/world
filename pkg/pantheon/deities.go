package pantheon

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/ironarachne/world/pkg/gender"
	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/words"
)

// Deity is a fictional god or goddess
type Deity struct {
	Name              string
	Domains           []Domain
	Appearance        string
	Gender            gender.Gender
	PersonalityTraits []string
	Relationships     []Relationship
	HolyItem          string
	HolySymbol        string
}

// SimplifiedDeity is a display version of deity
type SimplifiedDeity struct {
	Name        string   `json:"name"`
	Gender      string   `json:"gender"`
	Domains     []string `json:"domains"`
	Description string   `json:"description"`
}

func (deity Deity) getRandomHolyItem() (string, error) {
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

	holyItem, err := random.String(options)
	if err != nil {
		err = fmt.Errorf("Could not generate random holy item: %w", err)
		return "", err
	}

	return holyItem, nil
}

func (deity Deity) getRandomHolySymbol() (string, error) {
	options := []string{}

	if len(deity.Domains) == 0 {
		options = []string{
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
	} else {
		for _, d := range deity.Domains {
			options = append(options, d.HolySymbols...)
		}
	}

	holySymbol, err := random.String(options)
	if err != nil {
		err = fmt.Errorf("Could not generate random holy symbol: %w", err)
		return "", err
	}

	return holySymbol, nil
}

// GenerateDeity generates a random deity
func (pantheon Pantheon) GenerateDeity(lang language.Language) (Deity, error) {
	var deity Deity
	var domain Domain
	var allDomains []Domain

	domains := getAllDomains()

	numberOfDomains := rand.Intn(3) + 1

	for _, deity := range pantheon.Deities {
		for _, d := range deity.Domains {
			allDomains = append(allDomains, d)
		}
	}

	for i := 0; i < numberOfDomains; i++ {
		domain = getRandomDomain(domains)

		// Only add domain if it isn't already in Domains slice
		if !isDomainInSlice(domain, deity.Domains) && !isDomainInSlice(domain, allDomains) {
			deity.Domains = append(deity.Domains, domain)
			allDomains = append(allDomains, domain)
		}
	}

	appearances, err := getRandomGeneralAppearances(3)
	if err != nil {
		err = fmt.Errorf("Could not generate deity: %w", err)
		return Deity{}, err
	}
	appearances = append(appearances, getAllAppearancesForDomains(deity.Domains)...)

	appearance, err := random.String(appearances)
	if err != nil {
		err = fmt.Errorf("Could not generate deity: %w", err)
		return Deity{}, err
	}
	deity.Appearance = appearance
	deity.Gender = gender.Random()

	traits, err := deity.getRandomTraits()
	if err != nil {
		err = fmt.Errorf("Could not generate deity: %w", err)
		return Deity{}, err
	}
	deity.PersonalityTraits = traits

	holyItem, err := deity.getRandomHolyItem()
	if err != nil {
		err = fmt.Errorf("Could not generate deity: %w", err)
		return Deity{}, err
	}
	deity.HolyItem = holyItem
	holySymbol, err := deity.getRandomHolySymbol()
	if err != nil {
		err = fmt.Errorf("Could not generate deity: %w", err)
		return Deity{}, err
	}
	deity.HolySymbol = holySymbol

	name, err := lang.RandomGenderedName(deity.Gender.Name)
	if err != nil {
		err = fmt.Errorf("Could not generate deity: %w", err)
		return Deity{}, err
	}
	deity.Name = name

	return deity, nil
}

func randomDeityNameFromMap(deities map[string]Deity) string {
	names := []string{}
	for _, d := range deities {
		names = append(names, d.Name)
	}

	return names[rand.Intn(len(names))]
}

// Describe describes a deity
func (deity Deity) Describe() string {
	var relationship string

	description := deity.Name

	if len(deity.Domains) > 0 {
		description += " is the god"
		if deity.Gender.Name == "female" {
			description += "dess"
		}
		description += " of " + words.CombinePhrases(getDomainNames(deity.Domains)) + ". "
	} else {
		description += " is a god"
		if deity.Gender.Name == "female" {
			description += "dess"
		}
		description += " of no particular domain. "
	}

	description += strings.Title(deity.Gender.SubjectPronoun) + " is " + deity.Appearance + ". "

	description += strings.Title(deity.Gender.SubjectPronoun) + " is " + words.CombinePhrases(deity.PersonalityTraits) + ". "

	description += strings.Title(deity.Gender.PossessivePronoun) + " holy item is " + words.Pronoun(deity.HolyItem) + " " + deity.HolyItem + ", and "
	description += deity.Gender.PossessivePronoun + " holy symbol is " + words.Pronoun(deity.HolySymbol) + " " + deity.HolySymbol + ". "

	relationships := []string{}

	if len(deity.Relationships) > 0 {
		for _, r := range deity.Relationships {
			relationship = r.Descriptor + " " + r.Target
			relationships = append(relationships, relationship)
		}

		description += deity.Name + " " + words.CombinePhrases(relationships) + "."
	}

	return description
}

// Simplify returns a simplified version of a deity
func (deity Deity) Simplify() SimplifiedDeity {
	domains := getDomainNames(deity.Domains)
	description := deity.Describe()

	return SimplifiedDeity{
		Name:        deity.Name,
		Gender:      deity.Gender.Name,
		Domains:     domains,
		Description: description,
	}
}
