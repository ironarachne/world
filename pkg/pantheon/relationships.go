package pantheon

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/pantheon/deity"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/relationship"
	"github.com/ironarachne/world/pkg/words"
)

// GenerateRelationships generates relationships between deities
func (pantheon Pantheon) GenerateRelationships(ctx context.Context) ([]deity.Deity, error) {
	var modifiedDeities []deity.Deity
	var numberOfRelationships int
	var possibleDeities []deity.Deity
	var r relationship.Relationship
	var ir relationship.Relationship
	var rt relationship.Type
	var relationshipAllowed bool
	var allRelationships []relationship.Relationship
	var phrases []string

	rts := relationship.AllTypes()

	for _, d := range pantheon.Deities {
		possibleDeities = deity.Exclude(d, pantheon.Deities)

		numberOfRelationships = random.Intn(ctx, 3)
		for i := 0; i < numberOfRelationships; i++ {
			target, err := deity.Random(ctx, possibleDeities)
			rt = rts[random.Intn(ctx, len(rts))]
			descriptor, err := random.String(ctx, rt.Descriptors)
			if err != nil {
				err = fmt.Errorf("failed to generate relationship descriptor: %w", err)
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
			ir, err = relationship.GetInverse(ctx, r)
			if err != nil {
				err = fmt.Errorf("failed to generate inverse relationship: %w", err)
				return []deity.Deity{}, err
			}
			relationshipAllowed, err = relationship.AllowedIn(r, allRelationships)
			if err != nil {
				err = fmt.Errorf("failed to see if deity relationship is allowed: %w", err)
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
			err = fmt.Errorf("failed to find deity by name: %w", err)
			return []deity.Deity{}, err
		}

		d.Relationships = append(d.Relationships, r)

		modifiedDeities = deity.Replace(d, modifiedDeities)
	}

	for _, d := range modifiedDeities {
		if len(d.Relationships) > 0 {
			phrases = []string{}
			for _, r := range d.Relationships {
				phrases = append(phrases, r.Descriptor+" "+r.Target)
			}
			d.Description += d.Name + " " + words.CombinePhrases(phrases) + "."
			modifiedDeities = deity.Replace(d, modifiedDeities)
		}
	}

	return modifiedDeities, nil
}
