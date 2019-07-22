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
			Name:        "refined metal ore",
			Description: "refined metal ore",
			Tags: []string{
				"refined ore",
				"refined metal ore",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal ore",
					DescriptionTemplate: "refined {{.Resource.Origin}} ore",
				},
			},
		},
	}

	return patterns
}
