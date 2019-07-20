package resource

import "github.com/ironarachne/world/pkg/profession"

func getStone() []Pattern {
	producer := profession.ByName("mason")

	patterns := []Pattern{
		{
			Name: "stone block",
			Description: "a block of stone",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "stone",
					DescriptionTemplate: "{{.Resource.Name}} block",
				},
			},
		},
		{
			Name: "stone slab",
			Description: "a slab of stone",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "stone",
					DescriptionTemplate: "{{.Resource.Name}} slab",
				},
			},
		},
	}

	return patterns
}
