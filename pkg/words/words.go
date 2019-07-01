package words

import (
	"strings"

	"github.com/ironarachne/world/pkg/slices"
)

// CapitalizeFirst capitalizes the first letter of the string
func CapitalizeFirst(phrase string) string {
	firstLetter := string(phrase[0])
	rest := strings.TrimPrefix(phrase, firstLetter)

	result := strings.Title(firstLetter) + rest

	return result
}

// CombinePhrases combines phrases with "and" and commas
func CombinePhrases(phrases []string) string {
	combined := ""
	for index, i := range phrases {
		if index == len(phrases)-1 && len(phrases) > 2 {
			combined += "and " + i
		} else if index == len(phrases)-1 && len(phrases) == 2 {
			combined += " and " + i
		} else if index < len(phrases)-1 && len(phrases) > 2 {
			combined += i + ", "
		} else {
			combined += i
		}
	}

	return combined
}

// Pronoun returns the right singular pronoun based on starting letter of a word
func Pronoun(word string) string {
	if slices.StringIn(string(word[0]), []string{"a", "e", "i", "o", "u"}) {
		return "an"
	}

	return "a"
}
