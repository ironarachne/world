package conlang

// WordType is a type of word. This is used as rules to construct words.
type WordType struct {
	Name         string
	MaxSyllables int
	Prefix       string
	Suffix       string
	WordList     []string
}

func getAllWordTypes() []WordType {
	words := []WordType{}
	words = append(words, getAdjectives())
	words = append(words, getArticles())
	words = append(words, getInterjections())
	words = append(words, getNouns())
	words = append(words, getNumbers())
	words = append(words, getPrepositions())
	words = append(words, getPronouns())
	words = append(words, getQuestionWords())
	words = append(words, getVerbs())

	return words
}

func getAdjectives() WordType {
	return WordType{
		Name:         "adjective",
		MaxSyllables: 2,
		Prefix:       "",
		Suffix:       "",
		WordList: []string{
			"aromatic",
			"basted",
			"big",
			"bitter",
			"black",
			"blue",
			"brown",
			"chilled",
			"cold",
			"curried",
			"dark",
			"deep",
			"divine",
			"drunk",
			"empty",
			"evil",
			"familiar",
			"fat",
			"flat",
			"frail",
			"fried",
			"full",
			"good",
			"green",
			"grey",
			"honest",
			"hot",
			"light",
			"long",
			"loud",
			"mortal",
			"mysterious",
			"narrow",
			"old",
			"orange",
			"pungent",
			"purple",
			"quiet",
			"raw",
			"rectangular",
			"red",
			"roasted",
			"round",
			"salty",
			"savory",
			"shallow",
			"short",
			"smoked",
			"sober",
			"sour",
			"spicy",
			"spiral",
			"square",
			"steamed",
			"strange",
			"strong",
			"sturdy",
			"sweet",
			"tall",
			"thick",
			"thin",
			"warm",
			"weak",
			"white",
			"wide",
			"yellow",
			"young",
		},
	}
}

func getArticles() WordType {
	return WordType{
		Name:         "article",
		MaxSyllables: 1,
		Prefix:       "",
		Suffix:       "",
		WordList: []string{
			"a",
			"an",
			"the",
		},
	}
}

func getInterjections() WordType {
	return WordType{
		Name:         "interjection",
		MaxSyllables: 2,
		Prefix:       "",
		Suffix:       "",
		WordList: []string{
			"hello",
			"goodbye",
		},
	}
}

func getPrepositions() WordType {
	return WordType{
		Name:         "preposition",
		MaxSyllables: 1,
		Prefix:       "",
		Suffix:       "",
		WordList: []string{
			"and",
			"as",
			"from",
			"in",
			"of",
			"or",
			"to",
			"will",
			"with",
		},
	}
}

func getQuestionWords() WordType {
	return WordType{
		Name:         "question word",
		MaxSyllables: 1,
		Prefix:       "",
		Suffix:       "",
		WordList: []string{
			"who",
			"what",
			"where",
			"why",
			"how",
		},
	}
}

func getVerbs() WordType {
	return WordType{
		Name:         "verb",
		MaxSyllables: 2,
		Prefix:       "",
		Suffix:       "",
		WordList: []string{
			"to bake",
			"to be",
			"to belong",
			"to bite",
			"to break",
			"to burn",
			"to come",
			"to die",
			"to drink",
			"to eat",
			"to fall",
			"to fight",
			"to find",
			"to fish",
			"to fly",
			"to frown",
			"to go",
			"to growl",
			"to hate",
			"to have",
			"to hear",
			"to hide",
			"to hold",
			"to hunt",
			"to jump",
			"to kill",
			"to know",
			"to laugh",
			"to lie (down)",
			"to lie",
			"to live",
			"to lose",
			"to love",
			"to need",
			"to own",
			"to roast",
			"to run",
			"to see",
			"to sit",
			"to sleep",
			"to smell",
			"to smile",
			"to stand",
			"to strike",
			"to swallow",
			"to swim",
			"to taste",
			"to throw",
			"to walk",
			"to want",
		},
	}
}

func getNouns() WordType {
	return WordType{
		Name:         "noun",
		MaxSyllables: 3,
		Prefix:       "",
		Suffix:       "",
		WordList: []string{
			"afternoon",
			"ale",
			"all",
			"alligator",
			"arm",
			"ash",
			"axe",
			"bark",
			"bay",
			"beer",
			"bird",
			"blood",
			"boar",
			"bone",
			"breakfast",
			"breast",
			"castle",
			"cat",
			"cat",
			"chest",
			"chicken",
			"claw",
			"cloud",
			"coconut",
			"crime",
			"day",
			"dinner",
			"dog",
			"drink",
			"dungeon",
			"ear",
			"earth",
			"egg",
			"enemy",
			"evening",
			"eye",
			"feather",
			"fight",
			"finger",
			"fire",
			"fish",
			"flesh",
			"foot",
			"forest",
			"friend",
			"goose",
			"grease",
			"gulf",
			"hair",
			"hand",
			"hat",
			"hate",
			"head",
			"horn",
			"horse",
			"house",
			"inn",
			"island",
			"jaw",
			"lake",
			"leaf",
			"leg",
			"liver",
			"louse",
			"love",
			"lunch",
			"man",
			"many",
			"meal",
			"metal",
			"mine",
			"monster",
			"moon",
			"morning",
			"mountain",
			"mouth",
			"name",
			"neck",
			"night",
			"noodle",
			"nose",
			"ocean",
			"path",
			"pepper",
			"person",
			"pie",
			"pig",
			"rabbit",
			"rain",
			"rat",
			"river",
			"robe",
			"rock",
			"root",
			"salt",
			"sand",
			"seed",
			"skin",
			"sky",
			"smoke",
			"snake",
			"soup",
			"star",
			"stew",
			"stomach",
			"stone",
			"stream",
			"sun",
			"sword",
			"tail",
			"tavern",
			"thumb",
			"tongue",
			"tooth",
			"tree",
			"truth",
			"valley",
			"war",
			"water",
			"way",
			"wine",
			"woman",
		},
	}
}

func getNumbers() WordType {
	return WordType{
		Name:         "number",
		MaxSyllables: 1,
		Prefix:       "",
		Suffix:       "",
		WordList: []string{
			"zero",
			"one",
			"two",
			"three",
			"four",
			"five",
			"six",
			"seven",
			"eight",
			"nine",
			"ten",
		},
	}
}

func getPronouns() WordType {
	return WordType{
		Name:         "pronoun",
		MaxSyllables: 1,
		Prefix:       "",
		Suffix:       "",
		WordList: []string{
			"he",
			"her",
			"his",
			"I",
			"it",
			"me",
			"my",
			"she",
			"their",
			"they",
			"you",
		},
	}
}

// IsInWordList checks to see if a given word is in a word list
func IsInWordList(word string, wordList map[string]string) bool {
	for _, w := range wordList {
		if w == word {
			return true
		}
	}
	return false
}
