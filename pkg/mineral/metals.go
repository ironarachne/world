package mineral

import "github.com/ironarachne/world/pkg/resource"

// Metals returns all metals
func Metals() []Mineral {
	metals := []Mineral{
		Mineral{
			Name:         "copper",
			PluralName:   "copper",
			Hardness:     5,
			Malleability: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "copper",
					Origin:      "copper",
					Type:        "metal",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "copper",
					Origin:      "copper",
					Type:        "ore",
					Commonality: 5,
				},
			},
		},
		Mineral{
			Name:         "gold",
			PluralName:   "gold",
			Hardness:     2,
			Malleability: 8,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "gold",
					Origin:      "gold",
					Type:        "metal",
					Commonality: 1,
				},
				resource.Resource{
					Name:        "gold",
					Origin:      "gold",
					Type:        "ore",
					Commonality: 1,
				},
			},
		},
		Mineral{
			Name:         "iron",
			PluralName:   "iron",
			Hardness:     8,
			Malleability: 4,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "iron",
					Origin:      "iron",
					Type:        "metal",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "iron",
					Origin:      "iron",
					Type:        "ore",
					Commonality: 5,
				},
			},
		},
		Mineral{
			Name:         "tin",
			PluralName:   "tin",
			Hardness:     3,
			Malleability: 8,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "tin",
					Origin:      "tin",
					Type:        "metal",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "tin",
					Origin:      "tin",
					Type:        "ore",
					Commonality: 5,
				},
			},
		},
		Mineral{
			Name:         "lead",
			PluralName:   "lead",
			Hardness:     1,
			Malleability: 9,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "lead",
					Origin:      "lead",
					Type:        "metal",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "lead",
					Origin:      "lead",
					Type:        "ore",
					Commonality: 5,
				},
			},
		},
		Mineral{
			Name:         "nickel",
			PluralName:   "nickel",
			Hardness:     3,
			Malleability: 4,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "nickel",
					Origin:      "nickel",
					Type:        "metal",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "nickel",
					Origin:      "nickel",
					Type:        "ore",
					Commonality: 5,
				},
			},
		},
		Mineral{
			Name:         "silver",
			PluralName:   "silver",
			Hardness:     2,
			Malleability: 7,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "silver",
					Origin:      "silver",
					Type:        "metal",
					Commonality: 3,
				},
				resource.Resource{
					Name:        "silver",
					Origin:      "silver",
					Type:        "ore",
					Commonality: 3,
				},
			},
		},
	}

	return metals
}
