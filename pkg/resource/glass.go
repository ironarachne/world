package resource

import "github.com/ironarachne/world/pkg/profession"

func getGlass() []Pattern {
	producer := profession.ByName("glassmaker")

	patterns := []Pattern{
		{
			Name: "glass bottle",
			Description: "glass bottle",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "sand",
					DescriptionTemplate: "{{.Resource.Name}} glass bottle",
				},
			},
		},
		{
			Name: "glass",
			Description: "glass",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "sand",
					DescriptionTemplate: "{{.Resource.Name}} glass",
				},
			},
		},
	}

	return patterns
}