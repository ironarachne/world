package resource

import "github.com/ironarachne/world/pkg/profession"

func getMined() []Pattern {
	producer := profession.ByName("miner")

	patterns := []Pattern{
		{
			Name: "refined ore",
			Description: "refined ore",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "ore",
					DescriptionTemplate: "refined {{.Resource.Name}} ore",
				},
			},
		},
	}

	return patterns
}
