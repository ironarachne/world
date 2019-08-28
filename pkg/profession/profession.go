package profession

import (
	"math/rand"
)

// Profession is a person with a particular skillset that can make a resource
type Profession struct {
	Name        string
	Description string
	Tags        []string
}

// All returns all professions
func All() []Profession {
	var professions []Profession

	blacksmiths := blacksmiths()
	professions = append(professions, blacksmiths...)
	fighters := fighters()
	professions = append(professions, fighters...)
	finishers := finishers()
	professions = append(professions, finishers...)
	mages := mages()
	professions = append(professions, mages...)
	refiners := refiners()
	professions = append(professions, refiners...)
	social := social()
	professions = append(professions, social...)
	none := Profession{
		Name:        "none",
		Description: "no profession",
	}
	professions = append(professions, none)

	return professions
}

// ByName returns a profession by name
func ByName(name string) Profession {
	professions := All()

	for _, p := range professions {
		if p.Name == name {
			return p
		}
	}

	panic("Profession " + name + " not found!")
}

// ByTag returns a slice of professions that have the given tag
func ByTag(tag string) []Profession {
	var professions []Profession
	all := All()

	for _, p := range all {
		if p.HasTag(tag) {
			professions = append(professions, p)
		}
	}

	return professions
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
func Random() Profession {
	professions := All()

	profession := professions[rand.Intn(len(professions))]

	return profession
}

// RandomSet returns a random number of professions from a given set of professions
func RandomSet(max int, possible []Profession) []Profession {
	producers := []Profession{}
	producer := Profession{}

	for i := 0; i < max; i++ {
		producer = possible[rand.Intn(len(possible))]
		if !producer.InSlice(producers) {
			producers = append(producers, producer)
		}
	}

	return producers
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