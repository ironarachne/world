package mineral

import "github.com/ironarachne/world/pkg/resource"

// Metals returns all metals
func Metals() []Mineral {
	metals := []Mineral{
		{
			Name:         "copper",
			PluralName:   "copper",
			Hardness:     5,
			Malleability: 5,
			Resources: []resource.Resource{
				{
					Name:   "copper",
					Origin: "copper",
					Tags: []string{
						"metal ore",
						"ore",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:         "gold",
			PluralName:   "gold",
			Hardness:     2,
			Malleability: 8,
			Resources: []resource.Resource{
				{
					Name:   "gold",
					Origin: "gold",
					Tags: []string{
						"metal ore",
						"ore",
					},
					Commonality: 1,
				},
			},
		},
		{
			Name:         "iron",
			PluralName:   "iron",
			Hardness:     8,
			Malleability: 4,
			Resources: []resource.Resource{
				{
					Name:   "iron",
					Origin: "iron",
					Tags: []string{
						"metal ore",
						"ore",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:         "tin",
			PluralName:   "tin",
			Hardness:     3,
			Malleability: 8,
			Resources: []resource.Resource{
				{
					Name:   "tin",
					Origin: "tin",
					Tags: []string{
						"metal ore",
						"ore",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:         "lead",
			PluralName:   "lead",
			Hardness:     1,
			Malleability: 9,
			Resources: []resource.Resource{
				{
					Name:   "lead",
					Origin: "lead",
					Tags: []string{
						"metal ore",
						"ore",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:         "nickel",
			PluralName:   "nickel",
			Hardness:     3,
			Malleability: 4,
			Resources: []resource.Resource{
				{
					Name:   "nickel",
					Origin: "nickel",
					Tags: []string{
						"metal ore",
						"ore",
					},
					Commonality: 5,
				},
			},
		},
		{
			Name:         "silver",
			PluralName:   "silver",
			Hardness:     2,
			Malleability: 7,
			Resources: []resource.Resource{
				{
					Name:   "silver",
					Origin: "silver",
					Tags: []string{
						"metal ore",
						"ore",
					},
					Commonality: 3,
				},
			},
		},
	}

	return metals
}
