package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"text/template"

	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/random"
)

// Data is a collection of patterns
type Data struct {
	Patterns []Pattern `json:"patterns"`
}

// Pattern is a pattern for a resource
type Pattern struct {
	Name                 string                `json:"name" db:"name"`
	Description          string                `json:"description" db:"description"`
	Tags                 []string              `json:"tags" db:"tags"`
	Commonality          int                   `json:"commonality" db:"commonality"`
	Profession           profession.Profession `json:"profession" db:"profession"`
	Slots                []Slot                `json:"slots" db:"slots"`
	NameTemplate         string                `json:"name_template" db:"name_template"`
	MainMaterial         string                `json:"main_material" db:"main_material"`
	MainMaterialOverride string                `json:"main_material_override" db:"main_material_override"`
	OriginOverride       string                `json:"origin_override" db:"origin_override"`
	Value                int                   `json:"value" db:"value"`
}

// Slot is a slot for a resource requirement
type Slot struct {
	Name                string   `json:"name" db:"name"`
	RequiredTag         string   `json:"required_tag" db:"required_tag"`
	Resource            Resource `json:"resource" db:"resource"`
	DescriptionTemplate string   `json:"description_template" db:"description_template"`
	PossibleQuirks      []string `json:"possible_quirks" db:"possible_quirks"`
}

// All returns all predefined patterns from a JSON file on disk
func All() ([]Pattern, error) {
	var d Data

	jsonFile, err := os.Open(os.Getenv("WORLDAPI_DATA_PATH") + "/data/patterns.json")
	if err != nil {
		err = fmt.Errorf("could not open data file: %w", err)
		return []Pattern{}, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &d)

	all := d.Patterns

	if len(all) == 0 {
		err = fmt.Errorf("no patterns returned from database: patterns.json")
		return []Pattern{}, err
	}

	return all, nil
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
func GetPossibleProfessions(resources []Resource) ([]profession.Profession, error) {
	var possiblePatterns []Pattern
	possibleProducers, _ := profession.All()

	producers := []profession.Profession{}

	patterns := []Pattern{}
	allPatterns, err := All()
	if err != nil {
		err = fmt.Errorf("failed to get possible professions: %w", err)
		return []profession.Profession{}, err
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
			if !p.InSlice(producers) {
				producers = append(producers, p)
			}
		}
	}

	return producers, nil
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
		err = fmt.Errorf("failed to render resource name for pattern: %w", err)
		return "", err
	}
	err = tmpl.Execute(&tplOutput, pattern)
	if err != nil {
		err = fmt.Errorf("failed to render resource name for pattern: %w", err)
		return "", err
	}
	name := tplOutput.String()

	return name, nil
}

// RenderDescription turns a completed pattern into a string description
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
