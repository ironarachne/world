package resource

import "github.com/ironarachne/world/pkg/profession"

func getBreads() []Pattern {
	producer := profession.ByName("baker")

	patterns := []Pattern{
		{
			Name:        "bread",
			Description: "a loaf of simple bread",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "flour",
					DescriptionTemplate: "{{.Resource.Name}} bread",
				},
			},
		},
		{
			Name:        "roll",
			Description: "a small hand-sized roll",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "flour",
					DescriptionTemplate: "{{.Resource.Name}} roll",
				},
			},
		},
	}

	return patterns
}