/*
Package resource provides natural and manmade resources. It implements a system for dealing with such.
*/
package resource

import (
	"context"

	"github.com/ironarachne/world/pkg/random"
)

// Resource is a useful resource
type Resource struct {
	Name         string   `json:"name" db:"name"`
	Description  string   `json:"description" db:"description"`
	MainMaterial string   `json:"main_material" db:"main_material"`
	Origin       string   `json:"origin" db:"origin"`
	Tags         []string `json:"tags" db:"tags"`
	Commonality  int      `json:"commonality" db:"commonality"`
	Value        int      `json:"value" db:"value"`
}

// ByTag returns a slice of resources that have the given tag
func ByTag(tag string, from []Resource) []Resource {
	var resources []Resource

	for _, p := range from {
		if p.HasTag(tag) {
			resources = append(resources, p)
		}
	}

	return resources
}

// HasTag returns true if the resource has a given tag
func (resource Resource) HasTag(tag string) bool {
	for _, t := range resource.Tags {
		if t == tag {
			return true
		}
	}

	return false
}

// InSlice checks if a given resource is in a slice of resources
func (resource Resource) InSlice(resources []Resource) bool {
	for _, r := range resources {
		if r.Name == resource.Name {
			return true
		}
	}

	return false
}

// Random returns a random resource from a list
func Random(ctx context.Context, resources []Resource) Resource {
	if len(resources) == 1 {
		return resources[0]
	} else if len(resources) < 1 {
		panic("No resources given")
	}

	resource := resources[random.Intn(ctx, len(resources))]

	return resource
}

// RandomSet returns a slice of random elements of the given resources
func RandomSet(ctx context.Context, min int, max int, resources []Resource) []Resource {
	var result []Resource
	var resource Resource

	numberOfResources := random.Intn(ctx, max-min) + min
	if numberOfResources > len(resources) {
		numberOfResources = len(resources)
	}

	for i := 0; i < numberOfResources; i++ {
		resource = Random(ctx, resources)
		if !resource.InSlice(result) {
			result = append(result, resource)
		}
	}

	return result
}
