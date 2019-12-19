package resource

import "github.com/ironarachne/world/pkg/profession"

func getBrewed() []Pattern {
	producer, _ := profession.ByName("brewer")

	patterns := []Pattern{
		{
			Name:         "ale",
			NameTemplate: "{{.MainMaterial}} ale",
			Description:  "an ale",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "grain",
					DescriptionTemplate: "bottle of {{.Resource.MainMaterial}} ale",
				},
			},
			Value: 1,
		},
		{
			Name:         "beer",
			NameTemplate: "{{.MainMaterial}} beer",
			Description:  "a beer",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "grain",
					DescriptionTemplate: "bottle of {{.Resource.MainMaterial}} beer",
				},
			},
			Value: 1,
		},
		{
			Name:         "lager",
			NameTemplate: "{{.MainMaterial}} lager",
			Description:  "a lager",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "grain",
					DescriptionTemplate: "bottle of {{.Resource.MainMaterial}} lager",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
