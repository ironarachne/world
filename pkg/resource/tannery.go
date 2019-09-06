package resource

import "github.com/ironarachne/world/pkg/profession"

func getTannery() []Pattern {
	producer := profession.ByName("tanner")

	patterns := []Pattern{
		{
			Name:         "leather",
			NameTemplate: "{{.MainMaterial}} leather",
			Description:  "leather",
			Tags: []string{
				"leather",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "hide",
					DescriptionTemplate: "{{.Resource.MainMaterial}} leather",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
