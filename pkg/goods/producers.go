package goods

func qualityFromSkillLevel(level int) string {
	qualities := map[int]string{
		1:  "poor ",
		2:  "mediocre ",
		3:  "average ",
		4:  "",
		5:  "",
		6:  "excellent ",
		7:  "exceptional ",
		8:  "masterful ",
		9:  "wonderful ",
		10: "legendary ",
	}

	return qualities[level]
}
