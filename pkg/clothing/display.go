package clothing

import (
	"context"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/words"
)

// SimplifiedStyle is a simplified version of clothing style for display
type SimplifiedStyle struct {
	FemaleOutfit    []string `json:"female_outfit"`
	MaleOutfit      []string `json:"male_outfit"`
	CommonJewelry   []string `json:"common_jewelry"`
	CommonColors    []string `json:"common_colors"`
	DecorativeStyle string   `json:"decorative_style"`
}

// Describe creates a prose description for an item
func (item Item) Describe(ctx context.Context) string {
	description := ""

	if item.PrefixModifier != "" {
		if random.Intn(ctx, 100) > 50 {
			description += item.PrefixModifier + " " + item.MaterialType + " "
		} else {
			description += item.MaterialType + " " + item.PrefixModifier + " "
		}
	} else {
		description += item.MaterialType + " "
	}

	description += item.Name

	if item.SuffixModifier != "" {
		description += " " + item.SuffixModifier
	}

	return description
}

// Describe creates a prose description for a clothing style
func (style Style) Describe(ctx context.Context) string {
	female := []string{}
	male := []string{}

	for _, f := range style.FemaleOutfit {
		female = append(female, f.Describe(ctx))
	}

	for _, m := range style.MaleOutfit {
		male = append(male, m.Describe(ctx))
	}

	description := "The male outfit is: " + words.CombinePhrases(male) + ", and the female outfit is: " + words.CombinePhrases(female) + ". "
	description += "Common jewelry is " + words.CombinePhrases(style.CommonJewelry) + ". "
	description += "Common colors are " + words.CombinePhrases(style.CommonColors) + ". "
	description += "A common decorative element is " + style.DecorativeStyle + "."

	return description
}

// Simplify returns a simplified style for display
func (style Style) Simplify(ctx context.Context) SimplifiedStyle {
	female := []string{}
	male := []string{}

	for _, f := range style.FemaleOutfit {
		female = append(female, f.Describe(ctx))
	}

	for _, m := range style.MaleOutfit {
		male = append(male, m.Describe(ctx))
	}

	return SimplifiedStyle{
		FemaleOutfit:    female,
		MaleOutfit:      male,
		CommonJewelry:   style.CommonJewelry,
		CommonColors:    style.CommonColors,
		DecorativeStyle: style.DecorativeStyle,
	}
}
