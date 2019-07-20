package drink

import "github.com/ironarachne/world/pkg/resource"

// Pattern is a pattern for making a drink
type Pattern struct {
	Name         string
	RequiredBase string
	BaseStrength int
	Descriptors  []string
}

func getAllPatterns() []Pattern {
	patterns := []Pattern{
		{
			Name:         "ale",
			RequiredBase: "grain",
			BaseStrength: 1,
			Descriptors: []string{
				"bright gold in color",
				"bright yellow in color",
				"fizzy",
				"foamy",
				"golden",
				"pale",
			},
		},
		{
			Name:         "beer",
			RequiredBase: "grain",
			BaseStrength: 1,
			Descriptors: []string{
				"black",
				"brown",
				"dark",
				"fizzy",
				"flat",
				"foamy",
				"golden",
				"pale",
			},
		},
		{
			Name:         "brandy",
			RequiredBase: "fruit",
			BaseStrength: 3,
			Descriptors: []string{
				"amber in hue",
				"clear",
				"deep red in hue",
				"golden-hued",
				"orange in color",
				"pale red in color",
			},
		},
		{
			Name:         "liquor",
			RequiredBase: "grain",
			BaseStrength: 3,
			Descriptors: []string{
				"brown",
				"clear",
				"cloudy",
				"golden",
				"pale green in color",
				"pale",
				"tan-colored",
				"white",
				"yellow-hued",
			},
		},
		{
			Name:         "wine",
			RequiredBase: "fruit",
			BaseStrength: 2,
			Descriptors: []string{
				"dark red in color",
				"deep yellow",
				"light pink",
				"light red in color",
				"pale yellow",
				"red",
			},
		},
	}

	return patterns
}

func getValidPatterns(resources []resource.Resource) []Pattern {
	var patterns []Pattern

	all := getAllPatterns()

	for _, p := range all {
		if resource.IsTypeInResources(p.RequiredBase, resources) {
			patterns = append(patterns, p)
		}
	}

	return patterns
}
