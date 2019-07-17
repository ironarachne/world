package monster

// SimplifiedMonster is a monster simplified for display
type SimplifiedMonster struct {
	Name string
	PluralName string
	Description string
	NumAppearing int
	SizeCategory string
	Type string
}

// Simplify returns a simplified version of a monster
func (monster Monster) Simplify() SimplifiedMonster {
	numAppearing := monster.NumAppearing()

	sm := SimplifiedMonster{
		Name: monster.Name,
		PluralName: monster.PluralName,
		Description: monster.Description,
		NumAppearing: numAppearing,
		SizeCategory: monster.SizeCategory,
		Type: monster.Type,
	}

	return sm
}