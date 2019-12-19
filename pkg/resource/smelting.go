package resource

import "github.com/ironarachne/world/pkg/profession"

func getSmelting() []Pattern {
	producer, _ := profession.ByName("blacksmith")

	patterns := []Pattern{
		{
			Name:         "metal bar",
			NameTemplate: "{{.MainMaterial}} metal bar",
			Description:  "a metal bar",
			Tags: []string{
				"metal bar",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "refined hard metal ore",
					DescriptionTemplate: "{{.Resource.MainMaterial}} metal bar",
				},
			},
			Value: 2,
		},
		{
			Name:         "metal bar",
			NameTemplate: "brass metal bar",
			Description:  "a metal bar",
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
			OriginOverride:       "brass",
			MainMaterialOverride: "brass",
			Value:                5,
		},
		{
			Name:         "metal bar",
			NameTemplate: "bronze metal bar",
			Description:  "a metal bar",
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
			OriginOverride:       "bronze",
			MainMaterialOverride: "bronze",
			Value:                5,
		},
		{
			Name:         "metal ingot",
			NameTemplate: "{{.MainMaterial}} metal ingot",
			Description:  "a metal ingot",
			Tags: []string{
				"metal ingot",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "refined soft metal ore",
					DescriptionTemplate: "{{.Resource.MainMaterial}} metal ingot",
				},
			},
			Value: 2,
		},
		{
			Name:         "metal ingot",
			NameTemplate: "electrum metal ingot",
			Description:  "a metal ingot",
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
			OriginOverride:       "electrum",
			MainMaterialOverride: "electrum",
			Value:                5,
		},
	}

	return patterns
}
