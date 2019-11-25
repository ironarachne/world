/*
Package plant provides plant implementation of species.Species
*/
package plant

import (
	"github.com/ironarachne/world/pkg/species"
)

// All returns all predefined plants
func All() []species.Species {
	var plants []species.Species

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
