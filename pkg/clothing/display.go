package clothing

import "math/rand"

// SimplifiedStyle is a simplified version of clothing style for display
type SimplifiedStyle struct {
	FemaleOutfit    []string `json:"female_outfit"`
	MaleOutfit      []string `json:"male_outfit"`
	CommonJewelry   []string `json:"common_jewelry"`
	CommonColors    []string `json:"common_colors"`
	DecorativeStyle string   `json:"decorative_style"`
}

// Describe creates a prose description for an item
func (item Item) Describe() string {
	description := ""

	if item.PrefixModifier != "" {
		if rand.Intn(100) > 50 {
			description += item.PrefixModifier + " " + item.Material + " "
		} else {
			description += item.Material + " " + item.PrefixModifier + " "
		}
	} else {
		description += item.Material + " "
	}

	description += item.Name

	if item.SuffixModifier != "" {
		description += " " + item.SuffixModifier
	}

	return description
}

// Simplify returns a simplified style for display
func (style Style) Simplify() SimplifiedStyle {
	female := []string{}
	male := []string{}

	for _, f := range style.FemaleOutfit {
		female = append(female, f.Describe())
	}

	for _, m := range style.MaleOutfit {
		male = append(male, m.Describe())
	}

	return SimplifiedStyle{
		FemaleOutfit:    female,
		MaleOutfit:      male,
		CommonJewelry:   style.CommonJewelry,
		CommonColors:    style.CommonColors,
		DecorativeStyle: style.DecorativeStyle,
	}
}
