package clothing

import (
	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

// Style describes what kind of clothing the culture wears
type Style struct {
	FemaleOutfit    []Item
	MaleOutfit      []Item
	CommonJewelry   []string
	CommonColors    []string
	DecorativeStyle string
}

// GenerateStyle generates a random clothing style based on a climate
func GenerateStyle(originClimate climate.Climate) Style {
	style := Style{}

	style.FemaleOutfit = GenerateOutfit(originClimate.Temperature, getHides(originClimate), getFabrics(originClimate), "female")
	style.MaleOutfit = GenerateOutfit(originClimate.Temperature, getHides(originClimate), getFabrics(originClimate), "male")

	style.CommonJewelry = generateJewelry(originClimate)

	style.DecorativeStyle = randomDecorativeStyle(originClimate)
	style.CommonColors = randomColorSet()

	return style
}

func getFabrics(originClimate climate.Climate) []string {
	fabrics := []string{}

	for _, i := range originClimate.Plants {
		if i.Name == "flax" {
			fabrics = append(fabrics, "linen")
		} else if i.IsFiber {
			fabrics = append(fabrics, i.Name)
		}
	}

	for _, i := range originClimate.Animals {
		if i.GivesWool {
			fabrics = append(fabrics, i.Name+" wool")
		}
	}

	return fabrics
}

func getHides(originClimate climate.Climate) []string {
	hides := []string{}

	for _, i := range originClimate.Animals {
		if i.GivesFur {
			hides = append(hides, i.Name+" fur")
		} else if i.Name == "cow" {
			hides = append(hides, "leather")
		} else if i.AnimalType == "reptile" {
			hides = append(hides, i.Name+" skin")
		} else if i.GivesHide {
			hides = append(hides, i.Name+" hide")
		}
	}

	return hides
}

func randomDecorativeStyle(originClimate climate.Climate) string {
	styles := []string{
		"beads",
		"complex embroidery",
		"geometric shapes",
		"knotwork",
		"plain",
		"simple embroidery",
		"stripes",
		"stylized animals",
		"tassels",
	}

	if climate.IsTypeInResources("ivory", originClimate.Resources) {
		styles = append(styles, "ivory decorations")
	}

	if climate.IsTypeInResources("feathers", originClimate.Resources) {
		styles = append(styles, "feather decorations")
	}

	if originClimate.Temperature < 6 {
		styles = append(styles, "many layers")
	}

	return random.String(styles)
}

func randomColorSet() []string {
	var newColor string
	var colorSet []string
	var saturation string

	colors := []string{
		"red",
		"blue",
		"green",
		"black",
		"white",
		"yellow",
		"purple",
		"orange",
		"grey",
	}

	for i := 0; i < 3; i++ {
		newColor = random.String(colors)

		if !slices.StringIn(newColor, []string{"white", "black"}) {
			saturation = randomSaturation()
			newColor = saturation + " " + newColor
		}

		if !slices.StringIn(newColor, colorSet) {
			colorSet = append(colorSet, newColor)
		}
	}

	return colorSet
}

func randomSaturation() string {
	saturations := []string{
		"bright",
		"dark",
		"light",
		"moderate",
		"pastel",
		"subdued",
	}

	return random.String(saturations)
}

// Random generates a completely random clothing style
func Random() Style {
	originClimate := climate.Generate()

	return GenerateStyle(originClimate)
}
