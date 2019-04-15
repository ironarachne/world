package culture

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
)

// Appearance is a set of physical characteristics
type Appearance struct {
	MaxFemaleHeight  int
	MinFemaleHeight  int
	MaxMaleHeight    int
	MinMaleHeight    int
	MaxFemaleWeight  int
	MinFemaleWeight  int
	MaxMaleWeight    int
	MinMaleWeight    int
	FaceShape        string
	EyeShape         string
	NoseShape        string
	MouthShape       string
	FacialHairStyles []string
	HairColors       []string
	HairConsistency  string
	FemaleHairStyles []string
	MaleHairStyles   []string
	SkinColors       []string
	EyeColors        []string
	UniqueTrait      string
}

func (culture Culture) generateAppearance() Appearance {
	appearance := Appearance{}

	appearance.SkinColors = culture.calculateSkinColors()
	appearance.HairColors = culture.calculateHairColors()
	appearance.EyeColors = culture.calculateEyeColors()
	appearance.HairConsistency = culture.calculateHairConsistency()
	appearance.FemaleHairStyles = culture.calculateHairStyles()
	appearance.MaleHairStyles = culture.calculateHairStyles()
	appearance.FacialHairStyles = randomFacialHair()
	appearance.EyeShape = randomEyeShape()
	appearance.FaceShape = randomFaceShape()
	appearance.MouthShape = randomMouthShape()
	appearance.NoseShape = randomNoseShape()

	appearance.MinFemaleHeight = rand.Intn(24) + 52
	appearance.MaxFemaleHeight = appearance.MinFemaleHeight + rand.Intn(12) + 1
	appearance.MinMaleHeight = rand.Intn(24) + 52
	appearance.MaxMaleHeight = appearance.MinMaleHeight + rand.Intn(12) + 1
	appearance.MinFemaleWeight = rand.Intn(30) + 90
	appearance.MaxFemaleWeight = appearance.MinFemaleWeight + rand.Intn(80) + 1
	appearance.MinMaleWeight = rand.Intn(30) + 140
	appearance.MaxMaleWeight = appearance.MinMaleWeight + rand.Intn(100) + 1

	appearance.UniqueTrait = randomUniqueTrait()

	return appearance
}

func (culture Culture) calculateSkinColors() []string {
	colorGrades := map[int]string{
		0: "alabaster",
		1: "white",
		2: "light",
		3: "tan",
		4: "olive",
		5: "bronze",
		6: "brown",
		7: "dark",
		8: "black",
		9: "ebony",
	}

	baseLevel := rand.Intn(10)

	skinColors := []string{colorGrades[baseLevel]}

	if baseLevel == 0 {
		skinColors = append(skinColors, colorGrades[1])
		skinColors = append(skinColors, colorGrades[2])
	} else if baseLevel == 9 {
		skinColors = append(skinColors, colorGrades[7])
		skinColors = append(skinColors, colorGrades[8])
	} else {
		skinColors = append(skinColors, colorGrades[baseLevel-1])
		skinColors = append(skinColors, colorGrades[baseLevel+1])
	}

	return skinColors
}

func (culture Culture) calculateEyeColors() []string {
	var eyeColor string
	var eyeColors []string

	colors := []string{
		"amber",
		"blue",
		"brown",
		"gold",
		"green",
		"hazel",
		"grey",
		"pale",
	}

	if inSlice("black", culture.Appearance.SkinColors) {
		colors = []string{
			"amber",
			"brown",
		}
	} else if inSlice("alabaster", culture.Appearance.SkinColors) {
		colors = append(colors, "white")
	}

	for i := 0; i < 3; i++ {
		eyeColor = random.String(colors)
		if !inSlice(eyeColor, eyeColors) {
			eyeColors = append(eyeColors, eyeColor)
		}
	}

	return eyeColors
}

func (culture Culture) calculateHairColors() []string {
	var hairColor string
	var hairColors []string

	colors := []string{
		"black",
		"auburn",
		"red",
		"grey",
		"brown",
		"dark brown",
		"light brown",
	}

	if inSlice("black", culture.Appearance.SkinColors) {
		colors = []string{
			"black",
			"dark brown",
		}
	} else if inSlice("alabaster", culture.Appearance.SkinColors) {
		colors = append(colors, "white")
	} else if inSlice("white", culture.Appearance.SkinColors) {
		colors = append(colors, "blonde")
	}

	for i := 0; i < 2; i++ {
		hairColor = random.String(colors)
		if !inSlice(hairColor, hairColors) {
			hairColors = append(hairColors, hairColor)
		}
	}

	return hairColors
}

func (culture Culture) calculateHairConsistency() string {
	consistencies := []string{
		"thick",
	}

	chanceThin := rand.Intn(10)
	if chanceThin > 8 {
		consistencies = append(consistencies, "thin")
	}

	if inSlice("black", culture.Appearance.SkinColors) {
		consistencies = append(consistencies, "coarse")
	}

	return random.String(consistencies)
}

func (culture Culture) calculateHairStyles() []string {
	var styles []string
	var style string

	possibleStyles := []string{
		"shoulder-length",
		"short-cropped",
		"bald",
		"bowl-cut",
		"ponytail",
		"topknot",
		"braided",
		"back-length",
		"mohawk",
	}

	if culture.Appearance.HairConsistency == "coarse" {
		possibleStyles = append(possibleStyles, "long dreadlocks")
		possibleStyles = append(possibleStyles, "short dreadlocks")
		possibleStyles = append(possibleStyles, "tight braids")
	}

	for i := 0; i < 3; i++ {
		style = random.String(possibleStyles)
		if !inSlice(style, styles) {
			styles = append(styles, style)
		}
	}

	return styles
}

func randomFacialHair() []string {
	var styles []string
	var style string

	possibleStyles := []string{
		"full beard",
		"short goatee",
		"long goatee",
		"handlebar mustache",
		"long beard",
		"short beard",
		"mutton chops",
		"clean shaven",
		"trim mustache",
		"narrow mustache",
	}

	for i := 0; i < 3; i++ {
		style = random.String(possibleStyles)
		if !inSlice(style, styles) {
			styles = append(styles, style)
		}
	}

	return styles
}

func randomUniqueTrait() string {
	possibleTraits := []string{
		"body tattoos",
		"facial tattoos",
		"hand tattoos",
		"ritual scarification",
		"nose piercings",
		"tongue piercings",
	}

	chanceTrait := rand.Intn(10)
	if chanceTrait > 9 {
		return random.String(possibleTraits)
	}
	return "none"
}

func randomEyeShape() string {
	shapes := []string{
		"almond",
		"narrow",
		"oval",
		"round",
		"slanted",
	}

	return random.String(shapes)
}

func randomFaceShape() string {
	shapes := []string{
		"broad",
		"chiseled",
		"oval",
		"round",
		"square",
	}

	return random.String(shapes)
}

func randomMouthShape() string {
	shapes := []string{
		"full",
		"narrow",
		"thin",
		"wide",
	}

	return random.String(shapes)
}

func randomNoseShape() string {
	shapes := []string{
		"acquiline",
		"long",
		"narrow",
		"pointed",
		"thin",
		"wide",
	}

	return random.String(shapes)
}
