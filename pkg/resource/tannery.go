package resource

import "github.com/ironarachne/world/pkg/profession"

func getTannery() []Pattern {
	producer := profession.ByName("tanner")

	patterns := []Pattern{
		{
			Name: "leather",
			Description: "leather",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "hide",
					DescriptionTemplate: "{{.Resource.Name}} leather",
				},
			},
		},
	}

	return patterns
}