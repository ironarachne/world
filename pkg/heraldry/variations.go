package heraldry

import (
	"math/rand"

	svg "github.com/ajstarks/svgo"
)

func allVariations() []Variation {
	variations := []Variation{
		{
			Name: "barry",
		},
		{
			Name: "bendy",
		},
		{
			Name: "chequy",
		},
		{
			Name: "paly",
		},
	}

	return variations
}

func randomVariation() Variation {
	all := allVariations()
	return all[rand.Intn(len(all))]
}

func shallWeHaveAVariation() bool {
	someRandomValue := rand.Intn(100)
	if someRandomValue > 90 {
		return true
	}
	return false
}

func insertVariationPattern(canvas *svg.SVG, variation Variation) {
	switch variation.Name {
	case "barry":
		canvas.Pattern(variation.Name, 0, 0, 160, 160, "user")
		canvas.Rect(0, 0, 160, 40, "fill:"+variation.Tincture1.HexCode)
		canvas.Rect(0, 40, 160, 40, "fill:"+variation.Tincture2.HexCode)
		canvas.Rect(0, 80, 160, 40, "fill:"+variation.Tincture1.HexCode)
		canvas.Rect(0, 120, 160, 40, "fill:"+variation.Tincture2.HexCode)
	case "bendy":
		canvas.Pattern(variation.Name, 0, 0, 70, 70, "user")
		canvas.Rect(0, 0, 70, 70, "fill:"+variation.Tincture1.HexCode)
		canvas.Rotate(45.0)
		canvas.Rect(0, 0, 99, 25, "fill:"+variation.Tincture2.HexCode)
		canvas.Rect(0, -50, 99, 25, "fill:"+variation.Tincture2.HexCode)
		canvas.Gend()
	case "chequy":
		canvas.Pattern(variation.Name, 0, 0, 80, 80, "user")
		canvas.Rect(0, 0, 40, 40, "fill:"+variation.Tincture1.HexCode)
		canvas.Rect(40, 0, 40, 40, "fill:"+variation.Tincture2.HexCode)
		canvas.Rect(0, 40, 40, 40, "fill:"+variation.Tincture2.HexCode)
		canvas.Rect(40, 40, 40, 40, "fill:"+variation.Tincture1.HexCode)
	case "paly":
		canvas.Pattern(variation.Name, 0, 0, 160, 160, "user")
		canvas.Rect(0, 0, 40, 160, "fill:"+variation.Tincture1.HexCode)
		canvas.Rect(40, 0, 40, 160, "fill:"+variation.Tincture2.HexCode)
		canvas.Rect(80, 0, 40, 160, "fill:"+variation.Tincture1.HexCode)
		canvas.Rect(120, 0, 40, 160, "fill:"+variation.Tincture2.HexCode)
	}
	canvas.PatternEnd()
}
