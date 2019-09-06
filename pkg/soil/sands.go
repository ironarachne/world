package soil

import "github.com/ironarachne/world/pkg/resource"

// Sands returns all sands
func Sands() []Soil {
	soils := []Soil{
		{
			Name:               "black sand",
			UsedForPottery:     false,
			IsSand:             true,
			UsedForAgriculture: false,
			Resources: []resource.Resource{
				{
					Name:         "black sand",
					Origin:       "black sand",
					MainMaterial: "black sand",
					Tags: []string{
						"sand",
					},
					Commonality: 3,
					Value:       1,
				},
			},
		},
		{
			Name:               "continental sand",
			UsedForPottery:     false,
			IsSand:             true,
			UsedForAgriculture: false,
			Resources: []resource.Resource{
				{
					Name:         "continental sand",
					Origin:       "continental sand",
					MainMaterial: "continental sand",
					Tags: []string{
						"sand",
					},
					Commonality: 4,
					Value:       1,
				},
			},
		},
		{
			Name:               "coral sand",
			UsedForPottery:     false,
			IsSand:             true,
			UsedForAgriculture: false,
			Resources: []resource.Resource{
				{
					Name:         "coral sand",
					Origin:       "coral sand",
					MainMaterial: "coral sand",
					Tags: []string{
						"sand",
					},
					Commonality: 1,
					Value:       1,
				},
			},
		},
		{
			Name:               "dune sand",
			UsedForPottery:     false,
			IsSand:             true,
			UsedForAgriculture: false,
			Resources: []resource.Resource{
				{
					Name:         "dune sand",
					Origin:       "dune sand",
					MainMaterial: "dune sand",
					Tags: []string{
						"sand",
					},
					Commonality: 4,
					Value:       1,
				},
			},
		},
		{
			Name:               "green sand",
			UsedForPottery:     false,
			IsSand:             true,
			UsedForAgriculture: false,
			Resources: []resource.Resource{
				{
					Name:         "green sand",
					Origin:       "green sand",
					MainMaterial: "green sand",
					Tags: []string{
						"sand",
					},
					Commonality: 1,
					Value:       1,
				},
			},
		},
		{
			Name:               "red sand",
			UsedForPottery:     false,
			IsSand:             true,
			UsedForAgriculture: false,
			Resources: []resource.Resource{
				{
					Name:         "red sand",
					Origin:       "red sand",
					MainMaterial: "red sand",
					Tags: []string{
						"sand",
					},
					Commonality: 2,
					Value:       1,
				},
			},
		},
		{
			Name:               "volcanic sand",
			UsedForPottery:     false,
			IsSand:             true,
			UsedForAgriculture: false,
			Resources: []resource.Resource{
				{
					Name:         "volcanic sand",
					Origin:       "volcanic sand",
					MainMaterial: "volcanic sand",
					Tags: []string{
						"sand",
					},
					Commonality: 3,
					Value:       1,
				},
			},
		},
		{
			Name:               "white sand",
			UsedForPottery:     false,
			IsSand:             true,
			UsedForAgriculture: false,
			Resources: []resource.Resource{
				{
					Name:         "white sand",
					Origin:       "white sand",
					MainMaterial: "white sand",
					Tags: []string{
						"sand",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
	}

	return soils
}
