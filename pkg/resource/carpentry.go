package resource

import "github.com/ironarachne/world/pkg/profession"

func getCarpentry() []Pattern {
	producer, _ := profession.ByName("carpenter")

	patterns := []Pattern{
		{
			Name:         "barrel",
			NameTemplate: "{{.MainMaterial}} wood barrel",
			Description:  "a barrel",
			Tags: []string{
				"furniture",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.MainMaterial}} wood barrel",
				},
			},
			Value: 1,
		},
		{
			Name:         "bench",
			NameTemplate: "{{.MainMaterial}} wood bench",
			Description:  "a bench",
			Tags: []string{
				"furniture",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.MainMaterial}} wood bench",
				},
			},
			Value: 3,
		},
		{
			Name:         "bowl",
			NameTemplate: "{{.MainMaterial}} wood bowl",
			Description:  "a bowl",
			Tags: []string{
				"food vessel",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood",
					DescriptionTemplate: "{{.Resource.MainMaterial}} wood bowl",
				},
			},
			Value: 1,
		},
		{
			Name:         "cask",
			NameTemplate: "{{.MainMaterial}} wood cask",
			Description:  "a cask",
			Tags: []string{
				"furniture",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.MainMaterial}} wood cask",
				},
			},
			Value: 1,
		},
		{
			Name:         "chair",
			NameTemplate: "{{.MainMaterial}} wood chair",
			Description:  "a chair",
			Tags: []string{
				"furniture",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.MainMaterial}} wood chair",
				},
			},
			Value: 1,
		},
		{
			Name:         "table",
			NameTemplate: "{{.MainMaterial}} wood table",
			Description:  "a table",
			Tags: []string{
				"furniture",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.MainMaterial}} wood table",
				},
			},
			Value: 5,
		},
		{
			Name:         "tankard",
			NameTemplate: "{{.MainMaterial}} wood tankard",
			Description:  "a tankard",
			Tags: []string{
				"drinking vessel",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.MainMaterial}} wood tankard",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
