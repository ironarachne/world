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
			Name:        "metal bar",
			Description: "a metal bar",
			Tags: []string{
				"metal bar",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "metal 1",
					RequiredTag:         "refined copper metal ore",
					DescriptionTemplate: "brass metal bar",
				},
				{
					Name:                "metal 2",
					RequiredTag:         "refined zinc metal ore",
					DescriptionTemplate: "",
				},
			},
			OriginOverride: "brass",
		},
		{
			Name:        "metal bar",
			Description: "a metal bar",
			Tags: []string{
				"metal bar",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "metal 1",
					RequiredTag:         "refined copper metal ore",
					DescriptionTemplate: "bronze metal bar",
				},
				{
					Name:                "metal 2",
					RequiredTag:         "refined tin metal ore",
					DescriptionTemplate: "",
				},
			},
			OriginOverride: "bronze",
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
					RequiredTag:         "refined copper metal ore",
					DescriptionTemplate: "electrum metal ingot",
				},
				{
					Name:                "body",
					RequiredTag:         "refined gold metal ore",
					DescriptionTemplate: "",
				},
			},
			OriginOverride: "electrum",
		},
	}

	return patterns
}
