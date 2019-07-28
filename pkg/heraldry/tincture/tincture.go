package tincture

import (
	svg "github.com/ajstarks/svgo"
	"github.com/ironarachne/world/pkg/random"
	"math/rand"
)

// Tincture is a tincture
type Tincture struct {
	Type    string
	Name    string
	HexCode string
}

func getColorTinctures() []Tincture {
	colors := []Tincture{
		{
			Name: "azure",
			Type: "color",
			HexCode: "#FFFFFF",
		},
		{
			Name: "gules",
			Type: "color",
			HexCode: "#EE0000",
		},
		{
			Name: "purpure",
			Type: "color",
			HexCode: "#831CA6",
		},
		{
			Name: "sable",
			Type: "color",
			HexCode: "#000000",
		},
		{
			Name: "vert",
			Type: "color",
			HexCode: "#009900",
		},
	}

	return colors
}

func getFurTinctures() []Tincture {
	furs := []Tincture{
		{
			Name: "erminois",
			Type: "fur",
			HexCode: "url(#erminois)",
		},
		{
			Name: "ermine",
			Type: "fur",
			HexCode: "url(#ermine)",
		},
		{
			Name: "ermines",
			Type: "fur",
			HexCode: "url(#ermines)",
		},
		{
			Name: "pean",
			Type: "fur",
			HexCode: "url(#pean)",
		},
	}

	return furs
}

func getMetalTinctures() []Tincture {
	metals := []Tincture{
		{
			Name: "argent",
			Type: "metal",
			HexCode: "#FFFFFF",
		},
		{
			Name: "Or",
			Type: "metal",
			HexCode: "#FFEC00",
		},
	}

	return metals
}

func getStainTinctures() []Tincture {
	stains := []Tincture{
		{
			Name: "murrey",
			Type: "stain",
			HexCode: "#591334",
		},
		{
			Name: "sanguine",
			Type: "stain",
			HexCode: "#720101",
		},
		{
			Name: "tenné",
			Type: "stain",
			HexCode: "#7C4215",
		},
	}

	return stains
}

// InsertErmine adds a pattern to a canvas
func InsertErmine(canvas *svg.SVG, variant string) {
	fillColor := "#FFFFFF"
	objectColor := "#000000"
	if variant == "erminois" {
		fillColor = "#FFEC00"
	} else if variant == "ermines" {
		fillColor = "#000000"
		objectColor = "#FFFFFF"
	} else if variant == "pean" {
		fillColor = "#000000"
		objectColor = "#FFEC00"
	}

	canvas.Pattern(variant, 0, 0, 50, 50, "user")
	canvas.Rect(0, 0, 100, 100, "fill:"+fillColor)
	canvas.Scale(0.5)
	canvas.Path("M 31.999996,26.250003 C 31.999996,29.700001 29.199996,32.500001 25.749998,32.500001 C 22.299999,32.500001 19.500001,29.700001 19.500001,26.250003 C 19.500001,22.800004 22.299999,20.000006 25.749998,20.000006 C 29.199996,20.000006 31.999996,22.800004 31.999996,26.250003 z M 22.999995,37.250003 C 22.999995,40.700001 20.199995,43.500001 16.749997,43.500001 C 13.299998,43.500001 10.5,40.700001 10.5,37.250003 C 10.5,33.800004 13.299998,31.000006 16.749997,31.000006 C 20.199995,31.000006 22.999995,33.800004 22.999995,37.250003 z M 40.999996,37.250003 C 40.999996,40.700001 38.199996,43.500001 34.749998,43.500001 C 31.299999,43.500001 28.500001,40.700001 28.500001,37.250003 C 28.500001,33.800004 31.299999,31.000006 34.749998,31.000006 C 38.199996,31.000006 40.999996,33.800004 40.999996,37.250003 z M 23.921297,33.936504 L 26.873286,34.14009 C 26.975079,53.37892 41.633235,83.814953 41.633235,83.814953 L 32.420991,75.00988 C 29.876173,77.96187 25.346395,83.814953 25.346395,83.814953 C 25.346395,83.814953 22.089027,77.910973 19.747794,74.958983 L 9.2631404,83.71316 C 9.2631404,83.71316 23.717711,53.582505 23.921297,33.936504 z", "color:"+objectColor+";fill:"+objectColor+";fill-opacity:1;fill-rule:nonzero;stroke:none;stroke-width:8;stroke-linecap:square;stroke-linejoin:miter;marker:none;marker-start:none;marker-mid:none;marker-end:none;stroke-miterlimit:4;stroke-dashoffset:0;stroke-opacity:1;visibility:visible;display:inline;overflow:visible")
	canvas.Translate(50, 50)
	canvas.Path("M 31.999996,26.250003 C 31.999996,29.700001 29.199996,32.500001 25.749998,32.500001 C 22.299999,32.500001 19.500001,29.700001 19.500001,26.250003 C 19.500001,22.800004 22.299999,20.000006 25.749998,20.000006 C 29.199996,20.000006 31.999996,22.800004 31.999996,26.250003 z M 22.999995,37.250003 C 22.999995,40.700001 20.199995,43.500001 16.749997,43.500001 C 13.299998,43.500001 10.5,40.700001 10.5,37.250003 C 10.5,33.800004 13.299998,31.000006 16.749997,31.000006 C 20.199995,31.000006 22.999995,33.800004 22.999995,37.250003 z M 40.999996,37.250003 C 40.999996,40.700001 38.199996,43.500001 34.749998,43.500001 C 31.299999,43.500001 28.500001,40.700001 28.500001,37.250003 C 28.500001,33.800004 31.299999,31.000006 34.749998,31.000006 C 38.199996,31.000006 40.999996,33.800004 40.999996,37.250003 z M 23.921297,33.936504 L 26.873286,34.14009 C 26.975079,53.37892 41.633235,83.814953 41.633235,83.814953 L 32.420991,75.00988 C 29.876173,77.96187 25.346395,83.814953 25.346395,83.814953 C 25.346395,83.814953 22.089027,77.910973 19.747794,74.958983 L 9.2631404,83.71316 C 9.2631404,83.71316 23.717711,53.582505 23.921297,33.936504 z", "color:"+objectColor+";fill:"+objectColor+";fill-opacity:1;fill-rule:nonzero;stroke:none;stroke-width:8;stroke-linecap:square;stroke-linejoin:miter;marker:none;marker-start:none;marker-mid:none;marker-end:none;stroke-miterlimit:4;stroke-dashoffset:0;stroke-opacity:1;visibility:visible;display:inline;overflow:visible")
	canvas.Gend()
	canvas.Translate(50, -50)
	canvas.Path("M 31.999996,26.250003 C 31.999996,29.700001 29.199996,32.500001 25.749998,32.500001 C 22.299999,32.500001 19.500001,29.700001 19.500001,26.250003 C 19.500001,22.800004 22.299999,20.000006 25.749998,20.000006 C 29.199996,20.000006 31.999996,22.800004 31.999996,26.250003 z M 22.999995,37.250003 C 22.999995,40.700001 20.199995,43.500001 16.749997,43.500001 C 13.299998,43.500001 10.5,40.700001 10.5,37.250003 C 10.5,33.800004 13.299998,31.000006 16.749997,31.000006 C 20.199995,31.000006 22.999995,33.800004 22.999995,37.250003 z M 40.999996,37.250003 C 40.999996,40.700001 38.199996,43.500001 34.749998,43.500001 C 31.299999,43.500001 28.500001,40.700001 28.500001,37.250003 C 28.500001,33.800004 31.299999,31.000006 34.749998,31.000006 C 38.199996,31.000006 40.999996,33.800004 40.999996,37.250003 z M 23.921297,33.936504 L 26.873286,34.14009 C 26.975079,53.37892 41.633235,83.814953 41.633235,83.814953 L 32.420991,75.00988 C 29.876173,77.96187 25.346395,83.814953 25.346395,83.814953 C 25.346395,83.814953 22.089027,77.910973 19.747794,74.958983 L 9.2631404,83.71316 C 9.2631404,83.71316 23.717711,53.582505 23.921297,33.936504 z", "color:"+objectColor+";fill:"+objectColor+";fill-opacity:1;fill-rule:nonzero;stroke:none;stroke-width:8;stroke-linecap:square;stroke-linejoin:miter;marker:none;marker-start:none;marker-mid:none;marker-end:none;stroke-miterlimit:4;stroke-dashoffset:0;stroke-opacity:1;visibility:visible;display:inline;overflow:visible")
	canvas.Gend()
	canvas.Gend()
	canvas.PatternEnd()
}

// Random returns a random tincture
func Random() Tincture {
	var tinctureChances = map[string]int{
		"color": 20,
		"metal": 20,
		"stain": 3,
		"fur":   5,
	}
	typeOfTincture := random.StringFromThresholdMap(tinctureChances)

	if typeOfTincture == "metal" {
		return randomTinctureMetal()
	} else if typeOfTincture == "color" {
		return randomTinctureColor()
	} else if typeOfTincture == "stain" {
		return randomTinctureStain()
	}

	return randomTinctureFur()
}

func randomTinctureColor() Tincture {
	colors := getColorTinctures()
	t := colors[rand.Intn(len(colors))]
	return t
}

func randomTinctureFur() Tincture {
	furs := getFurTinctures()
	t := furs[rand.Intn(len(furs))]
	return t
}

func randomTinctureMetal() Tincture {
	metals := getMetalTinctures()
	t := metals[rand.Intn(len(metals))]
	return t
}

func randomTinctureStain() Tincture {
	stains := getStainTinctures()
	t := stains[rand.Intn(len(stains))]
	return t
}

// RandomComplementaryTincture returns a random tincture that complements the given one
func RandomComplementaryTincture(t Tincture) Tincture {
	var colorOrStainChances = map[string]int{
		"color": 9,
		"stain": 1,
	}
	availableTinctures := []Tincture{}

	colors := getColorTinctures()
	metals := getMetalTinctures()
	stains := getStainTinctures()

	if t.Type == "color" || t.Type == "stain" {
		typeOfTincture := random.StringFromThresholdMap(colorOrStainChances)
		if typeOfTincture == "color" {
			for _, color := range colors {
				if color.Name != t.Name {
					availableTinctures = append(availableTinctures, color)
				}
			}
		} else {
			for _, stain := range stains {
				if stain.Name != t.Name {
					availableTinctures = append(availableTinctures, stain)
				}
			}
		}
	} else {
		for _, metal := range metals {
			if metal.Name != t.Name {
				availableTinctures = append(availableTinctures, metal)
			}
		}
	}
	t2 := availableTinctures[rand.Intn(len(availableTinctures))]

	return t2
}

// RandomContrastingTincture returns a random tincture that contrasts with the given one
func RandomContrastingTincture(t Tincture) Tincture {
	var colorOrStainChances = map[string]int{
		"color": 9,
		"stain": 1,
	}

	t2 := Tincture{}
	if t.Type == "metal" {
		typeOfTincture := random.StringFromThresholdMap(colorOrStainChances)
		if typeOfTincture == "color" {
			t2 = randomTinctureColor()
		} else {
			t2 = randomTinctureStain()
		}
	} else {
		t2 = randomTinctureMetal()
	}

	return t2
}
