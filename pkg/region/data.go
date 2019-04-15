package region

var (
	classes = map[string]int{
		"barony":    3,
		"county":    4,
		"duchy":     2,
		"march":     2,
		"viscounty": 1,
	}

	classData = map[string]RegionClass{
		"barony": RegionClass{
			Name:             "Barony",
			RulerTitleFemale: "Baroness",
			RulerTitleMale:   "Baron",
			MinNumberOfTowns: 2,
			MaxNumberOfTowns: 4,
		},
		"county": RegionClass{
			Name:             "County",
			RulerTitleFemale: "Countess",
			RulerTitleMale:   "Count",
			MinNumberOfTowns: 1,
			MaxNumberOfTowns: 3,
		},
		"duchy": RegionClass{
			Name:             "Duchy",
			RulerTitleFemale: "Duchess",
			RulerTitleMale:   "Duke",
			MinNumberOfTowns: 3,
			MaxNumberOfTowns: 6,
		},
		"march": RegionClass{
			Name:             "March",
			RulerTitleFemale: "Marchioness",
			RulerTitleMale:   "Marquess",
			MinNumberOfTowns: 2,
			MaxNumberOfTowns: 4,
		},
		"viscounty": RegionClass{
			Name:             "Viscounty",
			RulerTitleFemale: "Viscountess",
			RulerTitleMale:   "Viscount",
			MinNumberOfTowns: 1,
			MaxNumberOfTowns: 2,
		},
	}
)
