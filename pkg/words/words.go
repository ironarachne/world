/*
Package words provides convenience methods for dealing with words and phrases
*/
package words

import (
	"math"
	"strings"

	"github.com/ironarachne/world/pkg/slices"
)

const (
	negativeSign = "minus"
	hundred      = "hundred"
)

var (
	numbers0To9 = []string{
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
	}

	numbers10To19 = []string{
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}

	numbersTenths = []string{
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}

	numbersMagnitudes = []string{
		"",
		"thousand",
		"million",
		"billion",
		"trillion",
		"quadrillion",
		"quintillion", // this is the largest magnitude a 64-bit integer can reach
	}
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
func NumberToWord(number int) string {
	var out strings.Builder
	if number < 0 {
		out.WriteString(negativeSign)
		out.WriteRune(' ')
		number = -number
	}
	switch {
	case number < 10:
		out.WriteString(numbers0To9[number])
	case number < 20:
		out.WriteString(numbers10To19[number-10])
	case number < 100:
		out.WriteString(numbersTenths[(number-20)/10])
		firstDigit := number % 10
		if firstDigit > 0 {
			out.WriteRune('-')
			out.WriteString(numbers0To9[firstDigit])
		}
	case number < 1000:
		hundredCount := number / 100
		out.WriteString(numbers0To9[hundredCount])
		out.WriteRune(' ')
		out.WriteString(hundred)
		rest := number % 100
		if rest > 0 {
			out.WriteRune(' ')
			out.WriteString(NumberToWord(rest))
		}
	default:
		var magnitude int
		for n := number / 1000; n > 0; n, magnitude = n/1000, magnitude+1 {
		}
		for ; magnitude > 0; magnitude-- {
			base := int(math.Pow(1000, float64(magnitude)))
			if base > number {
				continue
			}
			magnitudeCount := number / base
			out.WriteString(NumberToWord(magnitudeCount))
			out.WriteRune(' ')
			out.WriteString(numbersMagnitudes[magnitude])
			number = number % base
			if number >= 1000 {
				out.WriteRune(' ')
			}
		}
		if number > 0 {
			out.WriteRune(' ')
			out.WriteString(NumberToWord(number))
		}
	}
	return out.String()
}

// Pronoun returns the right singular pronoun based on starting letter of a word
func Pronoun(word string) string {
	if len(word) > 0 && slices.StringIn(string(word[0]), []string{"a", "e", "i", "o", "u"}) {
		return "an"
	}

	return "a"
}
