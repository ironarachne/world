package plant

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/species"
)

func getBushes() []species.Species {
	plants := []species.Species{
		{
			Name:           "blackberry bush",
			PluralName:     "blackberry bushes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "blackberry",
					Origin:       "blackberry bush",
					MainMaterial: "blackberry",
					Tags: []string{
						"berry",
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "blueberry bush",
			PluralName:     "blueberry bushes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "blueberry",
					Origin:       "blueberry bush",
					MainMaterial: "blueberry",
					Tags: []string{
						"berry",
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "raspberry bush",
			PluralName:     "raspberry bushes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "raspberry",
					Origin:       "raspberry bush",
					MainMaterial: "raspberry",
					Tags: []string{
						"berry",
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "strawberry bush",
			PluralName:     "strawberry bushes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "strawberry",
					Origin:       "strawberry bush",
					MainMaterial: "strawberry",
					Tags: []string{
						"berry",
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "boysenberry bush",
			PluralName:     "boysenberry bushes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "boysenberry",
					Origin:       "boysenberry bush",
					MainMaterial: "boysenberry",
					Tags: []string{
						"berry",
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "gooseberry bush",
			PluralName:     "gooseberry bushes",
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "gooseberry",
					Origin:       "gooseberry bush",
					MainMaterial: "gooseberry",
					Tags: []string{
						"berry",
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
	}

	return plants
}
