package resource

import (
	"math/rand"
)

// Resource is a useful resource
type Resource struct {
	Name        string
	Origin      string
	Type        string
	Commonality int
}

// IsTypeInResources checks to see if a given type is present in a collection of resources
func IsTypeInResources(typeName string, resources []Resource) bool {
	for _, i := range resources {
		if i.Type == typeName {
			return true
		}
	}

	return false
}

// ListOfType returns all resources in the given collection that match the type
func ListOfType(typeName string, resources []Resource) []Resource {
	filteredResources := []Resource{}

	for _, i := range resources {
		if i.Type == typeName {
			filteredResources = append(filteredResources, i)
		}
	}

	return filteredResources
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
func Random(resources []Resource) Resource {
	resource := resources[rand.Intn(len(resources))]

	return resource
}
