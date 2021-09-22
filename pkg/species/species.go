/*
Package species implements the backbone of all living entities in a world
*/
package species

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ironarachne/world/pkg/age"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/trait"
)

// Data is a struct containing a slice of Species
type Data struct {
	Species []Species `json:"species"`
}

// Species is a species of living thing
type Species struct {
	Name           string              `json:"name" db:"name"`
	PluralName     string              `json:"plural_name" db:"plural_name"`
	Adjective      string              `json:"adjective" db:"adjective"`
	Commonality    int                 `json:"commonality" db:"commonality"`
	PossibleTraits []trait.Template    `json:"possible_traits" db:"possible_traits"` // Traits that individuals of this species *might* have
	CommonTraits   []trait.Template    `json:"common_traits" db:"common_traits"`     // Traits that all members of this species share
	AgeCategories  []age.Category      `json:"age_categories" db:"age_categories"`
	MinHumidity    int                 `json:"min_humidity" db:"humidity_min"`
	MaxHumidity    int                 `json:"max_humidity" db:"humidity_max"`
	MinTemperature int                 `json:"min_temperature" db:"temperature_min"`
	MaxTemperature int                 `json:"max_temperature" db:"temperature_max"`
	Resources      []resource.Resource `json:"resources" db:"resources"` // These are resources that can be derived from this species
	Tags           []string            `json:"tags" db:"tags"`
}

// Load returns all predefined species of a given type from a JSON file on disk
func Load(fileName string) ([]Species, error) {
	var d Data

	jsonFile, err := os.Open(os.Getenv("WORLDAPI_DATA_PATH") + "/data/" + fileName + ".json")
	if err != nil {
		err = fmt.Errorf("could not open data file: %w", err)
		return []Species{}, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &d)
	if err != nil {
		err = fmt.Errorf("could not process species data: %w", err)
		return []Species{}, err
	}

	all := d.Species

	if len(all) == 0 {
		err = fmt.Errorf("no species returned from database: " + fileName + ".json")
		return []Species{}, err
	}

	return all, nil
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

// ByTagIn returns a slice of species that have at least one of the given tags
func ByTagIn(tags []string, from []Species) []Species {
	var filtered []Species

	for _, s := range from {
		for _, t := range tags {
			if s.HasTag(t) && !s.InSlice(filtered) {
				filtered = append(filtered, s)
			}
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

// HasResource returns true if the species has a given resource
func (s Species) HasResource(resource string) bool {
	for _, r := range s.Resources {
		if r.Name == resource {
			return true
		}
	}

	return false
}

// HasResourceWithTag returns true if the species has a resource with a given tag
func (s Species) HasResourceWithTag(resourceTag string) bool {
	for _, r := range s.Resources {
		if r.HasTag(resourceTag) {
			return true
		}
	}

	return false
}

// HasTag returns true if the species has a given tag
func (s Species) HasTag(tag string) bool {
	for _, t := range s.Tags {
		if t == tag {
			return true
		}
	}

	return false
}

// InSlice checks whether a given species is in a slice of speciess
func (s Species) InSlice(from []Species) bool {
	isIt := false
	for _, a := range from {
		if a.Name == s.Name {
			isIt = true
		}
	}

	return isIt
}

// Random returns a set number of randomly chosen speciess from a slice
func Random(ctx context.Context, amount int, from []Species) []Species {
	var s Species
	var result []Species

	if amount > len(from) {
		amount = len(from)
	}

	for i := 0; i < amount; i++ {
		s = from[random.Intn(ctx, len(from))]
		if !s.InSlice(result) {
			result = append(result, s)
		}
	}

	return result
}

// RandomWithResourceTag returns a random species that has a resource with the given tag
func RandomWithResourceTag(ctx context.Context, resourceTag string, from []Species) (Species, error) {
	filtered := []Species{}

	for _, s := range from {
		if s.HasResourceWithTag(resourceTag) {
			if !s.InSlice(filtered) {
				filtered = append(filtered, s)
			}
		}
	}

	if len(filtered) == 0 {
		err := fmt.Errorf("no species matching tag " + resourceTag + " was found")
		return Species{}, err
	}

	if len(filtered) == 1 {
		return filtered[0], nil
	}

	species := filtered[random.Intn(ctx, len(filtered))]

	return species, nil
}
