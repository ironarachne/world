/*
Package clothing provides methods and tools for generating
fantasy clothing styles.
*/
package clothing

import (
	"context"
	"fmt"
	"github.com/ironarachne/world/pkg/geography"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/slices"
)

const clothingError = "failed to generate clothing style: %w"

// Style describes what kind of clothing the culture wears
type Style struct {
	Description     string   `json:"description"`
	FemaleOutfit    []Item   `json:"female_outfit"`
	MaleOutfit      []Item   `json:"male_outfit"`
	CommonJewelry   []string `json:"common_jewelry"`
	CommonColors    []string `json:"common_colors"`
	DecorativeStyle string   `json:"decorative_style"`
}

// GenerateStyle generates a random clothing style based on a climate
func GenerateStyle(ctx context.Context, temperature int, resources []resource.Resource) (Style, error) {
	style := Style{}

	femaleOutfit, err := GenerateOutfit(ctx, temperature, "female")
	if err != nil {
		err = fmt.Errorf(clothingError, err)
		return Style{}, err
	}
	style.FemaleOutfit = femaleOutfit
	maleOutfit, err := GenerateOutfit(ctx, temperature, "male")
	if err != nil {
		err = fmt.Errorf(clothingError, err)
		return Style{}, err
	}
	style.MaleOutfit = maleOutfit

	jewelry, err := generateJewelry(ctx)
	if err != nil {
		err = fmt.Errorf(clothingError, err)
		return Style{}, err
	}
	style.CommonJewelry = jewelry

	decorativeStyle, err := randomDecorativeStyle(ctx, resources)
	if err != nil {
		err = fmt.Errorf(clothingError, err)
		return Style{}, err
	}
	style.DecorativeStyle = decorativeStyle

	colors, err := randomColorSet(ctx)
	if err != nil {
		err = fmt.Errorf(clothingError, err)
		return Style{}, err
	}
	style.CommonColors = colors
	style.Description = style.Describe(ctx)

	return style, nil
}

func randomDecorativeStyle(ctx context.Context, resources []resource.Resource) (string, error) {
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

	ivory := resource.ByTag("ivory", resources)
	if len(ivory) > 0 {
		styles = append(styles, "ivory decorations")
	}

	feathers := resource.ByTag("feather", resources)
	if len(feathers) > 0 {
		styles = append(styles, "feather decorations")
	}

	style, err := random.String(ctx, styles)
	if err != nil {
		err = fmt.Errorf("failed to generate clothing decorative style: %w", err)
		return "", err
	}

	return style, nil
}

func randomColorSet(ctx context.Context) ([]string, error) {
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
		newColor, err = random.String(ctx, colors)
		if err != nil {
			err = fmt.Errorf("Could not generate color set: %w", err)
			return []string{}, err
		}

		if !slices.StringIn(newColor, []string{"white", "black"}) {
			saturation, err = randomSaturation(ctx)
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

func randomSaturation(ctx context.Context) (string, error) {
	saturations := []string{
		"bright",
		"dark",
		"light",
		"moderate",
		"pastel",
		"subdued",
	}

	saturation, err := random.String(ctx, saturations)
	if err != nil {
		err = fmt.Errorf("Could not generate saturation: %w", err)
		return "", err
	}

	return saturation, nil
}

// Random generates a completely random clothing style
func Random(ctx context.Context) (Style, error) {
	area, err := geography.Generate(ctx)
	if err != nil {
		err = fmt.Errorf(clothingError, err)
		return Style{}, err
	}

	resources := area.GetResources()

	style, err := GenerateStyle(ctx, area.Region.Temperature, resources)
	if err != nil {
		err = fmt.Errorf(clothingError, err)
		return Style{}, err
	}

	return style, nil
}
