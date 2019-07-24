package resource

import "github.com/ironarachne/world/pkg/profession"

func getClothing() []Pattern {
	producer := profession.ByName("tailor")

	patterns := []Pattern{
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
		},
	}

	return patterns
}
