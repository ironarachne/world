package resource

import "github.com/ironarachne/world/pkg/profession"

func getWeaving() []Pattern {
	producer := profession.ByName("spinner")

	patterns := []Pattern{
		{
			Name:        "blanket",
			Description: "a large blanket",
			Tags: []string{
				"blanket",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.Origin}} blanket",
				},
			},
		},
		{
			Name:        "fabric",
			Description: "a bolt of fabric",
			Tags: []string{
				"fabric",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric fiber",
					DescriptionTemplate: "{{.Resource.Origin}}",
				},
			},
		},
		{
			Name:        "fabric",
			Description: "a bolt of fabric",
			Tags: []string{
				"fabric",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "cotton",
					DescriptionTemplate: "velveteen",
				},
			},
			OriginOverride: "velveteen",
		},
		{
			Name:        "thread",
			Description: "a spool of thread",
			Tags: []string{
				"thread",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric fiber",
					DescriptionTemplate: "{{.Resource.Origin}} thread",
				},
			},
		},
	}

	return patterns
}
