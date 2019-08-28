package resource

import "github.com/ironarachne/world/pkg/profession"

func getGlass() []Pattern {
	producer := profession.ByName("glassmaker")

	patterns := []Pattern{
		{
			Name:        "glass bottle",
			Description: "glass bottle",
			Tags: []string{
				"drinking vessel",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "sand",
					DescriptionTemplate: "{{.Resource.Origin}} glass bottle",
				},
			},
			Value: 1,
		},
		{
			Name:        "glass",
			Description: "glass",
			Tags: []string{
				"glass",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "sand",
					DescriptionTemplate: "{{.Resource.Origin}} glass",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
