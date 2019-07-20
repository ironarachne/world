package resource

import "github.com/ironarachne/world/pkg/profession"

func getWeapons() []Pattern {
	producer := profession.ByName("weaponsmith")

	patterns := []Pattern{
		{
			Name: "battle axe",
			Description: "a battle axe",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "blade",
					RequiredType: "metal",
					DescriptionTemplate: "battle axe with a {{.Resource.Name}} blade,",
				},
				{
					Name: "haft",
					RequiredType: "wood",
					DescriptionTemplate: " {{.Resource.Name}} haft,",
				},
				{
					Name: "handle",
					RequiredType: "hide",
					DescriptionTemplate: " and {{.Resource.Name}} wrapped handle",
				},
			},
		},
		{
			Name: "long sword",
			Description: "a long sword",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "blade",
					RequiredType: "metal",
					DescriptionTemplate: "long sword with a {{.Resource.Name}} blade,",
				},
				{
					Name: "hilt",
					RequiredType: "metal",
					DescriptionTemplate: " {{.Resource.Name}} hilt,",
				},
				{
					Name: "handle",
					RequiredType: "hide",
					DescriptionTemplate: " and {{.Resource.Name}} wrapped handle",
				},
			},
		},
		{
			Name: "short sword",
			Description: "a short sword",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "blade",
					RequiredType: "metal",
					DescriptionTemplate: "short sword with a {{.Resource.Name}} blade,",
				},
				{
					Name: "hilt",
					RequiredType: "metal",
					DescriptionTemplate: " {{.Resource.Name}} hilt,",
				},
				{
					Name: "handle",
					RequiredType: "hide",
					DescriptionTemplate: " and {{.Resource.Name}} wrapped handle",
				},
			},
		},
		{
			Name: "spear",
			Description: "a spear",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "blade",
					RequiredType: "metal",
					DescriptionTemplate: "spear with a {{.Resource.Name}} blade,",
				},
				{
					Name: "haft",
					RequiredType: "wood",
					DescriptionTemplate: " {{.Resource.Name}} haft,",
				},
				{
					Name: "handle",
					RequiredType: "hide",
					DescriptionTemplate: " and {{.Resource.Name}} wrapped handle",
				},
			},
		},
	}

	return patterns
}