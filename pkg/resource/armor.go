package resource

import (
	"github.com/ironarachne/world/pkg/profession"
)

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
	producer, _ := profession.ByName("armorsmith")
	patterns := []Pattern{
		{
			Name:         "breastplate",
			NameTemplate: "{{.MainMaterial}} breastplate",
			Description:  "a solid armor covering for the front of the torso",
			Tags: []string{
				"armor",
				"body armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} breastplate",
				},
				{
					Name:                "straps",
					RequiredTag:         "leather",
					DescriptionTemplate: " with {{.Resource.MainMaterial}} leather straps",
				},
			},
			Value: 10,
		},
		{
			Name:         "brigandine",
			NameTemplate: "{{.MainMaterial}} leather brigandine",
			Description:  "a leather garment lined with small oblong metal plates sewn into the fabric",
			Tags: []string{
				"armor",
				"body armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "leather",
					DescriptionTemplate: "{{.Resource.MainMaterial}} leather brigandine",
				},
				{
					Name:                "plates",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with {{.Resource.MainMaterial}} plates",
				},
			},
			Value: 6,
		},
		{
			Name:         "chain hauberk",
			NameTemplate: "{{.MainMaterial}} chain hauberk",
			Description:  "a garment made of chain maille that covers the arms and torso down to the knees",
			Tags: []string{
				"armor",
				"body armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "links",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} chain hauberk",
				},
			},
			Value: 8,
		},
		{
			Name:         "chain shirt",
			NameTemplate: "{{.MainMaterial}} chain shirt",
			Description:  "a shirt made of chain maille that covers the arms and torso down to the waist",
			Tags: []string{
				"armor",
				"body armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "links",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} chain shirt",
				},
			},
			Value: 6,
		},
		{
			Name:         "gauntlet",
			NameTemplate: "{{.MainMaterial}} gauntlet",
			Description:  "a glove made of metal plates that protects the hand, wrist, and part of the arm",
			Tags: []string{
				"armor",
				"hand armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} gauntlet",
				},
			},
			Value: 5,
		},
		{
			Name:         "gorget",
			NameTemplate: "{{.MainMaterial}} gorget",
			Description:  "a metal collar to protect the neck, clavicles, and sternum",
			Tags: []string{
				"armor",
				"neck armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} gorget",
				},
				{
					Name:                "straps",
					RequiredTag:         "leather",
					DescriptionTemplate: " with {{.Resource.MainMaterial}} leather straps",
				},
			},
			Value: 4,
		},
		{
			Name:         "greave",
			NameTemplate: "{{.MainMaterial}} greaves",
			Description:  "a pair of shaped metal plates that protects the lower leg",
			Tags: []string{
				"armor",
				"leg armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} greaves",
				},
			},
			Value: 5,
		},
		{
			Name:         "pauldron",
			NameTemplate: "{{.MainMaterial}} pauldrons",
			Description:  "a shoulder armor that covers the shoulder, armpit, and part of the back and chest",
			Tags: []string{
				"armor",
				"shoulder armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} pauldrons",
				},
			},
			Value: 5,
		},
		{
			Name:         "spaulder",
			NameTemplate: "{{.MainMaterial}} spaulder",
			Description:  "a metal shoulder protector that doesn't cover the armpit",
			Tags: []string{
				"armor",
				"shoulder armor",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} spaulder",
				},
			},
			Value: 5,
		},
	}

	return patterns
}

func getHelmets() []Pattern {
	producer, _ := profession.ByName("armorsmith")

	patterns := []Pattern{
		{
			Name:         "close helm",
			NameTemplate: "{{.MainMaterial}} close helm",
			Description:  "a helmet that fully covers the head, with a hinged face plate",
			Tags: []string{
				"armor",
				"helmet",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} close helm",
				},
				{
					Name:                "trim",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " with {{.Resource.MainMaterial}} trim",
				},
				{
					Name:                "padding",
					RequiredTag:         "fabric",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} padding inside",
				},
			},
			Value: 6,
		},
		{
			Name:         "nasal helm",
			NameTemplate: "{{.MainMaterial}} nasal helmet",
			Description:  "a conical helmet with a plate of metal over the nose",
			Tags: []string{
				"armor",
				"helmet",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} nasal helmet",
				},
				{
					Name:                "trim",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " with {{.Resource.MainMaterial}} trim",
				},
				{
					Name:                "padding",
					RequiredTag:         "fabric",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} padding",
				},
			},
			Value: 6,
		},
		{
			Name:         "pot helm",
			NameTemplate: "{{.MainMaterial}} pot helm",
			Description:  "a helmet shaped like a cylinder with a rectangular eye slit",
			Tags: []string{
				"armor",
				"helmet",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} pot helm",
				},
				{
					Name:                "trim",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " with {{.Resource.MainMaterial}} trim",
				},
				{
					Name:                "padding",
					RequiredTag:         "fabric",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} padding",
				},
			},
			Value: 5,
		},
		{
			Name:         "spangenhelm",
			NameTemplate: "{{.MainMaterial}} spangenhelm",
			Description:  "a helmet that covers the top of the head and has plates over the ears",
			Tags: []string{
				"armor",
				"helmet",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} spangenhelm",
				},
				{
					Name:                "trim",
					RequiredTag:         "metal ingot",
					DescriptionTemplate: " with {{.Resource.MainMaterial}} trim",
				},
				{
					Name:                "padding",
					RequiredTag:         "fabric",
					DescriptionTemplate: " and {{.Resource.MainMaterial}} padding",
				},
			},
			Value: 5,
		},
	}

	return patterns
}

func getShields() []Pattern {
	producer, _ := profession.ByName("armorsmith")

	patterns := []Pattern{
		{
			Name:         "buckler",
			NameTemplate: "{{.MainMaterial}} buckler",
			Description:  "a small round shield strapped to one arm",
			Tags: []string{
				"armor",
				"shield",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.MainMaterial}} buckler",
				},
				{
					Name:                "trim",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with {{.Resource.MainMaterial}} trim",
				},
			},
			Value: 5,
		},
		{
			Name:         "heater shield",
			NameTemplate: "{{.MainMaterial}} heater shield",
			Description:  "a shield with a pointed bottom and flat top",
			Tags: []string{
				"armor",
				"shield",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.MainMaterial}} heater shield",
				},
				{
					Name:                "trim",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with {{.Resource.MainMaterial}} trim",
				},
			},
			Value: 5,
		},
		{
			Name:         "kite shield",
			NameTemplate: "{{.MainMaterial}} kite shield",
			Description:  "a shield shaped like a kite",
			Tags: []string{
				"armor",
				"shield",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.MainMaterial}} kite shield",
				},
				{
					Name:                "trim",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with {{.Resource.MainMaterial}} trim",
				},
			},
			Value: 5,
		},
		{
			Name:         "tower shield",
			NameTemplate: "{{.MainMaterial}} tower shield",
			Description:  "a large, heavy shield that covers the entire body",
			Tags: []string{
				"armor",
				"shield",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood plank",
					DescriptionTemplate: "{{.Resource.MainMaterial}} tower shield",
				},
				{
					Name:                "trim",
					RequiredTag:         "metal bar",
					DescriptionTemplate: " with {{.Resource.MainMaterial}} trim",
				},
			},
			Value: 10,
		},
	}

	return patterns
}
