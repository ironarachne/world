package tree

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/slices"
)

// Tree is a tree
type Tree struct {
	Name           string
	PluralName     string
	MinHumidity    int
	MaxHumidity    int
	MinTemperature int
	MaxTemperature int
	Resources      []resource.Resource
	Tags           []string
}

// ByTag returns all trees that match the tag given
func ByTag(tag string, trees []Tree) []Tree {
	var result []Tree

	for _, t := range trees {
		if slices.StringIn(tag, t.Tags) {
			result = append(result, t)
		}
	}

	return result
}

// ExcludeTag removes all trees that match the tag from the given slice
func ExcludeTag(tag string, trees []Tree) []Tree {
	var result []Tree

	for _, t := range trees {
		if !slices.StringIn(tag, t.Tags) {
			result = append(result, t)
		}
	}

	return result
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

// RandomSubset returns a random subset of plants
func RandomSubset(amount int, from []Tree) []Tree {
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
