package resource

import "github.com/ironarachne/world/pkg/profession"

func getVehicles() []Pattern {
	producer, _ := profession.ByName("cartwright")

	patterns := []Pattern{
		{
			Name:         "cart wheel",
			NameTemplate: "{{.MainMaterial}} cart wheel",
			Description:  "a cart wheel",
			Tags: []string{
				"cart component",
				"wheel",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood log",
					DescriptionTemplate: "{{.Resource.MainMaterial}} cart wheel",
				},
			},
			Value: 2,
		},
		{
			Name:         "wagon wheel",
			NameTemplate: "{{.MainMaterial}} wagon wheel",
			Description:  "a wagon wheel",
			Tags: []string{
				"wagon component",
				"wheel",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood log",
					DescriptionTemplate: "{{.Resource.MainMaterial}} wagon wheel",
				},
			},
			Value: 2,
		},
		{
			Name:         "cart axle",
			NameTemplate: "{{.MainMaterial}} cart axle",
			Description:  "a cart axle",
			Tags: []string{
				"axle",
				"cart component",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood log",
					DescriptionTemplate: "{{.Resource.MainMaterial}} cart axle",
				},
			},
			Value: 2,
		},
		{
			Name:         "wagon axle",
			NameTemplate: "{{.MainMaterial}} wagon axle",
			Description:  "a wagon axle",
			Tags: []string{
				"axle",
				"wagon component",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood log",
					DescriptionTemplate: "{{.Resource.MainMaterial}} wagon axle",
				},
			},
			Value: 2,
		},
		{
			Name:         "cart",
			NameTemplate: "{{.MainMaterial}} cart",
			Description:  "a cart",
			Tags: []string{
				"land vehicle",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood log",
					DescriptionTemplate: "{{.Resource.MainMaterial}} cart",
				},
				{
					Name:                "wheels",
					RequiredTag:         "cart wheel",
					DescriptionTemplate: "",
				},
				{
					Name:                "axle",
					RequiredTag:         "cart axle",
					DescriptionTemplate: "",
				},
			},
			Value: 10,
		},
		{
			Name:         "wagon",
			NameTemplate: "{{.MainMaterial}} wagon",
			Description:  "a wagon",
			Tags: []string{
				"land vehicle",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "wood log",
					DescriptionTemplate: "{{.Resource.MainMaterial}} wagon",
				},
				{
					Name:                "wheels",
					RequiredTag:         "wagon wheel",
					DescriptionTemplate: "",
				},
				{
					Name:                "axle",
					RequiredTag:         "wagon axle",
					DescriptionTemplate: "",
				},
			},
			Value: 10,
		},
	}

	return patterns
}
