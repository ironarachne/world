package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/species"
)

func getUrsines() []species.Species {
	animals := []species.Species{
		{
			Name:       "black bear",
			PluralName: "black bears",
		},
		{
			Name:       "brown bear",
			PluralName: "brown bears",
		},
	}

	for _, a := range animals {
		a.MinHumidity = 0
		a.MaxHumidity = 10
		a.MinTemperature = 6
		a.MaxTemperature = 10
		a.Resources = []resource.Resource{
			{
				Name:         a.Name + " hide",
				Origin:       a.Name,
				MainMaterial: a.Name,
				Tags: []string{
					"hide",
				},
				Commonality: 5,
				Value:       5,
			},
			{
				Name:         a.Name + " teeth",
				Origin:       a.Name,
				MainMaterial: a.Name,
				Tags: []string{
					"teeth",
				},
				Commonality: 3,
				Value:       1,
			},
			{
				Name:         a.Name,
				Origin:       a.Name,
				MainMaterial: a.Name,
				Tags: []string{
					"meat",
				},
				Commonality: 3,
				Value:       1,
			},
			{
				Name:         a.Name,
				Origin:       a.Name,
				MainMaterial: a.Name,
				Tags: []string{
					"sinew",
				},
				Commonality: 5,
				Value:       1,
			},
		}
		a.Tags = []string{
			"animal",
			"bear",
			"omnivore",
		}
	}

	return animals
}
