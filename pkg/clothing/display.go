package clothing

// SimplifiedStyle is a simplified version of clothing style for display
type SimplifiedStyle struct {
	CommonFemaleItems []string `json:"common_female_items"`
	CommonMaleItems   []string `json:"common_male_items"`
	CommonJewelry     []string `json:"common_jewelry"`
	CommonColors      []string `json:"common_colors"`
	CommonMaterials   []string `json:"common_materials"`
	DecorativeStyle   string   `json:"decorative_style"`
}

// Simplify returns a simplified style for display
func (style Style) Simplify() SimplifiedStyle {
	female := []string{}
	male := []string{}

	for _, f := range style.CommonFemaleItems {
		female = append(female, f.Name)
	}

	for _, m := range style.CommonMaleItems {
		male = append(male, m.Name)
	}

	return SimplifiedStyle{
		CommonFemaleItems: female,
		CommonMaleItems:   male,
		CommonJewelry:     style.CommonJewelry,
		CommonColors:      style.CommonColors,
		CommonMaterials:   style.CommonMaterials,
		DecorativeStyle:   style.DecorativeStyle,
	}
}
