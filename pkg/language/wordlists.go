package language

// WordType is a type of word. This is used as rules to construct words.
type WordType struct {
	Name         string
	NumSyllables int
	Prefix       string
	Suffix       string
	WordList     []string
}

func getAllWordTypes() []WordType {
	words := []WordType{}
	words = append(words, getAdjectives())
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
		NumSyllables: 2,
		Prefix:       "",
		Suffix:       "",
		WordList: []string{
			"big",
			"black",
			"blue",
			"brown",
			"cold",
			"dark",
			"deep",
			"divine",
			"drunk",
			"empty",
			"evil",
			"familiar",
			"fat",
			"frail",
			"full",
			"good",
			"green",
			"grey",
			"honest",
			"hot",
			"light",
			"loud",
			"mortal",
			"mysterious",
			"narrow",
			"old",
			"orange",
			"purple",
			"quiet",
			"red",
			"shallow",
			"short",
			"sober",
			"strange",
			"strong",
			"sturdy",
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

func getInterjections() WordType {
	return WordType{
		Name:         "interjection",
		NumSyllables: 2,
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
		NumSyllables: 1,
		Prefix:       "",
		Suffix:       "",
		WordList: []string{
			"a",
			"and",
			"as",
			"from",
			"in",
			"of",
			"or",
			"to",
			"will",
		},
	}
}

func getQuestionWords() WordType {
	return WordType{
		Name:         "question word",
		NumSyllables: 1,
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
		NumSyllables: 2,
		Prefix:       "",
		Suffix:       "",
		WordList: []string{
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
			"to hear",
			"to hide",
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
			"to run",
			"to see",
			"to sit",
			"to sleep",
			"to smell",
			"to smile",
			"to stand",
			"to swallow",
			"to swim",
			"to taste",
			"to walk",
			"to want",
		},
	}
}

func getNouns() WordType {
	return WordType{
		Name:         "noun",
		NumSyllables: 3,
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
			"nose",
			"ocean",
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
			"sand",
			"seed",
			"skin",
			"sky",
			"smoke",
			"snake",
			"star",
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
			"valley",
			"war",
			"water",
			"wine",
			"woman",
		},
	}
}

func getNumbers() WordType {
	return WordType{
		Name:         "number",
		NumSyllables: 1,
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
		NumSyllables: 1,
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

func isInWordList(word string, wordList map[string]string) bool {
	for _, w := range wordList {
		if w == word {
			return true
		}
	}
	return false
}
