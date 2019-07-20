package resource

import "github.com/ironarachne/world/pkg/profession"

func getCarpentry() []Pattern {
	producer := profession.ByName("carpenter")

	patterns := []Pattern{
		{
			Name: "bench",
			Description: "a bench",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "wood",
					DescriptionTemplate: "{{.Resource.Name}} wood bench",
				},
			},
		},
		{
			Name: "bowl",
			Description: "a bowl",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "wood",
					DescriptionTemplate: "{{.Resource.Name}} wood bowl",
				},
			},
		},
		{
			Name: "chair",
			Description: "a chair",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "wood",
					DescriptionTemplate: "{{.Resource.Name}} wood chair",
				},
			},
		},
		{
			Name: "table",
			Description: "a table",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "wood",
					DescriptionTemplate: "{{.Resource.Name}} wood table",
				},
			},
		},
		{
			Name: "tankard",
			Description: "a tankard",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "wood",
					DescriptionTemplate: "{{.Resource.Name}} wood tankard",
				},
			},
		},
	}

	return patterns
}