package resource

import "github.com/ironarachne/world/pkg/profession"

func getWeaving() []Pattern {
	producer := profession.ByName("spinner")

	patterns := []Pattern{
		{
			Name:         "large blanket",
			NameTemplate: "large {{.MainMaterial}} blanket",
			Description:  "a large blanket",
			Tags: []string{
				"blanket",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "large {{.Resource.MainMaterial}} blanket",
				},
			},
			Value: 1,
		},
		{
			Name:         "small blanket",
			NameTemplate: "small {{.MainMaterial}} blanket",
			Description:  "a small blanket",
			Tags: []string{
				"blanket",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "small {{.Resource.MainMaterial}} blanket",
				},
			},
			Value: 1,
		},
		{
			Name:         "fabric",
			NameTemplate: "bolt of {{.MainMaterial}} fabric",
			Description:  "a bolt of fabric",
			Tags: []string{
				"fabric",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric fiber",
					DescriptionTemplate: "bolt of {{.Resource.MainMaterial}} fabric",
				},
			},
			Value: 1,
		},
		{
			Name:         "fabric",
			NameTemplate: "bolt of velveteen fabric",
			Description:  "a bolt of fabric",
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
			Name:         "thread",
			NameTemplate: "spool of {{.MainMaterial}} thread",
			Description:  "a spool of thread",
			Tags: []string{
				"thread",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric fiber",
					DescriptionTemplate: "spool of {{.Resource.MainMaterial}} thread",
				},
			},
			Value: 1,
		},
		{
			Name:         "quilt",
			NameTemplate: "{{.MainMaterial}} quilt",
			Description:  "a quilt",
			Tags: []string{
				"blanket",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fabric",
					DescriptionTemplate: "{{.Resource.MainMaterial}} quilt",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
