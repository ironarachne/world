package plant

import (
	"fmt"
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

	bushes := getBushes()
	plants = append(plants, bushes...)
	cactii := getCactii()
	plants = append(plants, cactii...)
	fibers := getFibers()
	plants = append(plants, fibers...)
	grains := getGrains()
	plants = append(plants, grains...)
	herbs := getHerbs()
	plants = append(plants, herbs...)
	melons := getMelons()
	plants = append(plants, melons...)
	squash := getSquash()
	plants = append(plants, squash...)
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

// RandomPlantWithResource returns a random plant with a resource of the specified type
func RandomPlantWithResource(resourceTag string) (Plant, error) {
	plants := All()
	filtered := []Plant{}

	for _, p := range plants {
		for _, r := range p.Resources {
			if r.HasTag(resourceTag) {
				if !p.InSlice(filtered) {
					filtered = append(filtered, p)
				}
			}
		}
	}

	if len(filtered) == 0 {
		err := fmt.Errorf("No plant matching tag " + resourceTag + " was found")
		return Plant{}, err
	}

	if len(filtered) == 1 {
		return filtered[0], nil
	}

	plant := filtered[rand.Intn(len(filtered))]

	return plant, nil
}
