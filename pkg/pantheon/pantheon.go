package pantheon

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/language"
)

// Pantheon is a nonhierarchical group of deities
type Pantheon struct {
	Deities map[string]Deity
}

// SimplifiedPantheon is a simplified version of the data for display
type SimplifiedPantheon struct {
	Deities []SimplifiedDeity `json:"deities"`
}

// Generate creates a random pantheon of deities
func Generate(minSize int, maxSize int, lang language.Language) Pantheon {
	var deity Deity
	var pantheon Pantheon
	numberOfDeities := minSize

	pantheon.Deities = make(map[string]Deity)

	if maxSize-minSize > 0 {
		numberOfDeities = rand.Intn(maxSize-minSize) + minSize
	}

	if numberOfDeities < 1 {
		return pantheon
	}

	for i := 0; i < numberOfDeities; i++ {
		deity = pantheon.GenerateDeity(lang)
		pantheon.Deities[deity.Name] = deity
	}

	if len(pantheon.Deities) > 1 {
		pantheon.Deities = pantheon.GenerateRelationships()
	}

	return pantheon
}

// Simplify returns a simplified pantheon for display
func (pantheon Pantheon) Simplify() SimplifiedPantheon {
	sp := SimplifiedPantheon{}
	simplifiedDeity := SimplifiedDeity{}

	for _, d := range pantheon.Deities {
		simplifiedDeity = d.Simplify()
		sp.Deities = append(sp.Deities, simplifiedDeity)
	}

	return sp
}
