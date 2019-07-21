package resource

import "github.com/ironarachne/world/pkg/profession"

func getBrewed() []Pattern {
	producer := profession.ByName("brewer")

	patterns := []Pattern{
		{
			Name:        "ale",
			Description: "an ale",
			Type:        "alcohol",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "grain",
					DescriptionTemplate: "{{.Resource.Name}} ale",
				},
			},
		},
		{
			Name:        "beer",
			Description: "a beer",
			Type:        "alcohol",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "grain",
					DescriptionTemplate: "{{.Resource.Name}} beer",
				},
			},
		},
		{
			Name:        "lager",
			Description: "an lager",
			Type:        "alcohol",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "grain",
					DescriptionTemplate: "{{.Resource.Name}} lager",
				},
			},
		},
	}

	return patterns
}
