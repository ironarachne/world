package resource

import (
	"bytes"
	"fmt"
	"math/rand"
	"text/template"

	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/random"
)

// Pattern is a pattern for a resource
type Pattern struct {
	Name                 string                `json:"name"`
	Description          string                `json:"description"`
	Tags                 []string              `json:"tags"`
	Commonality          int                   `json:"commonality"`
	Profession           profession.Profession `json:"profession"`
	Slots                []Slot                `json:"slots"`
	NameTemplate         string                `json:"name_template"`
	MainMaterial         string                `json:"main_material"`
	MainMaterialOverride string                `json:"main_material_override"`
	OriginOverride       string                `json:"origin_override"`
	Value                int                   `json:"value"`
}

// Slot is a slot for a resource requirement
type Slot struct {
	Name                string   `json:"name"`
	RequiredTag         string   `json:"required_tag"`
	Resource            Resource `json:"resource"`
	DescriptionTemplate string   `json:"description_template"`
	PossibleQuirks      []string `json:"possible_quirks"`
}

// AllPatterns returns all patterns
func AllPatterns() []Pattern {
	var patterns []Pattern
	armor := getArmor()
	beekeeping := getBeekeeping()
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
	vehicles := getVehicles()
	weapons := getWeapons()
	weaving := getWeaving()
	wine := getWine()

	patterns = append(patterns, armor...)
	patterns = append(patterns, beekeeping...)
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
	patterns = append(patterns, vehicles...)
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

// RenderName renders the name of a completed pattern
func (pattern Pattern) RenderName() (string, error) {
	var tplOutput bytes.Buffer

	tmpl, err := template.New(pattern.Name).Parse(pattern.NameTemplate)
	if err != nil {
		err = fmt.Errorf("Failed to render resource name for pattern: %w", err)
		return "", err
	}
	err = tmpl.Execute(&tplOutput, pattern)
	if err != nil {
		err = fmt.Errorf("Failed to render resource name for pattern: %w", err)
		return "", err
	}
	name := tplOutput.String()

	return name, nil
}

// Render turns a completed pattern into a string description
func (pattern Pattern) RenderDescription() (string, error) {
	description := ""

	for _, s := range pattern.Slots {
		slotDescription, err := s.describe()
		if err != nil {
			err = fmt.Errorf("Failed to render pattern: %w", err)
			return "", err
		}
		description += slotDescription
	}

	return description, nil
}

// ToResource transforms a completed pattern into a resource
func (pattern Pattern) ToResource() (Resource, error) {
	var resource Resource

	description, err := pattern.RenderDescription()
	if err != nil {
		err = fmt.Errorf("Failed to transform pattern into resource: %w", err)
		return Resource{}, err
	}

	mainMaterial := pattern.Slots[0].Resource.MainMaterial
	if pattern.MainMaterialOverride != "" {
		mainMaterial = pattern.MainMaterialOverride
	}
	pattern.MainMaterial = mainMaterial

	name, err := pattern.RenderName()
	if err != nil {
		err = fmt.Errorf("Failed to transform pattern into resource: %w", err)
		return Resource{}, err
	}

	resource.Name = name
	resource.Description = description
	resource.MainMaterial = mainMaterial
	resource.Origin = pattern.Slots[0].Resource.Origin
	if pattern.OriginOverride != "" {
		resource.Origin = pattern.OriginOverride
	}
	resource.Tags = pattern.Tags
	resource.Tags = append(resource.Tags, resource.Name)
	resource.Commonality = pattern.Commonality
	value := 0

	for _, s := range pattern.Slots {
		value += s.Resource.Value
	}
	value += pattern.Value
	resource.Value = value

	return resource, nil
}

func (slot Slot) describe() (string, error) {
	var tplOutput bytes.Buffer

	tmpl, err := template.New(slot.Name).Parse(slot.DescriptionTemplate)
	if err != nil {
		err = fmt.Errorf("Failed to describe slot: %w", err)
		return "", err
	}
	err = tmpl.Execute(&tplOutput, slot)
	if err != nil {
		err = fmt.Errorf("Failed to describe slot: %w", err)
		return "", err
	}
	name := tplOutput.String()

	if len(slot.PossibleQuirks) > 0 {
		quirkChance := rand.Intn(100)
		if quirkChance > 80 {
			quirk, err := random.String(slot.PossibleQuirks)
			if err != nil {
				err = fmt.Errorf("Failed to describe slot: %w", err)
				return "", err
			}
			name += " " + quirk
		}
	}

	return name, nil
}
