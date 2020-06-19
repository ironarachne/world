package organization

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/words"
)

// SimplifiedMember is a simplified version of an organization member
type SimplifiedMember struct {
	CharacterData character.SimplifiedCharacter `json:"character_data"`
	Rank          string                        `json:"rank"`
}

// SimplifiedOrganization is a simplified version of an organization
type SimplifiedOrganization struct {
	Name           string             `json:"name"`
	Type           string             `json:"type"`
	Blazon         string             `json:"blazon"`
	Device         string             `json:"device"`
	Description    string             `json:"description"`
	Leader         SimplifiedMember   `json:"leader"`
	NotableMembers []SimplifiedMember `json:"notable_members"`
}

// Simplify returns a simplified version of a member
func (mem Member) Simplify(ctx context.Context) (SimplifiedMember, error) {
	scd, err := mem.CharacterData.Simplify(ctx)
	if err != nil {
		err = fmt.Errorf("Could not simplify member: %w", err)
		return SimplifiedMember{}, err
	}
	result := SimplifiedMember{
		CharacterData: scd,
		Rank:          mem.Rank.Title,
	}

	return result, nil
}

// Simplify returns a simplified version of the organization
func (org Organization) Simplify(ctx context.Context) (SimplifiedOrganization, error) {
	sl, err := org.Leader.Simplify(ctx)
	if err != nil {
		err = fmt.Errorf("Could not simplify organization: %w", err)
		return SimplifiedOrganization{}, err
	}
	simplified := SimplifiedOrganization{
		Name:        org.Name,
		Type:        org.SizeClass.Name + " " + org.Type.Name,
		Blazon:      org.Heraldry.Blazon,
		Device:      org.Heraldry.ImageURL,
		Description: org.Describe(),
		Leader:      sl,
	}

	for _, n := range org.NotableMembers {
		sn, err := n.Simplify(ctx)
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
	description += "It's led by " + strings.Title(org.Leader.Rank.Title) + " " + org.Leader.CharacterData.FirstName + " " + org.Leader.CharacterData.LastName + ". "
	description += "The " + org.Type.Name + " is known for being " + org.PrimaryTrait + "."

	return description
}

// RandomSimplified returns a random simplified version of the organization
func RandomSimplified(ctx context.Context) (SimplifiedOrganization, error) {
	org, err := Random(ctx)
	if err != nil {
		err = fmt.Errorf("Could not simplify random organization: %w", err)
		return SimplifiedOrganization{}, err
	}

	os, err := org.Simplify(ctx)
	if err != nil {
		err = fmt.Errorf("Could not simplify organization: %w", err)
		return SimplifiedOrganization{}, err
	}

	return os, nil
}
