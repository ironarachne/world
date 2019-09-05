package pantheon

import (
	"fmt"
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
func Generate(minSize int, maxSize int, lang language.Language) (Pantheon, error) {
	var pantheon Pantheon
	numberOfDeities := minSize

	pantheon.Deities = make(map[string]Deity)

	if maxSize-minSize > 0 {
		numberOfDeities = rand.Intn(maxSize-minSize) + minSize
	}

	if numberOfDeities < 1 {
		return pantheon, nil
	}

	for i := 0; i < numberOfDeities; i++ {
		deity, err := pantheon.GenerateDeity(lang)
		if err != nil {
			err = fmt.Errorf("Could not generate pantheon: %w", err)
			return Pantheon{}, err
		}
		pantheon.Deities[deity.Name] = deity
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
