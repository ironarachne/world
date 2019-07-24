package resource

import "github.com/ironarachne/world/pkg/profession"

func getMined() []Pattern {
	producer := profession.ByName("miner")

	patterns := []Pattern{
		{
			Name:        "refined gem ore",
			Description: "refined gem ore",
			Tags: []string{
				"refined ore",
				"refined gem ore",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "gem ore",
					DescriptionTemplate: "refined {{.Resource.Origin}} ore",
				},
			},
		},
		{
			Name:        "refined soft metal ore",
			Description: "refined soft metal ore",
			Tags: []string{
				"refined ore",
				"refined soft metal ore",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "soft metal ore",
					DescriptionTemplate: "refined {{.Resource.Origin}} ore",
				},
			},
		},
		{
			Name:        "refined hard metal ore",
			Description: "refined hard metal ore",
			Tags: []string{
				"refined ore",
				"refined hard metal ore",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "hard metal ore",
					DescriptionTemplate: "refined {{.Resource.Origin}} ore",
				},
			},
		},
	}

	return patterns
}
