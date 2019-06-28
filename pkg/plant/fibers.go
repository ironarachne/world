package plant

import (
	"math/rand"
)

func getRandomFiber(from []Plant) Plant {
	var fibers []Plant

	for _, p := range from {
		if p.IsFiber {
			fibers = append(fibers, p)
		}
	}

	return fibers[rand.Intn(len(fibers))]
}

func getFibers() []Plant {
	fibers := []Plant{
		Plant{
			Name:           "cotton",
			PluralName:     "cotton",
			IsFiber:        true,
			IsFruit:        false,
			IsGrain:        false,
			IsHerb:         false,
			IsMedicine:     false,
			IsNut:          false,
			IsRoot:         false,
			IsSpice:        false,
			IsToxic:        false,
			IsVegetable:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
		},
		Plant{
			Name:           "flax",
			PluralName:     "flax",
			IsFiber:        true,
			IsFruit:        false,
			IsGrain:        false,
			IsHerb:         false,
			IsMedicine:     false,
			IsNut:          false,
			IsRoot:         false,
			IsSpice:        false,
			IsToxic:        false,
			IsVegetable:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
		},
		Plant{
			Name:           "hemp",
			PluralName:     "hemp",
			IsFiber:        true,
			IsFruit:        false,
			IsGrain:        false,
			IsHerb:         false,
			IsMedicine:     false,
			IsNut:          false,
			IsRoot:         false,
			IsSpice:        false,
			IsToxic:        false,
			IsVegetable:    false,
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 9,
		},
		Plant{
			Name:           "coir",
			PluralName:     "coir",
			IsFiber:        true,
			IsFruit:        false,
			IsGrain:        false,
			IsHerb:         false,
			IsMedicine:     false,
			IsNut:          false,
			IsRoot:         false,
			IsSpice:        false,
			IsToxic:        false,
			IsVegetable:    false,
			MinHumidity:    6,
			MaxHumidity:    10,
			MinTemperature: 8,
			MaxTemperature: 10,
		},
		Plant{
			Name:           "papyrus",
			PluralName:     "papyrus",
			IsFiber:        true,
			IsFruit:        false,
			IsGrain:        false,
			IsHerb:         false,
			IsMedicine:     false,
			IsNut:          false,
			IsRoot:         false,
			IsSpice:        false,
			IsToxic:        false,
			IsVegetable:    false,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 9,
		},
		Plant{
			Name:           "jute",
			PluralName:     "jute",
			IsFiber:        true,
			IsFruit:        false,
			IsGrain:        false,
			IsHerb:         false,
			IsMedicine:     false,
			IsNut:          false,
			IsRoot:         false,
			IsSpice:        false,
			IsToxic:        false,
			IsVegetable:    false,
			MinHumidity:    3,
			MaxHumidity:    10,
			MinTemperature: 2,
			MaxTemperature: 9,
		},
	}

	return fibers
}
