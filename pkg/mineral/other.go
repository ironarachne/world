package mineral

import "github.com/ironarachne/world/pkg/resource"

// OtherMinerals returns all minerals not in other categories
func OtherMinerals() []Mineral {
	others := []Mineral{
		{
			Name:         "coal",
			PluralName:   "coal",
			Hardness:     1,
			Malleability: 4,
			Resources: []resource.Resource{
				{
					Name:        "coal",
					Origin:      "coal",
					Type:        "ore",
					Commonality: 5,
				},
			},
		},
		{
			Name:         "salt",
			PluralName:   "salt",
			Hardness:     1,
			Malleability: 1,
			Resources: []resource.Resource{
				{
					Name:        "salt",
					Origin:      "salt",
					Type:        "spice",
					Commonality: 5,
				},
			},
		},
	}

	return others
}
