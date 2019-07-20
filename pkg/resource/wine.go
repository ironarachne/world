package resource

import "github.com/ironarachne/world/pkg/profession"

func getWine() []Pattern {
	producer := profession.ByName("vintner")

	patterns := []Pattern{
		{
			Name: "wine",
			Description: "wine",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "fruit",
					DescriptionTemplate: "{{.Resource.Name}} wine",
				},
			},
		},
	}

	return patterns
}