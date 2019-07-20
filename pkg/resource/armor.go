package resource

import "github.com/ironarachne/world/pkg/profession"

func getArmor() []Pattern {
	var patterns []Pattern

	bodyArmor := getBodyArmor()
	helmets := getHelmets()
	patterns = append(patterns, bodyArmor...)
	patterns = append(patterns, helmets...)

	return patterns
}

func getBodyArmor() []Pattern {
	producer := profession.ByName("armorsmith")

	patterns := []Pattern{
		{
			Name: "breastplate",
			Description: "a solid armor covering for the front of the torso",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "metal",
					DescriptionTemplate: "{{.Resource.Name}} breastplate",
				},
				{
					Name: "straps",
					RequiredType: "hide",
					DescriptionTemplate: " with {{.Resource.Name}} straps",
				},
			},
		},
		{
			Name: "chain hauberk",
			Description: "a garment made of chain maille that covers the arms and torso down to the knees",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "links",
					RequiredType: "metal",
					DescriptionTemplate: "{{.Resource.Name}} chain hauberk",
				},
			},
		},
		{
			Name: "chain shirt",
			Description: "a shirt made of chain maille that covers the arms and torso down to the waist",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "links",
					RequiredType: "metal",
					DescriptionTemplate: "{{.Resource.Name}} chain shirt",
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