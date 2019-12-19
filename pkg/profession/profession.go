/*
Package profession provides fantasy professions and metadata for them
*/
package profession

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // sqlite driver
)

// Profession is a person with a particular skillset that can make a resource
type Profession struct {
	ID          int      `json:"id" db:"id"`
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	Tags        []string `json:"tags"`
}

// All returns all professions
func All() ([]Profession, error) {
	var professions []Profession
	var all []Profession
	var result []Profession
	var tags []string

	db, err := sqlx.Connect("sqlite3", os.Getenv("WORLDAPI_DATA_PATH")+"/data.db")
	if err != nil {
		err = fmt.Errorf("could not open database: %w", err)
		return []Profession{}, nil
	}

	db.Select(&professions, "SELECT * FROM professions")

	for _, prof := range professions {
		db.Select(&tags, "SELECT tag FROM profession_tags WHERE profession_id = ?", prof.ID)
		prof.Tags = tags
		all = append(all, prof)
	}

	db.Close()

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

// Random returns a random profession
func Random() (Profession, error) {
	professions, err := All()
	if err != nil {
		err = fmt.Errorf("could not fetch professions: %w", err)
		return Profession{}, err
	}

	if len(professions) == 0 {
		err = fmt.Errorf("found no professions to select from")
		return Profession{}, err
	}

	if len(professions) == 1 {
		return professions[0], nil
	}

	profession := professions[rand.Intn(len(professions))]

	return profession, nil
}

// RandomSet returns a random number of professions from a given set of professions
func RandomSet(max int, possible []Profession) ([]Profession, error) {
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
		profession = possible[rand.Intn(len(possible))]
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
