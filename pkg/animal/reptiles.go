package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/species"
)

func getReptiles() []species.Species {
	reptiles := []species.Species{
		{
			Name:           "alligator",
			PluralName:     "alligators",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 5,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "alligator hide",
					Origin:       "alligator",
					MainMaterial: "alligator",
					Tags: []string{
						"hide",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:         "alligator fangs",
					Origin:       "alligator",
					MainMaterial: "alligator",
					Tags: []string{
						"fangs",
						"teeth",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:         "alligator",
					Origin:       "alligator",
					MainMaterial: "alligator",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
			},
			Tags: []string{
				"animal",
				"reptile",
				"carnivore",
			},
		},
		{
			Name:           "crocodile",
			PluralName:     "crocodiles",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 5,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "crocodile hide",
					Origin:       "crocodile",
					MainMaterial: "crocodile",
					Tags: []string{
						"hide",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:         "crocodile fangs",
					Origin:       "crocodile",
					MainMaterial: "crocodile",
					Tags: []string{
						"fangs",
						"teeth",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:         "crocodile",
					Origin:       "crocodile",
					MainMaterial: "crocodile",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
			},
			Tags: []string{
				"animal",
				"reptile",
				"carnivore",
			},
		},
		{
			Name:           "monitor lizard",
			PluralName:     "monitor lizards",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 5,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "monitor lizard hide",
					Origin:       "monitor lizard",
					MainMaterial: "monitor lizard",
					Tags: []string{
						"hide",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:         "monitor lizard fangs",
					Origin:       "monitor lizard",
					MainMaterial: "monitor lizard",
					Tags: []string{
						"fangs",
						"teeth",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:         "monitor lizard",
					Origin:       "monitor lizard",
					MainMaterial: "monitor lizard",
					Tags: []string{
						"meat",
					},
					Commonality: 4,
					Value:       1,
				},
				{
					Name:         "monitor lizard venom",
					Origin:       "monitor lizard",
					MainMaterial: "monitor lizard",
					Tags: []string{
						"venom",
					},
					Commonality: 4,
					Value:       1,
				},
			},
			Tags: []string{
				"animal",
				"reptile",
				"carnivore",
			},
		},
	}

	return reptiles
}
