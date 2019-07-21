package resource

import "github.com/ironarachne/world/pkg/profession"

func getPottery() []Pattern {
	producer := profession.ByName("potter")

	patterns := []Pattern{
		{
			Name:        "bowl",
			Description: "a clay bowl",
			Type:        "food vessel",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "clay",
					DescriptionTemplate: "{{.Resource.Name}} bowl",
				},
			},
		},
		{
			Name:        "mug",
			Description: "a clay mug",
			Type:        "drinking vessel",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "clay",
					DescriptionTemplate: "{{.Resource.Name}} mug",
				},
			},
		},
		{
			Name:        "pitcher",
			Description: "a clay pitcher",
			Type:        "liquid vessel",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "clay",
					DescriptionTemplate: "{{.Resource.Name}} pitcher",
				},
			},
		},
		{
			Name:        "vase",
			Description: "a clay vase",
			Type:        "vase",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "clay",
					DescriptionTemplate: "{{.Resource.Name}} vase",
				},
			},
		},
	}

	return patterns
}
