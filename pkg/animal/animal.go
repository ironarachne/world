package animal

import "math/rand"

// Animal is an animal
type Animal struct {
	Name           string
	PluralName     string
	AnimalType     string
	EatsAnimals    bool
	EatsPlants     bool
	GivesBone      bool
	GivesEggs      bool
	GivesFur       bool
	GivesHide      bool
	GivesHorn      bool
	GivesMeat      bool
	GivesMilk      bool
	GivesWool      bool
	IsAquatic      bool
	IsMount        bool
	IsPackAnimal   bool
	IsScavenger    bool
	IsVenomous     bool
	MinHumidity    int
	MaxHumidity    int
	MinTemperature int
	MaxTemperature int
}

// All returns all pre-defined animals
func All() []Animal {
	var animals []Animal

	birds := getBirds()
	animals = append(animals, birds...)
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
