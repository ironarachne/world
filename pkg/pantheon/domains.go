package pantheon

import (
	"math/rand"
)

// Domain is an area of control
type Domain struct {
	Name              string
	AppearanceTraits  []string
	PersonalityTraits []string
	HolyItems         []string
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

func getDomainNames(domains []Domain) []string {
	names := []string{}

	for _, d := range domains {
		names = append(names, d.Name)
	}

	return names
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
			HolyItems: []string{
				"feather",
				"crystal orb",
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
			HolyItems: []string{
				"talon",
				"fang",
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
			HolyItems: []string{
				"brush",
				"quill pen",
			},
			PersonalityTraits: []string{},
		},
		Domain{
			Name: "autumn",
			AppearanceTraits: []string{
				"leaf-haired",
			},
			HolyItems: []string{
				"leaf",
			},
			PersonalityTraits: []string{},
		},
		Domain{
			Name:             "balance",
			AppearanceTraits: []string{},
			HolyItems: []string{
				"scales",
			},
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
			HolyItems: []string{
				"eight-pointed star",
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
			HolyItems: []string{
				"cloak",
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
			HolyItems: []string{
				"torch",
			},
			PersonalityTraits: []string{},
		},
		Domain{
			Name: "death",
			AppearanceTraits: []string{
				"gaunt",
				"skeletal",
			},
			HolyItems: []string{
				"funerary mask",
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
			HolyItems: []string{
				"three-legged bowl",
				"cauldron",
			},
			PersonalityTraits: []string{},
		},
		Domain{
			Name: "destruction",
			AppearanceTraits: []string{
				"spined",
			},
			HolyItems: []string{
				"sword",
				"axe",
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
			HolyItems: []string{
				"mask",
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
			HolyItems: []string{
				"pick",
				"mallet",
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
			HolyItems: []string{
				"spiral dagger",
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
			HolyItems: []string{
				"mask",
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
			HolyItems: []string{
				"egg",
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
			HolyItems: []string{
				"torch",
				"brazier",
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
			HolyItems: []string{
				"fox tail",
				"fox paw",
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
			HolyItems: []string{
				"shield",
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
			HolyItems: []string{
				"sickle",
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
			HolyItems: []string{
				"wand",
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
			HolyItems: []string{
				"hoop",
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
			HolyItems: []string{
				"horse skull",
				"horse tooth",
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
			HolyItems: []string{
				"bow",
				"arrow",
				"spear",
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
			HolyItems: []string{
				"scales",
				"sword",
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
			HolyItems: []string{
				"book",
			},
			PersonalityTraits: []string{
				"absent-minded",
				"studious",
			},
		},
		Domain{
			Name:             "language",
			AppearanceTraits: []string{},
			HolyItems: []string{
				"book",
			},
			PersonalityTraits: []string{
				"social",
			},
		},
		Domain{
			Name:             "law",
			AppearanceTraits: []string{},
			HolyItems: []string{
				"book",
				"sword",
				"hammer",
			},
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
			HolyItems: []string{
				"egg",
				"seed",
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
			HolyItems: []string{
				"torch",
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
			HolyItems: []string{
				"rod",
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
			HolyItems: []string{},
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
			HolyItems: []string{
				"rabbit's foot",
				"dice",
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
			HolyItems: []string{
				"blunt sword",
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
			HolyItems: []string{
				"harp",
				"lute",
				"lyre",
				"drum",
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
			HolyItems: []string{
				"staff",
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
			HolyItems: []string{
				"crown",
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
			HolyItems: []string{
				"shell",
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
			HolyItems: []string{
				"shield",
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
			HolyItems: []string{
				"hoe",
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
			HolyItems: []string{
				"shield",
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
			HolyItems: []string{},
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
			HolyItems: []string{},
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
			HolyItems: []string{
				"staff",
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
			HolyItems: []string{
				"gauntlet",
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
			HolyItems: []string{
				"wreath",
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
			HolyItems: []string{
				"dagger",
				"coin purse",
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
			HolyItems: []string{
				"plate",
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
			HolyItems: []string{
				"plate",
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
			HolyItems: []string{
				"drum",
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
			HolyItems: []string{
				"sundial",
				"hourglass",
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
			HolyItems: []string{
				"scales",
				"coin",
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
			HolyItems: []string{
				"staff",
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
			HolyItems: []string{
				"two-faced mask",
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
			HolyItems: []string{
				"sword",
				"spear",
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
			HolyItems: []string{
				"flask",
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
			HolyItems: []string{
				"cloak",
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
			HolyItems: []string{
				"book",
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
