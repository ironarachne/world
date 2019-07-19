package plant

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/resource"
)

// Plant is a plant
type Plant struct {
	Name           string
	PluralName     string
	MinHumidity    int
	MaxHumidity    int
	MinTemperature int
	MaxTemperature int
	Resources      []resource.Resource
}

// All returns all predefined plants
func All() []Plant {
	var plants []Plant

	fibers := getFibers()
	plants = append(plants, fibers...)
	fruits := getFruits()
	plants = append(plants, fruits...)
	grains := getGrains()
	plants = append(plants, grains...)
	herbs := getHerbs()
	plants = append(plants, herbs...)
	vegetables := getVegetables()
	plants = append(plants, vegetables...)

	return plants
}

// InSlice checks to see if the given plant is in the slice
func (plant Plant) InSlice(plants []Plant) bool {
	isIt := false
	for _, a := range plants {
		if a.Name == plant.Name {
			isIt = true
		}
	}

	return isIt
}

// Random returns a random subset of plants
func Random(amount int, from []Plant) []Plant {
	var plant Plant

	plants := []Plant{}

	if amount > len(from) {
		amount = len(from)
	}

	for i := 1; i < amount; i++ {
		plant = from[rand.Intn(len(from))]
		if !plant.InSlice(plants) {
			plants = append(plants, plant)
		}
	}

	return plants
}

// RandomFabric returns a random fabric plant
func RandomFabric() Plant {
	var resources []resource.Resource

	fibers := getFibers()

	for _, p := range fibers {
		resources = append(resources, p.Resources...)
	}
	fabrics := resource.ListOfType("fabric", resources)

	fabric := fabrics[rand.Intn(len(fabrics))]

	for _, p := range fibers {
		if p.Name == fabric.Origin {
			return p
		}
	}

	panic("Lost my fabric!")
}
