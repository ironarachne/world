package monster

import "github.com/ironarachne/world/pkg/species"

func getAllDraconids() []species.Species {
	return []species.Species{
		{
			Name:           "dragon",
			PluralName:     "dragons",
			MinTemperature: 3,
			MaxTemperature: 10,
			MinHumidity:    1,
			MaxHumidity:    10,
			Tags: []string{
				"draconid",
			},
		},
		{
			Name:           "drake",
			PluralName:     "drakes",
			MinTemperature: 3,
			MaxTemperature: 10,
			MinHumidity:    1,
			MaxHumidity:    10,
			Tags: []string{
				"draconid",
			},
		},
		{
			Name:           "wyvern",
			PluralName:     "wyverns",
			MinTemperature: 3,
			MaxTemperature: 10,
			MinHumidity:    1,
			MaxHumidity:    10,
			Tags: []string{
				"draconid",
			},
		},
	}
}
