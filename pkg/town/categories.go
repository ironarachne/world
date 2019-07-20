package town

import "github.com/ironarachne/world/pkg/random"

// Category is a type of town
type Category struct {
	Name        string
	MinSize     int
	MaxSize     int
	MinExports  int
	MaxExports  int
	MinImports  int
	MaxImports  int
	Commonality int
}

func getAllCategories() []Category {
	categories := []Category{
		{
			Name:        "metropolis",
			MinSize:     100000,
			MaxSize:     1000000,
			MinExports:  50,
			MaxExports:  500,
			MinImports:  50,
			MaxImports:  100,
			Commonality: 1,
		},
		{
			Name:        "city",
			MinSize:     10000,
			MaxSize:     100000,
			MinExports:  20,
			MaxExports:  100,
			MinImports:  3,
			MaxImports:  6,
			Commonality: 3,
		},
		{
			Name:        "town",
			MinSize:     1000,
			MaxSize:     10000,
			MinExports:  10,
			MaxExports:  30,
			MinImports:  1,
			MaxImports:  3,
			Commonality: 5,
		},
		{
			Name:        "village",
			MinSize:     10,
			MaxSize:     100,
			MinExports:  1,
			MaxExports:  3,
			MinImports:  1,
			MaxImports:  2,
			Commonality: 10,
		},
	}

	return categories
}

func getCategoryByName(name string) Category {
	cats := getAllCategories()

	for _, c := range cats {
		if c.Name == name {
			return c
		}
	}

	return Category{}
}

func getRandomWeightedCategory() Category {
	categories := getAllCategories()

	weights := map[string]int{}

	for _, c := range categories {
		weights[c.Name] = c.Commonality
	}

	name := random.StringFromThresholdMap(weights)

	for _, c := range categories {
		if c.Name == name {
			return c
		}
	}

	return Category{}
}
