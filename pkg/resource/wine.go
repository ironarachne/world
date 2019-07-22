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
					DescriptionTemplate: "{{.Resource.Origin}} wine",
				},
			},
		},
	}

	return patterns
}
