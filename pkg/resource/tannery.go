package resource

import "github.com/ironarachne/world/pkg/profession"

func getTannery() []Pattern {
	producer := profession.ByName("tanner")

	patterns := []Pattern{
		{
			Name:        "leather",
			Description: "leather",
			Tags: []string{
				"leather",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "hide",
					DescriptionTemplate: "{{.Resource.Origin}} leather",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
