package resource

import "github.com/ironarachne/world/pkg/profession"

func getJewelry() []Pattern {
	producer := profession.ByName("jeweller")

	patterns := []Pattern{
		{
			Name:         "gemstone",
			NameTemplate: "{{.MainMaterial}}",
			Description:  "a gemstone",
			Tags: []string{
				"gem",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "refined gem ore",
					DescriptionTemplate: "{{.Resource.MainMaterial}}",
				},
			},
			Value: 10,
		},
		{
			Name:         "necklace",
			NameTemplate: "{{.MainMaterial}} necklace",
			Description:  "a necklace",
			Tags: []string{
				"necklace",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: "{{.Resource.MainMaterial}} necklace",
				},
			},
			Value: 10,
		},
		{
			Name:         "pendant",
			NameTemplate: "{{.MainMaterial}} pendant",
			Description:  "a pendant",
			Tags: []string{
				"pendant",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "gem",
					RequiredTag:         "gem",
					DescriptionTemplate: "{{.Resource.MainMaterial}} pendant",
				},
				{
					Name:                "thong",
					RequiredTag:         "leather",
					DescriptionTemplate: " on a {{.Resource.MainMaterial}} leather thong",
				},
			},
			Value: 10,
		},
		{
			Name:         "ring",
			NameTemplate: "{{.MainMaterial}} ring",
			Description:  "a ring",
			Tags: []string{
				"ring",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "band",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: "{{.Resource.MainMaterial}} ring",
				},
				{
					Name:                "gem",
					RequiredTag:         "gem",
					DescriptionTemplate: " set with {{.Resource.MainMaterial}}",
				},
			},
			Value: 10,
		},
		{
			Name:         "ring",
			NameTemplate: "{{.MainMaterial}} ring",
			Description:  "a ring",
			Tags: []string{
				"ring",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "band",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: "{{.Resource.MainMaterial}} ring",
				},
			},
			Value: 10,
		},
	}

	return patterns
}
