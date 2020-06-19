/*
Package relationship provides tools for dealing with relationships between various entities
*/
package relationship

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

// Relationship is a unidirectional relationship between entities
type Relationship struct {
	Origin     string `json:"origin"`
	Target     string `json:"target"`
	Descriptor string `json:"descriptor"`
	Type       string `json:"type"`
}

// Describe returns a string version of the relationship
func (r Relationship) Describe() string {
	description := r.Origin + " " + r.Descriptor + " " + r.Target

	return description
}

// InSlice checks to see if a relationship is present in a set of relationships
func InSlice(r Relationship, relationships []Relationship) bool {
	for _, i := range relationships {
		if r.Type == i.Type && r.Target == i.Target && r.Origin == i.Origin {
			return true
		}
	}
	return false
}

// GetTypeFromSet gets the named relationship type from the given set
func GetTypeFromSet(name string, set []Type) (Type, error) {
	for _, t := range set {
		if t.Name == name {
			return t, nil
		}
	}

	err := fmt.Errorf("Failed to find relationship type " + name + " in the given set")

	return Type{}, err
}

// GetTypeByName returns the specified type, if it exists
func GetTypeByName(name string) (Type, error) {
	all := AllTypes()

	for _, t := range all {
		if t.Name == name {
			return t, nil
		}
	}

	err := fmt.Errorf("Could not find named relationship type " + name)
	return Type{}, err
}

// GetInverse gets an inverse relationship for a given relationship
func GetInverse(ctx context.Context, r Relationship) (Relationship, error) {
	originType, err := GetTypeByName(r.Type)
	if err != nil {
		err = fmt.Errorf("Failed to get inverse relationship: %w", err)
		return Relationship{}, err
	}

	inverseType, err := GetTypeByName(originType.InverseName)
	if err != nil {
		err = fmt.Errorf("Failed to get inverse relationship: %w", err)
		return Relationship{}, err
	}

	descriptor, err := random.String(ctx, inverseType.Descriptors)
	if err != nil {
		err = fmt.Errorf("Failed to get inverse relationship: %w", err)
		return Relationship{}, err
	}

	relationship := Relationship{
		Origin:     r.Target,
		Target:     r.Origin,
		Descriptor: descriptor,
		Type:       inverseType.Name,
	}

	return relationship, nil
}

// AllowedIn checks to see if a given relationship is compatible with the given set of relationships
func AllowedIn(r Relationship, set []Relationship) (bool, error) {
	for _, s := range set {
		compat, err := IsCompatible(r, s)
		if err != nil {
			err = fmt.Errorf("Failed to check compatibility of relationship: %w", err)
			return false, err
		}
		if (r.Origin == s.Target || r.Target == s.Origin) && !compat {
			return false, nil
		}
	}

	return true, nil
}

// IsCompatible checks to see if the two relationships are compatible
func IsCompatible(r1 Relationship, r2 Relationship) (bool, error) {
	r1Type, err := GetTypeByName(r1.Type)
	if err != nil {
		err = fmt.Errorf("Failed to check opposite relationship: %w", err)
		return false, err
	}
	r2Type, err := GetTypeByName(r2.Type)
	if err != nil {
		err = fmt.Errorf("Failed to check opposite relationship: %w", err)
		return false, err
	}

	if !slices.StringIn(r1Type.Name, r2Type.Disallows) && !slices.StringIn(r2Type.Name, r1Type.Disallows) {
		return true, nil
	}

	return false, nil
}
