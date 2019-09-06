package resource

import "github.com/ironarachne/world/pkg/profession"

func getMined() []Pattern {
	producer := profession.ByName("miner")

	patterns := []Pattern{
		{
			Name:         "refined gem ore",
			NameTemplate: "refined {{.MainMaterial}} ore",
			Description:  "refined gem ore",
			Tags: []string{
				"refined ore",
				"refined gem ore",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "gem ore",
					DescriptionTemplate: "refined {{.Resource.MainMaterial}} ore",
				},
			},
			Value: 2,
		},
		{
			Name:         "refined soft metal ore",
			NameTemplate: "refined {{.MainMaterial}} ore",
			Description:  "refined soft metal ore",
			Tags: []string{
				"refined ore",
				"refined soft metal ore",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "soft metal ore",
					DescriptionTemplate: "refined {{.Resource.MainMaterial}} ore",
				},
			},
			Value: 2,
		},
		{
			Name:         "refined hard metal ore",
			NameTemplate: "refined {{.MainMaterial}} ore",
			Description:  "refined hard metal ore",
			Tags: []string{
				"refined ore",
				"refined hard metal ore",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "hard metal ore",
					DescriptionTemplate: "refined {{.Resource.MainMaterial}} ore",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
