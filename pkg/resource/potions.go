package resource

import "github.com/ironarachne/world/pkg/profession"

func getPotions() []Pattern {
	producer := profession.ByName("alchemist")

	patterns := []Pattern{
		{
			Name:        "healing potion",
			Description: "a healing potion",
			Tags: []string{
				"potion",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "herb",
					DescriptionTemplate: "healing potion",
				},
			},
		},
	}

	return patterns
}
