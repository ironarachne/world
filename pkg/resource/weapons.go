package resource

import "github.com/ironarachne/world/pkg/profession"

func getWeapons() []Pattern {
	producer := profession.ByName("weaponsmith")

	patterns := []Pattern{
		{
			Name:        "bastard sword",
			Description: "a bastard sword",
			Type:        "weapon",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredType:        "metal",
					DescriptionTemplate: "bastard sword with a {{.Resource.Name}} blade,",
				},
				{
					Name:                "hilt",
					RequiredType:        "metal",
					DescriptionTemplate: " {{.Resource.Name}} hilt,",
				},
				{
					Name:                "handle",
					RequiredType:        "hide",
					DescriptionTemplate: " and {{.Resource.Name}} wrapped handle",
				},
			},
		},
		{
			Name:        "battle axe",
			Description: "a battle axe",
			Type:        "weapon",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredType:        "metal",
					DescriptionTemplate: "battle axe with a {{.Resource.Name}} blade,",
				},
				{
					Name:                "haft",
					RequiredType:        "wood",
					DescriptionTemplate: " {{.Resource.Name}} haft,",
				},
				{
					Name:                "handle",
					RequiredType:        "hide",
					DescriptionTemplate: " and {{.Resource.Name}} wrapped handle",
				},
			},
		},
		{
			Name:        "long sword",
			Description: "a long sword",
			Type:        "weapon",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredType:        "metal",
					DescriptionTemplate: "long sword with a {{.Resource.Name}} blade,",
				},
				{
					Name:                "hilt",
					RequiredType:        "metal",
					DescriptionTemplate: " {{.Resource.Name}} hilt,",
				},
				{
					Name:                "handle",
					RequiredType:        "hide",
					DescriptionTemplate: " and {{.Resource.Name}} wrapped handle",
				},
			},
		},
		{
			Name:        "mace",
			Description: "a mace",
			Type:        "weapon",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredType:        "metal",
					DescriptionTemplate: "mace with a {{.Resource.Name}} head,",
				},
				{
					Name:                "haft",
					RequiredType:        "wood",
					DescriptionTemplate: " {{.Resource.Name}} haft,",
				},
				{
					Name:                "handle",
					RequiredType:        "hide",
					DescriptionTemplate: " and {{.Resource.Name}} wrapped handle",
				},
			},
		},
		{
			Name:        "morningstar",
			Description: "a morningstar",
			Type:        "weapon",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredType:        "metal",
					DescriptionTemplate: "morningstar with a {{.Resource.Name}} head,",
				},
				{
					Name:                "haft",
					RequiredType:        "wood",
					DescriptionTemplate: " {{.Resource.Name}} haft,",
				},
				{
					Name:                "handle",
					RequiredType:        "hide",
					DescriptionTemplate: " and {{.Resource.Name}} wrapped handle",
				},
			},
		},
		{
			Name:        "pike",
			Description: "a pike",
			Type:        "weapon",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredType:        "metal",
					DescriptionTemplate: "pike with a {{.Resource.Name}} blade,",
				},
				{
					Name:                "haft",
					RequiredType:        "wood",
					DescriptionTemplate: " {{.Resource.Name}} haft,",
				},
				{
					Name:                "handle",
					RequiredType:        "hide",
					DescriptionTemplate: " and {{.Resource.Name}} wrapped handle",
				},
			},
		},
		{
			Name:        "short sword",
			Description: "a short sword",
			Type:        "weapon",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredType:        "metal",
					DescriptionTemplate: "short sword with a {{.Resource.Name}} blade,",
				},
				{
					Name:                "hilt",
					RequiredType:        "metal",
					DescriptionTemplate: " {{.Resource.Name}} hilt,",
				},
				{
					Name:                "handle",
					RequiredType:        "hide",
					DescriptionTemplate: " and {{.Resource.Name}} wrapped handle",
				},
			},
		},
		{
			Name:        "spear",
			Description: "a spear",
			Type:        "weapon",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredType:        "metal",
					DescriptionTemplate: "spear with a {{.Resource.Name}} blade,",
				},
				{
					Name:                "haft",
					RequiredType:        "wood",
					DescriptionTemplate: " {{.Resource.Name}} haft,",
				},
				{
					Name:                "handle",
					RequiredType:        "hide",
					DescriptionTemplate: " and {{.Resource.Name}} wrapped handle",
				},
			},
		},
		{
			Name:        "two-handed sword",
			Description: "a two-handed sword",
			Type:        "weapon",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredType:        "metal",
					DescriptionTemplate: "two-handed sword with a {{.Resource.Name}} blade,",
				},
				{
					Name:                "hilt",
					RequiredType:        "metal",
					DescriptionTemplate: " {{.Resource.Name}} hilt,",
				},
				{
					Name:                "handle",
					RequiredType:        "hide",
					DescriptionTemplate: " and {{.Resource.Name}} wrapped handle",
				},
			},
		},
	}

	return patterns
}
