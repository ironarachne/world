package resource

import "github.com/ironarachne/world/pkg/profession"

func getMounts() []Pattern {
	producer := profession.ByName("animal trainer")

	patterns := []Pattern{
		{
			Name:        "mount",
			Description: "a riding mount",
			Tags: []string{
				"trained mount",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "animal",
					RequiredTag:         "mount",
					DescriptionTemplate: "riding {{.Resource.Origin}}",
				},
			},
			Value: 10,
		},
		{
			Name:        "pack animal",
			Description: "a beast of burden",
			Tags: []string{
				"trained pack animal",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "animal",
					RequiredTag:         "pack animal",
					DescriptionTemplate: "pack {{.Resource.Origin}}",
				},
			},
			Value: 5,
		},
		{
			Name:        "war mount",
			Description: "a riding mount specifically trained for battle",
			Tags: []string{
				"trained mount",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "animal",
					RequiredTag:         "mount",
					DescriptionTemplate: "war {{.Resource.Origin}}",
				},
			},
			Value: 100,
		},
	}

	return patterns
}
