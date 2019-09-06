package resource

import "github.com/ironarachne/world/pkg/profession"

func getWine() []Pattern {
	producer := profession.ByName("vintner")

	patterns := []Pattern{
		{
			Name:         "barley wine",
			NameTemplate: "bottle of {{.MainMaterial}} barley wine",
			Description:  "barley wine",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "barley",
					DescriptionTemplate: "bottle of {{.Resource.MainMaterial}} barley wine",
				},
			},
			Value: 1,
		},
		{
			Name:         "rice wine",
			NameTemplate: "bottle of rice wine",
			Description:  "rice wine",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "rice",
					DescriptionTemplate: "bottle of rice wine",
				},
			},
			Value: 1,
		},
		{
			Name:         "wine",
			NameTemplate: "bottle of {{.MainMaterial}} wine",
			Description:  "wine",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "fruit",
					DescriptionTemplate: "bottle of {{.Resource.MainMaterial}} wine",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
