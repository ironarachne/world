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
			Type:        "body armor",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "metal",
					DescriptionTemplate: "{{.Resource.Name}} breastplate",
				},
				{
					Name:                "straps",
					RequiredType:        "hide",
					DescriptionTemplate: " with {{.Resource.Name}} straps",
				},
			},
		},
		{
			Name:        "brigandine",
			Description: "a leather garment lined with small oblong metal plates sewn into the fabric",
			Type:        "body armor",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "hide",
					DescriptionTemplate: "{{.Resource.Name}} brigandine",
				},
				{
					Name:                "plates",
					RequiredType:        "metal",
					DescriptionTemplate: " with {{.Resource.Name}} plates",
				},
			},
		},
		{
			Name:        "chain hauberk",
			Description: "a garment made of chain maille that covers the arms and torso down to the knees",
			Type:        "body armor",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "links",
					RequiredType:        "metal",
					DescriptionTemplate: "{{.Resource.Name}} chain hauberk",
				},
			},
		},
		{
			Name:        "chain shirt",
			Description: "a shirt made of chain maille that covers the arms and torso down to the waist",
			Type:        "body armor",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "links",
					RequiredType:        "metal",
					DescriptionTemplate: "{{.Resource.Name}} chain shirt",
				},
			},
		},
		{
			Name:        "gauntlet",
			Description: "a glove made of metal plates that protects the hand, wrist, and part of the arm",
			Type:        "hand armor",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "metal",
					DescriptionTemplate: "{{.Resource.Name}} gauntlet",
				},
			},
		},
		{
			Name:        "gorget",
			Description: "a metal collar to protect the neck, clavicles, and sternum",
			Type:        "neck armor",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "metal",
					DescriptionTemplate: "{{.Resource.Name}} gorget",
				},
				{
					Name:                "straps",
					RequiredType:        "hide",
					DescriptionTemplate: " with {{.Resource.Name}} straps",
				},
			},
		},
		{
			Name:        "greave",
			Description: "a pair of shaped metal plates that protects the lower leg",
			Type:        "leg armor",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "metal",
					DescriptionTemplate: "{{.Resource.Name}} greave",
				},
			},
		},
		{
			Name:        "pauldron",
			Description: "a shoulder armor that covers the shoulder, armpit, and part of the back and chest",
			Type:        "shoulder armor",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "metal",
					DescriptionTemplate: "{{.Resource.Name}} pauldron",
				},
			},
		},
		{
			Name:        "spaulder",
			Description: "a metal shoulder protector that doesn't cover the armpit",
			Type:        "shoulder armor",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "metal",
					DescriptionTemplate: "{{.Resource.Name}} spaulder",
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
			Type:        "helmet",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "metal",
					DescriptionTemplate: "{{.Resource.Name}} close helm",
				},
				{
					Name:                "trim",
					RequiredType:        "metal",
					DescriptionTemplate: " with {{.Resource.Name}} trim",
				},
				{
					Name:                "padding",
					RequiredType:        "fabric",
					DescriptionTemplate: " and {{.Resource.Name}} padding inside",
				},
			},
		},
		{
			Name:        "nasal helm",
			Description: "a conical helmet with a plate of metal over the nose",
			Type:        "helmet",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "metal",
					DescriptionTemplate: "{{.Resource.Name}} nasal helmet",
				},
				{
					Name:                "trim",
					RequiredType:        "metal",
					DescriptionTemplate: " with {{.Resource.Name}} trim",
				},
				{
					Name:                "padding",
					RequiredType:        "fabric",
					DescriptionTemplate: " and {{.Resource.Name}} padding",
				},
			},
		},
		{
			Name:        "pot helm",
			Description: "a helmet shaped like a cylinder with a rectangular eye slit",
			Type:        "helmet",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "metal",
					DescriptionTemplate: "{{.Resource.Name}} pot helm",
				},
				{
					Name:                "trim",
					RequiredType:        "metal",
					DescriptionTemplate: " with {{.Resource.Name}} trim",
				},
				{
					Name:                "padding",
					RequiredType:        "fabric",
					DescriptionTemplate: " and {{.Resource.Name}} padding",
				},
			},
		},
		{
			Name:        "spangenhelm",
			Description: "a helmet that covers the top of the head and has plates over the ears",
			Type:        "helmet",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredType:        "metal",
					DescriptionTemplate: "{{.Resource.Name}} spangenhelm",
				},
				{
					Name:                "trim",
					RequiredType:        "metal",
					DescriptionTemplate: " with {{.Resource.Name}} trim",
				},
				{
					Name:                "padding",
					RequiredType:        "fabric",
					DescriptionTemplate: " and {{.Resource.Name}} padding",
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
			Name: "buckler",
			Description: "a small round shield strapped to one arm",
			Type: "shield",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "wood",
					DescriptionTemplate: "{{.Resource.Name}} buckler",
				},
				{
					Name: "trim",
					RequiredType: "metal",
					DescriptionTemplate: " with {{.Resource.Name}} trim",
				},
			},
		},
		{
			Name: "heater shield",
			Description: "a shield with a pointed bottom and flat top",
			Type: "shield",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "wood",
					DescriptionTemplate: "{{.Resource.Name}} heater shield",
				},
				{
					Name: "trim",
					RequiredType: "metal",
					DescriptionTemplate: " with {{.Resource.Name}} trim",
				},
			},
		},
		{
			Name: "kite shield",
			Description: "a shield shaped like a kite",
			Type: "shield",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "wood",
					DescriptionTemplate: "{{.Resource.Name}} kite shield",
				},
				{
					Name: "trim",
					RequiredType: "metal",
					DescriptionTemplate: " with {{.Resource.Name}} trim",
				},
			},
		},
		{
			Name: "tower shield",
			Description: "a large, heavy shield that covers the entire body",
			Type: "shield",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "wood",
					DescriptionTemplate: "{{.Resource.Name}} tower shield",
				},
				{
					Name: "trim",
					RequiredType: "metal",
					DescriptionTemplate: " with {{.Resource.Name}} trim",
				},
			},
		},
	}

	return patterns
}
