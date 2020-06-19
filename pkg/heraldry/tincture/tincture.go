package tincture

import (
	"context"
	"fmt"
	"image/color"
	"os"

	"github.com/fogleman/gg"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
)

// Tincture is a tincture
type Tincture struct {
	Type            string
	Name            string
	Color           color.RGBA
	PatternFileName string
	Tags            []string
	Commonality     int
}

// Fill passes a tincture to gg's Fill context operation
func (tincture Tincture) Fill(dc *gg.Context) error {
	dataPath := os.Getenv("WORLDAPI_DATA_PATH")

	if tincture.Type == "fur" {
		filePath := dataPath + "/images/patterns/" + tincture.PatternFileName
		im, err := gg.LoadPNG(filePath)
		if err != nil {
			err = fmt.Errorf("Failed to fill with tincture: %w", err)
			return err
		}
		pattern := gg.NewSurfacePattern(im, gg.RepeatBoth)
		dc.SetFillStyle(pattern)
	} else {
		dc.SetColor(tincture.Color)
	}
	dc.Fill()
	return nil
}

func getColorTinctures() []Tincture {
	colors := []Tincture{
		{
			Name:        "azure",
			Type:        "color",
			Color:       color.RGBA{0, 12, 177, 255},
			Commonality: 1000,
			Tags: []string{
				"color",
			},
		},
		{
			Name:        "gules",
			Type:        "color",
			Color:       color.RGBA{206, 31, 0, 255},
			Commonality: 1000,
			Tags: []string{
				"color",
			},
		},
		{
			Name:        "purpure",
			Type:        "color",
			Color:       color.RGBA{72, 0, 169, 255},
			Commonality: 1000,
			Tags: []string{
				"color",
			},
		},
		{
			Name:        "sable",
			Type:        "color",
			Color:       color.RGBA{0, 0, 0, 255},
			Commonality: 1000,
			Tags: []string{
				"color",
			},
		},
		{
			Name:        "vert",
			Type:        "color",
			Color:       color.RGBA{12, 108, 0, 255},
			Commonality: 1000,
			Tags: []string{
				"color",
			},
		},
	}

	return colors
}

func getFurTinctures() []Tincture {
	furs := []Tincture{
		{
			Name:            "erminois",
			Type:            "fur",
			Color:           color.RGBA{255, 241, 0, 255},
			Commonality:     2,
			PatternFileName: "erminois.png",
			Tags: []string{
				"fur",
				"metal",
			},
		},
		{
			Name:            "ermine",
			Type:            "fur",
			Color:           color.RGBA{255, 255, 255, 255},
			Commonality:     2,
			PatternFileName: "ermine.png",
			Tags: []string{
				"fur",
				"metal",
			},
		},
		{
			Name:            "ermines",
			Type:            "fur",
			Color:           color.RGBA{0, 0, 0, 255},
			Commonality:     2,
			PatternFileName: "ermines.png",
			Tags: []string{
				"fur",
				"color",
			},
		},
		{
			Name:            "pean",
			Type:            "fur",
			Color:           color.RGBA{0, 0, 0, 255},
			Commonality:     2,
			PatternFileName: "pean.png",
			Tags: []string{
				"fur",
				"color",
			},
		},
	}

	return furs
}

func getMetalTinctures() []Tincture {
	metals := []Tincture{
		{
			Name:        "argent",
			Type:        "metal",
			Color:       color.RGBA{255, 255, 255, 255},
			Commonality: 950,
			Tags: []string{
				"metal",
			},
		},
		{
			Name:        "Or",
			Type:        "metal",
			Color:       color.RGBA{255, 241, 0, 255},
			Commonality: 950,
			Tags: []string{
				"metal",
			},
		},
	}

	return metals
}

func getStainTinctures() []Tincture {
	stains := []Tincture{
		{
			Name:        "murrey",
			Type:        "stain",
			Color:       color.RGBA{108, 0, 73, 255},
			Commonality: 5,
			Tags: []string{
				"color",
				"stain",
			},
		},
		{
			Name:        "sanguine",
			Type:        "stain",
			Color:       color.RGBA{135, 17, 0, 255},
			Commonality: 5,
			Tags: []string{
				"color",
				"stain",
			},
		},
		{
			Name:        "tenné",
			Type:        "stain",
			Color:       color.RGBA{108, 81, 0, 255},
			Commonality: 5,
			Tags: []string{
				"color",
				"stain",
			},
		},
	}

	return stains
}

// All returns all tinctures
func All() []Tincture {
	tinctures := []Tincture{}

	metals := getMetalTinctures()
	furs := getFurTinctures()
	colors := getColorTinctures()
	stains := getStainTinctures()

	tinctures = append(tinctures, colors...)
	tinctures = append(tinctures, furs...)
	tinctures = append(tinctures, metals...)
	tinctures = append(tinctures, stains...)

	return tinctures
}

// ByTag returns all tinctures that have the specified tag
func ByTag(tinctures []Tincture, tag string) []Tincture {
	filtered := []Tincture{}

	for _, t := range tinctures {
		if slices.StringIn(tag, t.Tags) {
			filtered = append(filtered, t)
		}
	}

	return filtered
}

// ExcludeTag returns all tinctures that don't have the specified tag
func ExcludeTag(tinctures []Tincture, tag string) []Tincture {
	filtered := []Tincture{}

	for _, t := range tinctures {
		if !slices.StringIn(tag, t.Tags) {
			filtered = append(filtered, t)
		}
	}

	return filtered
}

// RandomWeighted returns a random tincture based on its commonality from the specified set
func RandomWeighted(ctx context.Context, tinctures []Tincture) (Tincture, error) {
	weights := map[string]int{}

	for _, t := range tinctures {
		weights[t.Name] = t.Commonality
	}

	name, err := random.StringFromThresholdMap(ctx, weights)
	if err != nil {
		err = fmt.Errorf("Failed to get random weighted tincture: %w", err)
		return Tincture{}, err
	}

	for _, t := range tinctures {
		if t.Name == name {
			return t, nil
		}
	}

	err = fmt.Errorf("Failed to get random weighted tincture matching " + name)
	return Tincture{}, err
}

// Random returns a random tincture from the specified set
func Random(ctx context.Context, tinctures []Tincture) (Tincture, error) {
	if len(tinctures) == 0 {
		err := fmt.Errorf("Tried to find random tincture from empty set")
		return Tincture{}, err
	} else if len(tinctures) == 1 {
		return tinctures[0], nil
	}

	t := tinctures[random.Intn(ctx, len(tinctures))]
	return t, nil
}

// RandomAll returns a random tincture
func RandomAll(ctx context.Context) (Tincture, error) {
	tinctures := All()

	t, err := RandomWeighted(ctx, tinctures)
	if err != nil {
		err = fmt.Errorf("Failed to get random tincture from all: %w", err)
		return Tincture{}, err
	}
	return t, nil
}

// RandomByTag returns a random tincture that has the specified tag
func RandomByTag(ctx context.Context, tag string) (Tincture, error) {
	all := All()
	tinctures := ByTag(all, tag)
	tincture, err := RandomWeighted(ctx, tinctures)
	if err != nil {
		err = fmt.Errorf("Failed to get random tincture by tag "+tag+": %w", err)
		return Tincture{}, err
	}
	return tincture, nil
}

// Complementary returns all tinctures that complement the given one
func Complementary(t Tincture, includeFurs bool) []Tincture {
	tinctures := []Tincture{}

	all := All()

	if slices.StringIn("color", t.Tags) {
		tinctures = ExcludeTag(all, "metal")
	} else {
		tinctures = ExcludeTag(all, "color")
	}

	if !includeFurs {
		tinctures = ExcludeTag(tinctures, "fur")
	}

	tinctures = Remove(t, tinctures)

	return tinctures
}

// Contrasting returns all tinctures that contrast the given one
func Contrasting(t Tincture, includeFurs bool) []Tincture {
	tinctures := []Tincture{}

	all := All()

	if slices.StringIn("color", t.Tags) {
		tinctures = ExcludeTag(all, "color")
	} else {
		tinctures = ExcludeTag(all, "metal")
	}

	if !includeFurs {
		tinctures = ExcludeTag(tinctures, "fur")
	}

	return tinctures
}

// RandomComplementary returns a random tincture that contrasts with the given one
func RandomComplementary(ctx context.Context, t Tincture, includeFurs bool) (Tincture, error) {
	complementary := Complementary(t, includeFurs)

	if len(complementary) == 0 {
		err := fmt.Errorf("No complementary tinctures!")
		return Tincture{}, err
	} else if len(complementary) == 1 {
		return complementary[0], nil
	}

	t2, err := RandomWeighted(ctx, complementary)
	if err != nil {
		err = fmt.Errorf("Failed to get random complementary tincture: %w", err)
		return Tincture{}, err
	}

	return t2, nil
}

// RandomContrasting returns a random tincture that contrasts with the given one
func RandomContrasting(ctx context.Context, t Tincture, includeFurs bool) (Tincture, error) {
	contrasting := Contrasting(t, includeFurs)

	if len(contrasting) == 0 {
		err := fmt.Errorf("No contrasting tinctures!")
		return Tincture{}, err
	} else if len(contrasting) == 1 {
		return contrasting[0], nil
	}

	t2, err := RandomWeighted(ctx, contrasting)
	if err != nil {
		err = fmt.Errorf("Failed to get random contrasting tincture: %w", err)
		return Tincture{}, err
	}

	return t2, nil
}

// Remove removes the given tincture from the given set
func Remove(tincture Tincture, tinctures []Tincture) []Tincture {
	filtered := []Tincture{}

	for _, t := range tinctures {
		if t.Name != tincture.Name {
			filtered = append(filtered, t)
		}
	}

	return filtered
}
