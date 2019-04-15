package town

var (
	townCategoryOptions = map[string]int{
		"city":    3,
		"town":    5,
		"village": 10,
	}

	townCategories = map[string]TownCategory{
		"city": TownCategory{
			"city",
			10000,
			100000,
			3,
			12,
			3,
			6,
		},
		"town": TownCategory{
			"town",
			1000,
			10000,
			1,
			3,
			1,
			3,
		},
		"village": TownCategory{
			"village",
			10,
			100,
			1,
			1,
			1,
			1,
		},
	}
)
