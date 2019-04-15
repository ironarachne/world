package pantheon

import (
	"math/rand"

	"github.com/ironarachne/naminglanguage"
	"github.com/ironarachne/world/pkg/random"
)

// Deity is a fictional god or goddess
type Deity struct {
	Name              string
	Domains           []string
	Appearance        string
	Gender            string
	PersonalityTraits []string
}

// Pantheon is a nonhierarchical group of deities
type Pantheon struct {
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
	var domain string
	var allDomains []string

	domains := getAllDomains()

	numberOfDomains := rand.Intn(3) + 1

	for _, deity := range pantheon.Deities {
		for _, d := range deity.Domains {
			allDomains = append(allDomains, d)
		}
	}

	for i := 0; i < numberOfDomains; i++ {
		domain = random.String(domains)

		// Only add domain if it isn't already in Domains slice
		if !inSlice(domain, deity.Domains) && !inSlice(domain, allDomains) {
			deity.Domains = append(deity.Domains, domain)
			allDomains = append(allDomains, domain)
		}
	}

	appearances := getAllAppearances()

	deity.Appearance = random.String(appearances)
	deity.Gender = getRandomGender()

	deity.PersonalityTraits = deity.getRandomTraits()

	deity.Name = randomName()

	return deity
}

// Generate creates a random pantheon of deities
func Generate(maxSize int) Pantheon {
	var deity Deity
	var pantheon Pantheon

	numberOfDeities := rand.Intn(maxSize) + 1

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

func inSlice(value string, slice []string) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}

func randomName() string {
	var person *naminglanguage.Person

	person = naminglanguage.GeneratePerson()

	return person.FirstName
}
