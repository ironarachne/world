package resource

import "github.com/ironarachne/world/pkg/profession"

func getDistilled() []Pattern {
	producer := profession.ByName("brewer")

	patterns := []Pattern{
		{
			Name:        "bourbon",
			Description: "a bourbon",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "corn",
					DescriptionTemplate: "bourbon",
				},
			},
		},
		{
			Name:        "brandy",
			Description: "a brandy",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wine",
					DescriptionTemplate: "brandy",
				},
			},
		},
		{
			Name:        "gin",
			Description: "a gin",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "juniper",
					DescriptionTemplate: "gin",
				},
			},
		},
		{
			Name:        "whiskey",
			Description: "a whiskey",
			Tags: []string{
				"alcohol",
				"beverage",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "rye",
					DescriptionTemplate: "{{.Resource.Origin}} whiskey",
				},
			},
		},
	}

	return patterns
}
