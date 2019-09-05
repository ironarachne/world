package organization

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/words"
)

// SimplifiedOrganization is a simplified version of an organization
type SimplifiedOrganization struct {
	Name           string                          `json:"name"`
	Type           string                          `json:"type"`
	Description    string                          `json:"description"`
	Leader         character.SimplifiedCharacter   `json:"leader"`
	NotableMembers []character.SimplifiedCharacter `json:"notable_members"`
}

// Simplify returns a simplified version of the organization
func (org Organization) Simplify() (SimplifiedOrganization, error) {
	sl, err := org.Leader.Simplify()
	if err != nil {
		err = fmt.Errorf("Could not simplify organization: %w", err)
		return SimplifiedOrganization{}, err
	}
	simplified := SimplifiedOrganization{
		Name:        org.Name,
		Type:        org.SizeClass.Name + " " + org.Type.Name,
		Description: org.Describe(),
		Leader:      sl,
	}

	for _, n := range org.NotableMembers {
		sn, err := n.Simplify()
		if err != nil {
			err = fmt.Errorf("Could not simplify organization: %w", err)
			return SimplifiedOrganization{}, err
		}
		simplified.NotableMembers = append(simplified.NotableMembers, sn)
	}

	return simplified, nil
}

// Describe constructs a string description of the organization
func (org Organization) Describe() string {
	description := "The " + org.Name + " is " + words.Pronoun(org.SizeClass.Name) + " " + org.SizeClass.Name + " " + org.Type.Name + " with " + strconv.Itoa(org.Size) + " members. "
	description += "It's led by " + strings.Title(org.Leader.Title) + " " + org.Leader.FirstName + " " + org.Leader.LastName + ". "
	description += "The " + org.Type.Name + " is known for being " + org.PrimaryTrait + "."

	return description
}

// RandomSimplified returns a random simplified version of the organization
func RandomSimplified() (SimplifiedOrganization, error) {
	org, err := Random()
	if err != nil {
		err = fmt.Errorf("Could not simplify random organization: %w", err)
		return SimplifiedOrganization{}, err
	}

	os, err := org.Simplify()
	if err != nil {
		err = fmt.Errorf("Could not simplify organization: %w", err)
		return SimplifiedOrganization{}, err
	}

	return os, nil
}
