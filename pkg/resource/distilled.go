package resource

import "github.com/ironarachne/world/pkg/profession"

func getDistilled() []Pattern {
	producer, _ := profession.ByName("brewer")

	patterns := []Pattern{
		{
			Name:         "bourbon",
			NameTemplate: "bottle of bourbon",
			Description:  "a bourbon",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "corn",
					DescriptionTemplate: "bottle of bourbon",
				},
			},
			Value: 1,
		},
		{
			Name:         "brandy",
			NameTemplate: "bottle of brandy",
			Description:  "a brandy",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wine",
					DescriptionTemplate: "bottle of brandy",
				},
			},
			Value: 1,
		},
		{
			Name:         "gin",
			NameTemplate: "bottle of gin",
			Description:  "a gin",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "juniper",
					DescriptionTemplate: "bottle of gin",
				},
			},
			Value: 1,
		},
		{
			Name:         "whiskey",
			NameTemplate: "bottle of {{.MainMaterial}} whiskey",
			Description:  "a whiskey",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "rye",
					DescriptionTemplate: "bottle of {{.Resource.MainMaterial}} whiskey",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
