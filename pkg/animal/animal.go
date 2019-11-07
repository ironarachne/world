package animal

import "github.com/ironarachne/world/pkg/species"

// All returns all pre-defined animals
func All() []species.Species {
	var animals []species.Species

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
