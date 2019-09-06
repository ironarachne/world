package mineral

import "github.com/ironarachne/world/pkg/resource"

// Stones returns all stones
func Stones() []Mineral {
	stones := []Mineral{
		{
			Name:         "basalt",
			PluralName:   "basalt",
			Hardness:     5,
			Malleability: 5,
			Resources: []resource.Resource{
				{
					Name:         "basalt",
					Origin:       "basalt",
					MainMaterial: "basalt",
					Tags: []string{
						"stone",
						"building stone",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:         "chalk",
			PluralName:   "chalk",
			Hardness:     1,
			Malleability: 4,
			Resources: []resource.Resource{
				{
					Name:         "chalk",
					Origin:       "chalk",
					MainMaterial: "chalk",
					Tags: []string{
						"stone",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:         "granite",
			PluralName:   "granite",
			Hardness:     6,
			Malleability: 2,
			Resources: []resource.Resource{
				{
					Name:         "granite",
					Origin:       "granite",
					MainMaterial: "granite",
					Tags: []string{
						"stone",
						"building stone",
					},
					Commonality: 5,
					Value:       3,
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
					Name:         "limestone",
					Origin:       "limestone",
					MainMaterial: "limestone",
					Tags: []string{
						"stone",
						"building stone",
					},
					Commonality: 10,
					Value:       1,
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
					Name:         "marble",
					Origin:       "marble",
					MainMaterial: "marble",
					Tags: []string{
						"stone",
						"building stone",
						"sculpture stone",
					},
					Commonality: 2,
					Value:       10,
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
					Name:         "sandstone",
					Origin:       "sandstone",
					MainMaterial: "sandstone",
					Tags: []string{
						"stone",
						"building stone",
						"sculpture stone",
					},
					Commonality: 10,
					Value:       1,
				},
			},
		},
		{
			Name:         "shale",
			PluralName:   "shale",
			Hardness:     1,
			Malleability: 4,
			Resources: []resource.Resource{
				{
					Name:         "shale",
					Origin:       "shale",
					MainMaterial: "shale",
					Tags: []string{
						"stone",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:         "slate",
			PluralName:   "slate",
			Hardness:     4,
			Malleability: 4,
			Resources: []resource.Resource{
				{
					Name:         "slate",
					Origin:       "slate",
					MainMaterial: "slate",
					Tags: []string{
						"stone",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:         "soapstone",
			PluralName:   "soapstone",
			Hardness:     1,
			Malleability: 4,
			Resources: []resource.Resource{
				{
					Name:         "soapstone",
					Origin:       "soapstone",
					MainMaterial: "soapstone",
					Tags: []string{
						"stone",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
	}

	return stones
}
