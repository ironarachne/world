package animal

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

// Animal is an animal
type Animal struct {
	Name           string
	PluralName     string
	AnimalType     string
	EatsAnimals    bool
	EatsPlants     bool
	IsAquatic      bool
	IsMount        bool
	IsPackAnimal   bool
	IsScavenger    bool
	MinHumidity    int
	MaxHumidity    int
	MinTemperature int
	MaxTemperature int
	Resources      []resource.Resource
	Size           size.Category
	Tags           []string
}

// All returns all pre-defined animals
func All() []Animal {
	var animals []Animal

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
func ByTag(tag string, from []Animal) []Animal {
	var animals []Animal

	for _, p := range from {
		if p.HasTag(tag) {
			animals = append(animals, p)
		}
	}

	return animals
}

// HasTag returns true if the animal has a given tag
func (animal Animal) HasTag(tag string) bool {
	for _, t := range animal.Tags {
		if t == tag {
			return true
		}
	}

	return false
}

// InSlice checks whether a given animal is in a slice of animals
func (animal Animal) InSlice(animals []Animal) bool {
	isIt := false
	for _, a := range animals {
		if a.Name == animal.Name {
			isIt = true
		}
	}

	return isIt
}

// Random returns a set number of randomly chosen animals from a slice
func Random(amount int, from []Animal) []Animal {
	var animal Animal

	animals := []Animal{}

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
