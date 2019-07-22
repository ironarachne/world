package resource

import "github.com/ironarachne/world/pkg/profession"

func getBowery() []Pattern {
	producer := profession.ByName("bowyer")

	patterns := []Pattern{
		{
			Name:        "arrow",
			Description: "an arrow",
			Tags: []string{
				"ammunition",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "head",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: "arrow with a {{.Resource.Origin}} head",
				},
				{
					Name:                "shaft",
					RequiredTag:         "wood",
					DescriptionTemplate: " and {{.Resource.Origin}} shaft",
				},
				{
					Name:                "fletching",
					RequiredTag:         "feather",
					DescriptionTemplate: " and {{.Resource.Origin}} fletching",
				},
			},
		},
		{
			Name:        "long bow",
			Description: "a long bow",
			Tags: []string{
				"bow",
				"ranged weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "limbs",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.Origin}} long bow",
				},
				{
					Name:                "string",
					RequiredTag:         "sinew",
					DescriptionTemplate: " strung with {{.Resource.Origin}} sinew",
				},
				{
					Name:                "grip",
					RequiredTag:         "leather",
					DescriptionTemplate: " with a {{.Resource.Origin}} wrapped grip",
				},
			},
		},
		{
			Name:        "short bow",
			Description: "a short bow",
			Tags: []string{
				"bow",
				"ranged weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "limbs",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.Origin}} short bow",
				},
				{
					Name:                "string",
					RequiredTag:         "sinew",
					DescriptionTemplate: " strung with {{.Resource.Origin}} sinew",
				},
				{
					Name:                "grip",
					RequiredTag:         "leather",
					DescriptionTemplate: " with a {{.Resource.Origin}} wrapped grip",
				},
			},
		},
	}

	return patterns
}
