/*
Package pantheon implements fantasy pantheons
*/
package pantheon

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/pantheon/deity"
	"github.com/ironarachne/world/pkg/pantheon/domain"
	"github.com/ironarachne/world/pkg/random"
)

const pantheonError = "failed to generate pantheon: %w"

// Pantheon is a non-hierarchical group of deities
type Pantheon struct {
	Deities []deity.Deity `json:"deities"`
}

// Generate creates a random pantheon of deities
func Generate(ctx context.Context, minSize int, maxSize int, lang language.Language) (Pantheon, error) {
	var pantheon Pantheon
	numberOfDeities := minSize

	if maxSize-minSize > 0 {
		numberOfDeities = random.Intn(ctx, maxSize-minSize) + minSize
	}

	if numberOfDeities < 1 {
		return pantheon, nil
	}

	possibleDomains, err := domain.All()
	if err != nil {
		err = fmt.Errorf(pantheonError, err)
		return Pantheon{}, err
	}

	for i := 0; i < numberOfDeities; i++ {
		d, err := deity.Generate(ctx, lang, possibleDomains)
		if err != nil {
			err = fmt.Errorf(pantheonError, err)
			return Pantheon{}, err
		}
		pantheon.Deities = append(pantheon.Deities, d)
		possibleDomains = domain.Exclude(d.Domains, possibleDomains)
	}

	if len(pantheon.Deities) > 1 {
		deities, err := pantheon.GenerateRelationships(ctx)
		if err != nil {
			err = fmt.Errorf(pantheonError, err)
			return Pantheon{}, err
		}
		pantheon.Deities = deities
	}

	return pantheon, nil
}
