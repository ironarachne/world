package language

import (
	"strings"
)

// RosettaStone generates a version of a common phrase in the language
func (language Language) RosettaStone() string {
	// Phrase: Hello! It is good to see you, friend.
	translation := ""

	translation += strings.Title(language.WordList["hello"]) + "! "
	translation += strings.Title(language.WordList["it"]) + " " + language.Conjugate(language.WordList["to be"], "simple present") + " "
	translation += language.WordList["good"] + " " + language.WordList["to see"] + " " + language.WordList["you"] + ", " + language.WordList["friend"] + "."

	return translation
}
