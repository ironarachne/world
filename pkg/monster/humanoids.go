package monster

import "github.com/ironarachne/world/pkg/random"

func getAllHumanoids() []Monster {
	return []Monster{
		Monster{
			Name:                 "bugbear",
			PluralName:           "bugbears",
			SizeCategory:         "medium",
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				num := random.DiceTotal(1, 4)
				return num
			},
		},
		Monster{
			Name:                 "bullywug",
			PluralName:           "bullywugs",
			SizeCategory:         "medium",
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				num := random.DiceTotal(1, 4)
				return num
			},
		},
		Monster{
			Name:                 "goblin",
			PluralName:           "goblins",
			SizeCategory:         "small",
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				num := random.DiceTotal(2, 4)
				return num
			},
		},
		Monster{
			Name:                 "gorgon",
			PluralName:           "gorgons",
			SizeCategory:         "medium",
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				num := random.DiceTotal(1, 3)
				return num
			},
		},
		Monster{
			Name:                 "hobgoblin",
			PluralName:           "hobgoblins",
			SizeCategory:         "medium",
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				num := random.DiceTotal(1, 4)
				return num
			},
		},
		Monster{
			Name:                 "kobold",
			PluralName:           "kobolds",
			SizeCategory:         "small",
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				num := random.DiceTotal(1, 6)
				return num
			},
		},
		Monster{
			Name:                 "orc",
			PluralName:           "orcs",
			SizeCategory:         "medium",
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				num := random.DiceTotal(1, 6)
				return num
			},
		},
		Monster{
			Name:                 "troglodyte",
			PluralName:           "troglodytes",
			SizeCategory:         "medium",
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				num := random.DiceTotal(1, 6)
				return num
			},
		},
		Monster{
			Name:                 "troll",
			PluralName:           "trolls",
			SizeCategory:         "medium",
			Type:                 "humanoid",
			IdealTemperature:     5,
			TemperatureTolerance: 5,
			NumAppearing: func() int {
				num := random.DiceTotal(1, 4)
				return num
			},
		},
	}
}
