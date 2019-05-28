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
func Generate(maxSize int, lang language.Language) Pantheon {
	var deity Deity
	var pantheon Pantheon

	pantheon.Deities = make(map[string]Deity)

	numberOfDeities := rand.Intn(maxSize) + 1

	for i := 0; i < numberOfDeities; i++ {
		deity = pantheon.GenerateDeity(lang)
		pantheon.Deities[deity.Name] = deity
	}

	if len(pantheon.Deities) > 1 {
		pantheon.Deities = pantheon.GenerateRelationships()
	}

	return pantheon
}

// GenerateForDisplay returns a simplified pantheon for display purposes
func GenerateForDisplay(maxSize int, lang language.Language) SimplifiedPantheon {
	pantheon := Generate(maxSize, lang)

	display := pantheon.simplify()

	return display
}

func (pantheon Pantheon) simplify() SimplifiedPantheon {
	sp := SimplifiedPantheon{}

	for _, d := range pantheon.Deities {
		sp.Deities = append(sp.Deities, d.simplify())
	}

	return sp
}
