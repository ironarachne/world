package tree

import (
	"github.com/ironarachne/world/pkg/resource"
)

// All returns all predefined trees
func All() []Tree {
	trees := []Tree{
		{
			Name:           "acacia",
			PluralName:     "acacias",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "acacia",
					Origin: "acacia",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "alder",
					Origin: "alder",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "apple",
					Origin: "apple",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "apple",
					Origin: "apple",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "ash",
					Origin: "ash",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "aspen",
					Origin: "aspen",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    6,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "avocado",
					Origin: "avocado",
					Tags: []string{
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "avocado",
					Origin: "avocado",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "balsa",
					Origin: "balsa",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    4,
			MaxHumidity:    10,
			MinTemperature: 7,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "banana",
					Origin: "banana",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "banana",
					Origin: "banana",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "birch",
					Origin: "birch",
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
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "black pine",
					Origin: "black pine",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "blackwood",
					Origin: "blackwood",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "boxwood",
					Origin: "boxwood",
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
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "cedar",
					Origin: "cedar",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "cherry",
					Origin: "cherry",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "cherry",
					Origin: "cherry",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    8,
			MinTemperature: 0,
			MaxTemperature: 8,
			Resources: []resource.Resource{
				{
					Name:   "cinnamon",
					Origin: "cinnamon",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "cinnamon",
					Origin: "cinnamon",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "coachwood",
					Origin: "coachwood",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "coconut",
					Origin: "coconut",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "coconut",
					Origin: "coconut",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "corkwood",
					Origin: "corkwood",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "cottonwood",
					Origin: "cottonwood",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "crabapple",
					Origin: "crabapple",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "crabapple",
					Origin: "crabapple",
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
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "cypress",
					Origin: "cypress",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "dogwood",
					Origin: "dogwood",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 5,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:   "elm",
					Origin: "elm",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "eucalyptus",
					Origin: "eucalyptus",
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
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "fir",
					Origin: "fir",
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
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "hemlock",
					Origin: "hemlock",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "ironwood",
					Origin: "ironwood",
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
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "juniper",
					Origin: "juniper",
					Tags: []string{
						"berry",
						"fruit",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "juniper",
					Origin: "juniper",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "kingwood",
					Origin: "kingwood",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "lacewood",
					Origin: "lacewood",
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
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "larch",
					Origin: "larch",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    9,
			MinTemperature: 0,
			MaxTemperature: 9,
			Resources: []resource.Resource{
				{
					Name:   "lemon",
					Origin: "lemon",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "lemon",
					Origin: "lemon",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    9,
			MinTemperature: 0,
			MaxTemperature: 9,
			Resources: []resource.Resource{
				{
					Name:   "lime",
					Origin: "lime",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "lime",
					Origin: "lime",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "mahogany",
					Origin: "mahogany",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    4,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "mango",
					Origin: "mango",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "mango",
					Origin: "mango",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 5,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:   "maple",
					Origin: "maple",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 5,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:   "oak",
					Origin: "oak",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "olive",
					Origin: "olive",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "olives",
					Origin: "olive",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    4,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "papaya",
					Origin: "papaya",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "papaya",
					Origin: "papaya",
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
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "pine",
					Origin: "pine",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    9,
			MinTemperature: 6,
			MaxTemperature: 9,
			Resources: []resource.Resource{
				{
					Name:   "pineapple",
					Origin: "pineapple",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "pineapple",
					Origin: "pineapple",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:   "palm",
					Origin: "palm",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "pear",
					Origin: "pear",
					Tags: []string{
						"wood",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:   "pear",
					Origin: "pear",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "poplar",
					Origin: "poplar",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "red oak",
					Origin: "red oak",
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
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "red pine",
					Origin: "red pine",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "rosewood",
					Origin: "rosewood",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "sandalwood",
					Origin: "sandalwood",
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
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "spruce",
					Origin: "spruce",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "teak",
					Origin: "teak",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "walnut",
					Origin: "walnut",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "white oak",
					Origin: "white oak",
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
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "white pine",
					Origin: "white pine",
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
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "willow",
					Origin: "willow",
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
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "yew",
					Origin: "yew",
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
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				{
					Name:   "yellow pine",
					Origin: "yellow pine",
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
