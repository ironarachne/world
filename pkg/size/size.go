package size

// Category is a general size category for creatures
type Category struct {
	Name  string
	Level int
}

// Larger returns a category larger than the current one, or the current one if it's the biggest
func (category Category) Larger() Category {
	categories := getAllSizeCategories()

	for _, c := range categories {
		if c.Level == category.Level+1 {
			return c
		}
	}

	return category
}

// Smaller returns a category smaller than the current one, or the current one if it's the smallest
func (category Category) Smaller() Category {
	categories := getAllSizeCategories()

	for _, c := range categories {
		if c.Level == category.Level-1 {
			return c
		}
	}

	return category
}

func getAllSizeCategories() []Category {
	return []Category{
		Category{
			Name:  "tiny",
			Level: 1,
		},
		Category{
			Name:  "small",
			Level: 2,
		},
		Category{
			Name:  "medium",
			Level: 3,
		},
		Category{
			Name:  "large",
			Level: 4,
		},
		Category{
			Name:  "huge",
			Level: 5,
		},
		Category{
			Name:  "gargantuan",
			Level: 6,
		},
	}
}

// GetCategoryByName returns a size category based on name
func GetCategoryByName(name string) Category {
	categories := getAllSizeCategories()

	for _, c := range categories {
		if c.Name == name {
			return c
		}
	}

	panic("Couldn't find size category for name " + name)
}
