package resource

import "github.com/ironarachne/world/pkg/profession"

func getBreads() []Pattern {
	producer := profession.ByName("baker")

	patterns := []Pattern{
		{
			Name:         "bread",
			NameTemplate: "{{.MainMaterial}} bread",
			Description:  "a loaf of simple bread",
			Tags: []string{
				"bread",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "flour",
					DescriptionTemplate: "{{.Resource.MainMaterial}} bread",
				},
			},
			Value: 1,
		},
		{
			Name:         "roll",
			NameTemplate: "{{.MainMaterial}} roll",
			Description:  "a small hand-sized roll",
			Tags: []string{
				"bread",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "flour",
					DescriptionTemplate: "{{.Resource.MainMaterial}} roll",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
