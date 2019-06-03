package culture

func (culture Culture) getViewOnMagic() string {
	view := "Magic is "

	common := false
	feared := false

	if culture.Attributes.Curiosity > 70 {
		view += "embraced, "
	} else if culture.Attributes.Curiosity > 40 {
		view += "part of life, "
	} else if culture.Attributes.Superstition > 80 {
		view += "widely feared, "
		feared = true
	} else if culture.Attributes.Superstition > 40 {
		view += "distrusted, "
		feared = true
	} else {
		view += "neither trusted nor hated, "
	}

	if culture.Attributes.MagicPrevalence > 70 {
		if feared {
			view += "but "
		} else {
			view += "and "
		}
		view += "mages are everywhere."
		common = true
	} else if culture.Attributes.MagicPrevalence > 50 {
		view += "and mages are seen occasionally in society."
	} else if culture.Attributes.MagicPrevalence > 30 {
		if feared {
			view += "and "
		} else {
			view += "but "
		}
		view += "mages are uncommon."
	} else {
		if feared {
			view += "and "
		} else {
			view += "but "
		}
		view += "mages are almost unheard of."
	}

	if culture.Attributes.MagicStrength > 80 {
		if common {
			view += " Powerful mages exist in abundance."
		} else {
			view += " Though rare, there are mages of incredible power."
		}
	} else if culture.Attributes.MagicStrength > 50 {
		if common {
			view += " Magic-users are powerful."
		} else {
			view += " Some powerful mages exist."
		}
	} else if culture.Attributes.MagicStrength > 30 {
		if common {
			view += " However, while common, magic-users are generally weak."
		} else {
			view += " Mages are not powerful, rarely rising above the level of hedge wizards."
		}
	} else {
		if common {
			view += " Magic-users are universally weak, sticking primarily to simple magic."
		} else {
			view += " Those few mages that exist are weak and have little real power."
		}
	}

	return view
}

func (culture Culture) getViews() []string {
	views := []string{}

	views = append(views, culture.getViewOnMagic())

	return views
}
