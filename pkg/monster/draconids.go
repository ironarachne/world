package monster

import (
	"github.com/ironarachne/world/pkg/dice"
	"github.com/ironarachne/world/pkg/size"
)

func getAllDraconids() []Monster {
	return []Monster{
		{
			Name:                 "dragon",
			PluralName:           "dragons",
			SizeCategory:         size.GetCategoryByName("gargantuan"),
			Type:                 "draconid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				return 1
			},
		},
		{
			Name:                 "drake",
			PluralName:           "drakes",
			SizeCategory:         size.GetCategoryByName("medium"),
			Type:                 "draconid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				d := dice.Dice{Number: 1, Sides: 2}
				num := dice.Roll(d)
				return num
			},
		},
		{
			Name:                 "wyvern",
			PluralName:           "wyverns",
			SizeCategory:         size.GetCategoryByName("large"),
			Type:                 "draconid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				return 1
			},
		},
	}
}
