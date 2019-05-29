package organization

import (
	"github.com/ironarachne/world/pkg/character"
)

// SimplifiedOrganization is a simplified version of an organization
type SimplifiedOrganization struct {
	Name           string                          `json:"name"`
	Type           string                          `json:"type"`
	Leader         character.SimplifiedCharacter   `json:"leader"`
	NotableMembers []character.SimplifiedCharacter `json:"notable_members"`
}

// Simplify returns a simplified version of the organization
func (org Organization) Simplify() SimplifiedOrganization {
	simplified := SimplifiedOrganization{
		Name:   org.Name,
		Type:   org.SizeClass.Name + " " + org.Type.Name,
		Leader: org.Leader.Simplify(),
	}

	for _, n := range org.NotableMembers {
		simplified.NotableMembers = append(simplified.NotableMembers, n.Simplify())
	}

	return simplified
}

// RandomSimplified returns a random simplified version of the organization
func RandomSimplified() SimplifiedOrganization {
	org := Random()

	return org.Simplify()
}
