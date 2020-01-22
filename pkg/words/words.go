/*
Package words provides convenience methods for dealing with words and phrases
*/
package words

import (
  "fmt"
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

// NumberToWord returns the word version of a number
func NumberToWord(number int) (string, error) {
  var err error
	result := ""

	if number < 0 {
		err = fmt.Errorf("can't turn negative number into word")
		return "", err
	}

	if number > 20 {
		err = fmt.Errorf("no support for numbers greater than 20")
		return "", err
	}

	switch number {
	case 0:
		result = "zero"
	case 1:
		result = "one"
	case 2:
		result = "two"
	case 3:
		result = "three"
	case 4:
		result = "four"
	case 5:
		result = "five"
	case 6:
		result = "six"
	case 7:
		result = "seven"
	case 8:
		result = "eight"
	case 9:
		result = "nine"
	case 10:
		result = "ten"
	case 11:
		result = "eleven"
	case 12:
		result = "twelve"
	case 13:
		result = "thirteen"
	case 14:
		result = "fourteen"
	case 15:
		result = "fifteen"
	case 16:
		result = "sixteen"
	case 17:
		result = "seventeen"
	case 18:
		result = "eighteen"
	case 19:
		result = "nineteen"
	case 20:
		result = "twenty"
	}

	return result, nil
}

// Pronoun returns the right singular pronoun based on starting letter of a word
func Pronoun(word string) string {
	if len(word) > 0 && slices.StringIn(string(word[0]), []string{"a", "e", "i", "o", "u"}) {
		return "an"
	}

	return "a"
}
