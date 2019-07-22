package resource

import "github.com/ironarachne/world/pkg/profession"

func getSmithing() []Pattern {
	producer := profession.ByName("blacksmith")

	patterns := []Pattern{
		{
			Name:        "fork",
			Description: "a fork",
			Tags: []string{
				"eating utensil",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.Origin}} fork",
				},
			},
		},
		{
			Name:        "hammer",
			Description: "a hammer",
			Tags: []string{
				"tool",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.Origin}} hammer",
				},
			},
		},
		{
			Name:        "knife",
			Description: "a knife",
			Tags: []string{
				"eating utensil",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.Origin}} knife",
				},
			},
		},
		{
			Name:        "hand saw",
			Description: "a hand saw",
			Tags: []string{
				"tool",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.Origin}} hand saw",
				},
			},
		},
		{
			Name:        "metal bar",
			Description: "a metal bar",
			Tags: []string{
				"metal bar",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "refined metal ore",
					DescriptionTemplate: "{{.Resource.Origin}} metal bar",
				},
			},
		},
		{
			Name:        "metal ingot",
			Description: "a metal ingot",
			Tags: []string{
				"metal ingot",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "refined metal ore",
					DescriptionTemplate: "{{.Resource.Origin}} metal ingot",
				},
			},
		},
		{
			Name:        "spoon",
			Description: "a spoon",
			Tags: []string{
				"eating utensil",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.Origin}} spoon",
				},
			},
		},
	}

	farrier := profession.ByName("farrier")

	farrierPatterns := []Pattern{
		{
			Name:        "horseshoe",
			Description: "a horseshoe",
			Tags: []string{
				"horseshoe",
			},
			Profession: farrier,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.Origin}} horseshoe",
				},
			},
		},
	}

	patterns = append(patterns, farrierPatterns...)

	return patterns
}
