package resource

import "github.com/ironarachne/world/pkg/profession"

func getWeapons() []Pattern {
	producer := profession.ByName("weaponsmith")

	patterns := []Pattern{
		{
			Name:        "bastard sword",
			Description: "a bastard sword",
			Tags: []string{
				"sword",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "bastard sword with a {{.Resource.Origin}} blade,",
				},
				{
					Name:                "hilt",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " {{.Resource.Origin}} hilt,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.Origin}} leather wrapped handle",
				},
			},
			Value: 10,
		},
		{
			Name:        "battle axe",
			Description: "a battle axe",
			Tags: []string{
				"axe",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "battle axe with a {{.Resource.Origin}} blade,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.Origin}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.Origin}} leather wrapped handle",
				},
			},
			Value: 10,
		},
		{
			Name:        "long sword",
			Description: "a long sword",
			Tags: []string{
				"sword",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "long sword with a {{.Resource.Origin}} blade,",
				},
				{
					Name:                "hilt",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " {{.Resource.Origin}} hilt,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.Origin}} leather wrapped handle",
				},
			},
			Value: 10,
		},
		{
			Name:        "mace",
			Description: "a mace",
			Tags: []string{
				"mace",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "mace with a {{.Resource.Origin}} head,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.Origin}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.Origin}} leather wrapped handle",
				},
			},
			Value: 6,
		},
		{
			Name:        "morningstar",
			Description: "a morningstar",
			Tags: []string{
				"mace",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "morningstar with a {{.Resource.Origin}} head,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.Origin}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.Origin}} leather wrapped handle",
				},
			},
			Value: 6,
		},
		{
			Name:        "pike",
			Description: "a pike",
			Tags: []string{
				"polearm",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "pike with a {{.Resource.Origin}} blade,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.Origin}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.Origin}} leather wrapped handle",
				},
			},
			Value: 10,
		},
		{
			Name:        "short sword",
			Description: "a short sword",
			Tags: []string{
				"sword",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "short sword with a {{.Resource.Origin}} blade,",
				},
				{
					Name:                "hilt",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " {{.Resource.Origin}} hilt,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.Origin}} leather wrapped handle",
				},
			},
			Value: 8,
		},
		{
			Name:        "spear",
			Description: "a spear",
			Tags: []string{
				"spear",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "spear with a {{.Resource.Origin}} blade,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.Origin}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.Origin}} leather wrapped handle",
				},
			},
			Value: 5,
		},
		{
			Name:        "short spear",
			Description: "a short spear",
			Tags: []string{
				"spear",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "short {{.Resource.Origin}} spear",
				},
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with a {{.Resource.Origin}} point",
				},
			},
			Value: 5,
		},
		{
			Name:        "two-handed sword",
			Description: "a two-handed sword",
			Tags: []string{
				"sword",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "two-handed sword with a {{.Resource.Origin}} blade,",
				},
				{
					Name:                "hilt",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " {{.Resource.Origin}} hilt,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.Origin}} leather wrapped handle",
				},
			},
			Value: 10,
		},
	}

	return patterns
}
