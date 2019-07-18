package monster

import (
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/size"
)

func getAllDraconids() []Monster {
	return []Monster{
		Monster{
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
		Monster{
			Name:                 "drake",
			PluralName:           "drakes",
			SizeCategory:         size.GetCategoryByName("medium"),
			Type:                 "draconid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				num := random.DiceTotal(1, 2)
				return num
			},
		},
		Monster{
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
