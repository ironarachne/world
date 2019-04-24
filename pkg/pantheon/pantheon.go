package pantheon

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/random"
)

// Deity is a fictional god or goddess
type Deity struct {
	Name              string
	Domains           []Domain
	Appearance        string
	Gender            string
	PersonalityTraits []string
}

// Pantheon is a nonhierarchical group of deities
type Pantheon struct {
	Language      language.Language
	Deities       []Deity
	Relationships []Relationship
}

// Relationship is a unidirectional relationship between deities
type Relationship struct {
	Origin     Deity
	Target     Deity
	Descriptor string
}

// GenerateDeity generates a random deity
func (pantheon Pantheon) GenerateDeity() Deity {
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

	appearances := getGeneralAppearances()
	appearances = append(appearances, getAllAppearancesForDomains(deity.Domains)...)

	deity.Appearance = random.String(appearances)
	deity.Gender = getRandomGender()

	deity.PersonalityTraits = deity.getRandomTraits()

	deity.Name = pantheon.Language.RandomName()

	return deity
}

// Generate creates a random pantheon of deities
func Generate(maxSize int, lang language.Language) Pantheon {
	var deity Deity
	var pantheon Pantheon

	numberOfDeities := rand.Intn(maxSize) + 1

	pantheon.Language = lang

	for i := 0; i < numberOfDeities; i++ {
		deity = pantheon.GenerateDeity()
		pantheon.Deities = append(pantheon.Deities, deity)
	}

	pantheon.Relationships = pantheon.GenerateRelationships()

	return pantheon
}

// GenerateRelationships generates relationships between deities
func (pantheon Pantheon) GenerateRelationships() []Relationship {
	var descriptor string
	var relationship Relationship
	var relationships []Relationship
	var target Deity

	descriptors := []string{
		"is parent to",
		"hates",
		"loves",
		"fears",
		"respects",
		"is amused by",
		"is chagrined by",
		"worries about",
		"is suspicious of",
	}

	for _, deity := range pantheon.Deities {
		target = pantheon.Deities[rand.Intn(len(pantheon.Deities))]
		descriptor = random.String(descriptors)
		relationship = Relationship{Origin: deity, Target: target, Descriptor: descriptor}

		if deity.Name != target.Name {
			relationships = append(relationships, relationship)
		}
	}

	return relationships
}

func getRandomGender() string {
	genders := []string{
		"female",
		"male",
		"none",
	}

	return random.String(genders)
}
