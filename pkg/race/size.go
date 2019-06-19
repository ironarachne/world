package race

// Size is a type of race size
type Size struct {
	Name       string
	HeightBase int
	WeightBase int
}

func getAllSizes() []Size {
	return []Size{
		Size{
			Name:       "tiny",
			HeightBase: 36,
			WeightBase: 35,
		},
		Size{
			Name:       "small",
			HeightBase: 48,
			WeightBase: 50,
		},
		Size{
			Name:       "medium",
			HeightBase: 60,
			WeightBase: 100,
		},
		Size{
			Name:       "large",
			HeightBase: 72,
			WeightBase: 150,
		},
		Size{
			Name:       "huge",
			HeightBase: 84,
			WeightBase: 200,
		},
		Size{
			Name:       "giant",
			HeightBase: 96,
			WeightBase: 300,
		},
	}
}

func getSizeByName(name string) Size {
	sizes := getAllSizes()

	for _, s := range sizes {
		if s.Name == name {
			return s
		}
	}

	panic("No such race size found: " + name)
}
