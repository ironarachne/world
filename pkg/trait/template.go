package trait

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
)

// Template is a template for building a trait
type Template struct {
	Name                string   `json:"name"`
	PossibleValues      []string `json:"possible_values"`
	PossibleDescriptors []string `json:"possible_descriptors"`
	Tags                []string `json:"tags"`
}

// RandomTemplate returns a random template from a slice
func RandomTemplate(templates []Template) (Template, error) {
	if len(templates) == 0 {
		err := fmt.Errorf("could not select random template from empty slice")
		return Template{}, err
	}

	if len(templates) == 1 {
		return templates[0], nil
	}

	template := templates[rand.Intn(len(templates))]

	return template, nil
}

// ToTrait turns a template into a trait
func (t Template) ToTrait() (Trait, error) {
	r := Trait{}
	r.Name = t.Name
	r.Tags = t.Tags
	value, err := random.String(t.PossibleValues)
	if err != nil {
		err = fmt.Errorf("Failed to turn template into trait: %w", err)
		return Trait{}, err
	}
	r.Value = value
	descriptor, err := random.String(t.PossibleDescriptors)
	if err != nil {
		err = fmt.Errorf("Failed to turn template into trait: %w", err)
		return Trait{}, err
	}
	r.Descriptor = descriptor

	return r, nil
}
