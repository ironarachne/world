package resource

import "github.com/ironarachne/world/pkg/profession"

func getCobbler() []Pattern {
	producer := profession.ByName("cobbler")

	patterns := []Pattern{
		{
			Name: "boots",
			Description: "boots",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "hide",
					DescriptionTemplate: "{{.Resource.Name}} boots",
				},
			},
		},
		{
			Name: "shoes",
			Description: "shoes",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "hide",
					DescriptionTemplate: "{{.Resource.Name}} shoes",
				},
			},
		},
	}

	return patterns
}