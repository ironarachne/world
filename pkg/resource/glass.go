package resource

import "github.com/ironarachne/world/pkg/profession"

func getGlass() []Pattern {
	producer, _ := profession.ByName("glassmaker")

	patterns := []Pattern{
		{
			Name:         "glass bottle",
			NameTemplate: "{{.MainMaterial}} glass bottle",
			Description:  "glass bottle",
			Tags: []string{
				"drinking vessel",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "sand",
					DescriptionTemplate: "{{.Resource.MainMaterial}} glass bottle",
				},
			},
			Value: 1,
		},
		{
			Name:         "glass",
			NameTemplate: "{{.MainMaterial}} glass",
			Description:  "glass",
			Tags: []string{
				"glass",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "sand",
					DescriptionTemplate: "{{.Resource.MainMaterial}} glass",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
