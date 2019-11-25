/*
Package tree implements trees
*/
package tree

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/species"
)

// All returns all predefined trees
func All() []species.Species {
	trees := []species.Species{
		{
			Name:           "acacia",
			PluralName:     "acacias",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"arid",
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "acacia",
					Origin:       "acacia",
					MainMaterial: "acacia",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "alder",
			PluralName:     "alders",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "alder",
					Origin:       "alder",
					MainMaterial: "alder",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "apple",
			PluralName:     "apples",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "apple",
					Origin:       "apple",
					MainMaterial: "apple",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "apple",
					Origin:       "apple",
					MainMaterial: "apple",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "ash",
			PluralName:     "ashes",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "ash",
					Origin:       "ash",
					MainMaterial: "ash",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "aspen",
			PluralName:     "aspens",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "aspen",
					Origin:       "aspen",
					MainMaterial: "aspen",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "avocado",
			PluralName:     "avocados",
			MinHumidity:    6,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "avocado",
					Origin:       "avocado",
					MainMaterial: "avocado",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "avocado",
					Origin:       "avocado",
					MainMaterial: "avocado",
					Tags: []string{
						"wood",
					},
					Commonality: 4,
					Value:       1,
				},
			},
		},
		{
			Name:           "balsa",
			PluralName:     "balsas",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "balsa",
					Origin:       "balsa",
					MainMaterial: "balsa",
					Tags: []string{
						"wood",
					},
					Commonality: 2,
					Value:       1,
				},
			},
		},
		{
			Name:           "banana",
			PluralName:     "bananas",
			MinHumidity:    4,
			MaxHumidity:    10,
			MinTemperature: 7,
			MaxTemperature: 10,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "banana",
					Origin:       "banana",
					MainMaterial: "banana",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "banana",
					Origin:       "banana",
					MainMaterial: "banana",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "birch",
			PluralName:     "birches",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "birch",
					Origin:       "birch",
					MainMaterial: "birch",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "black pine",
			PluralName:     "black pines",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"coniferous",
			},
			Resources: []resource.Resource{
				{
					Name:         "black pine",
					Origin:       "black pine",
					MainMaterial: "black pine",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "blackwood",
			PluralName:     "blackwoods",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "blackwood",
					Origin:       "blackwood",
					MainMaterial: "blackwood",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "boxwood",
			PluralName:     "boxwoods",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "boxwood",
					Origin:       "boxwood",
					MainMaterial: "boxwood",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "cedar",
			PluralName:     "cedars",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"coniferous",
			},
			Resources: []resource.Resource{
				{
					Name:         "cedar",
					Origin:       "cedar",
					MainMaterial: "cedar",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "cherry",
			PluralName:     "cherries",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "cherry",
					Origin:       "cherry",
					MainMaterial: "cherry",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "cherry",
					Origin:       "cherry",
					MainMaterial: "cherry",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "cinnamon",
			PluralName:     "cinnamons",
			MinHumidity:    2,
			MaxHumidity:    8,
			MinTemperature: 0,
			MaxTemperature: 8,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "cinnamon",
					Origin:       "cinnamon",
					MainMaterial: "cinnamon",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "cinnamon",
					Origin:       "cinnamon",
					MainMaterial: "cinnamon",
					Tags: []string{
						"spice",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "coachwood",
			PluralName:     "coachwoods",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "coachwood",
					Origin:       "coachwood",
					MainMaterial: "coachwood",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "coconut",
			PluralName:     "coconuts",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "coconut",
					Origin:       "coconut",
					MainMaterial: "coconut",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "coconut",
					Origin:       "coconut",
					MainMaterial: "coconut",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "corkwood",
			PluralName:     "corkwoods",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "corkwood",
					Origin:       "corkwood",
					MainMaterial: "corkwood",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "cottonwood",
			PluralName:     "cottonwoods",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "cottonwood",
					Origin:       "cottonwood",
					MainMaterial: "cottonwood",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "crabapple",
			PluralName:     "crabapples",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "crabapple",
					Origin:       "crabapple",
					MainMaterial: "crabapple",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "crabapple",
					Origin:       "crabapple",
					MainMaterial: "crabapple",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "cypress",
			PluralName:     "cypress",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"coniferous",
			},
			Resources: []resource.Resource{
				{
					Name:         "cypress",
					Origin:       "cypress",
					MainMaterial: "cypress",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "dogwood",
			PluralName:     "dogwoods",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "dogwood",
					Origin:       "dogwood",
					MainMaterial: "dogwood",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "elm",
			PluralName:     "elms",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 5,
			MaxTemperature: 7,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "elm",
					Origin:       "elm",
					MainMaterial: "elm",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "eucalyptus",
			PluralName:     "eucalyptuses",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "eucalyptus",
					Origin:       "eucalyptus",
					MainMaterial: "eucalyptus",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "fir",
			PluralName:     "firs",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"coniferous",
			},
			Resources: []resource.Resource{
				{
					Name:         "fir",
					Origin:       "fir",
					MainMaterial: "fir",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "hemlock",
			PluralName:     "hemlocks",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"coniferous",
			},
			Resources: []resource.Resource{
				{
					Name:         "hemlock",
					Origin:       "hemlock",
					MainMaterial: "hemlock",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "ironwood",
			PluralName:     "ironwoods",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"arid",
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "ironwood",
					Origin:       "ironwood",
					MainMaterial: "ironwood",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "juniper",
			PluralName:     "junipers",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"coniferous",
			},
			Resources: []resource.Resource{
				{
					Name:         "juniper",
					Origin:       "juniper",
					MainMaterial: "juniper",
					Tags: []string{
						"berry",
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "juniper",
					Origin:       "juniper",
					MainMaterial: "juniper",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "kingwood",
			PluralName:     "kingwoods",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "kingwood",
					Origin:       "kingwood",
					MainMaterial: "kingwood",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "lacewood",
			PluralName:     "lacewoods",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "lacewood",
					Origin:       "lacewood",
					MainMaterial: "lacewood",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "larch",
			PluralName:     "larches",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"coniferous",
			},
			Resources: []resource.Resource{
				{
					Name:         "larch",
					Origin:       "larch",
					MainMaterial: "larch",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "lemon",
			PluralName:     "lemons",
			MinHumidity:    2,
			MaxHumidity:    9,
			MinTemperature: 0,
			MaxTemperature: 9,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "lemon",
					Origin:       "lemon",
					MainMaterial: "lemon",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "lemon",
					Origin:       "lemon",
					MainMaterial: "lemon",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "lime",
			PluralName:     "limes",
			MinHumidity:    2,
			MaxHumidity:    9,
			MinTemperature: 0,
			MaxTemperature: 9,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "lime",
					Origin:       "lime",
					MainMaterial: "lime",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "lime",
					Origin:       "lime",
					MainMaterial: "lime",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "mahogany",
			PluralName:     "mahoganies",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "mahogany",
					Origin:       "mahogany",
					MainMaterial: "mahogany",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "mango",
			PluralName:     "mangoes",
			MinHumidity:    4,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "mango",
					Origin:       "mango",
					MainMaterial: "mango",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "mango",
					Origin:       "mango",
					MainMaterial: "mango",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "maple",
			PluralName:     "maples",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 5,
			MaxTemperature: 7,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "maple",
					Origin:       "maple",
					MainMaterial: "maple",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "oak",
			PluralName:     "oaks",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 5,
			MaxTemperature: 7,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "oak",
					Origin:       "oak",
					MainMaterial: "oak",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "olive",
			PluralName:     "olives",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "olive",
					Origin:       "olive",
					MainMaterial: "olive",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "olives",
					Origin:       "olive",
					MainMaterial: "olive",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "papaya",
			PluralName:     "papayas",
			MinHumidity:    4,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "papaya",
					Origin:       "papaya",
					MainMaterial: "papaya",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "papaya",
					Origin:       "papaya",
					MainMaterial: "papaya",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "pine",
			PluralName:     "pine",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"coniferous",
			},
			Resources: []resource.Resource{
				{
					Name:         "pine",
					Origin:       "pine",
					MainMaterial: "pine",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "pineapple",
			PluralName:     "pineapples",
			MinHumidity:    2,
			MaxHumidity:    9,
			MinTemperature: 6,
			MaxTemperature: 9,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "pineapple",
					Origin:       "pineapple",
					MainMaterial: "pineapple",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "pineapple",
					Origin:       "pineapple",
					MainMaterial: "pineapple",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "palm",
			PluralName:     "palms",
			MinHumidity:    2,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "palm",
					Origin:       "palm",
					MainMaterial: "palm",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "pear",
			PluralName:     "pears",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "pear",
					Origin:       "pear",
					MainMaterial: "pear",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "pear",
					Origin:       "pear",
					MainMaterial: "pear",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "poplar",
			PluralName:     "poplars",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "poplar",
					Origin:       "poplar",
					MainMaterial: "poplar",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "red oak",
			PluralName:     "red oaks",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "red oak",
					Origin:       "red oak",
					MainMaterial: "red oak",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "red pine",
			PluralName:     "red pines",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"coniferous",
			},
			Resources: []resource.Resource{
				{
					Name:         "red pine",
					Origin:       "red pine",
					MainMaterial: "red pine",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "rosewood",
			PluralName:     "rosewoods",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "rosewood",
					Origin:       "rosewood",
					MainMaterial: "rosewood",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "sandalwood",
			PluralName:     "sandalwoods",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "sandalwood",
					Origin:       "sandalwood",
					MainMaterial: "sandalwood",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "spruce",
			PluralName:     "spruces",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"coniferous",
			},
			Resources: []resource.Resource{
				{
					Name:         "spruce",
					Origin:       "spruce",
					MainMaterial: "spruce",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "teak",
			PluralName:     "teaks",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "teak",
					Origin:       "teak",
					MainMaterial: "teak",
					Tags: []string{
						"wood",
					},
					Commonality: 3,
					Value:       1,
				},
			},
		},
		{
			Name:           "walnut",
			PluralName:     "walnuts",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "walnut",
					Origin:       "walnut",
					MainMaterial: "walnut",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "white oak",
			PluralName:     "white oaks",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "white oak",
					Origin:       "white oak",
					MainMaterial: "white oak",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "white pine",
			PluralName:     "white pines",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"coniferous",
			},
			Resources: []resource.Resource{
				{
					Name:         "white pine",
					Origin:       "white pine",
					MainMaterial: "white pine",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "willow",
			PluralName:     "willows",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"deciduous",
			},
			Resources: []resource.Resource{
				{
					Name:         "willow",
					Origin:       "willow",
					MainMaterial: "willow",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "yew",
			PluralName:     "yews",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"coniferous",
			},
			Resources: []resource.Resource{
				{
					Name:         "yew",
					Origin:       "yew",
					MainMaterial: "yew",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
		{
			Name:           "yellow pine",
			PluralName:     "yellow pines",
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Tags: []string{
				"coniferous",
			},
			Resources: []resource.Resource{
				{
					Name:         "yellow pine",
					Origin:       "yellow pine",
					MainMaterial: "yellow pine",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
			},
		},
	}

	return trees
}
