package pantheon

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/language"
)

// Pantheon is a nonhierarchical group of deities
type Pantheon struct {
	Deities map[string]Deity
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

	pantheon.Deities = pantheon.GenerateRelationships()

	return pantheon
}
