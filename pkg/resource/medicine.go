package resource

import "github.com/ironarachne/world/pkg/profession"

func getMedicine() []Pattern {
	producer := profession.ByName("apothecary")

	patterns := []Pattern{
		{
			Name: "healing draught",
			Description: "a healing draught",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "herb",
					DescriptionTemplate: "healing draught",
				},
			},
		},
	}

	return patterns
}