package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/species"
)

func getSnakes() []species.Species {
	animals := []species.Species{
		{
			Name:       "asp",
			PluralName: "asps",
		},
		{
			Name:       "cobra",
			PluralName: "cobras",
		},
		{
			Name:       "rattlesnake",
			PluralName: "rattlesnakes",
		},
		{
			Name:       "black mamba",
			PluralName: "black mambas",
		},
		{
			Name:       "coral snake",
			PluralName: "coral snakes",
		},
		{
			Name:       "boa constrictor",
			PluralName: "boa constrictors",
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
				Commonality: 3,
				Value:       1,
			},
			{
				Name:         a.Name + " fangs",
				Origin:       a.Name,
				MainMaterial: a.Name,
				Tags: []string{
					"fangs",
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
				Name:         a.Name + " heart",
				Origin:       a.Name,
				MainMaterial: a.Name,
				Tags: []string{
					"meat",
				},
				Commonality: 1,
				Value:       1,
			},
			{
				Name:         a.Name + " venom",
				Origin:       a.Name,
				MainMaterial: a.Name,
				Tags: []string{
					"venom",
				},
				Commonality: 3,
				Value:       10,
			},
		}
		a.Tags = []string{
			"animal",
			"reptile",
			"snake",
			"carnivore",
		}
	}

	return animals
}
