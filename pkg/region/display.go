package region

import (
	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/organization"
	"github.com/ironarachne/world/pkg/town"
)

// SimplifiedRegion is a simplified version of a region
type SimplifiedRegion struct {
	Name string `json:"name"`
	Climate string `json:"climate"`
	Capital string `json:"capital"`
	Ruler character.SimplifiedCharacter `json:"ruler"`
	Towns []town.SimplifiedTown `json:"towns"`
	Organizations []organization.SimplifiedOrganization `json:"organizations"`
}

// Simplify returns a simplified version of a region
func (region Region) Simplify() SimplifiedRegion {
	simplified := SimplifiedRegion{
		Name: "The " + region.Class.Name + " of " + region.Name,
		Climate: region.Climate.Description,
		Capital: region.Capital,
		Ruler: region.Ruler.Simplify(),
	}

	for _, t := range region.Towns {
		simplified.Towns = append(simplified.Towns, t.Simplify())
	}

	for _, o := range region.Organizations {
		simplified.Organizations = append(simplified.Organizations, o.Simplify())
	}

	return simplified
}

// RandomSimplified generates a completely random region
func RandomSimplified() SimplifiedRegion {
	region := Random()

	return region.Simplify()
}