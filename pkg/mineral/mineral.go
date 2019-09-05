package mineral

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
)

// Mineral is a mineral
type Mineral struct {
	Name         string
	PluralName   string
	Hardness     int
	Malleability int
	Commonality  int
	Resources    []resource.Resource
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

	name := random.StringFromThresholdMap(names)

	for _, m := range from {
		if m.Name == name {
			return m, nil
		}
	}

	err := fmt.Errorf("Couldn't find named mineral")

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

// All returns all minerals
func All() []Mineral {
	var minerals []Mineral

	gems := Gems()
	metal := Metals()
	other := OtherMinerals()
	stone := Stones()

	minerals = append(minerals, gems...)
	minerals = append(minerals, metal...)
	minerals = append(minerals, other...)
	minerals = append(minerals, stone...)

	return minerals
}
