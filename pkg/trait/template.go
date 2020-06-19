package trait

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

// Template is a template for building a trait
type Template struct {
	Name                string   `json:"name" db:"name"`
	PossibleValues      []string `json:"possible_values" db:"possible_values"`
	PossibleDescriptors []string `json:"possible_descriptors" db:"possible_descriptors"`
	Tags                []string `json:"tags" db:"tags"`
}

// RandomTemplate returns a random template from a slice
func RandomTemplate(ctx context.Context, templates []Template) (Template, error) {
	if len(templates) == 0 {
		err := fmt.Errorf("could not select random template from empty slice")
		return Template{}, err
	}

	if len(templates) == 1 {
		return templates[0], nil
	}

	template := templates[random.Intn(ctx, len(templates))]

	return template, nil
}

// ToTrait turns a template into a trait
func (t Template) ToTrait(ctx context.Context) (Trait, error) {
	r := Trait{}
	r.Name = t.Name
	r.Tags = t.Tags
	value, err := random.String(ctx, t.PossibleValues)
	if err != nil {
		err = fmt.Errorf("Failed to turn template into trait: %w", err)
		return Trait{}, err
	}
	r.Value = value
	descriptor, err := random.String(ctx, t.PossibleDescriptors)
	if err != nil {
		err = fmt.Errorf("Failed to turn template into trait: %w", err)
		return Trait{}, err
	}
	r.Descriptor = descriptor

	return r, nil
}
