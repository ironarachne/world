package resource

import "github.com/ironarachne/world/pkg/profession"

func getBrewed() []Pattern {
	producer := profession.ByName("brewer")

	patterns := []Pattern{
		{
			Name:        "ale",
			Description: "an ale",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "grain",
					DescriptionTemplate: "{{.Resource.Origin}} ale",
				},
			},
		},
		{
			Name:        "beer",
			Description: "a beer",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "grain",
					DescriptionTemplate: "{{.Resource.Origin}} beer",
				},
			},
		},
		{
			Name:        "lager",
			Description: "an lager",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "grain",
					DescriptionTemplate: "{{.Resource.Origin}} lager",
				},
			},
		},
	}

	return patterns
}
