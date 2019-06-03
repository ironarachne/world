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
			Name:         "mount",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "mount",
			Need2:        "",
			Need3:        "",
			NameTemplate: "riding {{.Material1}}s",
		},
		Pattern{
			Name:         "pack animal",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "pack animal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "pack {{.Material1}}s",
		},
		Pattern{
			Name:         "war mount",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "mount",
			Need2:        "",
			Need3:        "",
			NameTemplate: "war {{.Material1}}s",
		},
	}

	return patterns
}

func getArmor() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "chestplates",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "metal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} chestplates",
		},
		Pattern{
			Name:         "full plate",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "metal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} full plate",
		},
		Pattern{
			Name:         "gauntlets",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "metal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} gauntlets",
		},
		Pattern{
			Name:         "greeves",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "metal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} greeves",
		},
		Pattern{
			Name:         "mail shirt",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "metal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} mail shirt",
		},
		Pattern{
			Name:         "mail hauberks",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "metal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} mail hauberks",
		},
		Pattern{
			Name:         "spaulders",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "metal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} spaulders",
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
		Pattern{
			Name:         "buns",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "grain",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} buns",
		},
		Pattern{
			Name:         "rolls",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "grain",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} rolls",
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
		Pattern{
			Name:         "lager",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "grain",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} lager",
		},
	}

	return patterns
}

func getCarpenterGoods() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "wood benches",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "wood",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} wood benches",
		},
		Pattern{
			Name:         "wooden bowls",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "wood",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} wooden bowls",
		},
		Pattern{
			Name:         "wood chairs",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "wood",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} wood chairs",
		},
		Pattern{
			Name:         "wooden mugs",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "wood",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} wooden mugs",
		},
		Pattern{
			Name:         "wood tables",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "wood",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} wood tables",
		},
		Pattern{
			Name:         "wood tankards",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "wood",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} wood tankards",
		},
	}

	return patterns
}

func getClothing() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "breeches",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "fabric",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} breeches",
		},
		Pattern{
			Name:         "gloves",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "fabric",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} gloves",
		},
		Pattern{
			Name:         "hoods",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "fabric",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} hoods",
		},
		Pattern{
			Name:         "pants",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "fabric",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} pants",
		},
		Pattern{
			Name:         "pantaloons",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "fabric",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} pantaloons",
		},
		Pattern{
			Name:         "shirts",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "fabric",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} shirts",
		},
		Pattern{
			Name:         "tunics",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "fabric",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} tunics",
		},
	}

	return patterns
}

func getCobblerGoods() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "boots",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "hide",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} boots",
		},
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

func getFlour() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "flour",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "grain",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} flour",
		},
	}

	return patterns
}

func getMasonGoods() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "stone blocks",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "stone",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} blocks",
		},
		Pattern{
			Name:         "stone slabs",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "stone",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} slabs",
		},
	}

	return patterns
}

func getMedicines() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "healing draughts",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "herb",
			Need2:        "",
			Need3:        "",
			NameTemplate: "healing draughts",
		},
		Pattern{
			Name:         "healing salves",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "herb",
			Need2:        "",
			Need3:        "",
			NameTemplate: "healing salves",
		},
	}

	return patterns
}

func getMinerGoods() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "ore",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "ore",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} ore",
		},
	}

	return patterns
}

func getPotions() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "healing potions",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "herb",
			Need2:        "",
			Need3:        "",
			NameTemplate: "healing potions",
		},
	}

	return patterns
}

func getPottery() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "jugs",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "clay",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} jugs",
		},
		Pattern{
			Name:         "mugs",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "clay",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} mugs",
		},
		Pattern{
			Name:         "pots",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "clay",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} pots",
		},
		Pattern{
			Name:         "urns",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "clay",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} urns",
		},
		Pattern{
			Name:         "vases",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "clay",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} vases",
		},
	}

	return patterns
}

func getTannerGoods() []Pattern {
	patterns := []Pattern{
		Pattern{
			Name:         "hides",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "hide",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material1}} hides",
		},
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
			Name:         "axes",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "metal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material}} axes",
		},
		Pattern{
			Name:         "maces",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "metal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material}} maces",
		},
		Pattern{
			Name:         "polearms",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "metal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material}} polearms",
		},
		Pattern{
			Name:         "spears",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "metal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material}} spears",
		},
		Pattern{
			Name:         "swords",
			Material1:    "",
			Material2:    "",
			Material3:    "",
			Need1:        "metal",
			Need2:        "",
			Need3:        "",
			NameTemplate: "{{.Material}} swords",
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
