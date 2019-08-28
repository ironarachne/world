package resource

import "github.com/ironarachne/world/pkg/profession"

func getCarpentry() []Pattern {
	producer := profession.ByName("carpenter")

	patterns := []Pattern{
		{
			Name:        "barrel",
			Description: "a barrel",
			Tags: []string{
				"furniture",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.Origin}} wood barrel",
				},
			},
			Value: 1,
		},
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
			Value: 3,
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
			Value: 1,
		},
		{
			Name:        "cask",
			Description: "a cask",
			Tags: []string{
				"furniture",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.Origin}} wood cask",
				},
			},
			Value: 1,
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
			Value: 1,
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
			Value: 5,
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
			Value: 1,
		},
	}

	return patterns
}
