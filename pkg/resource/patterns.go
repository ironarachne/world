package resource

import (
	"bytes"
	"text/template"

	"github.com/ironarachne/world/pkg/profession"
)

// Pattern is a pattern for a resource
type Pattern struct {
	Name           string
	Description    string
	Tags           []string
	Commonality    int
	Profession     profession.Profession
	Slots          []Slot
	OriginOverride string
}

// Slot is a slot for a resource requirement
type Slot struct {
	Name                string
	RequiredTag         string
	Resource            Resource
	DescriptionTemplate string
}

// AllPatterns returns all patterns
func AllPatterns() []Pattern {
	var patterns []Pattern
	armor := getArmor()
	bowery := getBowery()
	breads := getBreads()
	brewed := getBrewed()
	carpentry := getCarpentry()
	clothing := getClothing()
	cobbler := getCobbler()
	distilled := getDistilled()
	glass := getGlass()
	logging := getLogging()
	medicine := getMedicine()
	milled := getMilled()
	mined := getMined()
	mounts := getMounts()
	potions := getPotions()
	pottery := getPottery()
	smelting := getSmelting()
	smithing := getSmithing()
	stone := getStone()
	tannery := getTannery()
	weapons := getWeapons()
	weaving := getWeaving()
	wine := getWine()

	patterns = append(patterns, armor...)
	patterns = append(patterns, bowery...)
	patterns = append(patterns, breads...)
	patterns = append(patterns, brewed...)
	patterns = append(patterns, carpentry...)
	patterns = append(patterns, clothing...)
	patterns = append(patterns, cobbler...)
	patterns = append(patterns, distilled...)
	patterns = append(patterns, glass...)
	patterns = append(patterns, logging...)
	patterns = append(patterns, medicine...)
	patterns = append(patterns, milled...)
	patterns = append(patterns, mined...)
	patterns = append(patterns, mounts...)
	patterns = append(patterns, potions...)
	patterns = append(patterns, pottery...)
	patterns = append(patterns, smelting...)
	patterns = append(patterns, smithing...)
	patterns = append(patterns, stone...)
	patterns = append(patterns, tannery...)
	patterns = append(patterns, weapons...)
	patterns = append(patterns, weaving...)
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

	producers := []profession.Profession{}

	patterns := []Pattern{}
	allPatterns := AllPatterns()

	for _, p := range possibleProducers {
		patterns = []Pattern{}
		possiblePatterns = FindPatternsForProfession(p, allPatterns)

		for _, i := range possiblePatterns {
			if i.CanMake(resources) {
				patterns = append(patterns, i)
			}
		}

		if len(patterns) > 0 {
			if !p.InSlice(producers) {
				producers = append(producers, p)
			}
		}
	}

	return producers
}

// CanMake returns true if the pattern can be made with the resources given
func (pattern Pattern) CanMake(resources []Resource) bool {
	var matchingResources []Resource
	possible := true

	for _, s := range pattern.Slots {
		matchingResources = ByTag(s.RequiredTag, resources)
		if len(matchingResources) == 0 {
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

// ToResource transforms a completed pattern into a resource
func (pattern Pattern) ToResource() Resource {
	var resource Resource

	resource.Name = pattern.Render()
	resource.Origin = pattern.Slots[0].Resource.Origin
	if pattern.OriginOverride != "" {
		resource.Origin = pattern.OriginOverride
	}
	resource.Tags = pattern.Tags
	resource.Tags = append(resource.Tags, resource.Name)
	resource.Commonality = pattern.Commonality

	return resource
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
