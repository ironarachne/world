package religion

import (
	"github.com/ironarachne/world/pkg/pantheon"
)

// SimplifiedReligion is a simplified version of Religion for display
type SimplifiedReligion struct {
	Type string
	LeaderTitle string
	GatheringPlace string
	Pantheon pantheon.SimplifiedPantheon
}

// Simplify returns a simplified version of a Religion
func (religion Religion) Simplify() SimplifiedReligion {
	simplified := SimplifiedReligion{
		Type: religion.Class.Name,
		LeaderTitle: religion.Class.LeaderTitle,
		GatheringPlace: religion.GatheringPlace,
		Pantheon: religion.Pantheon.Simplify(),
	}
	return simplified
}