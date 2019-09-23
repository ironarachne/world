package organization

import (
	"math/rand"
)

// SizeClass is a size range
type SizeClass struct {
	Name    string
	MinSize int
	MaxSize int
}

func getAllSizeClasses() []SizeClass {
	classes := []SizeClass{
		{
			Name:    "tiny",
			MinSize: 2,
			MaxSize: 5,
		},
		{
			Name:    "small",
			MinSize: 6,
			MaxSize: 20,
		},
		{
			Name:    "medium",
			MinSize: 21,
			MaxSize: 50,
		},
		{
			Name:    "big",
			MinSize: 51,
			MaxSize: 100,
		},
		{
			Name:    "huge",
			MinSize: 101,
			MaxSize: 500,
		},
		{
			Name:    "massive",
			MinSize: 501,
			MaxSize: 1000,
		},
	}

	return classes
}

func getRandomSizeClass(min int, max int) SizeClass {
	classes := getAllSizeClasses()

	size := rand.Intn(max-min) + min

	for _, c := range classes {
		if size <= c.MaxSize && size >= c.MinSize {
			return c
		}
	}

	return classes[0]
}
