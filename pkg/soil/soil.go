/*
Package soil implements soil types
*/
package soil

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
)

// Data is a collection of soils
type Data struct {
	Soils []Soil `json:"soils"`
}

// Soil is a type of mineral
type Soil struct {
	Name               string              `json:"name"`
	UsedForPottery     bool                `json:"used_for_pottery"`
	IsSand             bool                `json:"is_sand"`
	UsedForAgriculture bool                `json:"used_for_agriculture"`
	Resources          []resource.Resource `json:"resources"`
	Tags               []string            `json:"tags"`
}

// All returns all predefined minerals from a JSON file on disk
func All() ([]Soil, error) {
	var d Data

	jsonFile, err := os.Open(os.Getenv("WORLDAPI_DATA_PATH") + "/data/soils.json")
	if err != nil {
		err = fmt.Errorf("could not open data file: %w", err)
		return []Soil{}, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &d)

	all := d.Soils

	if len(all) == 0 {
		err = fmt.Errorf("no soils returned from database: soils.json")
		return []Soil{}, err
	}

	return all, nil
}

// ByTag returns a set of soils that match a tag
func ByTag(tag string, from []Soil) ([]Soil, error) {
	var filtered []Soil

	for _, s := range from {
		if s.HasTag(tag) {
			filtered = append(filtered, s)
		}
	}

	return filtered, nil
}

// ByTagIn returns a slice of soils that have at least one of the given tags
func ByTagIn(tags []string, from []Soil) []Soil {
	var filtered []Soil

	for _, s := range from {
		for _, t := range tags {
			if s.HasTag(t) && !InSlice(s, filtered) {
				filtered = append(filtered, s)
			}
		}
	}

	return filtered
}

// HasTag returns true if the soil has a tag
func (s Soil) HasTag(tag string) bool {
	for _, t := range s.Tags {
		if t == tag {
			return true
		}
	}

	return false
}

// Random returns a random subset of soils
func Random(ctx context.Context, amount int, from []Soil) []Soil {
	var soil Soil

	soils := []Soil{}

	if amount > len(from) {
		amount = len(from)
	}

	if len(from) == 0 {
		return soils
	}

	for i := 0; i < amount; i++ {
		soil = from[random.Intn(ctx, len(from))]
		if !InSlice(soil, soils) {
			soils = append(soils, soil)
		}
	}

	return soils
}

// InSlice checks if a soil is in a slice
func InSlice(soil Soil, soils []Soil) bool {
	isIt := false
	for _, a := range soils {
		if a.Name == soil.Name {
			isIt = true
		}
	}

	return isIt
}
