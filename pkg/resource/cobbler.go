package resource

import "github.com/ironarachne/world/pkg/profession"

func getCobbler() []Pattern {
	producer, _ := profession.ByName("cobbler")

	patterns := []Pattern{
		{
			Name:         "boots",
			NameTemplate: "{{.MainMaterial}} leather boots",
			Description:  "boots",
			Tags: []string{
				"clothing",
				"footwear",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.MainMaterial}} leather boots",
				},
			},
			Value: 1,
		},
		{
			Name:         "boots",
			NameTemplate: "{{.MainMaterial}} hide boots",
			Description:  "boots",
			Tags: []string{
				"clothing",
				"footwear",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.MainMaterial}} hide boots",
				},
			},
			Value: 1,
		},
		{
			Name:         "sandals",
			NameTemplate: "{{.MainMaterial}} hide sandals",
			Description:  "sandals",
			Tags: []string{
				"clothing",
				"footwear",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.MainMaterial}} hide sandals",
				},
			},
			Value: 1,
		},
		{
			Name:         "shoes",
			NameTemplate: "{{.MainMaterial}} hide shoes",
			Description:  "shoes",
			Tags: []string{
				"clothing",
				"footwear",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.MainMaterial}} hide shoes",
				},
			},
			Value: 1,
		},
		{
			Name:         "slippers",
			NameTemplate: "{{.MainMaterial}} slippers",
			Description:  "slippers",
			Tags: []string{
				"clothing",
				"footwear",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.MainMaterial}} slippers",
				},
			},
			Value: 1,
		},
		{
			Name:         "turnshoes",
			NameTemplate: "{{.MainMaterial}} hide turnshoes",
			Description:  "turnshoes",
			Tags: []string{
				"clothing",
				"footwear",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.MainMaterial}} hide turnshoes",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
