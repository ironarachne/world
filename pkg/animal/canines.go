package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getCanines() []Species {
	animals := []Species{
		{
			Name:       "coyote",
			PluralName: "coyotes",
		},
		{
			Name:       "fox",
			PluralName: "foxes",
		},
		{
			Name:       "wolf",
			PluralName: "wolves",
		},
	}

	for _, a := range animals {
		a.MinHumidity = 0
		a.MaxHumidity = 10
		a.MinTemperature = 0
		a.MaxTemperature = 10
		a.Size = size.GetCategoryByName("medium")
		a.Resources = []resource.Resource{
			{
				Name:         a.Name + " hide",
				Origin:       a.Name,
				MainMaterial: a.Name,
				Tags: []string{
					"hide",
				},
				Commonality: 4,
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
				Commonality: 4,
				Value:       1,
			},
			{
				Name:         a.Name,
				Origin:       a.Name,
				MainMaterial: a.Name,
				Tags: []string{
					"meat",
				},
				Commonality: 4,
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
			"canine",
			"carnivore",
			"hide",
		}
	}

	return animals
}
