package plant

import "github.com/ironarachne/world/pkg/resource"

func getGrains() []Plant {
	grains := []Plant{
		Plant{
			Name:           "rye",
			PluralName:     "ryes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "rye",
					Origin:      "rye",
					Type:        "grain",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "rye",
					Origin:      "rye",
					Type:        "flour",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "oat",
			PluralName:     "oats",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "oat",
					Origin:      "oat",
					Type:        "grain",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "oat",
					Origin:      "oat",
					Type:        "flour",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "millet",
			PluralName:     "millets",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "millet",
					Origin:      "millet",
					Type:        "grain",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "millet",
					Origin:      "millet",
					Type:        "flour",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "barley",
			PluralName:     "barleys",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "barley",
					Origin:      "barley",
					Type:        "grain",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "barley",
					Origin:      "barley",
					Type:        "flour",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "wheat",
			PluralName:     "wheats",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "wheat",
					Origin:      "wheat",
					Type:        "grain",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "wheat",
					Origin:      "wheat",
					Type:        "flour",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "rice",
			PluralName:     "rices",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 3,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "rice",
					Origin:      "rice",
					Type:        "grain",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "rice",
					Origin:      "rice",
					Type:        "flour",
					Commonality: 5,
				},
			},
		},
	}

	return grains
}
