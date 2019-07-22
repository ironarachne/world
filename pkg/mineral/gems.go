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
					Name:   "diamond",
					Origin: "diamond",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 1,
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
					Name:   "garnet",
					Origin: "garnet",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
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
					Name:   "ruby",
					Origin: "ruby",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 2,
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
					Name:   "sapphire",
					Origin: "sapphire",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 4,
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
					Name:   "amethyst",
					Origin: "amethyst",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
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
					Name:   "agate",
					Origin: "agate",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
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
					Name:   "jade",
					Origin: "jade",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
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
					Name:   "jasper",
					Origin: "jasper",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
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
					Name:   "moonstone",
					Origin: "moonstone",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
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
					Name:   "opal",
					Origin: "opal",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
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
					Name:   "peridot",
					Origin: "peridot",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
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
					Name:   "topaz",
					Origin: "topaz",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
				},
			},
		},
	}

	return gems
}
