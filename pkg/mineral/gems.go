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
					Name:        "diamond",
					Origin:      "diamond",
					Type:        "gem",
					Commonality: 1,
				},
				{
					Name:        "diamond",
					Origin:      "diamond",
					Type:        "ore",
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
					Name:        "garnet",
					Origin:      "garnet",
					Type:        "gem",
					Commonality: 5,
				},
				{
					Name:        "garnet",
					Origin:      "garnet",
					Type:        "ore",
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
					Name:        "ruby",
					Origin:      "ruby",
					Type:        "gem",
					Commonality: 2,
				},
				{
					Name:        "ruby",
					Origin:      "ruby",
					Type:        "ore",
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
					Name:        "sapphire",
					Origin:      "sapphire",
					Type:        "gem",
					Commonality: 4,
				},
				{
					Name:        "sapphire",
					Origin:      "sapphire",
					Type:        "ore",
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
					Name:        "amethyst",
					Origin:      "amethyst",
					Type:        "gem",
					Commonality: 5,
				},
				{
					Name:        "amethyst",
					Origin:      "amethyst",
					Type:        "ore",
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
					Name:        "agate",
					Origin:      "agate",
					Type:        "gem",
					Commonality: 5,
				},
				{
					Name:        "agate",
					Origin:      "agate",
					Type:        "ore",
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
					Name:        "jade",
					Origin:      "jade",
					Type:        "gem",
					Commonality: 5,
				},
				{
					Name:        "jade",
					Origin:      "jade",
					Type:        "ore",
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
					Name:        "jasper",
					Origin:      "jasper",
					Type:        "gem",
					Commonality: 5,
				},
				{
					Name:        "jasper",
					Origin:      "jasper",
					Type:        "ore",
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
					Name:        "moonstone",
					Origin:      "moonstone",
					Type:        "gem",
					Commonality: 5,
				},
				{
					Name:        "moonstone",
					Origin:      "moonstone",
					Type:        "ore",
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
					Name:        "opal",
					Origin:      "opal",
					Type:        "gem",
					Commonality: 5,
				},
				{
					Name:        "opal",
					Origin:      "opal",
					Type:        "ore",
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
					Name:        "peridot",
					Origin:      "peridot",
					Type:        "gem",
					Commonality: 5,
				},
				{
					Name:        "peridot",
					Origin:      "peridot",
					Type:        "ore",
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
					Name:        "topaz",
					Origin:      "topaz",
					Type:        "gem",
					Commonality: 5,
				},
				{
					Name:        "topaz",
					Origin:      "topaz",
					Type:        "ore",
					Commonality: 5,
				},
			},
		},
	}

	return gems
}
