/*
Package pantheon implements fantasy pantheons
*/
package pantheon

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/pantheon/deity"
	"github.com/ironarachne/world/pkg/pantheon/domain"
)

// Pantheon is a nonhierarchical group of deities
type Pantheon struct {
	Deities []deity.Deity
}

// Generate creates a random pantheon of deities
func Generate(minSize int, maxSize int, lang language.Language) (Pantheon, error) {
	var pantheon Pantheon
	numberOfDeities := minSize

	if maxSize-minSize > 0 {
		numberOfDeities = rand.Intn(maxSize-minSize) + minSize
	}

	if numberOfDeities < 1 {
		return pantheon, nil
	}

	possibleDomains := domain.All()

	for i := 0; i < numberOfDeities; i++ {
		d, err := deity.Generate(lang, possibleDomains)
		if err != nil {
			err = fmt.Errorf("Could not generate pantheon: %w", err)
			return Pantheon{}, err
		}
		pantheon.Deities = append(pantheon.Deities, d)
		possibleDomains = domain.Exclude(d.Domains, possibleDomains)
	}

	if len(pantheon.Deities) > 1 {
		deities, err := pantheon.GenerateRelationships()
		if err != nil {
			err = fmt.Errorf("Could not generate pantheon: %w", err)
			return Pantheon{}, err
		}
		pantheon.Deities = deities
	}

	return pantheon, nil
}
