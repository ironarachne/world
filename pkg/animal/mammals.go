package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
	"github.com/ironarachne/world/pkg/species"
)

func getMammals() []Species {
	animals := []Species{
		{
			Name:           "beaver",
			PluralName:     "beavers",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 6,
			Resources: []resource.Resource{
				{
					Name:         "beaver hide",
					Origin:       "beaver",
					MainMaterial: "beaver",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "beaver teeth",
					Origin:       "beaver",
					MainMaterial: "beaver",
					Tags: []string{
						"teeth",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "beaver meat",
					Origin:       "beaver",
					MainMaterial: "beaver",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("small"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
				"hide",
			},
		},
		{
			Name:           "deer",
			PluralName:     "deer",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "deer hide",
					Origin:       "deer",
					MainMaterial: "deer",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "deer teeth",
					Origin:       "deer",
					MainMaterial: "deer",
					Tags: []string{
						"teeth",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "venison",
					Origin:       "deer",
					MainMaterial: "deer",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "deer antler",
					Origin:       "deer",
					MainMaterial: "deer",
					Tags: []string{
						"antler",
						"bone",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "deer sinew",
					Origin:       "deer",
					MainMaterial: "deer",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("medium"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
				"hide",
			},
		},
		{
			Name:           "squirrel",
			PluralName:     "squirrels",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "squirrel hide",
					Origin:       "squirrel",
					MainMaterial: "squirrel",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "squirrel",
					Origin:       "squirrel",
					MainMaterial: "squirrel",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("tiny"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
			},
		},
		{
			Name:           "camel",
			PluralName:     "camels",
			MinHumidity:    0,
			MaxHumidity:    6,
			MinTemperature: 5,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "camel hide",
					Origin:       "camel",
					MainMaterial: "camel",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "camel teeth",
					Origin:       "camel",
					MainMaterial: "camel",
					Tags: []string{
						"teeth",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "camel",
					Origin:       "camel",
					MainMaterial: "camel",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "camel milk",
					Origin:       "camel",
					MainMaterial: "camel",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "camel sinew",
					Origin:       "camel",
					MainMaterial: "camel",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("medium"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
				"hide",
			},
		},
		{
			Name:           "bison",
			PluralName:     "bison",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 6,
			Resources: []resource.Resource{
				{
					Name:         "bison hide",
					Origin:       "bison",
					MainMaterial: "bison",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "bison teeth",
					Origin:       "bison",
					MainMaterial: "bison",
					Tags: []string{
						"teeth",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "bison",
					Origin:       "bison",
					MainMaterial: "bison",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "bison bone",
					Origin:       "bison",
					MainMaterial: "bison",
					Tags: []string{
						"bone",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "bison sinew",
					Origin:       "bison",
					MainMaterial: "bison",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("medium"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
				"hide",
			},
		},
		{
			Name:           "cow",
			PluralName:     "cows",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "calf brains",
					Origin:       "cow",
					MainMaterial: "cow",
					Tags: []string{
						"brains",
						"meat",
					},
					Commonality: 2,
					Value:       5,
				},
				{
					Name:         "cow hide",
					Origin:       "cow",
					MainMaterial: "cow",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "cow teeth",
					Origin:       "cow",
					MainMaterial: "cow",
					Tags: []string{
						"teeth",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "beef",
					Origin:       "cow",
					MainMaterial: "cow",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "beef loin",
					Origin:       "cow",
					MainMaterial: "cow",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "beef ribs",
					Origin:       "cow",
					MainMaterial: "cow",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "cow bone",
					Origin:       "cow",
					MainMaterial: "cow",
					Tags: []string{
						"bone",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "cow sinew",
					Origin:       "cow",
					MainMaterial: "cow",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("medium"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
				"hide",
			},
		},
		{
			Name:           "elephant",
			PluralName:     "elephants",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 5,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "elephant hide",
					Origin:       "elephant",
					MainMaterial: "elephant",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "elephant milk",
					Origin:       "elephant",
					MainMaterial: "elephant",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "elephant",
					Origin:       "elephant",
					MainMaterial: "elephant",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "ivory",
					Origin:       "elephant",
					MainMaterial: "elephant",
					Tags: []string{
						"bone",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "elephant sinew",
					Origin:       "elephant",
					MainMaterial: "elephant",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("large"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
				"hide",
			},
		},
		{
			Name:           "goat",
			PluralName:     "goats",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "goat hide",
					Origin:       "goat",
					MainMaterial: "goat",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "goat's milk",
					Origin:       "goat",
					MainMaterial: "goat",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "goat",
					Origin:       "goat",
					MainMaterial: "goat",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "goat sinew",
					Origin:       "goat",
					MainMaterial: "goat",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("medium"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
				"hide",
			},
		},
		{
			Name:           "sheep",
			PluralName:     "sheep",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "wool",
					Origin:       "sheep",
					MainMaterial: "sheep",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "sheep's milk",
					Origin:       "sheep",
					MainMaterial: "sheep",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "lamb",
					Origin:       "sheep",
					MainMaterial: "sheep",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "sheep sinew",
					Origin:       "sheep",
					MainMaterial: "sheep",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("medium"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
				"wool",
			},
		},
		{
			Name:           "alpaca",
			PluralName:     "alpacas",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 4,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:         "alpaca wool",
					Origin:       "alpaca",
					MainMaterial: "alpaca",
					Tags: []string{
						"hide",
						"wool",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "alpaca milk",
					Origin:       "alpaca",
					MainMaterial: "alpaca",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "alpaca",
					Origin:       "alpaca",
					MainMaterial: "alpaca",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "alpaca sinew",
					Origin:       "alpaca",
					MainMaterial: "alpaca",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("medium"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
				"wool",
			},
		},
		{
			Name:           "llama",
			PluralName:     "llamas",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:         "llama wool",
					Origin:       "llama",
					MainMaterial: "llama",
					Tags: []string{
						"hide",
						"wool",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "llama",
					Origin:       "llama",
					MainMaterial: "llama",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "llama milk",
					Origin:       "llama",
					MainMaterial: "llama",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "llama sinew",
					Origin:       "llama",
					MainMaterial: "llama",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("medium"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
				"wool",
			},
		},
		{
			Name:           "hippopotamus",
			PluralName:     "hippopotamuses",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "hippopotamus hide",
					Origin:       "hippopotamus",
					MainMaterial: "hippopotamus",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "hippopotamus",
					Origin:       "hippopotamus",
					MainMaterial: "hippopotamus",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "hippopotamus milk",
					Origin:       "hippopotamus",
					MainMaterial: "hippopotamus",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("large"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
				"hide",
			},
		},
		{
			Name:           "antelope",
			PluralName:     "antelopes",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "antelope hide",
					Origin:       "antelope",
					MainMaterial: "antelope",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "antelope",
					Origin:       "antelope",
					MainMaterial: "antelope",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "antelope milk",
					Origin:       "antelope",
					MainMaterial: "antelope",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "antelope sinew",
					Origin:       "antelope",
					MainMaterial: "antelope",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("medium"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
				"hide",
			},
		},
		{
			Name:           "gazelle",
			PluralName:     "gazelles",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 6,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "gazelle hide",
					Origin:       "gazelle",
					MainMaterial: "gazelle",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "gazelle",
					Origin:       "gazelle",
					MainMaterial: "gazelle",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "gazelle milk",
					Origin:       "gazelle",
					MainMaterial: "gazelle",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "gazelle sinew",
					Origin:       "gazelle",
					MainMaterial: "gazelle",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("medium"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
				"hide",
			},
		},
		{
			Name:           "rabbit",
			PluralName:     "rabbits",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "rabbit hide",
					Origin:       "rabbit",
					MainMaterial: "rabbit",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "rabbit",
					Origin:       "rabbit",
					MainMaterial: "rabbit",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("tiny"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
			},
		},
		{
			Name:           "ermine",
			PluralName:     "ermines",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 6,
			Resources: []resource.Resource{
				{
					Name:         "ermine fur",
					Origin:       "ermine",
					MainMaterial: "ermine",
					Tags: []string{
						"fur",
					},
					Commonality: 5,
					Value:       10,
				},
				{
					Name:         "ermine",
					Origin:       "ermine",
					MainMaterial: "ermine",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("small"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
			},
		},
		{
			Name:           "mink",
			PluralName:     "minks",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 6,
			Resources: []resource.Resource{
				{
					Name:         "mink fur",
					Origin:       "mink",
					MainMaterial: "mink",
					Tags: []string{
						"fur",
					},
					Commonality: 5,
					Value:       10,
				},
				{
					Name:         "mink",
					Origin:       "mink",
					MainMaterial: "mink",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("small"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
			},
		},
		{
			Name:           "pig",
			PluralName:     "pigs",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				{
					Name:         "pig hide",
					Origin:       "pig",
					MainMaterial: "pig",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "pork",
					Origin:       "pig",
					MainMaterial: "pig",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "pork loin",
					Origin:       "pig",
					MainMaterial: "pig",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "pork intestine",
					Origin:       "pig",
					MainMaterial: "pig",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "pork ribs",
					Origin:       "pig",
					MainMaterial: "pig",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "bacon",
					Origin:       "pig",
					MainMaterial: "pig",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "pig sinew",
					Origin:       "pig",
					MainMaterial: "pig",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("medium"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
				"hide",
			},
		},
		{
			Name:           "raccoon",
			PluralName:     "raccoons",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 7,
			Resources: []resource.Resource{
				{
					Name:         "raccoon hide",
					Origin:       "raccoon",
					MainMaterial: "raccoon",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "raccoon",
					Origin:       "raccoon",
					MainMaterial: "raccoon",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("small"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
			},
		},
		{
			Name:           "reindeer",
			PluralName:     "reindeer",
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 4,
			Resources: []resource.Resource{
				{
					Name:         "reindeer hide",
					Origin:       "reindeer",
					MainMaterial: "reindeer",
					Tags: []string{
						"hide",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "reindeer",
					Origin:       "reindeer",
					MainMaterial: "reindeer",
					Tags: []string{
						"meat",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "reindeer milk",
					Origin:       "reindeer",
					MainMaterial: "reindeer",
					Tags: []string{
						"milk",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "reindeer bone",
					Origin:       "reindeer",
					MainMaterial: "reindeer",
					Tags: []string{
						"bone",
					},
					Commonality: 5,
					Value:       1,
				},
				{
					Name:         "reindeer sinew",
					Origin:       "reindeer",
					MainMaterial: "reindeer",
					Tags: []string{
						"sinew",
					},
					Commonality: 5,
					Value:       1,
				},
			},
			Size: size.GetCategoryByName("medium"),
			Tags: []string{
				"animal",
				"mammal",
				"herbivore",
				"hide",
			},
		},
	}

	return animals
}
