package resource

import "github.com/ironarachne/world/pkg/profession"

func getWeaving() []Pattern {
	producer := profession.ByName("spinner")

	patterns := []Pattern{
		{
			Name:        "large blanket",
			Description: "a large blanket",
			Tags: []string{
				"blanket",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "large {{.Resource.Origin}} blanket",
				},
			},
			Value: 1,
		},
		{
			Name:        "small blanket",
			Description: "a small blanket",
			Tags: []string{
				"blanket",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "small {{.Resource.Origin}} blanket",
				},
			},
			Value: 1,
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
					DescriptionTemplate: "bolt of {{.Resource.Origin}} fabric",
				},
			},
			Value: 1,
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
					DescriptionTemplate: "bolt of velveteen fabric",
				},
			},
			Value:          5,
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
					DescriptionTemplate: "spool of {{.Resource.Origin}} thread",
				},
			},
			Value: 1,
		},
		{
			Name:        "quilt",
			Description: "a quilt",
			Tags: []string{
				"blanket",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.Origin}} quilt",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
