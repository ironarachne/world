package race

import "github.com/ironarachne/world/pkg/trait"

// GetCommonPossibleTraits returns trait templates that can be used for all races
func GetCommonPossibleTraits() []trait.Template {
	templates := []trait.Template{
		{
			Name: "scar",
			PossibleValues: []string{
				"a faded scar",
				"an angry scar over one eye",
				"a cross-shaped scar on the left cheek",
				"ragged scars",
				"several light scars",
				"several battle scars",
				"some battle scars",
			},
			PossibleDescriptors: []string{
				"{{.Value}}",
			},
			Tags: []string{
				"appearance",
				"physical",
				"scar",
			},
		},
		{
			Name: "expression",
			PossibleValues: []string{
				"fierce",
				"rosy",
				"neutral",
				"muted",
				"serious",
				"faraway",
			},
			PossibleDescriptors: []string{
				"a {{.Value}} expression",
			},
			Tags: []string{
				"appearance",
				"physical",
				"expression",
			},
		},
		{
			Name: "cheeks",
			PossibleValues: []string{
				"rosy",
				"flushed",
				"ruddy",
				"pock-marked",
			},
			PossibleDescriptors: []string{
				"{{.Value}} cheeks",
			},
			Tags: []string{
				"appearance",
				"physical",
				"cheeks",
			},
		},
	}

	return templates
}
