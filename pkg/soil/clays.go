package soil

import "github.com/ironarachne/world/pkg/resource"

// Clays returns all clays
func Clays() []Soil {
	soils := []Soil{
		{
			Name:               "earthenware",
			UsedForPottery:     true,
			IsSand:             false,
			UsedForAgriculture: false,
			Resources: []resource.Resource{
				{
					Name:        "earthenware",
					Origin:      "earthenware",
					Type:        "clay",
					Commonality: 5,
				},
			},
		},
		{
			Name:               "porcelain",
			UsedForPottery:     true,
			IsSand:             false,
			UsedForAgriculture: false,
			Resources: []resource.Resource{
				{
					Name:        "porcelin",
					Origin:      "porcelin",
					Type:        "clay",
					Commonality: 5,
				},
			},
		},
		{
			Name:               "stoneware",
			UsedForPottery:     true,
			IsSand:             false,
			UsedForAgriculture: false,
			Resources: []resource.Resource{
				{
					Name:        "stoneware",
					Origin:      "stoneware",
					Type:        "clay",
					Commonality: 5,
				},
			},
		},
		{
			Name:               "terra cotta",
			UsedForPottery:     true,
			IsSand:             false,
			UsedForAgriculture: false,
			Resources: []resource.Resource{
				{
					Name:        "terra cotta",
					Origin:      "terra cotta",
					Type:        "clay",
					Commonality: 5,
				},
			},
		},
	}

	return soils
}
