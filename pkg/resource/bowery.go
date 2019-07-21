package resource

import "github.com/ironarachne/world/pkg/profession"

func getBowery() []Pattern {
	producer := profession.ByName("bowyer")

	patterns := []Pattern{
		{
			Name:        "arrow",
			Description: "an arrow",
			Type:        "ammunition",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "head",
					RequiredType:        "metal",
					DescriptionTemplate: "arrow with a {{.Resource.Name}} head",
				},
				{
					Name:                "shaft",
					RequiredType:        "wood",
					DescriptionTemplate: " and {{.Resource.Name}} shaft",
				},
				{
					Name:                "fletching",
					RequiredType:        "feather",
					DescriptionTemplate: " and {{.Resource.Name}} fletching",
				},
			},
		},
		{
			Name:        "long bow",
			Description: "a long bow",
			Type:        "ranged weapon",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "limbs",
					RequiredType:        "wood",
					DescriptionTemplate: "{{.Resource.Name}} long bow",
				},
				{
					Name:                "string",
					RequiredType:        "hide",
					DescriptionTemplate: " strung with {{.Resource.Name}} sinew",
				},
				{
					Name:                "grip",
					RequiredType:        "hide",
					DescriptionTemplate: " with a {{.Resource.Name}} wrapped grip",
				},
			},
		},
		{
			Name:        "short bow",
			Description: "a short bow",
			Type:        "ranged weapon",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "limbs",
					RequiredType:        "wood",
					DescriptionTemplate: "{{.Resource.Name}} short bow",
				},
				{
					Name:                "string",
					RequiredType:        "hide",
					DescriptionTemplate: " strung with {{.Resource.Name}} sinew",
				},
				{
					Name:                "grip",
					RequiredType:        "hide",
					DescriptionTemplate: " with a {{.Resource.Name}} wrapped grip",
				},
			},
		},
	}

	return patterns
}
