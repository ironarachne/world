package resource

import "github.com/ironarachne/world/pkg/profession"

func getBeekeeping() []Pattern {
	producer := profession.ByName("beekeeper")

	patterns := []Pattern{
		{
			Name:         "braggot",
			NameTemplate: "bottle of braggot",
			Description:  "bottle of braggot",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "honey",
					DescriptionTemplate: "bottle of braggot",
				},
			},
			Value: 1,
		},
		{
			Name:         "hippocras",
			NameTemplate: "bottle of hippocras",
			Description:  "bottle of hippocras",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "honey",
					DescriptionTemplate: "bottle of hippocras",
				},
			},
			Value: 1,
		},
		{
			Name:         "mead",
			NameTemplate: "bottle of mead",
			Description:  "bottle of mead",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "honey",
					DescriptionTemplate: "bottle of mead",
				},
			},
			Value: 1,
		},
		{
			Name:         "metheglin",
			NameTemplate: "bottle of metheglin",
			Description:  "bottle of metheglin",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "honey",
					DescriptionTemplate: "bottle of metheglin",
				},
			},
			Value: 1,
		},
		{
			Name:         "sealing wax",
			NameTemplate: "stick of sealing wax",
			Description:  "stick of sealing wax",
			Tags: []string{
				"writing materials",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wax",
					DescriptionTemplate: "stick of sealing wax",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
