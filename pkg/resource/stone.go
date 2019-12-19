package resource

import "github.com/ironarachne/world/pkg/profession"

func getStone() []Pattern {
	producer, _ := profession.ByName("mason")

	patterns := []Pattern{
		{
			Name:         "milling wheel",
			NameTemplate: "{{.MainMaterial}} milling wheel",
			Description:  "a milling wheel",
			Tags: []string{
				"construction component",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "building stone",
					DescriptionTemplate: "{{.Resource.MainMaterial}} milling wheel",
				},
			},
			Value: 15,
		},
		{
			Name:         "stone block",
			NameTemplate: "{{.MainMaterial}} block",
			Description:  "a block of stone",
			Tags: []string{
				"building block",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "building stone",
					DescriptionTemplate: "{{.Resource.MainMaterial}} block",
				},
			},
			Value: 3,
		},
		{
			Name:         "stone slab",
			NameTemplate: "{{.MainMaterial}} slab",
			Description:  "a slab of stone",
			Tags: []string{
				"building block",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "building stone",
					DescriptionTemplate: "{{.Resource.MainMaterial}} slab",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
