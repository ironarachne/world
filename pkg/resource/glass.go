package resource

import "github.com/ironarachne/world/pkg/profession"

func getGlass() []Pattern {
	producer := profession.ByName("glassmaker")

	patterns := []Pattern{
		{
			Name:        "glass bottle",
			Description: "glass bottle",
			Type:        "drinking vessel",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "sand",
					DescriptionTemplate: "{{.Resource.Name}} glass bottle",
				},
			},
		},
		{
			Name:        "glass",
			Description: "glass",
			Type:        "glass",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "sand",
					DescriptionTemplate: "{{.Resource.Name}} glass",
				},
			},
		},
	}

	return patterns
}
