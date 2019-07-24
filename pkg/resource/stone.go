package resource

import "github.com/ironarachne/world/pkg/profession"

func getStone() []Pattern {
	producer := profession.ByName("mason")

	patterns := []Pattern{
		{
			Name:        "stone block",
			Description: "a block of stone",
			Tags: []string{
				"building block",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "stone",
					DescriptionTemplate: "{{.Resource.Origin}}",
				},
			},
		},
		{
			Name:        "stone slab",
			Description: "a slab of stone",
			Tags: []string{
				"building block",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "stone",
					DescriptionTemplate: "{{.Resource.Origin}}",
				},
			},
		},
	}

	return patterns
}
