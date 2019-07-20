package organization

import "math/rand"

// SizeClass is a size range
type SizeClass struct {
	Name    string
	MinSize int
	MaxSize int
}

func getAllSizeClasses() []SizeClass {
	classes := []SizeClass{
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
	}

	return classes
}

func getRandomSizeClass() SizeClass {
	classes := getAllSizeClasses()

	return classes[rand.Intn(len(classes)-1)]
}
