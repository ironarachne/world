package monster

// SimplifiedMonster is a monster simplified for display
type SimplifiedMonster struct {
	Name         string `json:"name"`
	PluralName   string `json:"plural_name"`
	Description  string `json:"description"`
	NumAppearing int    `json:"number_appearing"`
	SizeCategory string `json:"size_category"`
	Type         string `json:"type"`
}

// Simplify returns a simplified version of a monster
func (monster Monster) Simplify() SimplifiedMonster {
	numAppearing := monster.NumAppearing()

	sm := SimplifiedMonster{
		Name:         monster.Name,
		PluralName:   monster.PluralName,
		Description:  monster.Description,
		NumAppearing: numAppearing,
		SizeCategory: monster.SizeCategory,
		Type:         monster.Type,
	}

	return sm
}
