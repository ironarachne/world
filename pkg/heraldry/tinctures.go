package heraldry

import svg "github.com/ajstarks/svgo"

var (
	Argent          = Tincture{"metal", "argent", "#FFFFFF"}
	Azure           = Tincture{"color", "azure", "#0000FF"}
	Erminois        = Tincture{"fur", "erminois", "url(#erminois)"}
	Ermine          = Tincture{"fur", "ermine", "url(#ermine)"}
	Ermines         = Tincture{"fur", "ermines", "url(#ermines)"}
	Pean            = Tincture{"fur", "pean", "url(#pean)"}
	Gules           = Tincture{"color", "gules", "#EE0000"}
	Murrey          = Tincture{"stain", "murrey", "#591334"}
	Or              = Tincture{"metal", "Or", "#FFEC00"}
	Purpure         = Tincture{"color", "purpure", "#831CA6"}
	Sable           = Tincture{"color", "sable", "#000000"}
	Sanguine        = Tincture{"stain", "sanguine", "#720101"}
	Tenne           = Tincture{"stain", "tenné", "#7C4215"}
	Vert            = Tincture{"color", "vert", "#009900"}
	Furs            = []Tincture{Erminois, Ermine, Ermines, Pean}
	Metals          = []Tincture{Or, Argent}
	Colors          = []Tincture{Sable, Gules, Vert, Azure, Purpure}
	Stains          = []Tincture{Murrey, Sanguine, Tenne}
	tinctureChances = map[string]int{
		"color": 20,
		"metal": 20,
		"stain": 3,
		"fur":   5,
	}
	colorOrStainChances = map[string]int{
		"color": 9,
		"stain": 1,
	}
)

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
