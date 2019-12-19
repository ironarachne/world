package resource

import "github.com/ironarachne/world/pkg/profession"

func getMounts() []Pattern {
	producer, _ := profession.ByName("animal trainer")

	patterns := []Pattern{
		{
			Name:         "mount",
			NameTemplate: "riding {{.MainMaterial}}",
			Description:  "a riding mount",
			Tags: []string{
				"trained mount",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "animal",
					RequiredTag:         "mount",
					DescriptionTemplate: "riding {{.Resource.MainMaterial}}",
				},
			},
			Value: 10,
		},
		{
			Name:         "pack animal",
			NameTemplate: "pack {{.MainMaterial}}",
			Description:  "a beast of burden",
			Tags: []string{
				"trained pack animal",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "animal",
					RequiredTag:         "pack animal",
					DescriptionTemplate: "pack {{.Resource.MainMaterial}}",
				},
			},
			Value: 5,
		},
		{
			Name:         "war mount",
			NameTemplate: "war {{.MainMaterial}}",
			Description:  "a riding mount specifically trained for battle",
			Tags: []string{
				"trained mount",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "animal",
					RequiredTag:         "mount",
					DescriptionTemplate: "war {{.Resource.MainMaterial}}",
				},
			},
			Value: 100,
		},
	}

	return patterns
}
