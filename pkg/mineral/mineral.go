/*
Package mineral provides minerals and tools for their usage
*/
package mineral

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
)

// Mineral is a mineral
type Mineral struct {
	Name         string              `json:"name"`
	PluralName   string              `json:"plural_name"`
	Hardness     int                 `json:"hardness"`
	Malleability int                 `json:"malleability"`
	Commonality  int                 `json:"commonality"`
	Resources    []resource.Resource `json:"resources"`
	Tags         []string            `json:"tags"`
}

// All returns all minerals
func All() []Mineral {
	var all []Mineral

	gems := Gems()
	metals := Metals()
	other := OtherMinerals()
	stones := Stones()

	all = append(all, gems...)
	all = append(all, metals...)
	all = append(all, other...)
	all = append(all, stones...)

	return all
}

// ByTag returns a slice of minerals that have the given tag
func ByTag(tag string, from []Mineral) []Mineral {
	var filtered []Mineral

	for _, s := range from {
		if s.HasTag(tag) {
			filtered = append(filtered, s)
		}
	}

	return filtered
}

// ByTagIn returns a slice of minerals that have at least one of the given tags
func ByTagIn(tags []string, from []Mineral) []Mineral {
	var filtered []Mineral

	for _, s := range from {
		for _, t := range tags {
			if s.HasTag(t) && !InSlice(s, filtered) {
				filtered = append(filtered, s)
			}
		}
	}

	return filtered
}

// HasTag returns true if the mineral has a given tag
func (m Mineral) HasTag(tag string) bool {
	for _, t := range m.Tags {
		if t == tag {
			return true
		}
	}

	return false
}

// Random returns a subset of random minerals from a slice
func Random(amount int, from []Mineral) []Mineral {
	var mineral Mineral

	minerals := []Mineral{}

	if amount > len(from) {
		amount = len(from)
	}

	for i := 0; i < amount; i++ {
		mineral = from[rand.Intn(len(from))]
		if !InSlice(mineral, minerals) {
			minerals = append(minerals, mineral)
		}
	}

	return minerals
}

// RandomWeighted returns a random mineral from a slice, considering weights
func RandomWeighted(from []Mineral) (Mineral, error) {
	names := make(map[string]int)

	for _, m := range from {
		names[m.Name] = m.Commonality
	}

	name, err := random.StringFromThresholdMap(names)
	if err != nil {
		err = fmt.Errorf("Failed to get random weighted mineral: %w", err)
		return Mineral{}, err
	}

	for _, m := range from {
		if m.Name == name {
			return m, nil
		}
	}

	err = fmt.Errorf("Couldn't find named mineral")

	return Mineral{}, err
}

// RandomWeightedSet returns a set of random weighted minerals
func RandomWeightedSet(amount int, from []Mineral) ([]Mineral, error) {
	var minerals []Mineral

	for i := 0; i < amount; i++ {
		mineral, err := RandomWeighted(from)
		if err != nil {
			err = fmt.Errorf("Could not generate minerals: %w", err)
			return []Mineral{}, err
		}
		if !InSlice(mineral, minerals) {
			minerals = append(minerals, mineral)
		} else {
			i--
		}
	}

	return minerals, nil
}

// InSlice checks to see if a mineral is in a slice
func InSlice(mineral Mineral, minerals []Mineral) bool {
	isIt := false
	for _, m := range minerals {
		if m.Name == mineral.Name {
			isIt = true
		}
	}

	return isIt
}
