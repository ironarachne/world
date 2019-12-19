package resource

import "github.com/ironarachne/world/pkg/profession"

func getSmithing() []Pattern {
	producer, _ := profession.ByName("blacksmith")

	patterns := []Pattern{
		{
			Name:         "fork",
			NameTemplate: "{{.MainMaterial}} fork",
			Description:  "a fork",
			Tags: []string{
				"eating utensil",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} fork",
				},
			},
			Value: 1,
		},
		{
			Name:         "hammer",
			NameTemplate: "{{.MainMaterial}} hammer",
			Description:  "a hammer",
			Tags: []string{
				"tool",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} hammer",
				},
			},
			Value: 1,
		},
		{
			Name:         "knife",
			NameTemplate: "{{.MainMaterial}} knife",
			Description:  "a knife",
			Tags: []string{
				"eating utensil",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} knife",
				},
			},
			Value: 1,
		},
		{
			Name:         "hand saw",
			NameTemplate: "{{.MainMaterial}} hand saw",
			Description:  "a hand saw",
			Tags: []string{
				"tool",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} hand saw",
				},
			},
			Value: 1,
		},
		{
			Name:         "spoon",
			NameTemplate: "{{.MainMaterial}} spoon",
			Description:  "a spoon",
			Tags: []string{
				"eating utensil",
			},
			Profession: producer,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} spoon",
				},
			},
			Value: 1,
		},
	}

	farrier, _ := profession.ByName("farrier")

	farrierPatterns := []Pattern{
		{
			Name:         "horseshoe",
			NameTemplate: "{{.MainMaterial}} horseshoe",
			Description:  "a horseshoe",
			Tags: []string{
				"horseshoe",
			},
			Profession: farrier,
			Slots: []Slot{
				{
					Name:                "body",
					RequiredTag:         "metal bar",
					DescriptionTemplate: "{{.Resource.MainMaterial}} horseshoe",
				},
			},
			Value: 1,
		},
	}

	patterns = append(patterns, farrierPatterns...)

	return patterns
}
