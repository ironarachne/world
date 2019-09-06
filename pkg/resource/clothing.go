package resource

import "github.com/ironarachne/world/pkg/profession"

func getClothing() []Pattern {
	producer := profession.ByName("tailor")

	patterns := []Pattern{
		{
			Name:         "backpack",
			NameTemplate: "{{.MainMaterial}} leather backpack",
			Description:  "a backpack",
			Tags: []string{
				"belt",
				"clothing",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.MainMaterial}} leather backpack",
				},
			},
			Value: 1,
		},
		{
			Name:         "small bag",
			NameTemplate: "small {{.MainMaterial}} bag",
			Description:  "a small bag",
			Tags: []string{
				"belt",
				"clothing",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "small {{.Resource.MainMaterial}} bag",
				},
			},
			Value: 1,
		},
		{
			Name:         "medium bag",
			NameTemplate: "medium {{.MainMaterial}} bag",
			Description:  "a medium bag",
			Tags: []string{
				"belt",
				"clothing",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "medium {{.Resource.MainMaterial}} bag",
				},
			},
			Value: 2,
		},
		{
			Name:         "large bag",
			NameTemplate: "large {{.MainMaterial}} bag",
			Description:  "a large bag",
			Tags: []string{
				"belt",
				"clothing",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "large {{.Resource.MainMaterial}} bag",
				},
			},
			Value: 1,
		},
		{
			Name:         "belt",
			NameTemplate: "{{.MainMaterial}} leather belt",
			Description:  "belt",
			Tags: []string{
				"belt",
				"clothing",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.MainMaterial}} leather belt",
				},
				{
					Name:                "buckle",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " with a {{.Resource.MainMaterial}} buckle",
				},
			},
			Value: 1,
		},
		{
			Name:         "blouse",
			NameTemplate: "{{.MainMaterial}} blouse",
			Description:  "a blouse with long sleeves",
			Tags: []string{
				"clothing",
				"shirt",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.MainMaterial}} blouse",
				},
			},
			Value: 1,
		},
		{
			Name:         "breeches",
			NameTemplate: "{{.MainMaterial}} breeches",
			Description:  "breeches",
			Tags: []string{
				"clothing",
				"pants",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.MainMaterial}} breeches",
				},
			},
			Value: 1,
		},
		{
			Name:         "dress",
			NameTemplate: "{{.MainMaterial}} dress",
			Description:  "a dress with long sleeves",
			Tags: []string{
				"clothing",
				"dress",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.MainMaterial}} dress",
				},
			},
			Value: 1,
		},
		{
			Name:         "pants",
			NameTemplate: "{{.MainMaterial}} pants",
			Description:  "pants",
			Tags: []string{
				"clothing",
				"pants",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.MainMaterial}} pants",
				},
			},
			Value: 1,
		},
		{
			Name:         "shirt",
			NameTemplate: "{{.MainMaterial}} shirt",
			Description:  "a shirt with long sleeves",
			Tags: []string{
				"clothing",
				"shirt",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.MainMaterial}} shirt",
				},
			},
			Value: 1,
		},
		{
			Name:         "tunic",
			NameTemplate: "{{.MainMaterial}} tunic",
			Description:  "a tunic with long sleeves that reaches down to the knees",
			Tags: []string{
				"clothing",
				"shirt",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.MainMaterial}} tunic",
				},
			},
			Value: 1,
		},
		{
			Name:         "short tunic",
			NameTemplate: "short {{.MainMaterial}} tunic",
			Description:  "a tunic with long sleeves that reaches down to the waist",
			Tags: []string{
				"clothing",
				"shirt",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "short {{.Resource.MainMaterial}} tunic",
				},
			},
			Value: 1,
		},
		{
			Name:         "long tunic",
			NameTemplate: "long {{.MainMaterial}} tunic",
			Description:  "a tunic with long sleeves that reaches down to the knees",
			Tags: []string{
				"clothing",
				"shirt",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "long {{.Resource.MainMaterial}} tunic",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
