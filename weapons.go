package resource

import "github.com/ironarachne/world/pkg/profession"

func getWeapons() []Pattern {
	producer := profession.ByName("weaponsmith")

	patterns := []Pattern{
		{
			Name:         "bastard sword",
			NameTemplate: "{{.MainMaterial}} bastard sword",
			Description:  "a bastard sword",
			Tags: []string{
				"sword",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "bastard sword with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "hilt",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " {{.Resource.MainMaterial}} hilt,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 10,
		},
		{
			Name:         "battle axe",
			NameTemplate: "{{.MainMaterial}} battle axe",
			Description:  "a battle axe",
			Tags: []string{
				"axe",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "battle axe with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 10,
		},
		{
			Name:         "long sword",
			NameTemplate: "{{.MainMaterial}} long sword",
			Description:  "a long sword",
			Tags: []string{
				"sword",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "long sword with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "hilt",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " {{.Resource.MainMaterial}} hilt,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 10,
		},
		{
			Name:         "mace",
			NameTemplate: "{{.MainMaterial}} mace",
			Description:  "a mace",
			Tags: []string{
				"mace",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "mace with a {{.Resource.MainMaterial}} head,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 6,
		},
		{
			Name:         "morningstar",
			NameTemplate: "{{.MainMaterial}} morningstar",
			Description:  "a morningstar",
			Tags: []string{
				"mace",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "morningstar with a {{.Resource.MainMaterial}} head,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 6,
		},
		{
			Name:         "pike",
			NameTemplate: "{{.MainMaterial}} pike",
			Description:  "a pike",
			Tags: []string{
				"polearm",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "pike with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 10,
		},
		{
			Name:         "short sword",
			NameTemplate: "{{.MainMaterial}} short sword",
			Description:  "a short sword",
			Tags: []string{
				"sword",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "short sword with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "hilt",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " {{.Resource.MainMaterial}} hilt,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 8,
		},
		{
			Name:         "spear",
			NameTemplate: "{{.MainMaterial}} spear",
			Description:  "a spear",
			Tags: []string{
				"spear",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "spear with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 5,
		},
		{
			Name:         "short spear",
			NameTemplate: "{{.MainMaterial}} short spear",
			Description:  "a short spear",
			Tags: []string{
				"spear",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "short {{.Resource.MainMaterial}} spear",
				},
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with a {{.Resource.MainMaterial}} point",
				},
			},
			Value: 5,
		},
		{
			Name:         "two-handed sword",
			NameTemplate: "{{.MainMaterial}} two-handed sword",
			Description:  "a two-handed sword",
			Tags: []string{
				"sword",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "two-handed sword with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "hilt",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " {{.Resource.MainMaterial}} hilt,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 10,
		},
		{
			Name:         "club",
			NameTemplate: "{{.MainMaterial}} club",
			Description:  "a club",
			Tags: []string{
				"mace",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "club with a {{.Resource.MainMaterial}} head,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 0.1,
		},
		{
			Name:         "crescent blade",
			NameTemplate: "{{.MainMaterial}} crescent blade",
			Description:  "a crescent blade",
			Tags: []string{
				"sword",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "crescent blade with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "hilt",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " {{.Resource.MainMaterial}} hilt,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 80,
		},
		{
			Name:         "dagger",
			NameTemplate: "{{.MainMaterial}} dagger",
			Description:  "a dagger",
			Tags: []string{
				"sword",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "dagger with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "hilt",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " {{.Resource.MainMaterial}} hilt,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 2,
		},
		{
			Name:         "flail",
			NameTemplate: "{{.MainMaterial}} flail",
			Description:  "a flail",
			Tags: []string{
				"flail",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "head",
					RequiredTag:         "metal head",
					DescriptionTemplate: "flail with a {{.Resource.MainMaterial}} head,",
				},
				{
					Name:                "chain",
					RequiredTag:         "metal chain",
					DescriptionTemplate: " {{.Resource.MainMaterial}} chain,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 10,
		},
		{
			Name:         "glaive",
			NameTemplate: "{{.MainMaterial}} glaive",
			Description:  "a glaive",
			Tags: []string{
				"polearm",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "glaive with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 20,
		},
		{
			Name:         "greataxe",
			NameTemplate: "{{.MainMaterial}} greataxe",
			Description:  "a greataxe",
			Tags: []string{
				"axe",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "greataxe with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 30,
		},
		{
			Name:         "great bow",
			NameTemplate: "{{.MainMaterial}} great bow",
			Description:  "a great bow",
			Tags: []string{
				"bow",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "limb",
					RequiredTag:         "bow limb",
					DescriptionTemplate: "great bow with {{.Resource.MainMaterial}} limbs,",
				},
				{
					Name:                "string",
					RequiredTag:         "string",
					DescriptionTemplate: "simple {{.Resource.MainMaterial}} string,",
				},
				{
					Name:                "grip",
					RequiredTag:         "grip",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} grip",
				},
			},
			Value: 100,
		},
		{
			Name:         "greatclub",
			NameTemplate: "{{.MainMaterial}} greatclub",
			Description:  "a greatclub",
			Tags: []string{
				"mace",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "greatclub with a {{.Resource.MainMaterial}} head,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 0.2,
		},
		{
			Name:         "greatsword",
			NameTemplate: "{{.MainMaterial}} greatsword",
			Description:  "a greatsword",
			Tags: []string{
				"sword",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "greatsword with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "hilt",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " {{.Resource.MainMaterial}} hilt,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 50,
		},
		{
			Name:         "halberd",
			NameTemplate: "{{.MainMaterial}} halberd",
			Description:  "a halberd",
			Tags: []string{
				"polearm",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "halberd with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 20,
		},
		{
			Name:         "handaxe",
			NameTemplate: "{{.MainMaterial}} handaxe",
			Description:  "a handaxe",
			Tags: []string{
				"axe",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "handaxe with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 5,
		},
		{
			Name:         "hand crossbow",
			NameTemplate: "{{.MainMaterial}} hand crossbow",
			Description:  "a hand crossbow",
			Tags: []string{
				"bow",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "limb",
					RequiredTag:         "bow limb",
					DescriptionTemplate: "hand crossbow with {{.Resource.MainMaterial}} limbs,",
				},
				{
					Name:                "string",
					RequiredTag:         "string",
					DescriptionTemplate: "simple {{.Resource.MainMaterial}} string,",
				},
				{
					Name:                "grip",
					RequiredTag:         "grip",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} grip",
				},
			},
			Value: 75,
		},
		{
			Name:         "heavy crossbow",
			NameTemplate: "{{.MainMaterial}} heavy crossbow",
			Description:  "a heavy crossbow",
			Tags: []string{
				"bow",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "limb",
					RequiredTag:         "bow limb",
					DescriptionTemplate: "heavy crossbow with {{.Resource.MainMaterial}} limbs,",
				},
				{
					Name:                "string",
					RequiredTag:         "string",
					DescriptionTemplate: "simple {{.Resource.MainMaterial}} string,",
				},
				{
					Name:                "grip",
					RequiredTag:         "grip",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} grip",
				},
			},
			Value: 50,
		},
		{
			Name:         "javelin",
			NameTemplate: "{{.MainMaterial}} javelin",
			Description:  "a javelin",
			Tags: []string{
				"spear",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "javelin with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 0.5,
		},
		{
			Name:         "lance",
			NameTemplate: "{{.MainMaterial}} lance",
			Description:  "a lance",
			Tags: []string{
				"spear",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "lance with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 10,
		},
		{
			Name:         "light crossbow",
			NameTemplate: "{{.MainMaterial}} light crossbow",
			Description:  "a light crossbow",
			Tags: []string{
				"bow",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "limb",
					RequiredTag:         "bow limb",
					DescriptionTemplate: "light crossbow with {{.Resource.MainMaterial}} limbs,",
				},
				{
					Name:                "string",
					RequiredTag:         "string",
					DescriptionTemplate: "simple {{.Resource.MainMaterial}} string,",
				},
				{
					Name:                "grip",
					RequiredTag:         "grip",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} grip",
				},
			},
			Value: 25,
		},
		{
			Name:         "light hammer",
			NameTemplate: "{{.MainMaterial}} light hammer",
			Description:  "a light hammer",
			Tags: []string{
				"mace",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "light hammer with a {{.Resource.MainMaterial}} head,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 2,
		},
		{
			Name:         "longbow",
			NameTemplate: "{{.MainMaterial}} longbow",
			Description:  "a longbow",
			Tags: []string{
				"bow",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "limb",
					RequiredTag:         "bow limb",
					DescriptionTemplate: "longbow with {{.Resource.MainMaterial}} limbs,",
				},
				{
					Name:                "string",
					RequiredTag:         "string",
					DescriptionTemplate: "simple {{.Resource.MainMaterial}} string,",
				},
				{
					Name:                "grip",
					RequiredTag:         "grip",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} grip",
				},
			},
			Value: 50,
		},
		{
			Name:         "maul",
			NameTemplate: "{{.MainMaterial}} maul",
			Description:  "a maul",
			Tags: []string{
				"mace",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "maul with a {{.Resource.MainMaterial}} head,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 10,
		},
		{
			Name:         "quarterstaff",
			NameTemplate: "{{.MainMaterial}} quarterstaff",
			Description:  "a quarterstaff",
			Tags: []string{
				"staff",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "quarterstaff with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 0.2,
		},
		{
			Name:         "rapier",
			NameTemplate: "{{.MainMaterial}} rapier",
			Description:  "a rapier",
			Tags: []string{
				"sword",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "rapier with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "hilt",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " {{.Resource.MainMaterial}} hilt,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 25,
		},
		{
			Name:         "scimitar",
			NameTemplate: "{{.MainMaterial}} scimitar",
			Description:  "a scimitar",
			Tags: []string{
				"sword",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "scimitar with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "hilt",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " {{.Resource.MainMaterial}} hilt,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 25,
		},
		{
			Name:         "shortbow",
			NameTemplate: "{{.MainMaterial}} shortbow",
			Description:  "a shortbow",
			Tags: []string{
				"bow",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "limb",
					RequiredTag:         "bow limb",
					DescriptionTemplate: "shortbow with {{.Resource.MainMaterial}} limbs,",
				},
				{
					Name:                "string",
					RequiredTag:         "string",
					DescriptionTemplate: "simple {{.Resource.MainMaterial}} string,",
				},
				{
					Name:                "grip",
					RequiredTag:         "grip",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} grip",
				},
			},
			Value: 25,
		},
		{
			Name:         "sickle",
			NameTemplate: "{{.MainMaterial}} sickle",
			Description:  "a sickle",
			Tags: []string{
				"sword",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "sickle with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "hilt",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " {{.Resource.MainMaterial}} hilt,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 1,
		},
		{
			Name:         "spiked chain",
			NameTemplate: "{{.MainMaterial}} spiked chain",
			Description:  "a spiked chain",
			Tags: []string{
				"flail",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "head",
					RequiredTag:         "metal head",
					DescriptionTemplate: "spiked chain with a {{.Resource.MainMaterial}} head,",
				},
				{
					Name:                "chain",
					RequiredTag:         "metal chain",
					DescriptionTemplate: " {{.Resource.MainMaterial}} chain,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 50,
		},
		{
			Name:         "trident",
			NameTemplate: "{{.MainMaterial}} trident",
			Description:  "a trident",
			Tags: []string{
				"spear",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "trident with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 5,
		},
		{
			Name:         "warhammer",
			NameTemplate: "{{.MainMaterial}} warhammer",
			Description:  "a warhammer",
			Tags: []string{
				"mace",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "warhammer with a {{.Resource.MainMaterial}} head,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 15,
		},
		{
			Name:         "war pick",
			NameTemplate: "{{.MainMaterial}} war pick",
			Description:  "a war pick",
			Tags: []string{
				"polearm",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "war pick with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 5,
		},
		{
			Name:         "war scythe",
			NameTemplate: "{{.MainMaterial}} war scythe",
			Description:  "a war scythe",
			Tags: []string{
				"polearm",
				"weapon",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "blade",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "war scythe with a {{.Resource.MainMaterial}} blade,",
				},
				{
					Name:                "haft",
					RequiredTag:         "wood plank",
					DescriptionTemplate: " {{.Resource.MainMaterial}} haft,",
				},
				{
					Name:                "handle",
					RequiredTag:         "leather",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} leather wrapped handle",
				},
			},
			Value: 25,
		},
	}

	return patterns
}
