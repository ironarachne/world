package soil

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/resource"
)

// Soil is a type of mineral
type Soil struct {
	Name               string
	UsedForPottery     bool
	IsSand             bool
	UsedForAgriculture bool
	Resources          []resource.Resource
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
		soil = from[rand.Intn(len(from)-1)]
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
