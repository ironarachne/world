package climate

import "math/rand"

// Soil is a type of mineral
type Soil struct {
	Name               string
	UsedForPottery     bool
	IsSand             bool
	UsedForAgriculture bool
}

func (climate Climate) getFilteredSoils() []Soil {
	soils := []Soil{}
	agSoils := getAgriculturalSoils()
	clays := getClays()
	sands := getSands()

	if climate.Humidity < 2 && climate.Temperature >= 9 {
		soils = append(soils, sands...)
	} else {
		soils = append(soils, agSoils...)

		if climate.HasLakes {
			soils = append(soils, clays...)
		}
		if climate.HasOcean {
			soils = append(soils, sands...)
		}
	}

	return soils
}

func getRandomSoils(amount int, from []Soil) []Soil {
	var soil Soil

	soils := []Soil{}

	if amount > len(from) {
		amount = len(from)
	}

	if len(from) == 0 {
		return soils
	}

	for i := 0; i < amount; i++ {
		soil = from[rand.Intn(len(from)-1)]
		if !isSoilInSlice(soil, soils) {
			soils = append(soils, soil)
		}
	}

	return soils
}

func isSoilInSlice(soil Soil, soils []Soil) bool {
	isIt := false
	for _, a := range soils {
		if a.Name == soil.Name {
			isIt = true
		}
	}

	return isIt
}

func getAllSoils() []Soil {
	clays := getClays()
	sands := getSands()
	agSoils := getAgriculturalSoils()

	soils := append(clays, sands...)
	soils = append(soils, agSoils...)

	return soils
}

func getAgriculturalSoils() []Soil {
	soils := []Soil{
		Soil{
			Name:               "loam",
			UsedForPottery:     false,
			IsSand:             false,
			UsedForAgriculture: true,
		},
		Soil{
			Name:               "silt",
			UsedForPottery:     false,
			IsSand:             false,
			UsedForAgriculture: true,
		},
	}

	return soils
}

func getClays() []Soil {
	soils := []Soil{
		Soil{
			Name:               "earthenware",
			UsedForPottery:     true,
			IsSand:             false,
			UsedForAgriculture: false,
		},
		Soil{
			Name:               "porcelain",
			UsedForPottery:     true,
			IsSand:             false,
			UsedForAgriculture: false,
		},
		Soil{
			Name:               "stoneware",
			UsedForPottery:     true,
			IsSand:             false,
			UsedForAgriculture: false,
		},
		Soil{
			Name:               "terra cotta",
			UsedForPottery:     true,
			IsSand:             false,
			UsedForAgriculture: false,
		},
	}

	return soils
}

func getSands() []Soil {
	soils := []Soil{
		Soil{
			Name:               "black sand",
			UsedForPottery:     false,
			IsSand:             true,
			UsedForAgriculture: false,
		},
		Soil{
			Name:               "continental sand",
			UsedForPottery:     false,
			IsSand:             true,
			UsedForAgriculture: false,
		},
		Soil{
			Name:               "coral sand",
			UsedForPottery:     false,
			IsSand:             true,
			UsedForAgriculture: false,
		},
		Soil{
			Name:               "dune sand",
			UsedForPottery:     false,
			IsSand:             true,
			UsedForAgriculture: false,
		},
		Soil{
			Name:               "green sand",
			UsedForPottery:     false,
			IsSand:             true,
			UsedForAgriculture: false,
		},
		Soil{
			Name:               "red sand",
			UsedForPottery:     false,
			IsSand:             true,
			UsedForAgriculture: false,
		},
		Soil{
			Name:               "volcanic sand",
			UsedForPottery:     false,
			IsSand:             true,
			UsedForAgriculture: false,
		},
		Soil{
			Name:               "white sand",
			UsedForPottery:     false,
			IsSand:             true,
			UsedForAgriculture: false,
		},
	}

	return soils
}
