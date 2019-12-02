/*
Package trait provides tools for dealing with descriptive traits of creatures or objects
*/
package trait

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/ironarachne/world/pkg/slices"
)

// Trait is a trait of something that can be used for description
type Trait struct {
	Name       string   `json:"name"`       // The name of the trait
	Value      string   `json:"value"`      // The value of the trait
	Descriptor string   `json:"descriptor"` // The pattern by which a trait is described
	Tags       []string `json:"tags"`
}

// ByTag returns all traits from the slice that match the tag
func ByTag(tag string, traits []Trait) []Trait {
	result := []Trait{}

	for _, t := range traits {
		if slices.StringIn(tag, t.Tags) {
			result = append(result, t)
		}
	}

	return result
}

// InSlice checks to see if a given trait is in the slice
func InSlice(trait Trait, traits []Trait) bool {
	for _, t := range traits {
		if t.Name == trait.Name {
			return true
		}
	}

	return false
}

// ToString renders a trait to a string
func (t Trait) ToString() (string, error) {
	var tplOutput bytes.Buffer

	tmpl, err := template.New(t.Name).Parse(t.Descriptor)
	if err != nil {
		err = fmt.Errorf("Failed to describe trait: %w", err)
		return "", err
	}
	err = tmpl.Execute(&tplOutput, t)
	if err != nil {
		err = fmt.Errorf("Failed to describe trait: %w", err)
		return "", err
	}
	name := tplOutput.String()

	return name, nil
}
