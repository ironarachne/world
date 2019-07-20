package soil

import "github.com/ironarachne/world/pkg/resource"

// Agricultural returns all agricultural soils
func Agricultural() []Soil {
	soils := []Soil{
		{
			Name:               "loam",
			UsedForPottery:     false,
			IsSand:             false,
			UsedForAgriculture: true,
			Resources: []resource.Resource{
				{
					Name:        "loam",
					Origin:      "loam",
					Type:        "soil",
					Commonality: 5,
				},
			},
		},
		{
			Name:               "silt",
			UsedForPottery:     false,
			IsSand:             false,
			UsedForAgriculture: true,
			Resources: []resource.Resource{
				{
					Name:        "loam",
					Origin:      "loam",
					Type:        "soil",
					Commonality: 5,
				},
			},
		},
	}

	return soils
}
