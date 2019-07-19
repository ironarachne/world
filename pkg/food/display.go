package food

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
