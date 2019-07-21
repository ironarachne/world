package resource

import "github.com/ironarachne/world/pkg/profession"

func getCarpentry() []Pattern {
	producer := profession.ByName("carpenter")

	patterns := []Pattern{
		{
			Name:        "bench",
			Description: "a bench",
			Type:        "furniture",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "wood",
					DescriptionTemplate: "{{.Resource.Name}} wood bench",
				},
			},
		},
		{
			Name:        "bowl",
			Description: "a bowl",
			Type:        "food vessel",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "wood",
					DescriptionTemplate: "{{.Resource.Name}} wood bowl",
				},
			},
		},
		{
			Name:        "chair",
			Description: "a chair",
			Type:        "furniture",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "wood",
					DescriptionTemplate: "{{.Resource.Name}} wood chair",
				},
			},
		},
		{
			Name:        "table",
			Description: "a table",
			Type:        "furniture",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "wood",
					DescriptionTemplate: "{{.Resource.Name}} wood table",
				},
			},
		},
		{
			Name:        "tankard",
			Description: "a tankard",
			Type:        "drinking vessel",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "wood",
					DescriptionTemplate: "{{.Resource.Name}} wood tankard",
				},
			},
		},
	}

	return patterns
}
