package language

import (
	"github.com/ironarachne/world/pkg/words"
)

// SimplifiedLanguage is a simplified version of Language meant for display
type SimplifiedLanguage struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

// Describe describes a Language as a string
func (language Language) Describe() string {
	description := language.Name + " is " + words.Pronoun(language.Descriptors[0]) + " " + words.CombinePhrases(language.Descriptors) + " language. "
	description += "It uses " + words.Pronoun(language.WritingSystem.Classification) + " " + language.WritingSystem.Classification + " writing system, "
	description += "with characters composed of " + language.WritingSystem.StrokeType + " strokes."

	return description
}

// Simplify creates a simplified version of a language for display
func (language Language) Simplify() SimplifiedLanguage {
	return SimplifiedLanguage{
		Name: language.Name,
		Description: language.Describe(),
	}
}