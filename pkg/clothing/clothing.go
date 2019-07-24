package clothing

import (
	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
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

	hides := getHides(originClimate)
	fabrics := getFabrics(originClimate)

	style.FemaleOutfit = GenerateOutfit(originClimate.Temperature, hides, fabrics, "female")
	style.MaleOutfit = GenerateOutfit(originClimate.Temperature, hides, fabrics, "male")

	style.CommonJewelry = generateJewelry(originClimate)

	style.DecorativeStyle = randomDecorativeStyle(originClimate)
	style.CommonColors = randomColorSet()

	return style
}

func getFabrics(originClimate climate.Climate) []string {
	resources := resource.ByTag("fabric fiber", originClimate.Resources)
	fabrics := []string{}

	for _, r := range resources {
		fabrics = append(fabrics, r.Name)
	}

	return fabrics
}

func getHides(originClimate climate.Climate) []string {
	resources := resource.ByTag("hide", originClimate.Resources)
	hides := []string{}

	for _, i := range resources {
		hides = append(hides, i.Name)
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

	ivory := resource.ByTag("ivory", originClimate.Resources)
	if len(ivory) > 0 {
		styles = append(styles, "ivory decorations")
	}

	feathers := resource.ByTag("feather", originClimate.Resources)
	if len(feathers) > 0 {
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
