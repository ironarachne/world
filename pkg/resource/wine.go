package resource

import "github.com/ironarachne/world/pkg/profession"

func getWine() []Pattern {
	producer := profession.ByName("vintner")

	patterns := []Pattern{
		{
			Name:        "wine",
			Description: "wine",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fruit",
					DescriptionTemplate: "bottle of {{.Resource.Origin}} wine",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
