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
					Name:         "raw diamond ore",
					Origin:       "diamond",
					MainMaterial: "diamond",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 1,
					Value:       100,
				},
			},
			Tags: []string{
				"gem",
			},
		},
		{
			Name:         "garnet",
			PluralName:   "garnets",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:         "raw garnet ore",
					Origin:       "garnet",
					MainMaterial: "garnet",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       10,
				},
			},
			Tags: []string{
				"gem",
			},
		},
		{
			Name:         "ruby",
			PluralName:   "rubies",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:         "raw ruby ore",
					Origin:       "ruby",
					MainMaterial: "ruby",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 2,
					Value:       50,
				},
			},
			Tags: []string{
				"gem",
			},
		},
		{
			Name:         "sapphire",
			PluralName:   "sapphires",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:         "raw sapphire ore",
					Origin:       "sapphire",
					MainMaterial: "sapphire",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 4,
					Value:       40,
				},
			},
			Tags: []string{
				"gem",
			},
		},
		{
			Name:         "amethyst",
			PluralName:   "amethysts",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:         "raw amethyst ore",
					Origin:       "amethyst",
					MainMaterial: "amethyst",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       10,
				},
			},
			Tags: []string{
				"gem",
			},
		},
		{
			Name:         "agate",
			PluralName:   "agates",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:         "raw agate ore",
					Origin:       "agate",
					MainMaterial: "agate",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       5,
				},
			},
			Tags: []string{
				"gem",
			},
		},
		{
			Name:         "jade",
			PluralName:   "jade",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:         "raw jade ore",
					Origin:       "jade",
					MainMaterial: "jade",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       15,
				},
			},
			Tags: []string{
				"gem",
			},
		},
		{
			Name:         "jasper",
			PluralName:   "jaspers",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:         "raw jasper ore",
					Origin:       "jasper",
					MainMaterial: "jasper",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       5,
				},
			},
			Tags: []string{
				"gem",
			},
		},
		{
			Name:         "moonstone",
			PluralName:   "moonstones",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:         "raw moonstone ore",
					Origin:       "moonstone",
					MainMaterial: "moonstone",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       5,
				},
			},
			Tags: []string{
				"gem",
			},
		},
		{
			Name:         "opal",
			PluralName:   "opals",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:         "raw opal ore",
					Origin:       "opal",
					MainMaterial: "opal",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       5,
				},
			},
			Tags: []string{
				"gem",
			},
		},
		{
			Name:         "peridot",
			PluralName:   "peridots",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:         "raw peridot ore",
					Origin:       "peridot",
					MainMaterial: "peridot",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       5,
				},
			},
			Tags: []string{
				"gem",
			},
		},
		{
			Name:         "topaz",
			PluralName:   "topazes",
			Hardness:     6,
			Malleability: 3,
			Resources: []resource.Resource{
				{
					Name:         "raw topaz ore",
					Origin:       "topaz",
					MainMaterial: "topaz",
					Tags: []string{
						"gem ore",
						"ore",
					},
					Commonality: 5,
					Value:       5,
				},
			},
			Tags: []string{
				"gem",
			},
		},
	}

	return gems
}
