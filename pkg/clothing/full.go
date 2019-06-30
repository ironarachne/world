package clothing

func getRandomDress() Item {
	template := ItemTemplate{
		Name:         "dress",
		Type:         "full",
		MaterialType: "fabric",
		PrefixModifiers: []string{
			"form-fitting",
			"long-sleeved",
			"loose",
			"short-sleeved",
			"sleeveless",
			"thin-strapped",
			"tight-sleeved",
		},
		SuffixModifiers: []string{
			"with a tight top and loose skirt",
			"with a tight top and flared skirt",
			"with short sleeves",
			"with long sleeves",
			"with ribbed sleeves",
			"with sleeves flared at the wrists",
			"with tight sleeves",
			"with no sleeves",
			"with tapered sleeves",
		},
		IdealHeatLevel: "cold",
	}

	dress := getItemFromTemplate(template)

	return dress
}

func getRandomRobe() Item {
	template := ItemTemplate{
		Name:         "robe",
		Type:         "full",
		MaterialType: "fabric",
		PrefixModifiers: []string{
			"long-sleeved",
			"loose",
			"short-sleeved",
			"sleeveless",
			"tight-sleeved",
		},
		SuffixModifiers: []string{
			"with short sleeves",
			"with long sleeves",
			"with ribbed sleeves",
			"with sleeves flared at the wrists",
			"with tight sleeves",
			"with no sleeves",
			"with tapered sleeves",
		},
		IdealHeatLevel: "cold",
	}

	robe := getItemFromTemplate(template)

	return robe
}
