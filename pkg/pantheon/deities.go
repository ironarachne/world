package pantheon

import (
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
}

// SimplifiedDeity is a display version of deity
type SimplifiedDeity struct {
	Name        string
	Gender string
	Domains     []string
	Description string
}

// GenerateDeity generates a random deity
func (pantheon Pantheon) GenerateDeity(lang language.Language) Deity {
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
	deity.Gender = gender.Random()

	deity.PersonalityTraits = deity.getRandomTraits()

	deity.Name = lang.RandomName()

	return deity
}

func randomDeityNameFromMap(deities map[string]Deity) string {
	names := []string{}
	for _, d := range deities {
		names = append(names, d.Name)
	}

	return names[rand.Intn(len(names)-1)]
}

func (deity Deity) simplify() SimplifiedDeity {
	var relationship string

	sd := SimplifiedDeity{
		Name: deity.Name,
		Gender: deity.Gender.Name,
	}

	for _, d := range deity.Domains {
		sd.Domains = append(sd.Domains, d.Name)
	}

	description := deity.Name

	if len(sd.Domains) > 0 {
		description += " is the god"
		if deity.Gender.Name == "female" {
			description += "dess"
		}
		description += " of " + words.CombinePhrases(sd.Domains) + ". "
	} else {
		description += " is a god"
		if deity.Gender.Name == "female" {
			description += "dess"
		}
		description += " of no particular domain. "
	}

	description += strings.Title(deity.Gender.SubjectPronoun) + " is " + deity.Appearance + ". "

	description += strings.Title(deity.Gender.SubjectPronoun) + " is " + words.CombinePhrases(deity.PersonalityTraits) + ". "

	relationships := []string{}

	for _, r := range deity.Relationships {
		relationship = r.Descriptor + " " + r.Target
		relationships = append(relationships, relationship)
	}

	description += deity.Name + " " + words.CombinePhrases(relationships) + "."

	sd.Description = description

	return sd
}
