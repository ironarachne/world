package monster

import "github.com/ironarachne/world/pkg/random"

func getAllGiants() []Monster {
	return []Monster{
		Monster{
			Name:                 "cyclops",
			PluralName:           "cyclopses",
			SizeCategory:         "huge",
			Type:                 "giant",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				return 1
			},
		},
		Monster{
			Name:                 "ettin",
			PluralName:           "ettins",
			SizeCategory:         "large",
			Type:                 "giant",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				num := random.DiceTotal(1, 2)
				return num
			},
		},
		Monster{
			Name:                 "giant",
			PluralName:           "giants",
			SizeCategory:         "huge",
			Type:                 "giant",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				return 1
			},
		},
		Monster{
			Name:                 "ogre",
			PluralName:           "ogres",
			SizeCategory:         "large",
			Type:                 "giant",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				return 1
			},
		},
	}
}
