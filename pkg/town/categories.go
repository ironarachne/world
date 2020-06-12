package town

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/random"
)

// Category is a type of town
type Category struct {
	Name                 string `json:"name"`
	MinSize              int    `json:"min_size"`
	MaxSize              int    `json:"max_size"`
	MinExports           int    `json:"min_exports"`
	MaxExports           int    `json:"max_exports"`
	MinImports           int    `json:"min_imports"`
	MaxImports           int    `json:"max_imports"`
	ProductionIterations int    `json:"production_iterations"`
	Commonality          int    `json:"commonality"`
}

func getAllCategories() []Category {
	categories := []Category{
		{
			Name:                 "metropolis",
			MinSize:              1000000,
			MaxSize:              3000000,
			MinExports:           50,
			MaxExports:           500,
			MinImports:           50,
			MaxImports:           100,
			ProductionIterations: 12,
			Commonality:          1,
		},
		{
			Name:                 "city",
			MinSize:              100000,
			MaxSize:              999999,
			MinExports:           20,
			MaxExports:           100,
			MinImports:           3,
			MaxImports:           6,
			ProductionIterations: 7,
			Commonality:          3,
		},
		{
			Name:                 "borough",
			MinSize:              10000,
			MaxSize:              99999,
			MinExports:           20,
			MaxExports:           100,
			MinImports:           3,
			MaxImports:           6,
			ProductionIterations: 5,
			Commonality:          3,
		},
		{
			Name:                 "town",
			MinSize:              1000,
			MaxSize:              9999,
			MinExports:           10,
			MaxExports:           30,
			MinImports:           1,
			MaxImports:           3,
			ProductionIterations: 4,
			Commonality:          15,
		},
		{
			Name:                 "village",
			MinSize:              100,
			MaxSize:              999,
			MinExports:           1,
			MaxExports:           4,
			MinImports:           1,
			MaxImports:           3,
			ProductionIterations: 3,
			Commonality:          20,
		},
		{
			Name:                 "hamlet",
			MinSize:              10,
			MaxSize:              99,
			MinExports:           1,
			MaxExports:           3,
			MinImports:           1,
			MaxImports:           2,
			ProductionIterations: 1,
			Commonality:          10,
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

func getRandomWeightedCategory(ctx context.Context) (Category, error) {
	categories := getAllCategories()

	weights := map[string]int{}

	for _, c := range categories {
		weights[c.Name] = c.Commonality
	}

	name, err := random.StringFromThresholdMap(ctx, weights)
	if err != nil {
		err = fmt.Errorf("failed to get random weighted town category: %w", err)
		return Category{}, err
	}

	for _, c := range categories {
		if c.Name == name {
			return c, nil
		}
	}

	err = fmt.Errorf("failed to get random weighted town category")
	return Category{}, err
}
