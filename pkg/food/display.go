package food

// SimplifiedStyle is a simplified version of style for display
type SimplifiedStyle struct {
	CommonDesserts   []string `json:"common_desserts"`
	CommonMainDishes []string `json:"common_main_dishes"`
	EatingTraits     []string `json:"eating_traits"`
	Breads           []string `json:"breads"`
}

// Simplify returns a simplified style
func (style Style) Simplify() SimplifiedStyle {
	return SimplifiedStyle{
		CommonDesserts:   style.CommonDesserts,
		CommonMainDishes: style.CommonMainDishes,
		EatingTraits:     style.EatingTraits,
		Breads:           style.Breads,
	}
}
