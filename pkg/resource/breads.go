package resource

import "github.com/ironarachne/world/pkg/profession"

func getBreads() []Pattern {
	producer := profession.ByName("baker")

	patterns := []Pattern{
		{
			Name:        "bread",
			Description: "a loaf of simple bread",
			Tags: []string{
				"bread",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "flour",
					DescriptionTemplate: "{{.Resource.Origin}} bread",
				},
			},
			Value: 1,
		},
		{
			Name:        "roll",
			Description: "a small hand-sized roll",
			Tags: []string{
				"bread",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "flour",
					DescriptionTemplate: "{{.Resource.Origin}} roll",
				},
			},
			Value: 1,
		},
	}

	return patterns
}