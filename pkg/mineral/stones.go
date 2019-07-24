package mineral

import "github.com/ironarachne/world/pkg/resource"

// Stones returns all stones
func Stones() []Mineral {
	stones := []Mineral{
		{
			Name:         "granite",
			PluralName:   "granite",
			Hardness:     6,
			Malleability: 2,
			Resources: []resource.Resource{
				{
					Name:   "granite",
					Origin: "granite",
					Tags: []string{
						"stone",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:         "limestone",
			PluralName:   "limestone",
			Hardness:     4,
			Malleability: 4,
			Resources: []resource.Resource{
				{
					Name:   "limestone",
					Origin: "limestone",
					Tags: []string{
						"stone",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:         "marble",
			PluralName:   "marble",
			Hardness:     5,
			Malleability: 2,
			Resources: []resource.Resource{
				{
					Name:   "marble",
					Origin: "marble",
					Tags: []string{
						"stone",
					},
					Commonality: 2,
				},
			},
		},
		{
			Name:         "sandstone",
			PluralName:   "sandstone",
			Hardness:     4,
			Malleability: 4,
			Resources: []resource.Resource{
				{
					Name:   "sandstone",
					Origin: "sandstone",
					Tags: []string{
						"stone",
					},
					Commonality: 5,
				},
			},
		},
	}

	return stones
}
