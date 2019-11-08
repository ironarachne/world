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
					Name:         "loam",
					Origin:       "loam",
					MainMaterial: "loam",
					Tags: []string{
						"soil",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Tags: []string{
				"soil",
			},
		},
		{
			Name:               "silt",
			UsedForPottery:     false,
			IsSand:             false,
			UsedForAgriculture: true,
			Resources: []resource.Resource{
				{
					Name:         "silt",
					Origin:       "silt",
					MainMaterial: "silt",
					Tags: []string{
						"soil",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Tags: []string{
				"soil",
			},
		},
	}

	return soils
}
