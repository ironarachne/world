package monster

import (
	"github.com/ironarachne/world/pkg/dice"
	"github.com/ironarachne/world/pkg/size"
)

func getAllHumanoids() []Monster {
	return []Monster{
		{
			Name:                 "bugbear",
			PluralName:           "bugbears",
			SizeCategory:         size.GetCategoryByName("medium"),
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				d := dice.Dice{Number: 1, Sides: 4}
				num := dice.Roll(d)
				return num
			},
		},
		{
			Name:                 "bullywug",
			PluralName:           "bullywugs",
			SizeCategory:         size.GetCategoryByName("medium"),
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				d := dice.Dice{Number: 1, Sides: 4}
				num := dice.Roll(d)
				return num
			},
		},
		{
			Name:                 "goblin",
			PluralName:           "goblins",
			SizeCategory:         size.GetCategoryByName("small"),
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				d := dice.Dice{Number: 2, Sides: 4}
				num := dice.Roll(d)
				return num
			},
		},
		{
			Name:                 "gorgon",
			PluralName:           "gorgons",
			SizeCategory:         size.GetCategoryByName("medium"),
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				d := dice.Dice{Number: 1, Sides: 3}
				num := dice.Roll(d)
				return num
			},
		},
		{
			Name:                 "hobgoblin",
			PluralName:           "hobgoblins",
			SizeCategory:         size.GetCategoryByName("medium"),
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				d := dice.Dice{Number: 1, Sides: 4}
				num := dice.Roll(d)
				return num
			},
		},
		{
			Name:                 "kobold",
			PluralName:           "kobolds",
			SizeCategory:         size.GetCategoryByName("small"),
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				d := dice.Dice{Number: 1, Sides: 6}
				num := dice.Roll(d)
				return num
			},
		},
		{
			Name:                 "orc",
			PluralName:           "orcs",
			SizeCategory:         size.GetCategoryByName("medium"),
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				d := dice.Dice{Number: 1, Sides: 6}
				num := dice.Roll(d)
				return num
			},
		},
		{
			Name:                 "troglodyte",
			PluralName:           "troglodytes",
			SizeCategory:         size.GetCategoryByName("medium"),
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				d := dice.Dice{Number: 1, Sides: 6}
				num := dice.Roll(d)
				return num
			},
		},
		{
			Name:                 "troll",
			PluralName:           "trolls",
			SizeCategory:         size.GetCategoryByName("medium"),
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				d := dice.Dice{Number: 1, Sides: 4}
				num := dice.Roll(d)
				return num
			},
		},
	}
}
