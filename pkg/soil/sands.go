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
					Name:        "black sand",
					Origin:      "black sand",
					Type:        "sand",
					Commonality: 3,
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
					Name:        "continental sand",
					Origin:      "continental sand",
					Type:        "sand",
					Commonality: 4,
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
					Name:        "coral sand",
					Origin:      "coral sand",
					Type:        "sand",
					Commonality: 1,
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
					Name:        "dune sand",
					Origin:      "dune sand",
					Type:        "sand",
					Commonality: 4,
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
					Name:        "green sand",
					Origin:      "green sand",
					Type:        "sand",
					Commonality: 1,
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
					Name:        "red sand",
					Origin:      "red sand",
					Type:        "sand",
					Commonality: 2,
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
					Name:        "volcanic sand",
					Origin:      "volcanic sand",
					Type:        "sand",
					Commonality: 3,
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
					Name:        "white sand",
					Origin:      "white sand",
					Type:        "sand",
					Commonality: 5,
				},
			},
		},
	}

	return soils
}
