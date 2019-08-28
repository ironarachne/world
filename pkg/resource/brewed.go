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
					DescriptionTemplate: "bottle of {{.Resource.Origin}} ale",
				},
			},
			Value: 1,
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
					DescriptionTemplate: "bottle of {{.Resource.Origin}} beer",
				},
			},
			Value: 1,
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
					DescriptionTemplate: "bottle of {{.Resource.Origin}} lager",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
