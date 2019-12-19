package resource

import "github.com/ironarachne/world/pkg/profession"

func getPottery() []Pattern {
	producer, _ := profession.ByName("potter")

	patterns := []Pattern{
		{
			Name:         "bowl",
			NameTemplate: "{{.MainMaterial}} bowl",
			Description:  "a clay bowl",
			Tags: []string{
				"food vessel",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "clay",
					DescriptionTemplate: "{{.Resource.MainMaterial}} bowl",
				},
			},
			Value: 1,
		},
		{
			Name:         "mug",
			NameTemplate: "{{.MainMaterial}} mug",
			Description:  "a clay mug",
			Tags: []string{
				"drinking vessel",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "clay",
					DescriptionTemplate: "{{.Resource.MainMaterial}} mug",
				},
			},
			Value: 1,
		},
		{
			Name:         "pitcher",
			NameTemplate: "{{.MainMaterial}} pitcher",
			Description:  "a clay pitcher",
			Tags: []string{
				"liquid vessel",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "clay",
					DescriptionTemplate: "{{.Resource.MainMaterial}} pitcher",
				},
			},
			Value: 1,
		},
		{
			Name:         "vase",
			NameTemplate: "{{.MainMaterial}} vase",
			Description:  "a clay vase",
			Tags: []string{
				"vase",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "clay",
					DescriptionTemplate: "{{.Resource.MainMaterial}} vase",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
