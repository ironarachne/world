package town

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/character"
)

// SimplifiedTown is a simpler version of a town
type SimplifiedTown struct {
	Name            string                        `json:"name"`
	Population      int                           `json:"population"`
	BuildingStyle   string                        `json:"building_style"`
	Climate         string                        `json:"climate"`
	DominantCulture string                        `json:"dominant_culture"`
	Category        string                        `json:"category"`
	Mayor           character.SimplifiedCharacter `json:"mayor"`
	Exports         []string                      `json:"exports"`
	Imports         []string                      `json:"imports"`
}

// RandomSimplified generates a random simplified town
func RandomSimplified(ctx context.Context) (SimplifiedTown, error) {
	town, err := Random(ctx)
	if err != nil {
		err = fmt.Errorf("Could not simplify town: %w", err)
		return SimplifiedTown{}, err
	}

	st, err := town.Simplify(ctx)
	if err != nil {
		err = fmt.Errorf("Could not simplify town: %w", err)
		return SimplifiedTown{}, err
	}

	return st, nil
}

// Simplify returns the simplified version of a town
func (town Town) Simplify(ctx context.Context) (SimplifiedTown, error) {
	mayor, err := town.Mayor.Simplify(ctx)
	if err != nil {
		err = fmt.Errorf("Could not simplify town: %w", err)
		return SimplifiedTown{}, err
	}
	simplified := SimplifiedTown{
		Name:            town.Name,
		Population:      town.Population,
		BuildingStyle:   town.BuildingStyle.Simplify().Description,
		Climate:         town.Geography.Biome.Name,
		DominantCulture: town.Culture.Name,
		Category:        town.Category.Name,
		Mayor:           mayor,
	}

	for _, e := range town.Exports {
		simplified.Exports = append(simplified.Exports, e.Quality+e.Name)
	}

	for _, i := range town.Imports {
		simplified.Imports = append(simplified.Imports, i.Name)
	}

	return simplified, nil
}
