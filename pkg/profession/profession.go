/*
Package profession provides fantasy professions and metadata for them
*/
package profession

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ironarachne/world/pkg/random"
)

// Data is a struct containing a slice of professions
type Data struct {
	Professions []Profession `json:"professions"`
}

// Profession is a person with a particular skillset that can make a resource
type Profession struct {
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	Tags        []string `json:"tags"`
}

// All returns all professions
func All() ([]Profession, error) {
	var professions Data
	var result []Profession

	jsonFile, err := os.Open(os.Getenv("WORLDAPI_DATA_PATH") + "/data/professions.json")
	if err != nil {
		err = fmt.Errorf("could not open data file: %w", err)
		return []Profession{}, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &professions)

	all := professions.Professions

	if len(all) == 0 {
		err = fmt.Errorf("no professions returned from database")
		return []Profession{}, err
	}

	none := Profession{
		Name:        "none",
		Description: "no profession",
	}
	all = append(all, none)

	for _, p := range all {
		p.Tags = append(p.Tags, p.Name)
		result = append(result, p)
	}

	return result, nil
}

// ByName returns a profession by name
func ByName(name string) (Profession, error) {
	professions, err := All()
	if err != nil {
		err = fmt.Errorf("could not fetch professions: %w", err)
		return Profession{}, nil
	}

	for _, p := range professions {
		if p.Name == name {
			return p, nil
		}
	}

	err = fmt.Errorf("could not find profession by name: %s", name)
	return Profession{}, nil
}

// ByTag returns a slice of professions that have the given tag
func ByTag(tag string) ([]Profession, error) {
	var professions []Profession
	all, err := All()
	if err != nil {
		err = fmt.Errorf("could not fetch professions: %w", err)
		return []Profession{}, nil
	}

	for _, p := range all {
		if p.HasTag(tag) {
			professions = append(professions, p)
		}
	}

	return professions, nil
}

// HasTag returns true if the profession has a given tag
func (profession Profession) HasTag(tag string) bool {
	for _, t := range profession.Tags {
		if t == tag {
			return true
		}
	}

	return false
}

// Random returns a single random profession from all professions
func Random(ctx context.Context) (Profession, error) {
	all, err := All()
	if err != nil {
		err = fmt.Errorf("could not fetch professions: %w", err)
		return Profession{}, err
	}

	professions, err := RandomSet(ctx, 1, all)
	if err != nil {
		err = fmt.Errorf("could not get random profession: %w", err)
		return Profession{}, err
	}

	return professions[0], nil
}

// RandomSet returns a random number of professions from a given set of professions
func RandomSet(ctx context.Context, max int, possible []Profession) ([]Profession, error) {
	professions := []Profession{}
	profession := Profession{}

	if len(possible) == 0 {
		err := fmt.Errorf("no possible professions given")
		return []Profession{}, err
	}

	if max > len(possible) {
		max = len(possible)
	}

	for i := 0; i < max; i++ {
		profession = possible[random.Intn(ctx, len(possible))]
		if !profession.InSlice(professions) {
			professions = append(professions, profession)
		}
	}

	return professions, nil
}

// InSlice returns true if the given profession is in the slice
func (profession Profession) InSlice(professions []Profession) bool {
	for _, p := range professions {
		if p.Name == profession.Name {
			return true
		}
	}

	return false
}

// Names returns a slice of just the names of a list of professions
func Names(professions []Profession) []string {
	var names []string

	for _, p := range professions {
		names = append(names, p.Name)
	}

	return names
}
