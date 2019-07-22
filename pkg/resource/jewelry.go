package resource

import "github.com/ironarachne/world/pkg/profession"

func getJewelry() []Pattern {
	producer := profession.ByName("jeweller")

	patterns := []Pattern{
		{
			Name:        "gemstone",
			Description: "a gemstone",
			Tags: []string{
				"gem",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "refined gem ore",
					DescriptionTemplate: "{{.Resource.Origin}}",
				},
			},
		},
	}

	return patterns
}
