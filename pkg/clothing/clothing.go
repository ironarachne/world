package clothing

import (
	"fmt"

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
func GenerateStyle(originClimate climate.Climate) (Style, error) {
	style := Style{}

	hides := getHides(originClimate)
	fabrics := getFabrics(originClimate)

	femaleOutfit, err := GenerateOutfit(originClimate.Temperature, hides, fabrics, "female")
	if err != nil {
		err = fmt.Errorf("Could not generate clothing style: %w", err)
		return Style{}, err
	}
	style.FemaleOutfit = femaleOutfit
	maleOutfit, err := GenerateOutfit(originClimate.Temperature, hides, fabrics, "male")
	if err != nil {
		err = fmt.Errorf("Could not generate clothing style: %w", err)
		return Style{}, err
	}
	style.MaleOutfit = maleOutfit

	jewelry, err := generateJewelry(originClimate)
	if err != nil {
		err = fmt.Errorf("Could not generate clothing style: %w", err)
		return Style{}, err
	}
	style.CommonJewelry = jewelry

	decorativeStyle, err := randomDecorativeStyle(originClimate)
	if err != nil {
		err = fmt.Errorf("Could not generate clothing style: %w", err)
		return Style{}, err
	}
	style.DecorativeStyle = decorativeStyle

	colors, err := randomColorSet()
	if err != nil {
		err = fmt.Errorf("Could not generate clothing style: %w", err)
		return Style{}, err
	}
	style.CommonColors = colors

	return style, nil
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

func randomDecorativeStyle(originClimate climate.Climate) (string, error) {
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

	style, err := random.String(styles)
	if err != nil {
		err = fmt.Errorf("Could not generate clothing decorative style: %w", err)
		return "", err
	}

	return style, nil
}

func randomColorSet() ([]string, error) {
	var colorSet []string
	var err error
	var newColor string
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
		newColor, err = random.String(colors)
		if err != nil {
			err = fmt.Errorf("Could not generate color set: %w", err)
			return []string{}, err
		}

		if !slices.StringIn(newColor, []string{"white", "black"}) {
			saturation, err = randomSaturation()
			if err != nil {
				err = fmt.Errorf("Could not generate color set: %w", err)
				return []string{}, err
			}
			newColor = saturation + " " + newColor
		}

		if !slices.StringIn(newColor, colorSet) {
			colorSet = append(colorSet, newColor)
		}
	}

	return colorSet, nil
}

func randomSaturation() (string, error) {
	saturations := []string{
		"bright",
		"dark",
		"light",
		"moderate",
		"pastel",
		"subdued",
	}

	saturation, err := random.String(saturations)
	if err != nil {
		err = fmt.Errorf("Could not generate saturation: %w", err)
		return "", err
	}

	return saturation, nil
}

// Random generates a completely random clothing style
func Random() (Style, error) {
	originClimate, err := climate.Generate()
	if err != nil {
		err = fmt.Errorf("Could not generate clothing style: %w", err)
		return Style{}, err
	}

	style, err := GenerateStyle(originClimate)
	if err != nil {
		err = fmt.Errorf("Could not generate clothing style: %w", err)
		return Style{}, err
	}

	return style, nil
}
