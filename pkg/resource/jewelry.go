package resource

import "github.com/ironarachne/world/pkg/profession"

func getJewelry() []Pattern {
	producer := profession.ByName("jeweller")

	patterns := []Pattern{
		{
			Name:        "gemstone",
			Description: "a gemstone",
			Tags: []string{
				"gem",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "refined gem ore",
					DescriptionTemplate: "{{.Resource.Origin}}",
				},
			},
			Value: 10,
		},
		{
			Name:        "necklace",
			Description: "a necklace",
			Tags: []string{
				"necklace",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: "{{.Resource.Origin}} necklace",
				},
			},
			Value: 10,
		},
		{
			Name:        "pendant",
			Description: "a pendant",
			Tags: []string{
				"pendant",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "gem",
					RequiredTag:         "gem",
					DescriptionTemplate: "{{.Resource.Origin}} pendant",
				},
				{
					Name:                "thong",
					RequiredTag:         "leather",
					DescriptionTemplate: " on a {{.Resource.Origin}} leather thong",
				},
			},
			Value: 10,
		},
		{
			Name:        "ring",
			Description: "a ring",
			Tags: []string{
				"ring",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "band",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: "{{.Resource.Origin}} ring",
				},
				{
					Name:                "gem",
					RequiredTag:         "gem",
					DescriptionTemplate: " set with {{.Resource.Origin}}",
				},
			},
			Value: 10,
		},
		{
			Name:        "ring",
			Description: "a ring",
			Tags: []string{
				"ring",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "band",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: "{{.Resource.Origin}} ring",
				},
			},
			Value: 10,
		},
	}

	return patterns
}
