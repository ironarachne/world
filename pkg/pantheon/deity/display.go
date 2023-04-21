package deity

import (
	"github.com/ironarachne/world/pkg/words"
)

// SimplifiedDeity is a display version of deity
type SimplifiedDeity struct {
	Name        string   `json:"name"`
	Gender      string   `json:"gender"`
	Domains     []string `json:"domains"`
	Description string   `json:"description"`
}

func (deity Deity) getDomainNames() []string {
	names := []string{}

	for _, d := range deity.Domains {
		names = append(names, d.Name)
	}

	return names
}

// Describe describes a deity
func (deity Deity) Describe() string {
	var relationship string

	description := deity.Name

	if len(deity.Domains) > 0 {
		description += " is the god"
		if deity.Gender.Name == "female" {
			description += "dess"
		}
		description += " of " + words.CombinePhrases(deity.getDomainNames()) + ". "
	} else {
		description += " is a god"
		if deity.Gender.Name == "female" {
			description += "dess"
		}
		description += " of no particular domain. "
	}

	description += words.Title(deity.Gender.SubjectPronoun) + " is " + deity.Appearance + ". "

	description += words.Title(deity.Gender.SubjectPronoun) + " is " + words.CombinePhrases(deity.PersonalityTraits) + ". "

	description += words.Title(deity.Gender.PossessivePronoun) + " holy item is " + words.Pronoun(deity.HolyItem) + " " + deity.HolyItem + ", and "
	description += deity.Gender.PossessivePronoun + " holy symbol is " + words.Pronoun(deity.HolySymbol) + " " + deity.HolySymbol + ". "

	relationships := []string{}

	if len(deity.Relationships) > 0 {
		for _, r := range deity.Relationships {
			relationship = r.Descriptor + " " + r.Target
			relationships = append(relationships, relationship)
		}

		description += deity.Name + " " + words.CombinePhrases(relationships) + "."
	}

	return description
}

// Simplify returns a simplified version of a deity
func (deity Deity) Simplify() SimplifiedDeity {
	domains := deity.getDomainNames()
	description := deity.Describe()

	return SimplifiedDeity{
		Name:        deity.Name,
		Gender:      deity.Gender.Name,
		Domains:     domains,
		Description: description,
	}
}
