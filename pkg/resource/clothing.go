package resource

import "github.com/ironarachne/world/pkg/profession"

func getClothing() []Pattern {
	producer := profession.ByName("tailor")

	patterns := []Pattern{
		{
			Name:        "belt",
			Description: "belt",
			Type:        "belt",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "hide",
					DescriptionTemplate: "{{.Resource.Name}} belt",
				},
				{
					Name:                "buckle",
					RequiredType:        "metal",
					DescriptionTemplate: " with a {{.Resource.Name}} buckle",
				},
			},
		},
		{
			Name:        "pants",
			Description: "pants",
			Type:        "pants",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "fabric",
					DescriptionTemplate: "{{.Resource.Name}} pants",
				},
			},
		},
		{
			Name:        "shirt",
			Description: "a shirt with long sleeves",
			Type:        "shirt",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "fabric",
					DescriptionTemplate: "{{.Resource.Name}} shirt",
				},
			},
		},
		{
			Name:        "tunic",
			Description: "a tunic with long sleeves that reaches down to the knees",
			Type:        "shirt",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "fabric",
					DescriptionTemplate: "{{.Resource.Name}} tunic",
				},
			},
		},
	}

	return patterns
}
