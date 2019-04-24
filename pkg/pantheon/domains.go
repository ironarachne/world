package pantheon

import (
	"math/rand"
)

// Domain is an area of control
type Domain struct {
	Name              string
	AppearanceTraits  []string
	PersonalityTraits []string
}

func getAllAppearancesForDomains(domains []Domain) []string {
	var appearances []string

	for _, d := range domains {
		appearances = append(appearances, d.AppearanceTraits...)
	}

	return appearances
}

func getAllPersonalitiesForDomains(domains []Domain) []string {
	var personalities []string

	for _, d := range domains {
		personalities = append(personalities, d.PersonalityTraits...)
	}

	return personalities
}

func getRandomAppearanceFromDomains(domains []Domain) string {
	var possibleAppearances []string

	for _, d := range domains {
		possibleAppearances = append(possibleAppearances, d.AppearanceTraits...)
	}

	return possibleAppearances[rand.Intn(len(possibleAppearances)-1)]
}

func getRandomPersonalityFromDomains(domains []Domain) string {
	var possiblePersonalities []string

	for _, d := range domains {
		possiblePersonalities = append(possiblePersonalities, d.PersonalityTraits...)
	}

	return possiblePersonalities[rand.Intn(len(possiblePersonalities)-1)]
}

func getRandomDomain(domains []Domain) Domain {
	return domains[rand.Intn(len(domains)-1)]
}

func isDomainInSlice(domain Domain, domains []Domain) bool {
	for _, d := range domains {
		if domain.Name == d.Name {
			return true
		}
	}
	return false
}

func getAllDomains() []Domain {
	domains := []Domain{
		Domain{
			Name: "air",
			AppearanceTraits: []string{
				"bird-headed",
				"ethereal",
				"wind-blown",
				"winged",
			},
			PersonalityTraits: []string{
				"mercurial",
			},
		},
		Domain{
			Name: "animals",
			AppearanceTraits: []string{
				"dirty",
				"dog-headed",
				"bird-headed",
			},
			PersonalityTraits: []string{
				"brutish",
				"short-tempered",
				"skittish",
			},
		},
		Domain{
			Name: "art",
			AppearanceTraits: []string{
				"rainbow-hued",
			},
			PersonalityTraits: []string{},
		},
		Domain{
			Name: "autumn",
			AppearanceTraits: []string{
				"leaf-haired",
			},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:             "balance",
			AppearanceTraits: []string{},
			PersonalityTraits: []string{
				"careful",
				"measured",
			},
		},
		Domain{
			Name: "chaos",
			AppearanceTraits: []string{
				"blurred",
				"shifting",
			},
			PersonalityTraits: []string{
				"unpredictable",
				"wild",
			},
		},
		Domain{
			Name: "darkness",
			AppearanceTraits: []string{
				"dark as night",
				"shadowy",
			},
			PersonalityTraits: []string{},
		},
		Domain{
			Name: "dawn",
			AppearanceTraits: []string{
				"bright",
				"glowing",
				"radiant",
			},
			PersonalityTraits: []string{},
		},
		Domain{
			Name: "death",
			AppearanceTraits: []string{
				"gaunt",
				"skeletal",
			},
			PersonalityTraits: []string{},
		},
		Domain{
			Name: "demons",
			AppearanceTraits: []string{
				"black-eyed",
				"ember-eyed",
				"horned",
			},
			PersonalityTraits: []string{},
		},
		Domain{
			Name: "destruction",
			AppearanceTraits: []string{
				"spined",
			},
			PersonalityTraits: []string{
				"wrathful",
			},
		},
		Domain{
			Name: "dusk",
			AppearanceTraits: []string{
				"dark-skinned",
				"twilight-eyed",
			},
			PersonalityTraits: []string{
				"brooding",
			},
		},
		Domain{
			Name: "earth",
			AppearanceTraits: []string{
				"stone-skinned",
			},
			PersonalityTraits: []string{
				"grounded",
				"methodical",
				"practical",
			},
		},
		Domain{
			Name:              "evil",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "fear",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "fertility",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "fire",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "foxes",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "good",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "harvests",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "healing",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "hope",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "horses",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "hunting",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "justice",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "knowledge",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "language",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "law",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "life",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "light",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "lightning",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "love",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "luck",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "mercy",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "music",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "nature",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "nature",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "nobility",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "oceans",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "persistence",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "plants",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "protection",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "revenge",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "sky",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "spring",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "strength",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "summer",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "thieves",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "the moon",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "the sun",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "thunder",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "time",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "trade",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "travel",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "trickery",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "war",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "water",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "winter",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:              "wisdom",
			AppearanceTraits:  []string{},
			PersonalityTraits: []string{},
		},
	}

	return domains
}
