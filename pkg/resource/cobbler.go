package resource

import "github.com/ironarachne/world/pkg/profession"

func getCobbler() []Pattern {
	producer := profession.ByName("cobbler")

	patterns := []Pattern{
		{
			Name:        "boots",
			Description: "boots",
			Tags: []string{
				"clothing",
				"footwear",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.Origin}} leather boots",
				},
			},
			Value: 1,
		},
		{
			Name:        "boots",
			Description: "boots",
			Tags: []string{
				"clothing",
				"footwear",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.Origin}} hide boots",
				},
			},
			Value: 1,
		},
		{
			Name:        "sandals",
			Description: "sandals",
			Tags: []string{
				"clothing",
				"footwear",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.Origin}} hide sandals",
				},
			},
			Value: 1,
		},
		{
			Name:        "shoes",
			Description: "shoes",
			Tags: []string{
				"clothing",
				"footwear",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.Origin}} hide shoes",
				},
			},
			Value: 1,
		},
		{
			Name:        "slippers",
			Description: "slippers",
			Tags: []string{
				"clothing",
				"footwear",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.Origin}} slippers",
				},
			},
			Value: 1,
		},
		{
			Name:        "turnshoes",
			Description: "turnshoes",
			Tags: []string{
				"clothing",
				"footwear",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.Origin}} hide turnshoes",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
