package food

import (
	"github.com/ironarachne/world/pkg/words"
)

// SimplifiedStyle is a simplified version of style for display
type SimplifiedStyle struct {
	CommonDishes []string `json:"common_dishes"`
	EatingTraits []string `json:"eating_traits"`
	Breads       []string `json:"breads"`
}

// Simplify returns a simplified style
func (style Style) Simplify() SimplifiedStyle {
	return SimplifiedStyle{
		CommonDishes: style.CommonDishes,
		EatingTraits: style.EatingTraits,
		Breads:       style.Breads,
	}
}

// Describe returns a description of a drink
func (drink Drink) Describe() string {
	description := drink.Strength + " " + drink.Name + " made with " + words.CombinePhrases(drink.Ingredients)

	return description
}
