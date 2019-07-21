package resource

import "github.com/ironarachne/world/pkg/profession"

func getSmithing() []Pattern {
	producer := profession.ByName("blacksmith")

	patterns := []Pattern{
		{
			Name: "fork",
			Description: "a fork",
			Type: "eating utensil",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "metal",
					DescriptionTemplate: "{{.Resource.Name}} fork",
				},
			},
		},
		{
			Name: "hammer",
			Description: "a hammer",
			Type: "tool",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "metal",
					DescriptionTemplate: "{{.Resource.Name}} hammer",
				},
			},
		},
		{
			Name: "knife",
			Description: "a knife",
			Type: "eating utensil",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "metal",
					DescriptionTemplate: "{{.Resource.Name}} knife",
				},
			},
		},
		{
			Name: "hand saw",
			Description: "a hand saw",
			Type: "tool",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "metal",
					DescriptionTemplate: "{{.Resource.Name}} hand saw",
				},
			},
		},
		{
			Name: "spoon",
			Description: "a spoon",
			Type: "eating utensil",
			Profession: producer,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "metal",
					DescriptionTemplate: "{{.Resource.Name}} spoon",
				},
			},
		},
	}

	farrier := profession.ByName("farrier")

	farrierPatterns := []Pattern{
		{
			Name: "horseshoe",
			Description: "a horseshoe",
			Type: "horseshoe",
			Profession: farrier,
			Slots: []Slot{
				{
					Name: "body",
					RequiredType: "metal",
					DescriptionTemplate: "{{.Resource.Name}} horseshoe",
				},
			},
		},
	}

	patterns = append(patterns, farrierPatterns...)

	return patterns
}
