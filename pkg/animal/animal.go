package animal

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
	"github.com/ironarachne/world/pkg/species"
)

// All returns all pre-defined animals
func All() []Species {
	var animals []Species

	birds := getBirds()
	animals = append(animals, birds...)
	gameBirds := getGameBirds()
	animals = append(animals, gameBirds...)
	raptors := getRaptors()
	animals = append(animals, raptors...)
	cats := getBigCats()
	animals = append(animals, cats...)
	canines := getCanines()
	animals = append(animals, canines...)
	equines := getEquines()
	animals = append(animals, equines...)
	mammals := getMammals()
	animals = append(animals, mammals...)
	reptiles := getReptiles()
	animals = append(animals, reptiles...)

	return animals
}

// ByTag returns a slice of animals that have the given tag
func ByTag(tag string, from []Species) []Species {
	var animals []Species

	for _, p := range from {
		if p.HasTag(tag) {
			animals = append(animals, p)
		}
	}

	return animals
}

// HasTag returns true if the animal has a given tag
func (animal Species) HasTag(tag string) bool {
	for _, t := range animal.Tags {
		if t == tag {
			return true
		}
	}

	return false
}

// InSlice checks whether a given animal is in a slice of animals
func (animal Species) InSlice(animals []Species) bool {
	isIt := false
	for _, a := range animals {
		if a.Name == animal.Name {
			isIt = true
		}
	}

	return isIt
}

// Random returns a set number of randomly chosen animals from a slice
func Random(amount int, from []Species) []Species {
	var animal Species

	animals := []Species{}

	if amount > len(from) {
		amount = len(from)
	}

	for i := 0; i < amount; i++ {
		animal = from[rand.Intn(len(from))]
		if !animal.InSlice(animals) {
			animals = append(animals, animal)
		}
	}

	return animals
}
