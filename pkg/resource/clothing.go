package resource

import "github.com/ironarachne/world/pkg/profession"

func getClothing() []Pattern {
	producer := profession.ByName("tailor")

	patterns := []Pattern{
		{
			Name:        "backpack",
			Description: "a backpack",
			Tags: []string{
				"belt",
				"clothing",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.Origin}} leather backpack",
				},
			},
			Value: 1,
		},
		{
			Name:        "small bag",
			Description: "a small bag",
			Tags: []string{
				"belt",
				"clothing",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "small {{.Resource.Origin}} bag",
				},
			},
			Value: 1,
		},
		{
			Name:        "medium bag",
			Description: "a medium bag",
			Tags: []string{
				"belt",
				"clothing",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "medium {{.Resource.Origin}} bag",
				},
			},
			Value: 2,
		},
		{
			Name:        "large bag",
			Description: "a large bag",
			Tags: []string{
				"belt",
				"clothing",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "large {{.Resource.Origin}} bag",
				},
			},
			Value: 1,
		},
		{
			Name:        "belt",
			Description: "belt",
			Tags: []string{
				"belt",
				"clothing",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.Origin}} leather belt",
				},
				{
					Name:                "buckle",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " with a {{.Resource.Origin}} buckle",
				},
			},
			Value: 1,
		},
		{
			Name:        "blouse",
			Description: "a blouse with long sleeves",
			Tags: []string{
				"clothing",
				"shirt",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.Origin}} blouse",
				},
			},
			Value: 1,
		},
		{
			Name:        "breeches",
			Description: "breeches",
			Tags: []string{
				"clothing",
				"pants",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.Origin}} breeches",
				},
			},
			Value: 1,
		},
		{
			Name:        "dress",
			Description: "a dress with long sleeves",
			Tags: []string{
				"clothing",
				"dress",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.Origin}} dress",
				},
			},
			Value: 1,
		},
		{
			Name:        "pants",
			Description: "pants",
			Tags: []string{
				"clothing",
				"pants",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.Origin}} pants",
				},
			},
			Value: 1,
		},
		{
			Name:        "shirt",
			Description: "a shirt with long sleeves",
			Tags: []string{
				"clothing",
				"shirt",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.Origin}} shirt",
				},
			},
			Value: 1,
		},
		{
			Name:        "tunic",
			Description: "a tunic with long sleeves that reaches down to the knees",
			Tags: []string{
				"clothing",
				"shirt",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.Origin}} tunic",
				},
			},
			Value: 1,
		},
		{
			Name:        "short tunic",
			Description: "a tunic with long sleeves that reaches down to the waist",
			Tags: []string{
				"clothing",
				"shirt",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "short {{.Resource.Origin}} tunic",
				},
			},
			Value: 1,
		},
		{
			Name:        "long tunic",
			Description: "a tunic with long sleeves that reaches down to the knees",
			Tags: []string{
				"clothing",
				"shirt",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "long {{.Resource.Origin}} tunic",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
