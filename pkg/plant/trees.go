package plant

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/resource"
)

// Tree is a tree
type Tree struct {
	Name           string
	PluralName     string
	IsDeciduous    bool
	IsConiferous   bool
	MinHumidity    int
	MaxHumidity    int
	MinTemperature int
	MaxTemperature int
	Resources      []resource.Resource
}

// InSlice checks if a tree is in the given slice
func (tree Tree) InSlice(trees []Tree) bool {
	isIt := false
	for _, a := range trees {
		if a.Name == tree.Name {
			isIt = true
		}
	}

	return isIt
}

// RandomTrees returns a random subset of plants
func RandomTrees(amount int, from []Tree) []Tree {
	var tree Tree

	trees := []Tree{}

	if amount > len(from) {
		amount = len(from)
	}

	for i := 1; i < amount; i++ {
		tree = from[rand.Intn(len(from))]
		if !tree.InSlice(trees) {
			trees = append(trees, tree)
		}
	}

	return trees
}

// AllTrees returns all predefined trees
func AllTrees() []Tree {
	trees := []Tree{
		Tree{
			Name:           "acacia",
			PluralName:     "acacias",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "acacia",
					Origin:      "acacia",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "alder",
			PluralName:     "alders",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "alder",
					Origin:      "alder",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "apple",
			PluralName:     "apples",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "apple",
					Origin:      "apple",
					Type:        "wood",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "apple",
					Origin:      "apple",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "ash",
			PluralName:     "ashes",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "ash",
					Origin:      "ash",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "aspen",
			PluralName:     "aspens",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "aspen",
					Origin:      "aspen",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "balsa",
			PluralName:     "balsas",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "balsa",
					Origin:      "balsa",
					Type:        "wood",
					Commonality: 2,
				},
			},
		},
		Tree{
			Name:           "banana",
			PluralName:     "bananas",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    4,
			MaxHumidity:    10,
			MinTemperature: 7,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "banana",
					Origin:      "banana",
					Type:        "wood",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "banana",
					Origin:      "banana",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "birch",
			PluralName:     "birches",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "birch",
					Origin:      "birch",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "black pine",
			PluralName:     "black pines",
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "black pine",
					Origin:      "black pine",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "blackwood",
			PluralName:     "blackwoods",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "blackwood",
					Origin:      "blackwood",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "boxwood",
			PluralName:     "boxwoods",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "boxwood",
					Origin:      "boxwood",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "cedar",
			PluralName:     "cedars",
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "cedar",
					Origin:      "cedar",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "cherry",
			PluralName:     "cherries",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "cherry",
					Origin:      "cherry",
					Type:        "wood",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "cherry",
					Origin:      "cherry",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "cinnamon",
			PluralName:     "cinnamons",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    8,
			MinTemperature: 0,
			MaxTemperature: 8,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "cinnamon",
					Origin:      "cinnamon",
					Type:        "wood",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "cinnamon",
					Origin:      "cinnamon",
					Type:        "spice",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "coachwood",
			PluralName:     "coachwoods",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "coachwood",
					Origin:      "coachwood",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "coconut",
			PluralName:     "coconuts",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "coconut",
					Origin:      "coconut",
					Type:        "wood",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "coconut",
					Origin:      "coconut",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "corkwood",
			PluralName:     "corkwoods",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "corkwood",
					Origin:      "corkwood",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "cottonwood",
			PluralName:     "cottonwoods",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "cottonwood",
					Origin:      "cottonwood",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "crabapple",
			PluralName:     "crabapples",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "crabapple",
					Origin:      "crabapple",
					Type:        "wood",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "crabapple",
					Origin:      "crabapple",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "cypress",
			PluralName:     "cypress",
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "cypress",
					Origin:      "cypress",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "dogwood",
			PluralName:     "dogwoods",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "dogwood",
					Origin:      "dogwood",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "elm",
			PluralName:     "elms",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 5,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "elm",
					Origin:      "elm",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "eucalyptus",
			PluralName:     "eucalyptuses",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "eucalyptus",
					Origin:      "eucalyptus",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "fir",
			PluralName:     "firs",
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "fir",
					Origin:      "fir",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "hemlock",
			PluralName:     "hemlocks",
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "hemlock",
					Origin:      "hemlock",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "ironwood",
			PluralName:     "ironwoods",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "ironwood",
					Origin:      "ironwood",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "juniper",
			PluralName:     "junipers",
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "juniper",
					Origin:      "juniper",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "kingwood",
			PluralName:     "kingwoods",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "kingwood",
					Origin:      "kingwood",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "lacewood",
			PluralName:     "lacewoods",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "lacewood",
					Origin:      "lacewood",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "larch",
			PluralName:     "larches",
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "larch",
					Origin:      "larch",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "lemon",
			PluralName:     "lemons",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    9,
			MinTemperature: 0,
			MaxTemperature: 9,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "lemon",
					Origin:      "lemon",
					Type:        "wood",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "lemon",
					Origin:      "lemon",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "lime",
			PluralName:     "limes",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    9,
			MinTemperature: 0,
			MaxTemperature: 9,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "lime",
					Origin:      "lime",
					Type:        "wood",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "lime",
					Origin:      "lime",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "mahogany",
			PluralName:     "mahoganies",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "mahogany",
					Origin:      "mahogany",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "mango",
			PluralName:     "mangoes",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    4,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "mango",
					Origin:      "mango",
					Type:        "wood",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "mango",
					Origin:      "mango",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "maple",
			PluralName:     "maples",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 5,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "maple",
					Origin:      "maple",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "oak",
			PluralName:     "oaks",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 5,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "oak",
					Origin:      "oak",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "olive",
			PluralName:     "olives",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "olive",
					Origin:      "olive",
					Type:        "wood",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "olives",
					Origin:      "olive",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "papaya",
			PluralName:     "papayas",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    4,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "papaya",
					Origin:      "papaya",
					Type:        "wood",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "papaya",
					Origin:      "papaya",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "pine",
			PluralName:     "pine",
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "pine",
					Origin:      "pine",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "pineapple",
			PluralName:     "pineapples",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    9,
			MinTemperature: 6,
			MaxTemperature: 9,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "pineapple",
					Origin:      "pineapple",
					Type:        "wood",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "pineapple",
					Origin:      "pineapple",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "palm",
			PluralName:     "palms",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "palm",
					Origin:      "palm",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "pear",
			PluralName:     "pears",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "pear",
					Origin:      "pear",
					Type:        "wood",
					Commonality: 5,
				},
				resource.Resource{
					Name:        "pear",
					Origin:      "pear",
					Type:        "fruit",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "poplar",
			PluralName:     "poplars",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "poplar",
					Origin:      "poplar",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "red oak",
			PluralName:     "red oaks",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "red oak",
					Origin:      "red oak",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "red pine",
			PluralName:     "red pines",
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "red pine",
					Origin:      "red pine",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "rosewood",
			PluralName:     "rosewoods",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "rosewood",
					Origin:      "rosewood",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "sandalwood",
			PluralName:     "sandalwoods",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "sandalwood",
					Origin:      "sandalwood",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "spruce",
			PluralName:     "spruces",
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "spruce",
					Origin:      "spruce",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "teak",
			PluralName:     "teaks",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "teak",
					Origin:      "teak",
					Type:        "wood",
					Commonality: 3,
				},
			},
		},
		Tree{
			Name:           "walnut",
			PluralName:     "walnuts",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "walnut",
					Origin:      "walnut",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "white oak",
			PluralName:     "white oaks",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "white oak",
					Origin:      "white oak",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "white pine",
			PluralName:     "white pines",
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "white pine",
					Origin:      "white pine",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "willow",
			PluralName:     "willows",
			IsDeciduous:    true,
			IsConiferous:   false,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "willow",
					Origin:      "willow",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "yew",
			PluralName:     "yews",
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "yew",
					Origin:      "yew",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
		Tree{
			Name:           "yellow pine",
			PluralName:     "yellow pines",
			IsDeciduous:    false,
			IsConiferous:   true,
			MinHumidity:    2,
			MaxHumidity:    6,
			MinTemperature: 0,
			MaxTemperature: 5,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "yellow pine",
					Origin:      "yellow pine",
					Type:        "wood",
					Commonality: 5,
				},
			},
		},
	}

	return trees
}
