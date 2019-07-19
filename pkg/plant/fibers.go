package plant

import (
	"github.com/ironarachne/world/pkg/resource"
)

func getFibers() []Plant {
	fibers := []Plant{
		Plant{
			Name:           "cotton",
			PluralName:     "cotton",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "cotton",
					Origin:      "cotton",
					Type:        "fabric",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "cotton thread",
					Origin:      "cotton",
					Type:        "thread",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "flax",
			PluralName:     "flax",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "linen",
					Origin:      "flax",
					Type:        "fabric",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "linen thread",
					Origin:      "flax",
					Type:        "thread",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "hemp",
			PluralName:     "hemp",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 9,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "hemp",
					Origin:      "hemp",
					Type:        "fabric",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "hemp thread",
					Origin:      "hemp",
					Type:        "thread",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "coir",
			PluralName:     "coir",
			MinHumidity:    6,
			MaxHumidity:    10,
			MinTemperature: 8,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "coir seed",
					Origin:      "coir",
					Type:        "seed",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "papyrus",
			PluralName:     "papyrus",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 9,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "papyrus",
					Origin:      "papyrus",
					Type:        "paper",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "jute",
			PluralName:     "jute",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 9,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "jute",
					Origin:      "jute",
					Type:        "fabric",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "jute thread",
					Origin:      "jute",
					Type:        "thread",
					Commonality: 5,
				},
			},
		},
		Plant{
			Name:           "ramie",
			PluralName:     "ramie",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "ramie",
					Origin:      "ramie",
					Type:        "fabric",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "ramie thread",
					Origin:      "ramie",
					Type:        "thread",
					Commonality: 5,
				},
			},
		},
	}

	return fibers
}
