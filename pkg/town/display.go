package town

import (
	"strconv"
)

// SimplifiedTown is a simpler version of a town
type SimplifiedTown struct {
	Name string `json:"name"`
	Population int `json:"population"`
	Climate string `json:"climate"`
	DominantCulture string `json:"dominant_culture"`
	Category string `json:"category"`
	MayorName string `json:"mayor_name"`
	MayorDescription string `json:"mayor_description"`
	Producers []string `json:"producers"`
	Exports []string `json:"exports"`
	Imports []string `json:"imports"`
}

// RandomSimplified generates a random simplified town
func RandomSimplified() SimplifiedTown {
	town := Random()

	return town.Simplify()
}

// Simplify returns the simplified version of a town
func (town Town) Simplify() SimplifiedTown {
	simplified := SimplifiedTown{
		Name: town.Name,
		Population: town.Population,
		Climate: town.Climate.Description,
		DominantCulture: town.Culture.Name,
		Category: town.Category.Name,
		MayorName: town.Mayor.FirstName + " " + town.Mayor.LastName,
		MayorDescription: strconv.Itoa(town.Mayor.Age) + "-year-old " + town.Mayor.Gender.Noun + " who is " + town.Mayor.PositiveTraits[0] + " but " + town.Mayor.NegativeTraits[0],
	}

	for _, p := range town.NotableProducers {
		simplified.Producers = append(simplified.Producers, p.Skill() + " " + p.Name)
	}

	for _, e := range town.Exports {
		simplified.Exports = append(simplified.Exports, e.Quality + " " + e.Name)
	}

	for _, i := range town.Imports {
		simplified.Imports = append(simplified.Imports, i.Name)
	}

	return simplified
}