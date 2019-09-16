package pantheon

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/pantheon/deity"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/relationship"
)

// GenerateRelationships generates relationships between deities
func (pantheon Pantheon) GenerateRelationships() ([]deity.Deity, error) {
	var modifiedDeities []deity.Deity
	var numberOfRelationships int
	var possibleDeities []deity.Deity
	var r relationship.Relationship
	var ir relationship.Relationship
	var rt relationship.Type
	var relationshipAllowed bool

	rts := relationship.AllTypes()

	allRelationships := []relationship.Relationship{}

	for _, d := range pantheon.Deities {
		possibleDeities = deity.Exclude(d, pantheon.Deities)

		numberOfRelationships = rand.Intn(3)
		for i := 0; i < numberOfRelationships; i++ {
			target, err := deity.Random(possibleDeities)
			rt = rts[rand.Intn(len(rts))]
			descriptor, err := random.String(rt.Descriptors)
			if err != nil {
				err = fmt.Errorf("Could not generate deity relationships: %w", err)
				return []deity.Deity{}, err
			}

			if rt.Name == "parent" {
				if d.Gender.Name == "male" {
					descriptor = "is the father of"
				} else if d.Gender.Name == "female" {
					descriptor = "is the mother of"
				}
			} else if rt.Name == "child" {
				if d.Gender.Name == "male" {
					descriptor = "is the son of"
				} else if d.Gender.Name == "female" {
					descriptor = "is the daughter of"
				}
			}
			r = relationship.Relationship{Origin: d.Name, Target: target.Name, Descriptor: descriptor, Type: rt.Name}
			ir, err = relationship.GetInverse(r)
			if err != nil {
				err = fmt.Errorf("Could not generate deity relationships: %w", err)
				return []deity.Deity{}, err
			}
			relationshipAllowed, err = relationship.AllowedIn(r, allRelationships)
			if err != nil {
				err = fmt.Errorf("Could not generate deity relationships: %w", err)
				return []deity.Deity{}, err
			}
			if !relationship.InSlice(r, allRelationships) && relationshipAllowed {
				allRelationships = append(allRelationships, r)
				allRelationships = append(allRelationships, ir)
			}
		}
	}

	modifiedDeities = pantheon.Deities

	for _, r := range allRelationships {
		d, err := deity.ByName(r.Origin, modifiedDeities)
		if err != nil {
			err = fmt.Errorf("Could not generate deity relationships: %w", err)
			return []deity.Deity{}, err
		}

		d.Relationships = append(d.Relationships, r)
		modifiedDeities = deity.Replace(d, modifiedDeities)
	}

	return modifiedDeities, nil
}
