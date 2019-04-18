package goods

// TradeGood is a trade good entry
type TradeGood struct {
	Name    string
	Quality string
	Amount  int
}

// Pattern is a structure for building a trade good
type Pattern struct {
	Name         string
	Material1    string
	Material2    string
	Material3    string
	Need1        string
	Need2        string
	Need3        string
	NameTemplate string
}

func getAnimalTrainerGoods() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "pack animal",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "pack animal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "pack {{.Material1}}",
		},
	}

	return patterns
}

func getArmor() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "chestplate",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "metal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} chestplate",
		},
	}

	return patterns
}

func getBreads() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "bread",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "grain",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} bread",
		},
	}

	return patterns
}

func getBrewed() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "ale",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "grain",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} ale",
		},
		Pattern{
			Name:         "beer",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "grain",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} beer",
		},
	}

	return patterns
}

func getCarpenterGoods() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "bench",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "wood",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} bench",
		},
		Pattern{
			Name:         "chair",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "wood",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} chair",
		},
	}

	return patterns
}

func getClothing() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "shirt",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "fabric",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} shirt",
		},
		Pattern{
			Name:         "tunic",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "fabric",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} tunic",
		},
	}

	return patterns
}

func getCobblerGoods() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "shoes",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "hide",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} shoes",
		},
	}

	return patterns
}

func getMasonGoods() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "stone block",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "stone",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} block",
		},
	}

	return patterns
}

func getMedicines() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "healing salve",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "herb",
			Need2:        "",
			Need3:        "",
			NameTemplate: "healing salve",
		},
	}

	return patterns
}

func getPotions() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "healing potion",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "herb",
			Need2:        "",
			Need3:        "",
			NameTemplate: "healing potion",
		},
	}

	return patterns
}

func getTannerGoods() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "leather",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "hide",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} leather",
		},
	}

	return patterns
}

func getWeapons() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "axe",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "metal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material}} axe",
		},
		Pattern{
			Name:         "sword",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "metal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material}} sword",
		},
	}

	return patterns
}

func getWine() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "wine",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "fruit",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} wine",
		},
	}

	return patterns
}
