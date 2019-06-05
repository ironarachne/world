package religion

import (
	"github.com/ironarachne/world/pkg/pantheon"
)

// SimplifiedReligion is a simplified version of Religion for display
type SimplifiedReligion struct {
	Type           string                      `json:"type"`
	LeaderTitle    string                      `json:"leader_title"`
	GatheringPlace string                      `json:"gathering_place"`
	Pantheon       pantheon.SimplifiedPantheon `json:"pantheon"`
}

// Simplify returns a simplified version of a Religion
func (religion Religion) Simplify() SimplifiedReligion {
	simplified := SimplifiedReligion{
		Type:           religion.Class.Name,
		LeaderTitle:    religion.Class.LeaderTitle,
		GatheringPlace: religion.GatheringPlace,
		Pantheon:       religion.Pantheon.Simplify(),
	}
	return simplified
}
