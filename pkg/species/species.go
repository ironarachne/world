package species

import (
	"github.com/ironarachne/world/pkg/age"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/trait"
	"math/rand"
)

// Species is a species of living thing
type Species struct {
	Name           string
	PluralName     string
	Adjective      string
	Commonality    int
	PossibleTraits []trait.Template // Traits that individuals of this species *might* have
	CommonTraits   []trait.Template // Traits that all members of this species share
	AgeCategories  []age.Category
	MinHumidity    int
	MaxHumidity    int
	MinTemperature int
	MaxTemperature int
	Resources      []resource.Resource // These are resources that can be derived from this species
	Tags           []string
}

// ByResource returns a slice of species that have the given resource
func ByResource(resource string, from []Species) []Species {
	var filtered []Species

	for _, s := range from {
		if s.HasResource(resource) {
			filtered = append(filtered, s)
		}
	}

	return filtered
}

// ByTag returns a slice of species that have the given tag
func ByTag(tag string, from []Species) []Species {
	var filtered []Species

	for _, s := range from {
		if s.HasTag(tag) {
			filtered = append(filtered, s)
		}
	}

	return filtered
}

// ExcludeTag returns a slice without a given tag
func ExcludeTag(tag string, from []Species) []Species {
	var filtered []Species

	for _, s := range from {
		if !s.HasTag(tag) {
			filtered = append(filtered, s)
		}
	}

	return filtered
}

// FilterHumidity filters a slice of Species by humidity
func FilterHumidity(humidity int, from []Species) []Species {
	var filtered []Species

	for _, s := range from {
		if s.MinHumidity <= humidity && s.MaxHumidity >= humidity {
			filtered = append(filtered, s)
		}
	}

	return filtered
}

// FilterTemperature filters a slice of Species by temperature
func FilterTemperature(temperature int, from []Species) []Species {
	var filtered []Species

	for _, s := range from {
		if s.MinTemperature <= temperature && s.MaxTemperature >= temperature {
			filtered = append(filtered, s)
		}
	}

	return filtered
}

// HasResource returns true if the animal has a given resource
func (s Species) HasResource(resource string) bool {
	for _, r := range s.Resources {
		if r.Name == resource {
			return true
		}
	}

	return false
}

// HasTag returns true if the animal has a given tag
func (s Species) HasTag(tag string) bool {
	for _, t := range s.Tags {
		if t == tag {
			return true
		}
	}

	return false
}

// InSlice checks whether a given animal is in a slice of animals
func (s Species) InSlice(from []Species) bool {
	isIt := false
	for _, a := range from {
		if a.Name == s.Name {
			isIt = true
		}
	}

	return isIt
}

// Random returns a set number of randomly chosen animals from a slice
func Random(amount int, from []Species) []Species {
	var s Species
	var result []Species

	if amount > len(from) {
		amount = len(from)
	}

	for i := 0; i < amount; i++ {
		s = from[rand.Intn(len(from))]
		if !s.InSlice(result) {
			result = append(result, s)
		}
	}

	return result
}