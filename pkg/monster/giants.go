package monster

import "github.com/ironarachne/world/pkg/species"

func getAllGiants() []species.Species {
	return []species.Species{
		{
			Name:           "cyclops",
			PluralName:     "cyclopses",
			MinTemperature: 1,
			MaxTemperature: 10,
			MinHumidity:    1,
			MaxHumidity:    10,
			Tags: []string{
				"giant",
			},
		},
		{
			Name:           "ettin",
			PluralName:     "ettins",
			MinTemperature: 1,
			MaxTemperature: 10,
			MinHumidity:    1,
			MaxHumidity:    10,
			Tags: []string{
				"giant",
			},
		},
		{
			Name:           "giant",
			PluralName:     "giants",
			MinTemperature: 1,
			MaxTemperature: 10,
			MinHumidity:    1,
			MaxHumidity:    10,
			Tags: []string{
				"giant",
			},
		},
		{
			Name:           "ogre",
			PluralName:     "ogres",
			MinTemperature: 1,
			MaxTemperature: 10,
			MinHumidity:    1,
			MaxHumidity:    10,
			Tags: []string{
				"giant",
			},
		},
	}
}
