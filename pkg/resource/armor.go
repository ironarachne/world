package resource

import "github.com/ironarachne/world/pkg/profession"

func getArmor() []Pattern {
	var patterns []Pattern

	bodyArmor := getBodyArmor()
	helmets := getHelmets()
	shields := getShields()
	patterns = append(patterns, bodyArmor...)
	patterns = append(patterns, helmets...)
	patterns = append(patterns, shields...)

	return patterns
}

func getBodyArmor() []Pattern {
	producer := profession.ByName("armorsmith")

	patterns := []Pattern{
		{
			Name:        "breastplate",
			Description: "a solid armor covering for the front of the torso",
			Tags: []string{
				"armor",
				"body armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.Origin}} breastplate",
				},
				{
					Name:                "straps",
					RequiredTag:         "leather",
					DescriptionTemplate: " with {{.Resource.Origin}} leather straps",
				},
			},
		},
		{
			Name:        "brigandine",
			Description: "a leather garment lined with small oblong metal plates sewn into the fabric",
			Tags: []string{
				"armor",
				"body armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.Origin}} leather brigandine",
				},
				{
					Name:                "plates",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with {{.Resource.Origin}} plates",
				},
			},
		},
		{
			Name:        "chain hauberk",
			Description: "a garment made of chain maille that covers the arms and torso down to the knees",
			Tags: []string{
				"armor",
				"body armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "links",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: "{{.Resource.Origin}} chain hauberk",
				},
			},
		},
		{
			Name:        "chain shirt",
			Description: "a shirt made of chain maille that covers the arms and torso down to the waist",
			Tags: []string{
				"armor",
				"body armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "links",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: "{{.Resource.Origin}} chain shirt",
				},
			},
		},
		{
			Name:        "gauntlet",
			Description: "a glove made of metal plates that protects the hand, wrist, and part of the arm",
			Tags: []string{
				"armor",
				"hand armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.Origin}} gauntlet",
				},
			},
		},
		{
			Name:        "gorget",
			Description: "a metal collar to protect the neck, clavicles, and sternum",
			Tags: []string{
				"armor",
				"neck armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.Origin}} gorget",
				},
				{
					Name:                "straps",
					RequiredTag:         "leather",
					DescriptionTemplate: " with {{.Resource.Origin}} leather straps",
				},
			},
		},
		{
			Name:        "greave",
			Description: "a pair of shaped metal plates that protects the lower leg",
			Tags: []string{
				"armor",
				"leg armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.Origin}} greave",
				},
			},
		},
		{
			Name:        "pauldron",
			Description: "a shoulder armor that covers the shoulder, armpit, and part of the back and chest",
			Tags: []string{
				"armor",
				"shoulder armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.Origin}} pauldron",
				},
			},
		},
		{
			Name:        "spaulder",
			Description: "a metal shoulder protector that doesn't cover the armpit",
			Tags: []string{
				"armor",
				"spaulder",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.Origin}} spaulder",
				},
			},
		},
	}

	return patterns
}

func getHelmets() []Pattern {
	producer := profession.ByName("armorsmith")

	patterns := []Pattern{
		{
			Name:        "close helm",
			Description: "a helmet that fully covers the head, with a hinged face plate",
			Tags: []string{
				"armor",
				"helmet",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.Origin}} close helm",
				},
				{
					Name:                "trim",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with {{.Resource.Origin}} trim",
				},
				{
					Name:                "padding",
					RequiredTag:         "fabric",
					DescriptionTemplate: " and {{.Resource.Origin}} padding inside",
				},
			},
		},
		{
			Name:        "nasal helm",
			Description: "a conical helmet with a plate of metal over the nose",
			Tags: []string{
				"armor",
				"helmet",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.Origin}} nasal helmet",
				},
				{
					Name:                "trim",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with {{.Resource.Origin}} trim",
				},
				{
					Name:                "padding",
					RequiredTag:         "fabric",
					DescriptionTemplate: " and {{.Resource.Origin}} padding",
				},
			},
		},
		{
			Name:        "pot helm",
			Description: "a helmet shaped like a cylinder with a rectangular eye slit",
			Tags: []string{
				"armor",
				"helmet",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.Origin}} pot helm",
				},
				{
					Name:                "trim",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with {{.Resource.Origin}} trim",
				},
				{
					Name:                "padding",
					RequiredTag:         "fabric",
					DescriptionTemplate: " and {{.Resource.Origin}} padding",
				},
			},
		},
		{
			Name:        "spangenhelm",
			Description: "a helmet that covers the top of the head and has plates over the ears",
			Tags: []string{
				"armor",
				"helmet",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.Origin}} spangenhelm",
				},
				{
					Name:                "trim",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with {{.Resource.Origin}} trim",
				},
				{
					Name:                "padding",
					RequiredTag:         "fabric",
					DescriptionTemplate: " and {{.Resource.Origin}} padding",
				},
			},
		},
	}

	return patterns
}

func getShields() []Pattern {
	producer := profession.ByName("armorsmith")

	patterns := []Pattern{
		{
			Name:        "buckler",
			Description: "a small round shield strapped to one arm",
			Tags: []string{
				"armor",
				"shield",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.Origin}} buckler",
				},
				{
					Name:                "trim",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with {{.Resource.Origin}} trim",
				},
			},
		},
		{
			Name:        "heater shield",
			Description: "a shield with a pointed bottom and flat top",
			Tags: []string{
				"armor",
				"shield",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.Origin}} heater shield",
				},
				{
					Name:                "trim",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with {{.Resource.Origin}} trim",
				},
			},
		},
		{
			Name:        "kite shield",
			Description: "a shield shaped like a kite",
			Tags: []string{
				"armor",
				"shield",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.Origin}} kite shield",
				},
				{
					Name:                "trim",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with {{.Resource.Origin}} trim",
				},
			},
		},
		{
			Name:        "tower shield",
			Description: "a large, heavy shield that covers the entire body",
			Tags: []string{
				"armor",
				"shield",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.Origin}} tower shield",
				},
				{
					Name:                "trim",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with {{.Resource.Origin}} trim",
				},
			},
		},
	}

	return patterns
}
