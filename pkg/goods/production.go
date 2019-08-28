package goods

import (
	"github.com/ironarachne/world/pkg/profession"
	"github.com/ironarachne/world/pkg/resource"
)

// Produce generates a set of resources from what can be produced
func Produce(professions []profession.Profession, resources []resource.Resource) []resource.Resource {
	var filledPattern resource.Pattern
	var newResource resource.Resource
	var patterns []resource.Pattern
	var possiblePatterns []resource.Pattern
	var produced []resource.Resource
	var resourceForSlot resource.Resource
	var resourcesForSlot []resource.Resource

	allPatterns := resource.AllPatterns()

	for _, p := range professions {
		possiblePatterns = []resource.Pattern{}

		patterns = resource.FindPatternsForProfession(p, allPatterns)
		for _, n := range patterns {
			if n.CanMake(resources) {
				possiblePatterns = append(possiblePatterns, n)
			}
		}
		if len(possiblePatterns) > 0 {
			for _, n := range possiblePatterns {
				filledPattern = resource.Pattern{
					Name:        n.Name,
					Description: n.Description,
					Profession:  n.Profession,
					Tags:        n.Tags,
					Slots:       []resource.Slot{},
					Value:       n.Value,
				}
				for _, s := range n.Slots {
					resourcesForSlot = resource.ByTag(s.RequiredTag, resources)
					resourceForSlot = resource.Random(resourcesForSlot)
					s.Resource = resourceForSlot
					filledPattern.Slots = append(filledPattern.Slots, s)
				}
				newResource = filledPattern.ToResource()
				produced = append(produced, newResource)
			}
		}
	}

	return produced
}
