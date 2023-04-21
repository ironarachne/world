package words

import (
	"strconv"
	"testing"
)

func TestCombinePhrases(t *testing.T) {
	tests := []struct {
		input []string
		want  string
	}{
		{
			input: []string{"one", "two", "three"},
			want:  "one, two, and three",
		},
		{
			input: []string{"one", "two"},
			want:  "one and two",
		},
		{
			input: []string{"one"},
			want:  "one",
		},
		{
			input: []string{},
			want:  "",
		},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(len(test.input)), func(t *testing.T) {
			got := CombinePhrases(test.input)

			if got != test.want {
				t.Errorf("CombinePhrases() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestNumberToWord(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{
			input: 0,
			want:  "zero",
		},
		{
			input: 5,
			want:  "five",
		},
		{
			input: 13,
			want:  "thirteen",
		},
		{
			input: 30,
			want:  "thirty",
		},
		{
			input: 56,
			want:  "fifty-six",
		},
		{
			input: 400,
			want:  "four hundred",
		},
		{
			input: 740,
			want:  "seven hundred forty",
		},
		{
			input: 839,
			want:  "eight hundred thirty-nine",
		},
		{
			input: 1000,
			want:  "one thousand",
		},
		{
			input: 3543,
			want:  "three thousand five hundred forty-three",
		},
		{
			input: 123_456,
			want:  "one hundred twenty-three thousand four hundred fifty-six",
		},
		{
			input: 5_000_000_000,
			want:  "five billion",
		},
		{
			input: 5_000_003_000,
			want:  "five billion three thousand",
		},
		{
			input: 1_000_000_000_000_000_000,
			want:  "one quintillion",
		},
		{
			input: 1<<63 - 1,
			want: "nine quintillion two hundred twenty-three quadrillion three hundred seventy-two " +
				"trillion thirty-six billion eight hundred fifty-four million seven hundred " +
				"seventy-five thousand eight hundred seven",
		},
		{
			input: -999,
			want:  "minus nine hundred ninety-nine",
		},
	}
	for _, test := range tests {
		t.Run(strconv.Itoa(test.input), func(t *testing.T) {
			got := NumberToWord(test.input)
			if !(test.want == got) {
				t.Errorf("want: %s\ngot : %s", test.want, got)
			}
		})
	}
}

func TestPronoun(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "apple",
			want: "an",
		},
		{
			input: "heart",
			want: "a",
		},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := Pronoun(test.input)
			if !(test.want == got) {
				t.Errorf("want: %s\ngot : %s", test.want, got)
			}
		})
	}
}

func TestTitle(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "hello world",
			want:  "Hello World",
		},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := Title(test.input)
			if !(test.want == got) {
				t.Errorf("want: %s\ngot : %s", test.want, got)
			}
		})
	}
}
