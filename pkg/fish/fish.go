package fish

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/slices"
)

// Fish is a fish
type Fish struct {
	Name             string
	PluralName       string
	IdealTemperature int
	LivesInLakes     bool
	LivesInOceans    bool
	LivesInRivers    bool
	Resources        []resource.Resource
	Tags             []string
}

// ByTag returns all fish from the given slice that match the tag
func ByTag(tag string, fish []Fish) []Fish {
	var result []Fish

	for _, f := range fish {
		if slices.StringIn(tag, f.Tags) {
			result = append(result, f)
		}
	}

	return result
}

// ExcludeTag returns a filtered slice of fish without the given tag
func ExcludeTag(tag string, fish []Fish) []Fish {
	var result []Fish

	for _, f := range fish {
		if !slices.StringIn(tag, f.Tags) {
			result = append(result, f)
		}
	}

	return result
}

// All returns all predefined fish
func All() []Fish {
	fish := []Fish{}

	fish = append(fish, getLakeFish()...)
	fish = append(fish, getOceanFish()...)
	fish = append(fish, getRiverFish()...)

	for _, f := range fish {
		f.Resources = []resource.Resource{
			{
				Name:         f.Name,
				Origin:       f.Name,
				MainMaterial: f.Name,
				Tags: []string{
					"meat",
				},
				Commonality: 4,
				Value:       1,
			},
		}
	}

	return fish
}

// InSlice checks to see if a fish is in the given slice
func InSlice(fish Fish, fishes []Fish) bool {
	isIt := false
	for _, a := range fishes {
		if a.Name == fish.Name {
			isIt = true
		}
	}

	return isIt
}

// Random returns a random fish from a slice
func Random(amount int, from []Fish) []Fish {
	var fish Fish

	randomFish := []Fish{}

	for i := 0; i < amount; i++ {
		fish = from[rand.Intn(len(from))]
		if !InSlice(fish, randomFish) {
			randomFish = append(randomFish, fish)
		} else {
			i--
		}
	}

	return randomFish
}
