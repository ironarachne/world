package pantheon

import (
	"math/rand"

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

	appearances := getRandomGeneralAppearances(3)
	appearances = append(appearances, getAllAppearancesForDomains(deity.Domains)...)

	deity.Appearance = random.String(appearances)
	deity.Gender = getRandomGender()

	deity.PersonalityTraits = deity.getRandomTraits()

	deity.Name = pantheon.Language.RandomName()

	return deity
}

func getRandomGender() string {
	genders := map[string]int{
		"female": 6,
		"male":   5,
		"none":   1,
	}

	return random.StringFromThresholdMap(genders)
}
