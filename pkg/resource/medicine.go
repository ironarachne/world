package resource

import "github.com/ironarachne/world/pkg/profession"

func getMedicine() []Pattern {
	producer := profession.ByName("apothecary")

	patterns := []Pattern{
		{
			Name:         "healing draught",
			NameTemplate: "healing draught",
			Description:  "a healing draught",
			Tags: []string{
				"medicine",
				"potion",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "herb",
					DescriptionTemplate: "healing draught",
				},
			},
			Value: 5,
		},
	}

	return patterns
}
