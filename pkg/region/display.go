package region

import (
	"fmt"

	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/organization"
	"github.com/ironarachne/world/pkg/town"
)

// SimplifiedRegion is a simplified version of a region
type SimplifiedRegion struct {
	Name            string                                `json:"name"`
	Climate         string                                `json:"climate"`
	Capital         string                                `json:"capital"`
	DominantCulture culture.Culture                       `json:"dominant_culture"`
	Ruler           character.SimplifiedCharacter         `json:"ruler"`
	RulingHouse     organization.SimplifiedOrganization   `json:"ruling_house"`
	Towns           []town.SimplifiedTown                 `json:"towns"`
	Organizations   []organization.SimplifiedOrganization `json:"organizations"`
}

// Simplify returns a simplified version of a region
func (region Region) Simplify() (SimplifiedRegion, error) {
	sc := region.Culture
	sr, err := region.RulingBody.Leader.CharacterData.Simplify()
	if err != nil {
		err = fmt.Errorf("Could not generate simplified region: %w", err)
		return SimplifiedRegion{}, err
	}
	so, err := region.RulingBody.Simplify()
	if err != nil {
		err = fmt.Errorf("Could not generate simplified region: %w", err)
		return SimplifiedRegion{}, err
	}
	simplified := SimplifiedRegion{
		Name:            "The " + region.Class.Name + " of " + region.Name,
		Climate:         region.Geography.Biome.Name,
		Capital:         region.Capital,
		DominantCulture: sc,
		Ruler:           sr,
		RulingHouse:     so,
	}

	for _, t := range region.Towns {
		st, err := t.Simplify()
		if err != nil {
			err = fmt.Errorf("Could not generate simplified region: %w", err)
			return SimplifiedRegion{}, err
		}
		simplified.Towns = append(simplified.Towns, st)
	}

	for _, o := range region.Organizations {
		so, err := o.Simplify()
		if err != nil {
			err = fmt.Errorf("Could not generate simplified region: %w", err)
			return SimplifiedRegion{}, err
		}
		simplified.Organizations = append(simplified.Organizations, so)
	}

	return simplified, nil
}

// RandomSimplified generates a completely random region
func RandomSimplified() (SimplifiedRegion, error) {
	region, err := Random()
	if err != nil {
		err = fmt.Errorf("Could not generate simplified region: %w", err)
		return SimplifiedRegion{}, err
	}

	rs, err := region.Simplify()
	if err != nil {
		err = fmt.Errorf("Could not generate simplified region: %w", err)
		return SimplifiedRegion{}, err
	}

	return rs, nil
}
