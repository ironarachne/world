package resource

import "github.com/ironarachne/world/pkg/profession"

func getMounts() []Pattern {
	producer := profession.ByName("animal trainer")

	patterns := []Pattern{
		{
			Name:        "mount",
			Description: "a riding mount",
			Type:        "trained mount",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "animal",
					RequiredType:        "mount",
					DescriptionTemplate: "riding {{.Resource.Name}}",
				},
			},
		},
		{
			Name:        "pack animal",
			Description: "a beast of burden",
			Type:        "trained pack animal",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "animal",
					RequiredType:        "pack animal",
					DescriptionTemplate: "pack {{.Resource.Name}}",
				},
			},
		},
		{
			Name:        "war mount",
			Description: "a riding mount specifically trained for battle",
			Type:        "trained mount",
			Profession:  producer,
			Slots: []Slot{
				{
					Name:                "animal",
					RequiredType:        "mount",
					DescriptionTemplate: "war {{.Resource.Name}}",
				},
			},
		},
	}

	return patterns
}
