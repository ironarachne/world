package resource

import "github.com/ironarachne/world/pkg/profession"

func getCarpentry() []Pattern {
	producer := profession.ByName("carpenter")

	patterns := []Pattern{
		{
			Name:        "bench",
			Description: "a bench",
			Tags: []string{
				"furniture",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.Origin}} wood bench",
				},
			},
		},
		{
			Name:        "bowl",
			Description: "a bowl",
			Tags: []string{
				"food vessel",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood",
					DescriptionTemplate: "{{.Resource.Origin}} wood bowl",
				},
			},
		},
		{
			Name:        "chair",
			Description: "a chair",
			Tags: []string{
				"furniture",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.Origin}} wood chair",
				},
			},
		},
		{
			Name:        "table",
			Description: "a table",
			Tags: []string{
				"furniture",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.Origin}} wood table",
				},
			},
		},
		{
			Name:        "tankard",
			Description: "a tankard",
			Tags: []string{
				"drinking vessel",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.Origin}} wood tankard",
				},
			},
		},
	}

	return patterns
}
