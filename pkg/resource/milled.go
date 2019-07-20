package resource

import "github.com/ironarachne/world/pkg/profession"

func getMilled() []Pattern {
	producer := profession.ByName("miller")

	patterns := []Pattern{
		{
			Name: "flour",
			Description: "flour",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "grain",
					DescriptionTemplate: "{{.Resource.Name}} flour",
				},
			},
		},
	}

	return patterns
}