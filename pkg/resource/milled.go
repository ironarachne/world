package resource

import "github.com/ironarachne/world/pkg/profession"

func getMilled() []Pattern {
	producer, _ := profession.ByName("miller")

	patterns := []Pattern{
		{
			Name:         "flour",
			NameTemplate: "bag of {{.MainMaterial}} flour",
			Description:  "flour",
			Tags: []string{
				"flour",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "grain",
					DescriptionTemplate: "bag of {{.Resource.MainMaterial}} flour",
				},
			},
			Value: 1,
		},
	}

	return patterns
}
