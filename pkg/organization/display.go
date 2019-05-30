package organization

import (
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
func (org Organization) Simplify() SimplifiedOrganization {
	simplified := SimplifiedOrganization{
		Name:        org.Name,
		Type:        org.SizeClass.Name + " " + org.Type.Name,
		Description: org.Describe(),
		Leader:      org.Leader.Simplify(),
	}

	for _, n := range org.NotableMembers {
		simplified.NotableMembers = append(simplified.NotableMembers, n.Simplify())
	}

	return simplified
}

// Describe constructs a string description of the organization
func (org Organization) Describe() string {
	description := "The " + org.Name + " is " + words.Pronoun(org.SizeClass.Name) + " " + org.SizeClass.Name + " " + org.Type.Name + " with " + strconv.Itoa(org.Size) + " members. "
	description += "It's led by " + strings.Title(org.Leader.Title) + " " + org.Leader.FirstName + " " + org.Leader.LastName + ". "
	description += "The " + org.Type.Name + " is known for being " + org.PrimaryTrait + "."

	return description
}

// RandomSimplified returns a random simplified version of the organization
func RandomSimplified() SimplifiedOrganization {
	org := Random()

	return org.Simplify()
}
