package pattern

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/ironarachne/world/pkg/resource"

	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/random"
)

// Data is a collection of patterns
type Data struct {
	Patterns []Pattern `json:"patterns"`
}

// Pattern is a pattern for a resource
type Pattern struct {
	Name                 string   `json:"name" db:"name"`
	Description          string   `json:"description" db:"description"`
	Tags                 []string `json:"tags" db:"tags"`
	Commonality          int      `json:"commonality" db:"commonality"`
	ProfessionName       string   `json:"profession_name" db:"profession_name"`
	Slots                []Slot   `json:"slots" db:"slots"`
	NameTemplate         string   `json:"name_template" db:"name_template"`
	MainMaterial         string   `json:"main_material" db:"main_material"`
	MainMaterialOverride string   `json:"main_material_override" db:"main_material_override"`
	OriginOverride       string   `json:"origin_override" db:"origin_override"`
	Value                int      `json:"value" db:"value"`
}

// Slot is a slot for a resource requirement
type Slot struct {
	Name                string            `json:"name" db:"name"`
	RequiredTag         string            `json:"required_tag" db:"required_tag"`
	Resource            resource.Resource `json:"resource" db:"resource"`
	DescriptionTemplate string            `json:"description_template" db:"description_template"`
	PossibleQuirks      []string          `json:"possible_quirks" db:"possible_quirks"`
}

// All returns all predefined patterns from JSON files on disk
func All() ([]Pattern, error) {
	var all []Pattern
	var err error
	var patterns []Pattern

	patternFiles := []string{
		"armor",
		"carpentry",
		"clothing",
		"drinks",
		"foods",
		"miscellaneous",
		"mounts",
		"smithing",
		"tools",
		"vehicles",
		"weapons",
	}

	for _, f := range patternFiles {
		patterns, err = LoadFromFile("patterns/" + f)
		if err != nil {
			err = fmt.Errorf("could not open data file: %w", err)
			return []Pattern{}, err
		}
		all = append(all, patterns...)
	}

	return all, nil
}

// ForProfession returns all patterns from the set that a profession can make
func ForProfession(prof profession.Profession, from []Pattern) []Pattern {
	var patterns []Pattern

	for _, p := range from {
		if p.ProfessionName == prof.Name {
			patterns = append(patterns, p)
		}
	}

	return patterns
}

// GetPossibleProfessions gets all possible professions for a given set of resources
func GetPossibleProfessions(resources []resource.Resource) ([]profession.Profession, error) {
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
		possiblePatterns = ForProfession(p, allPatterns)

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

// LoadFromFile loads patterns from the given JSON file
func LoadFromFile(name string) ([]Pattern, error) {
	var d Data

	jsonFile, err := os.Open(os.Getenv("WORLDAPI_DATA_PATH") + "/data/" + name + ".json")
	if err != nil {
		err = fmt.Errorf("could not open data file: %w", err)
		return []Pattern{}, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &d)
	if err != nil {
		err = fmt.Errorf("could not unmarshal JSON data: %w", err)
		return []Pattern{}, err
	}

	patterns := d.Patterns

	if len(patterns) == 0 {
		err = fmt.Errorf("no patterns returned from database: " + name + ".json")
		return []Pattern{}, err
	}

	return patterns, nil
}

// CanMake returns true if the pattern can be made with the resources given
func (pattern Pattern) CanMake(resources []resource.Resource) bool {
	var matchingResources []resource.Resource
	possible := true

	for _, s := range pattern.Slots {
		matchingResources = resource.ByTag(s.RequiredTag, resources)
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
func (pattern Pattern) RenderDescription(ctx context.Context) (string, error) {
	description := ""

	for _, s := range pattern.Slots {
		slotDescription, err := s.describe(ctx)
		if err != nil {
			err = fmt.Errorf("Failed to render pattern: %w", err)
			return "", err
		}
		description += slotDescription
	}

	return description, nil
}

// ToResource transforms a completed pattern into a resource
func (pattern Pattern) ToResource(ctx context.Context) (resource.Resource, error) {
	var res resource.Resource

	description, err := pattern.RenderDescription(ctx)
	if err != nil {
		err = fmt.Errorf("Failed to transform pattern into resource: %w", err)
		return resource.Resource{}, err
	}

	mainMaterial := pattern.Slots[0].Resource.MainMaterial
	if pattern.MainMaterialOverride != "" {
		mainMaterial = pattern.MainMaterialOverride
	}
	pattern.MainMaterial = mainMaterial

	name, err := pattern.RenderName()
	if err != nil {
		err = fmt.Errorf("Failed to transform pattern into resource: %w", err)
		return resource.Resource{}, err
	}

	res.Name = name
	res.Description = description
	res.MainMaterial = mainMaterial
	res.Origin = pattern.Slots[0].Resource.Origin
	if pattern.OriginOverride != "" {
		res.Origin = pattern.OriginOverride
	}
	res.Tags = pattern.Tags
	res.Tags = append(res.Tags, res.Name)
	res.Commonality = pattern.Commonality
	value := 0

	for _, s := range pattern.Slots {
		value += s.Resource.Value
	}
	value += pattern.Value
	res.Value = value

	return res, nil
}

func (slot Slot) describe(ctx context.Context) (string, error) {
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
		quirkChance := random.Intn(ctx, 100)
		if quirkChance > 80 {
			quirk, err := random.String(ctx, slot.PossibleQuirks)
			if err != nil {
				err = fmt.Errorf("Failed to describe slot: %w", err)
				return "", err
			}
			name += " " + quirk
		}
	}

	return name, nil
}
