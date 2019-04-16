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
		Category{
			Name:        "city",
			MinSize:     10000,
			MaxSize:     100000,
			MinExports:  3,
			MaxExports:  12,
			MinImports:  3,
			MaxImports:  6,
			Commonality: 3,
		},
		Category{
			Name:        "town",
			MinSize:     1000,
			MaxSize:     10000,
			MinExports:  1,
			MaxExports:  3,
			MinImports:  1,
			MaxImports:  3,
			Commonality: 5,
		},
		Category{
			Name:        "village",
			MinSize:     10,
			MaxSize:     100,
			MinExports:  1,
			MaxExports:  1,
			MinImports:  1,
			MaxImports:  1,
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
