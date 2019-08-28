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
			Value: 1,
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
			Value: 1,
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
			Value: 1,
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
			Value: 1,
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
			Value: 1,
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
			Value: 1,
		},
	}

	patterns = append(patterns, farrierPatterns...)

	return patterns
}
