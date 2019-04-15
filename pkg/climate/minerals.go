package climate

import "math/rand"

// Mineral is a mineral
type Mineral struct {
	Name         string
	PluralName   string
	Hardness     int
	HasOre       bool
	IsEdible     bool
	IsGem        bool
	IsMetal      bool
	IsPrecious   bool
	IsStone      bool
	Malleability int
}

func getRandomMinerals(amount int, from []Mineral) []Mineral {
	var mineral Mineral

	minerals := []Mineral{}

	if amount > len(from) {
		amount = len(from)
	}

	for i := 0; i < amount; i++ {
		mineral = from[rand.Intn(len(from)-1)]
		if !isMineralInSlice(mineral, minerals) {
			minerals = append(minerals, mineral)
		}
	}

	return minerals
}

func isMineralInSlice(mineral Mineral, minerals []Mineral) bool {
	isIt := false
	for _, m := range minerals {
		if m.Name == mineral.Name {
			isIt = true
		}
	}

	return isIt
}

func getAllMinerals() []Mineral {
	var minerals []Mineral

	gems := getGems()
	metal := getMetals()
	other := getOtherMinerals()
	stone := getStones()

	minerals = append(minerals, gems...)
	minerals = append(minerals, metal...)
	minerals = append(minerals, other...)
	minerals = append(minerals, stone...)

	return minerals
}

func getCommonMetals() []Mineral {
	var commonMetals []Mineral

	metals := getMetals()

	for _, m := range metals {
		if !m.IsPrecious {
			commonMetals = append(commonMetals, m)
		}
	}

	return commonMetals
}

func getPreciousMetals() []Mineral {
	var preciousMetals []Mineral

	metals := getMetals()

	for _, m := range metals {
		if m.IsPrecious {
			preciousMetals = append(preciousMetals, m)
		}
	}

	return preciousMetals
}

func getGems() []Mineral {
	gems := []Mineral{
		Mineral{
			Name:         "diamond",
			PluralName:   "diamonds",
			Hardness:     10,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        true,
			IsMetal:      false,
			IsPrecious:   true,
			IsStone:      false,
			Malleability: 1,
		},
		Mineral{
			Name:         "garnet",
			PluralName:   "garnets",
			Hardness:     6,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        true,
			IsMetal:      false,
			IsPrecious:   true,
			IsStone:      false,
			Malleability: 3,
		},
		Mineral{
			Name:         "ruby",
			PluralName:   "rubies",
			Hardness:     6,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        true,
			IsMetal:      false,
			IsPrecious:   true,
			IsStone:      false,
			Malleability: 3,
		},
		Mineral{
			Name:         "sapphire",
			PluralName:   "sapphires",
			Hardness:     6,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        true,
			IsMetal:      false,
			IsPrecious:   true,
			IsStone:      false,
			Malleability: 3,
		},
		Mineral{
			Name:         "amethyst",
			PluralName:   "amethysts",
			Hardness:     6,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        true,
			IsMetal:      false,
			IsPrecious:   true,
			IsStone:      false,
			Malleability: 3,
		},
		Mineral{
			Name:         "agate",
			PluralName:   "agates",
			Hardness:     6,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        true,
			IsMetal:      false,
			IsPrecious:   true,
			IsStone:      false,
			Malleability: 3,
		},
		Mineral{
			Name:         "jade",
			PluralName:   "jade",
			Hardness:     6,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        true,
			IsMetal:      false,
			IsPrecious:   true,
			IsStone:      false,
			Malleability: 3,
		},
		Mineral{
			Name:         "jasper",
			PluralName:   "jaspers",
			Hardness:     6,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        true,
			IsMetal:      false,
			IsPrecious:   true,
			IsStone:      false,
			Malleability: 3,
		},
		Mineral{
			Name:         "moonstone",
			PluralName:   "moonstones",
			Hardness:     6,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        true,
			IsMetal:      false,
			IsPrecious:   true,
			IsStone:      false,
			Malleability: 3,
		},
		Mineral{
			Name:         "opal",
			PluralName:   "opals",
			Hardness:     6,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        true,
			IsMetal:      false,
			IsPrecious:   true,
			IsStone:      false,
			Malleability: 3,
		},
		Mineral{
			Name:         "peridot",
			PluralName:   "peridots",
			Hardness:     6,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        true,
			IsMetal:      false,
			IsPrecious:   true,
			IsStone:      false,
			Malleability: 3,
		},
		Mineral{
			Name:         "topaz",
			PluralName:   "topazes",
			Hardness:     6,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        true,
			IsMetal:      false,
			IsPrecious:   true,
			IsStone:      false,
			Malleability: 3,
		},
	}

	return gems
}

func getMetals() []Mineral {
	metals := []Mineral{
		Mineral{
			Name:         "copper",
			PluralName:   "copper",
			Hardness:     5,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        false,
			IsMetal:      true,
			IsPrecious:   true,
			IsStone:      false,
			Malleability: 5,
		},
		Mineral{
			Name:         "gold",
			PluralName:   "gold",
			Hardness:     2,
			HasOre:       true,
			IsEdible:     true,
			IsGem:        false,
			IsMetal:      true,
			IsPrecious:   true,
			IsStone:      false,
			Malleability: 8,
		},
		Mineral{
			Name:         "iron",
			PluralName:   "iron",
			Hardness:     8,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        false,
			IsMetal:      true,
			IsPrecious:   false,
			IsStone:      false,
			Malleability: 4,
		},
		Mineral{
			Name:         "tin",
			PluralName:   "tin",
			Hardness:     3,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        false,
			IsMetal:      true,
			IsPrecious:   false,
			IsStone:      false,
			Malleability: 8,
		},
		Mineral{
			Name:         "lead",
			PluralName:   "lead",
			Hardness:     1,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        false,
			IsMetal:      true,
			IsPrecious:   false,
			IsStone:      false,
			Malleability: 9,
		},
		Mineral{
			Name:         "nickel",
			PluralName:   "nickel",
			Hardness:     3,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        false,
			IsMetal:      true,
			IsPrecious:   false,
			IsStone:      false,
			Malleability: 4,
		},
		Mineral{
			Name:         "silver",
			PluralName:   "silver",
			Hardness:     2,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        false,
			IsMetal:      true,
			IsPrecious:   true,
			IsStone:      false,
			Malleability: 7,
		},
	}

	return metals
}

func getOtherMinerals() []Mineral {
	others := []Mineral{
		Mineral{
			Name:         "coal",
			PluralName:   "coal",
			Hardness:     1,
			HasOre:       true,
			IsEdible:     false,
			IsGem:        false,
			IsMetal:      false,
			IsPrecious:   false,
			IsStone:      false,
			Malleability: 4,
		},
		Mineral{
			Name:         "salt",
			PluralName:   "salt",
			Hardness:     1,
			HasOre:       false,
			IsEdible:     true,
			IsGem:        false,
			IsMetal:      false,
			IsPrecious:   false,
			IsStone:      false,
			Malleability: 1,
		},
	}

	return others
}

func getStones() []Mineral {
	stones := []Mineral{
		Mineral{
			Name:         "granite",
			PluralName:   "granite",
			Hardness:     6,
			HasOre:       false,
			IsEdible:     false,
			IsGem:        false,
			IsMetal:      false,
			IsPrecious:   false,
			IsStone:      true,
			Malleability: 2,
		},
		Mineral{
			Name:         "limestone",
			PluralName:   "limestone",
			Hardness:     4,
			HasOre:       false,
			IsEdible:     false,
			IsGem:        false,
			IsMetal:      false,
			IsPrecious:   false,
			IsStone:      true,
			Malleability: 4,
		},
		Mineral{
			Name:         "marble",
			PluralName:   "marble",
			Hardness:     5,
			HasOre:       false,
			IsEdible:     false,
			IsGem:        false,
			IsMetal:      false,
			IsPrecious:   false,
			IsStone:      true,
			Malleability: 2,
		},
		Mineral{
			Name:         "sandstone",
			PluralName:   "sandstone",
			Hardness:     4,
			HasOre:       false,
			IsEdible:     false,
			IsGem:        false,
			IsMetal:      false,
			IsPrecious:   false,
			IsStone:      true,
			Malleability: 4,
		},
	}

	return stones
}
