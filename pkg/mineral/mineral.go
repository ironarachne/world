/*
Package mineral provides minerals and tools for their usage
*/
package mineral

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
)

// Data is a collection of minerals
type Data struct {
	Minerals []Mineral `json:"minerals"`
}

// Mineral is a mineral
type Mineral struct {
	Name         string              `json:"name"`
	PluralName   string              `json:"plural_name"`
	Hardness     int                 `json:"hardness"`
	Malleability int                 `json:"malleability"`
	Commonality  int                 `json:"commonality"`
	Resources    []resource.Resource `json:"resources"`
	Tags         []string            `json:"tags"`
}

// All returns all predefined minerals from a JSON file on disk
func All() ([]Mineral, error) {
	var d Data

	jsonFile, err := os.Open(os.Getenv("WORLDAPI_DATA_PATH") + "/data/minerals.json")
	if err != nil {
		err = fmt.Errorf("could not open data file: %w", err)
		return []Mineral{}, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &d)

	all := d.Minerals

	if len(all) == 0 {
		err = fmt.Errorf("no minerals returned from database: minerals.json")
		return []Mineral{}, err
	}

	return all, nil
}

// ByTag returns a slice of minerals that have the given tag
func ByTag(tag string, from []Mineral) []Mineral {
	var filtered []Mineral

	for _, s := range from {
		if s.HasTag(tag) {
			filtered = append(filtered, s)
		}
	}

	return filtered
}

// ByTagIn returns a slice of minerals that have at least one of the given tags
func ByTagIn(tags []string, from []Mineral) []Mineral {
	var filtered []Mineral

	for _, s := range from {
		for _, t := range tags {
			if s.HasTag(t) && !InSlice(s, filtered) {
				filtered = append(filtered, s)
			}
		}
	}

	return filtered
}

// HasTag returns true if the mineral has a given tag
func (m Mineral) HasTag(tag string) bool {
	for _, t := range m.Tags {
		if t == tag {
			return true
		}
	}

	return false
}

// Random returns a subset of random minerals from a slice
func Random(ctx context.Context, amount int, from []Mineral) []Mineral {
	var mineral Mineral

	minerals := []Mineral{}

	if amount > len(from) {
		amount = len(from)
	}

	for i := 0; i < amount; i++ {
		mineral = from[random.Intn(ctx, len(from))]
		if !InSlice(mineral, minerals) {
			minerals = append(minerals, mineral)
		}
	}

	return minerals
}

// RandomWeighted returns a random mineral from a slice, considering weights
func RandomWeighted(ctx context.Context, from []Mineral) (Mineral, error) {
	names := make(map[string]int)

	for _, m := range from {
		names[m.Name] = m.Commonality
	}

	name, err := random.StringFromThresholdMap(ctx, names)
	if err != nil {
		err = fmt.Errorf("Failed to get random weighted mineral: %w", err)
		return Mineral{}, err
	}

	for _, m := range from {
		if m.Name == name {
			return m, nil
		}
	}

	err = fmt.Errorf("Couldn't find named mineral")

	return Mineral{}, err
}

// RandomWeightedSet returns a set of random weighted minerals
func RandomWeightedSet(ctx context.Context, amount int, from []Mineral) ([]Mineral, error) {
	var minerals []Mineral

	for i := 0; i < amount; i++ {
		mineral, err := RandomWeighted(ctx, from)
		if err != nil {
			err = fmt.Errorf("Could not generate minerals: %w", err)
			return []Mineral{}, err
		}
		if !InSlice(mineral, minerals) {
			minerals = append(minerals, mineral)
		} else {
			i--
		}
	}

	return minerals, nil
}

// InSlice checks to see if a mineral is in a slice
func InSlice(mineral Mineral, minerals []Mineral) bool {
	isIt := false
	for _, m := range minerals {
		if m.Name == mineral.Name {
			isIt = true
		}
	}

	return isIt
}
