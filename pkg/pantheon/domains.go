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
	HolySymbols       []string
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
		{
			Name: "air",
			AppearanceTraits: []string{
				"bird-headed",
				"ethereal",
				"wind-blown",
				"winged",
			},
			HolyItems: []string{
				"feather",
			},
			HolySymbols: []string{
				"cloud",
			},
			PersonalityTraits: []string{
				"mercurial",
			},
		},
		{
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
			HolySymbols: []string{
				"dog's face",
				"beast",
			},
			PersonalityTraits: []string{
				"brutish",
				"short-tempered",
				"skittish",
			},
		},
		{
			Name: "art",
			AppearanceTraits: []string{
				"rainbow-hued",
			},
			HolyItems: []string{
				"brush",
				"quill pen",
			},
			HolySymbols: []string{
				"paint palette",
				"spiral",
			},
			PersonalityTraits: []string{},
		},
		{
			Name: "autumn",
			AppearanceTraits: []string{
				"leaf-haired",
			},
			HolyItems: []string{
				"leaf",
			},
			HolySymbols: []string{
				"leaf",
				"sun setting on the horizon",
			},
			PersonalityTraits: []string{},
		},
		{
			Name: "balance",
			AppearanceTraits: []string{
				"grey-eyed",
			},
			HolyItems: []string{
				"scale",
			},
			HolySymbols: []string{
				"circle divided in half",
			},
			PersonalityTraits: []string{
				"careful",
				"measured",
			},
		},
		{
			Name: "chaos",
			AppearanceTraits: []string{
				"perpetually blurry",
				"shifting",
			},
			HolyItems: []string{
				"eight-pointed star",
			},
			HolySymbols: []string{
				"eight-pointed star",
			},
			PersonalityTraits: []string{
				"unpredictable",
				"wild",
			},
		},
		{
			Name: "darkness",
			AppearanceTraits: []string{
				"dark as night",
				"shadowy",
			},
			HolyItems: []string{
				"cloak",
			},
			HolySymbols: []string{
				"black circle",
				"eclipse",
			},
			PersonalityTraits: []string{},
		},
		{
			Name: "dawn",
			AppearanceTraits: []string{
				"bright",
				"glowing",
				"radiant",
			},
			HolyItems: []string{
				"torch",
			},
			HolySymbols: []string{
				"radiant sun rising",
			},
			PersonalityTraits: []string{},
		},
		{
			Name: "death",
			AppearanceTraits: []string{
				"gaunt",
				"skeletal",
			},
			HolyItems: []string{
				"funerary mask",
				"sickle",
				"scythe",
				"skull",
			},
			HolySymbols: []string{
				"skull",
				"scythe",
			},
			PersonalityTraits: []string{},
		},
		{
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
			HolySymbols: []string{
				"upside-down five-pointed star",
				"evil eye",
			},
			PersonalityTraits: []string{},
		},
		{
			Name: "destruction",
			AppearanceTraits: []string{
				"spined",
			},
			HolyItems: []string{
				"sword",
				"axe",
			},
			HolySymbols: []string{
				"skull",
				"spiked gauntlet",
			},
			PersonalityTraits: []string{
				"wrathful",
			},
		},
		{
			Name: "dusk",
			AppearanceTraits: []string{
				"dark-skinned",
				"twilight-eyed",
			},
			HolyItems: []string{
				"mask",
			},
			HolySymbols: []string{
				"sun setting on the horizon",
				"half-circle",
			},
			PersonalityTraits: []string{
				"brooding",
			},
		},
		{
			Name: "earth",
			AppearanceTraits: []string{
				"stone-skinned",
			},
			HolyItems: []string{
				"pick",
				"mallet",
			},
			HolySymbols: []string{
				"mountain",
			},
			PersonalityTraits: []string{
				"grounded",
				"methodical",
				"practical",
			},
		},
		{
			Name: "evil",
			AppearanceTraits: []string{
				"black-eyed",
				"horned",
				"spider-bodied",
			},
			HolyItems: []string{
				"spiral dagger",
			},
			HolySymbols: []string{
				"evil eye",
				"spider",
			},
			PersonalityTraits: []string{
				"brutal",
				"cruel",
				"vicious",
			},
		},
		{
			Name: "fear",
			AppearanceTraits: []string{
				"shadowy",
				"nebulous",
				"unnaturally thin",
			},
			HolyItems: []string{
				"mask",
			},
			HolySymbols: []string{
				"screaming face",
			},
			PersonalityTraits: []string{
				"cruel",
			},
		},
		{
			Name: "fertility",
			AppearanceTraits: []string{
				"entrancing",
				"hauntingly beautiful",
			},
			HolyItems: []string{
				"egg",
			},
			HolySymbols: []string{
				"pregnant woman",
				"penis",
				"vagina",
			},
			PersonalityTraits: []string{
				"playful",
				"seductive",
			},
		},
		{
			Name: "fire",
			AppearanceTraits: []string{
				"black skinned, with firey veins like magma",
				"ember-eyed",
				"wreathed in flames",
			},
			HolyItems: []string{
				"torch",
				"brazier",
			},
			HolySymbols: []string{
				"fire",
			},
			PersonalityTraits: []string{
				"easily upsettable",
				"firey",
				"passionate",
				"quick to anger",
			},
		},
		{
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
			HolySymbols: []string{
				"fox's face",
			},
			PersonalityTraits: []string{
				"playful",
				"tricksy",
			},
		},
		{
			Name: "good",
			AppearanceTraits: []string{
				"taciturn",
			},
			HolyItems: []string{
				"shield",
			},
			HolySymbols: []string{
				"radiant sun",
			},
			PersonalityTraits: []string{
				"jovial",
				"tolerant",
			},
		},
		{
			Name: "harvests",
			AppearanceTraits: []string{
				"stocky",
			},
			HolyItems: []string{
				"scythe",
				"sickle",
			},
			HolySymbols: []string{
				"sickle",
				"scythe",
			},
			PersonalityTraits: []string{
				"festive",
				"hard-working",
			},
		},
		{
			Name: "healing",
			AppearanceTraits: []string{
				"green-haloed",
			},
			HolyItems: []string{
				"wand",
			},
			HolySymbols: []string{
				"wreath of flowers",
				"open hand",
			},
			PersonalityTraits: []string{
				"caring",
				"compassionate",
				"patient",
			},
		},
		{
			Name: "hope",
			AppearanceTraits: []string{
				"haloed",
			},
			HolyItems: []string{
				"hoop",
			},
			HolySymbols: []string{
				"open hand",
				"circle",
			},
			PersonalityTraits: []string{
				"patient",
			},
		},
		{
			Name: "horses",
			AppearanceTraits: []string{
				"horse-bodied",
				"horse-headed",
			},
			HolyItems: []string{
				"horse skull",
				"horse tooth",
			},
			HolySymbols: []string{
				"horse's face",
				"horse",
			},
			PersonalityTraits: []string{
				"ebullient",
			},
		},
		{
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
			HolySymbols: []string{
				"bow and arrow",
				"beast",
			},
			PersonalityTraits: []string{
				"patient",
				"subtle",
			},
		},
		{
			Name: "justice",
			AppearanceTraits: []string{
				"blind",
				"blindfolded",
			},
			HolyItems: []string{
				"scale",
				"sword",
			},
			HolySymbols: []string{
				"blind man",
				"blind woman",
				"scale",
				"sword",
			},
			PersonalityTraits: []string{
				"callous",
				"even-tempered",
				"fair",
			},
		},
		{
			Name: "knowledge",
			AppearanceTraits: []string{
				"old and wrinkled",
				"bespectacled",
			},
			HolyItems: []string{
				"book",
			},
			HolySymbols: []string{
				"book",
				"open eye",
				"pyramid",
			},
			PersonalityTraits: []string{
				"absent-minded",
				"studious",
			},
		},
		{
			Name: "language",
			AppearanceTraits: []string{
				"covered in mouths",
				"large-mouthed",
				"always speaking",
			},
			HolyItems: []string{
				"book",
				"scroll",
			},
			HolySymbols: []string{
				"speaking face",
				"scroll",
				"book",
			},
			PersonalityTraits: []string{
				"social",
			},
		},
		{
			Name: "law",
			AppearanceTraits: []string{
				"grim-faced",
			},
			HolyItems: []string{
				"book",
				"sword",
				"hammer",
			},
			HolySymbols: []string{
				"hammer",
				"hammer overlapping a book",
				"crossed sword and hammer",
				"sword",
				"sword overlapping a book",
			},
			PersonalityTraits: []string{
				"fair",
				"stern",
			},
		},
		{
			Name: "life",
			AppearanceTraits: []string{
				"haloed",
				"surrounded by living things",
			},
			HolyItems: []string{
				"egg",
				"seed",
			},
			HolySymbols: []string{
				"ouroboros",
				"circle",
			},
			PersonalityTraits: []string{
				"full of life",
				"joyful",
			},
		},
		{
			Name: "light",
			AppearanceTraits: []string{
				"glowing",
				"white-haired",
			},
			HolyItems: []string{
				"torch",
			},
			HolySymbols: []string{
				"beams of light coming from the sky",
				"radiant sun",
			},
			PersonalityTraits: []string{
				"serious",
			},
		},
		{
			Name: "lightning",
			AppearanceTraits: []string{
				"lightning-eyed",
				"surrounded by crackling sparks",
			},
			HolyItems: []string{
				"rod",
			},
			HolySymbols: []string{
				"lightning bolt",
			},
			PersonalityTraits: []string{
				"emotional",
				"unpredictable",
				"violent",
			},
		},
		{
			Name: "love",
			AppearanceTraits: []string{
				"beautiful",
				"gorgeous",
				"pretty",
				"unabashedly naked",
			},
			HolyItems: []string{},
			HolySymbols: []string{
				"heart",
			},
			PersonalityTraits: []string{
				"joyful",
				"lustful",
				"passionate",
			},
		},
		{
			Name: "luck",
			AppearanceTraits: []string{
				"always playing with cards",
				"perpetually grinning",
			},
			HolyItems: []string{
				"rabbit's foot",
				"dice",
				"playing cards",
			},
			HolySymbols: []string{
				"grinning face",
				"face split into a grinning side and a frowning side",
			},
			PersonalityTraits: []string{
				"adventurous",
			},
		},
		{
			Name: "mercy",
			AppearanceTraits: []string{
				"plain",
			},
			HolyItems: []string{
				"blunt sword",
			},
			HolySymbols: []string{
				"open hand",
				"wreath of vines",
			},
			PersonalityTraits: []string{
				"compassionate",
				"fair",
			},
		},
		{
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
			HolySymbols: []string{
				"encircled heart",
				"harp",
				"lute",
				"lyre",
			},
			PersonalityTraits: []string{
				"passionate",
				"reflective",
			},
		},
		{
			Name: "nature",
			AppearanceTraits: []string{
				"barefoot",
				"leaf-haired",
				"bark-skinned",
			},
			HolyItems: []string{
				"staff",
			},
			HolySymbols: []string{
				"tree",
				"flower",
			},
			PersonalityTraits: []string{
				"down-to-earth",
				"moody",
				"practical",
			},
		},
		{
			Name: "nobility",
			AppearanceTraits: []string{
				"wearing a crown",
				"as if carved from marble",
			},
			HolyItems: []string{
				"crown",
			},
			HolySymbols: []string{
				"crown",
			},
			PersonalityTraits: []string{
				"arrogant",
			},
		},
		{
			Name: "oceans",
			AppearanceTraits: []string{
				"surrounded by seaweed",
				"surrounded by a bubble of seawater when on land",
				"fish-tailed",
			},
			HolyItems: []string{
				"shell",
			},
			HolySymbols: []string{
				"wave",
			},
			PersonalityTraits: []string{
				"commanding",
				"mercurial",
				"passionate",
			},
		},
		{
			Name: "persistence",
			AppearanceTraits: []string{
				"heavily scarred",
				"perpetually wounded",
			},
			HolyItems: []string{
				"shield",
			},
			HolySymbols: []string{
				"man standing resolute",
				"fist",
			},
			PersonalityTraits: []string{
				"enduring",
				"persistent",
			},
		},
		{
			Name: "plants",
			AppearanceTraits: []string{
				"wreathed in vines",
				"green-skinned",
				"covered in leaves",
			},
			HolyItems: []string{
				"hoe",
			},
			HolySymbols: []string{
				"plant",
				"tree",
			},
			PersonalityTraits: []string{
				"patient",
			},
		},
		{
			Name: "protection",
			AppearanceTraits: []string{
				"heavily armored with no skin visible",
				"enormous, with metallic skin",
			},
			HolyItems: []string{
				"shield",
			},
			HolySymbols: []string{
				"shield",
				"gauntlet",
			},
			PersonalityTraits: []string{
				"enduring",
				"protective",
				"strong-willed",
			},
		},
		{
			Name: "revenge",
			AppearanceTraits: []string{
				"bald",
				"reed-thin",
			},
			HolyItems: []string{},
			HolySymbols: []string{
				"evil eye",
				"broken circle",
			},
			PersonalityTraits: []string{
				"brooding",
				"perpetually angry",
				"perpetually vexed",
				"vengeful",
			},
		},
		{
			Name: "sky",
			AppearanceTraits: []string{
				"sitting on a cloud",
			},
			HolyItems: []string{},
			HolySymbols: []string{
				"cloud",
			},
			PersonalityTraits: []string{
				"wistful",
			},
		},
		{
			Name: "spring",
			AppearanceTraits: []string{
				"perpetually damp",
				"clothed in flowers",
				"wearing a crown of flowers",
			},
			HolyItems: []string{
				"staff",
			},
			HolySymbols: []string{
				"wreath of flowers",
				"flower",
			},
			PersonalityTraits: []string{
				"energetic",
				"frisky",
			},
		},
		{
			Name: "strength",
			AppearanceTraits: []string{
				"heavily-muscled",
				"powerfully built",
			},
			HolyItems: []string{
				"gauntlet",
			},
			HolySymbols: []string{
				"muscled arm",
				"fist",
				"tower",
			},
			PersonalityTraits: []string{
				"blunt",
				"boisterous",
				"courageous",
			},
		},
		{
			Name: "summer",
			AppearanceTraits: []string{
				"orange-haired",
				"red-haired",
				"perpetually sunburnt",
			},
			HolyItems: []string{
				"wreath",
			},
			HolySymbols: []string{
				"radiant sun",
				"fire",
			},
			PersonalityTraits: []string{
				"warm",
			},
		},
		{
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
			HolySymbols: []string{
				"dagger",
				"pierced circle",
			},
			PersonalityTraits: []string{
				"avaricious",
			},
		},
		{
			Name: "the moon",
			AppearanceTraits: []string{
				"cloaked in silver",
				"naked",
				"pale",
			},
			HolyItems: []string{
				"plate",
			},
			HolySymbols: []string{
				"crescent moon",
			},
			PersonalityTraits: []string{
				"sad",
				"thoughtful",
			},
		},
		{
			Name: "the sun",
			AppearanceTraits: []string{
				"glowing",
				"ringed in a firey halo",
			},
			HolyItems: []string{
				"plate",
			},
			HolySymbols: []string{
				"radiant sun",
			},
			PersonalityTraits: []string{
				"courageous",
				"warm",
			},
		},
		{
			Name: "thunder",
			AppearanceTraits: []string{
				"barrel-chested",
			},
			HolyItems: []string{
				"drum",
				"hammer",
			},
			HolySymbols: []string{
				"lightning bolt",
				"lightning bolt over a hammer",
				"hammer",
				"drum",
			},
			PersonalityTraits: []string{
				"booming",
				"loud",
			},
		},
		{
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
			HolySymbols: []string{
				"hourglass",
				"ouroboros",
			},
			PersonalityTraits: []string{
				"aloof",
				"unknowable",
				"wise",
			},
		},
		{
			Name: "trade",
			AppearanceTraits: []string{
				"fat",
				"rotund",
				"round",
				"well-fed",
			},
			HolyItems: []string{
				"scale",
				"coin",
			},
			HolySymbols: []string{
				"coin",
				"scale",
				"book and scale",
			},
			PersonalityTraits: []string{
				"calculating",
				"greedy",
				"measured",
			},
		},
		{
			Name: "travel",
			AppearanceTraits: []string{
				"dusty",
			},
			HolyItems: []string{
				"staff",
			},
			HolySymbols: []string{
				"man walking",
				"road disappearing into the horizon",
			},
			PersonalityTraits: []string{
				"wandering",
			},
		},
		{
			Name: "trickery",
			AppearanceTraits: []string{
				"cloaked in night",
				"shadowy",
			},
			HolyItems: []string{
				"two-faced mask",
			},
			HolySymbols: []string{
				"face bearing an evil half and a good half",
			},
			PersonalityTraits: []string{
				"cunning",
				"mischievious",
				"puzzling",
				"tricky",
			},
		},
		{
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
			HolySymbols: []string{
				"sword",
				"crossed pair of axes",
			},
			PersonalityTraits: []string{
				"aggressive",
				"bellicose",
				"belligerent",
			},
		},
		{
			Name: "water",
			AppearanceTraits: []string{
				"blue-skinned",
				"hairless",
				"half fish",
			},
			HolyItems: []string{
				"flask",
			},
			HolySymbols: []string{
				"wave",
				"fountain",
			},
			PersonalityTraits: []string{
				"mercurial",
				"moody",
			},
		},
		{
			Name: "winter",
			AppearanceTraits: []string{
				"completely white",
				"covered in frost",
				"blue-skinned",
				"white eyes",
			},
			HolyItems: []string{
				"cloak",
			},
			HolySymbols: []string{
				"snowflake",
				"icicle",
			},
			PersonalityTraits: []string{
				"cold",
				"relentless",
				"unforgiving",
			},
		},
		{
			Name: "wisdom",
			AppearanceTraits: []string{
				"frail",
				"always leaning on a staff",
				"old and wizened",
			},
			HolyItems: []string{
				"book",
			},
			HolySymbols: []string{
				"book",
				"scroll",
				"book and scroll",
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
