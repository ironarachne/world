package resource

import "github.com/ironarachne/world/pkg/profession"

func getMilled() []Pattern {
	producer := profession.ByName("miller")

	patterns := []Pattern{
		{
			Name:        "flour",
			Description: "flour",
			Tags: []string{
				"flour",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:        "grain",
					DescriptionTemplate: "{{.Resource.Origin}} flour",
				},
			},
		},
	}

	return patterns
}
