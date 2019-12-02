/*
Package soil implements soil types
*/
package soil

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/resource"
)

// Soil is a type of mineral
type Soil struct {
	Name               string              `json:"name"`
	UsedForPottery     bool                `json:"used_for_pottery"`
	IsSand             bool                `json:"is_sand"`
	UsedForAgriculture bool                `json:"used_for_agriculture"`
	Resources          []resource.Resource `json:"resources"`
	Tags               []string            `json:"tags"`
}

// All returns all soils
func All() []Soil {
	clays := Clays()
	sands := Sands()
	agSoils := Agricultural()

	soils := append(clays, sands...)
	soils = append(soils, agSoils...)

	return soils
}

// ByTag returns a set of soils that match a tag
func ByTag(tag string) []Soil {
	var filtered []Soil

	soils := All()

	for _, s := range soils {
		if s.HasTag(tag) {
			filtered = append(filtered, s)
		}
	}

	return filtered
}

// ByTagIn returns a slice of soils that have at least one of the given tags
func ByTagIn(tags []string, from []Soil) []Soil {
	var filtered []Soil

	for _, s := range from {
		for _, t := range tags {
			if s.HasTag(t) && !InSlice(s, filtered) {
				filtered = append(filtered, s)
			}
		}
	}

	return filtered
}

// HasTag returns true if the soil has a tag
func (s Soil) HasTag(tag string) bool {
	for _, t := range s.Tags {
		if t == tag {
			return true
		}
	}

	return false
}

// Random returns a random subset of soils
func Random(amount int, from []Soil) []Soil {
	var soil Soil

	soils := []Soil{}

	if amount > len(from) {
		amount = len(from)
	}

	if len(from) == 0 {
		return soils
	}

	for i := 0; i < amount; i++ {
		soil = from[rand.Intn(len(from))]
		if !InSlice(soil, soils) {
			soils = append(soils, soil)
		}
	}

	return soils
}

// InSlice checks if a soil is in a slice
func InSlice(soil Soil, soils []Soil) bool {
	isIt := false
	for _, a := range soils {
		if a.Name == soil.Name {
			isIt = true
		}
	}

	return isIt
}
