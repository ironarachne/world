package resource

import "github.com/ironarachne/world/pkg/profession"

func getBowery() []Pattern {
	producer := profession.ByName("bowyer")

	patterns := []Pattern{
		{
			Name:        "arrow",
			Description: "a quiver of arrows",
			Tags: []string{
				"ammunition",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "head",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "quiver of arrows with {{.Resource.Origin}} heads",
				},
				{
					Name:                "shaft",
					RequiredTag:         "wood",
					DescriptionTemplate: ", {{.Resource.Origin}} shafts",
				},
				{
					Name:                "fletching",
					RequiredTag:         "feather",
					DescriptionTemplate: ", and {{.Resource.Origin}}-feather fletching",
				},
			},
			Value: 1,
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
			Value: 7,
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
			Value: 5,
		},
		{
			Name:        "crossbow",
			Description: "a crossbow",
			Tags: []string{
				"crossbow",
				"ranged weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "limbs",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.Origin}} crossbow",
				},
				{
					Name:                "bar",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with a {{.Resource.Origin}} bar",
				},
			},
			Value: 8,
		},
	}

	return patterns
}
