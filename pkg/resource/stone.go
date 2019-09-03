package resource

import "github.com/ironarachne/world/pkg/profession"

func getStone() []Pattern {
	producer := profession.ByName("mason")

	patterns := []Pattern{
		{
			Name:        "milling wheel",
			Description: "a milling wheel",
			Tags: []string{
				"construction component",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "building stone",
					DescriptionTemplate: "{{.Resource.Origin}} milling wheel",
				},
			},
			Value: 15,
		},
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
					RequiredTag:         "building stone",
					DescriptionTemplate: "{{.Resource.Origin}} block",
				},
			},
			Value: 3,
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
					RequiredTag:         "building stone",
					DescriptionTemplate: "{{.Resource.Origin}} slab",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
