package resource

import (
	"github.com/ironarachne/world/pkg/profession"
)

func getLogging() []Pattern {
	producer := profession.ByName("logger")

	patterns := []Pattern{
		{
			Name:        "wood log",
			Description: "a trimmed wood log",
			Tags: []string{
				"wood log",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood",
					DescriptionTemplate: "a bundle of {{.Resource.Origin}} logs",
				},
			},
			Value: 1,
		},
		{
			Name:        "wood plank",
			Description: "a rectangular length of wood",
			Tags: []string{
				"wood plank",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood",
					DescriptionTemplate: "a bundle of {{.Resource.Origin}} planks",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
