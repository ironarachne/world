package resource

import "github.com/ironarachne/world/pkg/profession"

func getSmelting() []Pattern {
	producer := profession.ByName("blacksmith")

	patterns := []Pattern{
		{
			Name:        "metal bar",
			Description: "a metal bar",
			Tags: []string{
				"metal bar",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "refined hard metal ore",
					DescriptionTemplate: "{{.Resource.Origin}} metal bar",
				},
			},
		},
		{
			Name:        "metal ingot",
			Description: "a metal ingot",
			Tags: []string{
				"metal ingot",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "refined soft metal ore",
					DescriptionTemplate: "{{.Resource.Origin}} metal ingot",
				},
			},
		},
	}

	return patterns
}
