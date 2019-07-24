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

// Birds returns all birds
func Birds() []Animal {
	var animals []Animal

	birds := getBirds()
	animals = append(animals, birds...)
	gameBirds := getGameBirds()
	animals = append(animals, gameBirds...)
	raptors := getRaptors()
	animals = append(animals, raptors...)

	return animals
}

// Land returns all land animals
func Land() []Animal {
	var animals []Animal

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
		animal = from[rand.Intn(len(from)-1)]
		if !animal.InSlice(animals) {
			animals = append(animals, animal)
		}
	}

	return animals
}
