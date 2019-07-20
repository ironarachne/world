package resource

import (
	"bytes"
	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/slices"
	"text/template"
)

// Pattern is a pattern for a resource
type Pattern struct {
	Name        string
	Description string
	Profession  profession.Profession
	Slots       []Slot
}

// Slot is a slot for a resource requirement
type Slot struct {
	Name                string
	RequiredType        string
	Resource            Resource
	DescriptionTemplate string
}

// AllPatterns returns all patterns
func AllPatterns() []Pattern {
	var patterns []Pattern
	armor := getArmor()
	breads := getBreads()
	brewed := getBrewed()
	carpentry := getCarpentry()
	clothing := getClothing()
	cobbler := getCobbler()
	glass := getGlass()
	medicine := getMedicine()
	milled := getMilled()
	mined := getMined()
	mounts := getMounts()
	potions := getPotions()
	pottery := getPottery()
	tannery := getTannery()
	weapons := getWeapons()
	wine := getWine()

	patterns = append(patterns, armor...)
	patterns = append(patterns, breads...)
	patterns = append(patterns, brewed...)
	patterns = append(patterns, carpentry...)
	patterns = append(patterns, clothing...)
	patterns = append(patterns, cobbler...)
	patterns = append(patterns, glass...)
	patterns = append(patterns, medicine...)
	patterns = append(patterns, milled...)
	patterns = append(patterns, mined...)
	patterns = append(patterns, mounts...)
	patterns = append(patterns, potions...)
	patterns = append(patterns, pottery...)
	patterns = append(patterns, tannery...)
	patterns = append(patterns, weapons...)
	patterns = append(patterns, wine...)

	return patterns
}

// FindPatternsForProfession returns all patterns from the set that a profession can make
func FindPatternsForProfession(prof profession.Profession, from []Pattern) []Pattern {
	var patterns []Pattern

	for _, p := range from {
		if p.Profession.Name == prof.Name {
			patterns = append(patterns, p)
		}
	}

	return patterns
}

// GetPossibleProfessions gets all possible professions for a given set of resources
func GetPossibleProfessions(resources []Resource) []profession.Profession {
	var possiblePatterns []Pattern
	possibleProducers := profession.All()

	availableTypes := []string{}
	producers := []profession.Profession{}

	patterns := []Pattern{}
	allPatterns := AllPatterns()

	for _, r := range resources {
		if !slices.StringIn(r.Type, availableTypes) {
			availableTypes = append(availableTypes, r.Type)
		}
	}

	for _, p := range possibleProducers {
		patterns = []Pattern{}
		possiblePatterns = FindPatternsForProfession(p, allPatterns)

		for _, i := range possiblePatterns {
			if i.CanMake(resources) {
				patterns = append(patterns, i)
			}
		}

		if len(patterns) > 0 {
			producers = append(producers, p)
		}
	}

	return producers
}

// CanMake returns true if the pattern can be made with the resources given
func (pattern Pattern) CanMake(resources []Resource) bool {
	possible := true

	for _, s := range pattern.Slots {
		if !IsTypeInResources(s.RequiredType, resources) {
			possible = false
		}
	}

	return possible
}

// Render turns a completed pattern into a string description
func (pattern Pattern) Render() string {
	var slotDescription string
	description := ""

	for _, s := range pattern.Slots {
		slotDescription = s.describe()
		description += slotDescription
	}

	return description
}

func (slot Slot) describe() string {
	var tplOutput bytes.Buffer

	tmpl, err := template.New(slot.Name).Parse(slot.DescriptionTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&tplOutput, slot)
	if err != nil {
		panic(err)
	}
	name := tplOutput.String()

	return name
}