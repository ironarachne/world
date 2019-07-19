package mineral

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/resource"
)

// Mineral is a mineral
type Mineral struct {
	Name         string
	PluralName   string
	Hardness     int
	Malleability int
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
		mineral = from[rand.Intn(len(from)-1)]
		if !InSlice(mineral, minerals) {
			minerals = append(minerals, mineral)
		}
	}

	return minerals
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
