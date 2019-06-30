package buildings

// SimplifiedBuildingStyle is a simplified version of a building style for display
type SimplifiedBuildingStyle struct {
	Description string `json:"description"`
}

// Describe describes a building style
func (style BuildingStyle) Describe() string {
	description := "Buildings have walls made of " + style.WallStyle + ", and roofs are " + style.RoofStyle + ". "
	description += "Doors are " + style.DoorStyle + ". Windows are " + style.WindowStyle + ". "
	description += "For decoration, " + style.Decorations + ". "

	return description
}

// Simplify turns a BuildingStyle into a SimplifiedBuildingStyle
func (style BuildingStyle) Simplify() SimplifiedBuildingStyle {
	return SimplifiedBuildingStyle{
		Description: style.Describe(),
	}
}
