package insect

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/resource"
)

// Insect is a freaking bug
type Insect struct {
	Name           string
	PluralName     string
	MinHumidity    int
	MaxHumidity    int
	MinTemperature int
	MaxTemperature int
	Tags           []string
	Resources      []resource.Resource
}

// All returns all insects
func All() []Insect {
	arachnids := getArachnids()
	beetles := getBeetles()

	insects := getInsects()

	insects = append(insects, arachnids...)
	insects = append(insects, beetles...)

	return insects
}

// ByTag returns a slice of insects that have the given tag
func ByTag(tag string, from []Insect) []Insect {
	var insects []Insect

	for _, p := range from {
		if p.HasTag(tag) {
			insects = append(insects, p)
		}
	}

	return insects
}

// HasTag returns true if the insect has a given tag
func (insect Insect) HasTag(tag string) bool {
	for _, t := range insect.Tags {
		if t == tag {
			return true
		}
	}

	return false
}

// InSlice checks whether a given insect is in a slice of insects
func (insect Insect) InSlice(insects []Insect) bool {
	isIt := false
	for _, a := range insects {
		if a.Name == insect.Name {
			isIt = true
		}
	}

	return isIt
}

// RandomSubset returns a set number of randomly chosen insects from a slice
func RandomSubset(amount int, from []Insect) []Insect {
	var insect Insect

	insects := []Insect{}

	if amount > len(from) {
		amount = len(from)
	}

	for i := 0; i < amount; i++ {
		insect = from[rand.Intn(len(from)-1)]
		if !insect.InSlice(insects) {
			insects = append(insects, insect)
		}
	}

	return insects
}
