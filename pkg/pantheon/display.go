package pantheon

import "github.com/ironarachne/world/pkg/pantheon/deity"

// SimplifiedPantheon is a simplified version of the data for display
type SimplifiedPantheon struct {
	Deities []deity.SimplifiedDeity `json:"deities"`
}

// Simplify returns a simplified pantheon for display
func (pantheon Pantheon) Simplify() SimplifiedPantheon {
	sp := SimplifiedPantheon{}
	simplifiedDeity := deity.SimplifiedDeity{}

	for _, d := range pantheon.Deities {
		simplifiedDeity = d.Simplify()
		sp.Deities = append(sp.Deities, simplifiedDeity)
	}

	return sp
}
