package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/species"
)

func getBigCats() []species.Species {
	animals := []species.Species{
		{
			Name:       "bobcat",
			PluralName: "bobcats",
		},
		{
			Name:       "cheetah",
			PluralName: "cheetahs",
		},
		{
			Name:       "cougar",
			PluralName: "cougars",
		},
		{
			Name:       "jaguar",
			PluralName: "jaguars",
		},
		{
			Name:       "leopard",
			PluralName: "leopards",
		},
		{
			Name:       "lion",
			PluralName: "lions",
		},
		{
			Name:       "tiger",
			PluralName: "tigers",
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
				Value:       4,
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
				Commonality: 5,
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
			"big cat",
			"carnivore",
			"mammal",
		}
	}

	return animals
}
