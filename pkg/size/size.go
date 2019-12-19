/*
Package size implements a system for size categories of creatures
*/
package size

// Category is a general size category for creatures
type Category struct {
	Name       string `json:"name" db:"name"`
	Level      int    `json:"level" db:"level"`
	BaseHeight int    `json:"base_height" db:"base_height"` // Base height in inches
	BaseWeight int    `json:"base_weight" db:"base_weight"` // Base weight in pounds
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
		{
			Name:       "fine",
			Level:      1,
			BaseHeight: 3,
			BaseWeight: 1,
		},
		{
			Name:       "diminutive",
			Level:      2,
			BaseHeight: 9,
			BaseWeight: 1,
		},
		{
			Name:       "tiny",
			Level:      3,
			BaseHeight: 20,
			BaseWeight: 6,
		},
		{
			Name:       "small",
			Level:      4,
			BaseHeight: 36,
			BaseWeight: 30,
		},
		{
			Name:       "medium",
			Level:      5,
			BaseHeight: 72,
			BaseWeight: 200,
		},
		{
			Name:       "large",
			Level:      6,
			BaseHeight: 144,
			BaseWeight: 1000,
		},
		{
			Name:       "huge",
			Level:      7,
			BaseHeight: 288,
			BaseWeight: 14000,
		},
		{
			Name:       "gargantuan",
			Level:      8,
			BaseHeight: 576,
			BaseWeight: 142000,
		},
		{
			Name:       "colossal",
			Level:      9,
			BaseHeight: 768,
			BaseWeight: 250000,
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
