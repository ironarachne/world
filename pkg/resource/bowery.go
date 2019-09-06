package resource

import "github.com/ironarachne/world/pkg/profession"

func getBowery() []Pattern {
	producer := profession.ByName("bowyer")

	patterns := []Pattern{
		{
			Name:         "arrow",
			NameTemplate: "a quiver of arrows",
			Description:  "a quiver of arrows",
			Tags: []string{
				"ammunition",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "head",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "quiver of arrows with {{.Resource.MainMaterial}} heads",
				},
				{
					Name:                "shaft",
					RequiredTag:         "wood",
					DescriptionTemplate: ", {{.Resource.MainMaterial}} shafts",
				},
				{
					Name:                "fletching",
					RequiredTag:         "feather",
					DescriptionTemplate: ", and {{.Resource.MainMaterial}}-feather fletching",
				},
			},
			Value: 1,
		},
		{
			Name:         "long bow",
			NameTemplate: "{{.MainMaterial}} long bow",
			Description:  "a long bow",
			Tags: []string{
				"bow",
				"ranged weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "limbs",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.MainMaterial}} long bow",
				},
				{
					Name:                "string",
					RequiredTag:         "sinew",
					DescriptionTemplate: " strung with {{.Resource.MainMaterial}} sinew",
				},
				{
					Name:                "grip",
					RequiredTag:         "leather",
					DescriptionTemplate: " with a {{.Resource.MainMaterial}} wrapped grip",
				},
			},
			Value: 7,
		},
		{
			Name:         "short bow",
			NameTemplate: "{{.MainMaterial}} short bow",
			Description:  "a short bow",
			Tags: []string{
				"bow",
				"ranged weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "limbs",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.MainMaterial}} short bow",
				},
				{
					Name:                "string",
					RequiredTag:         "sinew",
					DescriptionTemplate: " strung with {{.Resource.MainMaterial}} sinew",
				},
				{
					Name:                "grip",
					RequiredTag:         "leather",
					DescriptionTemplate: " with a {{.Resource.MainMaterial}} wrapped grip",
				},
			},
			Value: 5,
		},
		{
			Name:         "crossbow",
			NameTemplate: "{{.MainMaterial}} crossbow",
			Description:  "a crossbow",
			Tags: []string{
				"crossbow",
				"ranged weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "limbs",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.MainMaterial}} crossbow",
				},
				{
					Name:                "bar",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with a {{.Resource.MainMaterial}} bar",
				},
			},
			Value: 8,
		},
	}

	return patterns
}
