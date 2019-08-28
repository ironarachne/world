package mineral

import "github.com/ironarachne/world/pkg/resource"

// Gems returns all gems
func Gems() []Mineral {
	gems := []Mineral{
		{
			Name:         "diamond",
			PluralName:   "diamonds",
			Hardness:     10,
			Malleability: 1,
			Resources: []resource.Resource{
				{
					Name:   "raw diamond ore",
					Origin: "diamond",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 1,
					Value:       100,
				},
			},
		},
		{
			Name:         "garnet",
			PluralName:   "garnets",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:   "raw garnet ore",
					Origin: "garnet",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       10,
				},
			},
		},
		{
			Name:         "ruby",
			PluralName:   "rubies",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:   "raw ruby ore",
					Origin: "ruby",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 2,
					Value:       50,
				},
			},
		},
		{
			Name:         "sapphire",
			PluralName:   "sapphires",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:   "raw sapphire ore",
					Origin: "sapphire",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 4,
					Value:       40,
				},
			},
		},
		{
			Name:         "amethyst",
			PluralName:   "amethysts",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:   "raw amethyst ore",
					Origin: "amethyst",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       10,
				},
			},
		},
		{
			Name:         "agate",
			PluralName:   "agates",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:   "raw agate ore",
					Origin: "agate",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       5,
				},
			},
		},
		{
			Name:         "jade",
			PluralName:   "jade",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:   "raw jade ore",
					Origin: "jade",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       15,
				},
			},
		},
		{
			Name:         "jasper",
			PluralName:   "jaspers",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:   "raw jasper ore",
					Origin: "jasper",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       5,
				},
			},
		},
		{
			Name:         "moonstone",
			PluralName:   "moonstones",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:   "raw moonstone ore",
					Origin: "moonstone",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       5,
				},
			},
		},
		{
			Name:         "opal",
			PluralName:   "opals",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:   "raw opal ore",
					Origin: "opal",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       5,
				},
			},
		},
		{
			Name:         "peridot",
			PluralName:   "peridots",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:   "raw peridot ore",
					Origin: "peridot",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       5,
				},
			},
		},
		{
			Name:         "topaz",
			PluralName:   "topazes",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:   "raw topaz ore",
					Origin: "topaz",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       5,
				},
			},
		},
	}

	return gems
}
