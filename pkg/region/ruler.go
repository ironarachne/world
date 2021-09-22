package region

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/organization"
)

func (region Region) generateRulingBody(ctx context.Context) (organization.Organization, error) {
	var members []organization.Member

	rulingBody, err := organization.GenerateNobleHouse(ctx, region.Culture)
	if err != nil {
		err = fmt.Errorf("Could not generate region ruler: %w", err)
		return organization.Organization{}, err
	}

	ruler := rulingBody.Leader

	rulerTitle := region.Class.RulerTitleFemale
	if ruler.CharacterData.Gender.Name == "male" {
		rulerTitle = region.Class.RulerTitleMale
	}
	ruler.CharacterData.Title = rulerTitle
	ruler.Rank.Title = rulerTitle
	ranks := rulingBody.Type.Ranks
	ranks[0].Title = rulerTitle
	rulingBody.Type.Ranks = ranks
	rulingBody.Leader = ruler

	for _, m := range rulingBody.NotableMembers {
		if m.Rank.Title == "Spouse" {
			if m.CharacterData.Gender.Name == "male" {
				m.Rank.Title = region.Class.RulerTitleMale
				m.CharacterData.Title = region.Class.RulerTitleMale
			} else {
				m.Rank.Title = region.Class.RulerTitleFemale
				m.CharacterData.Title = region.Class.RulerTitleFemale
			}
		}
		members = append(members, m)
	}

	rulingBody.NotableMembers = members

	return rulingBody, nil
}
