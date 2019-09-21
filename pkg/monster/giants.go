package monster

import (
	"github.com/ironarachne/world/pkg/dice"
	"github.com/ironarachne/world/pkg/size"
)

func getAllGiants() []Monster {
	return []Monster{
		{
			Name:                 "cyclops",
			PluralName:           "cyclopses",
			SizeCategory:         size.GetCategoryByName("huge"),
			Type:                 "giant",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				return 1
			},
		},
		{
			Name:                 "ettin",
			PluralName:           "ettins",
			SizeCategory:         size.GetCategoryByName("large"),
			Type:                 "giant",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				d := dice.Dice{Number: 1, Sides: 2}
				num := dice.Roll(d)
				return num
			},
		},
		{
			Name:                 "giant",
			PluralName:           "giants",
			SizeCategory:         size.GetCategoryByName("huge"),
			Type:                 "giant",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				return 1
			},
		},
		{
			Name:                 "ogre",
			PluralName:           "ogres",
			SizeCategory:         size.GetCategoryByName("large"),
			Type:                 "giant",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				return 1
			},
		},
	}
}
