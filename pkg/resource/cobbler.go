package resource

import "github.com/ironarachne/world/pkg/profession"

func getCobbler() []Pattern {
	producer := profession.ByName("cobbler")

	patterns := []Pattern{
		{
			Name:        "boots",
			Description: "boots",
			Tags: []string{
				"clothing",
				"footwear",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.Origin}} boots",
				},
			},
		},
		{
			Name:        "shoes",
			Description: "shoes",
			Tags: []string{
				"clothing",
				"footwear",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.Origin}} shoes",
				},
			},
		},
	}

	return patterns
}
