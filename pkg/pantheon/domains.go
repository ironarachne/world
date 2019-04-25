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
			Name: "evil",
			AppearanceTraits: []string{
				"black-eyed",
				"horned",
				"spider-bodied",
			},
			PersonalityTraits: []string{
				"brutal",
				"cruel",
				"vicious",
			},
		},
		Domain{
			Name: "fear",
			AppearanceTraits: []string{
				"shadowy",
				"nebulous",
				"unnaturally thin",
			},
			PersonalityTraits: []string{
				"cruel",
			},
		},
		Domain{
			Name: "fertility",
			AppearanceTraits: []string{
				"entrancing",
				"hauntingly beautiful",
			},
			PersonalityTraits: []string{
				"playful",
				"seductive",
			},
		},
		Domain{
			Name: "fire",
			AppearanceTraits: []string{
				"black skin shot through with firey lines",
				"ember-eyed",
				"fire-for-hair",
			},
			PersonalityTraits: []string{
				"easily upsettable",
				"firey",
				"passionate",
				"quick to anger",
			},
		},
		Domain{
			Name: "foxes",
			AppearanceTraits: []string{
				"fox-eared",
				"fox-headed",
				"fox-bodied",
			},
			PersonalityTraits: []string{
				"playful",
				"tricksy",
			},
		},
		Domain{
			Name: "good",
			AppearanceTraits: []string{
				"taciturn",
			},
			PersonalityTraits: []string{
				"jovial",
				"tolerant",
			},
		},
		Domain{
			Name: "harvests",
			AppearanceTraits: []string{
				"stocky",
			},
			PersonalityTraits: []string{
				"festive",
				"hard-working",
			},
		},
		Domain{
			Name: "healing",
			AppearanceTraits: []string{
				"green-haloed",
			},
			PersonalityTraits: []string{
				"caring",
				"compassionate",
				"patient",
			},
		},
		Domain{
			Name: "hope",
			AppearanceTraits: []string{
				"haloed",
			},
			PersonalityTraits: []string{
				"patient",
			},
		},
		Domain{
			Name: "horses",
			AppearanceTraits: []string{
				"horse-bodied",
				"horse-headed",
			},
			PersonalityTraits: []string{
				"ebullient",
			},
		},
		Domain{
			Name: "hunting",
			AppearanceTraits: []string{
				"narrow-featured",
				"tightly muscled",
			},
			PersonalityTraits: []string{
				"patient",
				"subtle",
			},
		},
		Domain{
			Name: "justice",
			AppearanceTraits: []string{
				"blindfolded",
			},
			PersonalityTraits: []string{
				"callous",
				"even-tempered",
				"fair",
			},
		},
		Domain{
			Name: "knowledge",
			AppearanceTraits: []string{
				"old and wrinkled",
				"bespectacled",
			},
			PersonalityTraits: []string{
				"absent-minded",
				"studious",
			},
		},
		Domain{
			Name:             "language",
			AppearanceTraits: []string{},
			PersonalityTraits: []string{
				"social",
			},
		},
		Domain{
			Name:             "law",
			AppearanceTraits: []string{},
			PersonalityTraits: []string{
				"fair",
				"stern",
			},
		},
		Domain{
			Name: "life",
			AppearanceTraits: []string{
				"haloed",
				"surrounded by living things",
			},
			PersonalityTraits: []string{
				"full of life",
				"joyful",
			},
		},
		Domain{
			Name: "light",
			AppearanceTraits: []string{
				"glowing",
				"white-haired",
			},
			PersonalityTraits: []string{
				"serious",
			},
		},
		Domain{
			Name: "lightning",
			AppearanceTraits: []string{
				"lightning-eyed",
				"surrounded by crackling sparks",
			},
			PersonalityTraits: []string{
				"emotional",
				"unpredictable",
				"violent",
			},
		},
		Domain{
			Name: "love",
			AppearanceTraits: []string{
				"beautiful",
				"gorgeous",
				"pretty",
				"unabashedly naked",
			},
			PersonalityTraits: []string{
				"joyful",
				"lustful",
				"passionate",
			},
		},
		Domain{
			Name: "luck",
			AppearanceTraits: []string{
				"always playing with cards",
				"perpetually grinning",
			},
			PersonalityTraits: []string{
				"adventurous",
			},
		},
		Domain{
			Name: "mercy",
			AppearanceTraits: []string{
				"plain",
			},
			PersonalityTraits: []string{
				"compassionate",
				"fair",
			},
		},
		Domain{
			Name: "music",
			AppearanceTraits: []string{
				"entracing",
				"harpstrings for hair",
			},
			PersonalityTraits: []string{
				"passionate",
				"reflective",
			},
		},
		Domain{
			Name: "nature",
			AppearanceTraits: []string{
				"barefoot",
				"leaf-haired",
				"bark-skinned",
			},
			PersonalityTraits: []string{
				"down-to-earth",
				"moody",
				"practical",
			},
		},
		Domain{
			Name: "nobility",
			AppearanceTraits: []string{
				"wears a crown",
			},
			PersonalityTraits: []string{
				"arrogant",
			},
		},
		Domain{
			Name: "oceans",
			AppearanceTraits: []string{
				"fishtailed",
				"seaweed for hair",
				"eyes like the ocean",
			},
			PersonalityTraits: []string{
				"commanding",
				"mercurial",
				"passionate",
			},
		},
		Domain{
			Name: "persistence",
			AppearanceTraits: []string{
				"heavily scarred",
				"perpetually wounded",
			},
			PersonalityTraits: []string{
				"enduring",
				"persistent",
			},
		},
		Domain{
			Name: "plants",
			AppearanceTraits: []string{
				"leaves for hair",
				"pale green skin",
				"vines for hair",
			},
			PersonalityTraits: []string{
				"patient",
			},
		},
		Domain{
			Name: "protection",
			AppearanceTraits: []string{
				"armored",
				"enormous",
			},
			PersonalityTraits: []string{
				"enduring",
				"protective",
				"strong-willed",
			},
		},
		Domain{
			Name: "revenge",
			AppearanceTraits: []string{
				"bald",
				"reed-thin",
			},
			PersonalityTraits: []string{
				"brooding",
				"perpetually angry",
				"perpetually vexed",
				"vengeful",
			},
		},
		Domain{
			Name: "sky",
			AppearanceTraits: []string{
				"sits on a cloud",
			},
			PersonalityTraits: []string{
				"wistful",
			},
		},
		Domain{
			Name: "spring",
			AppearanceTraits: []string{
				"damp",
				"clothed in flowers",
				"wears a crown of flowers",
			},
			PersonalityTraits: []string{
				"energetic",
				"frisky",
			},
		},
		Domain{
			Name: "strength",
			AppearanceTraits: []string{
				"heavily-muscled",
				"powerfully built",
			},
			PersonalityTraits: []string{
				"blunt",
				"boisterous",
				"courageous",
			},
		},
		Domain{
			Name: "summer",
			AppearanceTraits: []string{
				"orange-haired",
				"red-haired",
				"perpetually sunburnt",
			},
			PersonalityTraits: []string{
				"warm",
			},
		},
		Domain{
			Name: "thieves",
			AppearanceTraits: []string{
				"clad in dark leather",
				"cloaked",
				"thin",
			},
			PersonalityTraits: []string{
				"avaricious",
			},
		},
		Domain{
			Name: "the moon",
			AppearanceTraits: []string{
				"cloaked in silver",
				"naked",
				"pale",
			},
			PersonalityTraits: []string{
				"sad",
				"thoughtful",
			},
		},
		Domain{
			Name: "the sun",
			AppearanceTraits: []string{
				"glowing",
				"ringed in a firey halo",
			},
			PersonalityTraits: []string{
				"courageous",
				"warm",
			},
		},
		Domain{
			Name: "thunder",
			AppearanceTraits: []string{
				"barrel-chested",
			},
			PersonalityTraits: []string{
				"booming",
				"loud",
			},
		},
		Domain{
			Name: "time",
			AppearanceTraits: []string{
				"ancient-looking",
				"wizened",
				"youthful",
			},
			PersonalityTraits: []string{
				"aloof",
				"unknowable",
				"wise",
			},
		},
		Domain{
			Name: "trade",
			AppearanceTraits: []string{
				"fat",
				"rotund",
				"round",
				"well-fed",
			},
			PersonalityTraits: []string{
				"calculating",
				"greedy",
				"measured",
			},
		},
		Domain{
			Name: "travel",
			AppearanceTraits: []string{
				"dusty",
			},
			PersonalityTraits: []string{
				"wandering",
			},
		},
		Domain{
			Name: "trickery",
			AppearanceTraits: []string{
				"cloaked in night",
				"shadowy",
			},
			PersonalityTraits: []string{
				"cunning",
				"mischievious",
				"puzzling",
				"tricky",
			},
		},
		Domain{
			Name: "war",
			AppearanceTraits: []string{
				"bloody",
				"crowned with knives",
				"spined",
			},
			PersonalityTraits: []string{
				"aggressive",
				"bellicose",
				"belligerent",
			},
		},
		Domain{
			Name: "water",
			AppearanceTraits: []string{
				"blue-skinned",
				"hairless",
				"half fish",
			},
			PersonalityTraits: []string{
				"mercurial",
				"moody",
			},
		},
		Domain{
			Name: "winter",
			AppearanceTraits: []string{
				"completely white",
				"covered in frost",
				"pale blue skin",
				"white eyes",
			},
			PersonalityTraits: []string{
				"cold",
				"relentless",
				"unforgiving",
			},
		},
		Domain{
			Name: "wisdom",
			AppearanceTraits: []string{
				"frail",
				"leans on a staff",
				"old and wizened",
			},
			PersonalityTraits: []string{
				"careful",
				"knowing",
				"wise",
			},
		},
	}

	return domains
}
