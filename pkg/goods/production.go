package goods

import (
	"context"
	"fmt"
	"github.com/ironarachne/world/pkg/pattern"

	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/resource"
)

// Produce generates a set of resources from what can be produced
func Produce(ctx context.Context, professions []profession.Profession, resources []resource.Resource) ([]resource.Resource, error) {
	var filledPattern pattern.Pattern
	var patterns []pattern.Pattern
	var possiblePatterns []pattern.Pattern
	var produced []resource.Resource
	var resourceForSlot resource.Resource
	var resourcesForSlot []resource.Resource

	allPatterns, err := pattern.All()
	if err != nil {
		err = fmt.Errorf("Failed to produce resources: %w", err)
		return []resource.Resource{}, err
	}

	for _, p := range professions {
		possiblePatterns = []pattern.Pattern{}

		patterns = pattern.ForProfession(p, allPatterns)
		for _, n := range patterns {
			if n.CanMake(resources) {
				possiblePatterns = append(possiblePatterns, n)
			}
		}
		if len(possiblePatterns) > 0 {
			for _, n := range possiblePatterns {
				filledPattern = n
				filledPattern.Slots = []pattern.Slot{}

				for _, s := range n.Slots {
					resourcesForSlot = resource.ByTag(s.RequiredTag, resources)
					resourceForSlot = resource.Random(ctx, resourcesForSlot)
					s.Resource = resourceForSlot
					filledPattern.Slots = append(filledPattern.Slots, s)
				}
				newResource, err := filledPattern.ToResource(ctx)
				if err != nil {
					err = fmt.Errorf("Failed to produce resources: %w", err)
					return []resource.Resource{}, err
				}
				produced = append(produced, newResource)
			}
		}
	}

	return produced, nil
}
