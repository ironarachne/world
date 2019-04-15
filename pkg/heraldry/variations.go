package heraldry

import (
	"math/rand"

	svg "github.com/ajstarks/svgo"
)

var (
	Barry               = Variation{"barry", Tincture{}, Tincture{}}
	Bendy               = Variation{"bendy", Tincture{}, Tincture{}}
	Chequy              = Variation{"chequy", Tincture{}, Tincture{}}
	Paly                = Variation{"paly", Tincture{}, Tincture{}}
	availableVariations = []Variation{Barry, Bendy, Chequy, Paly}
)

func randomVariation() Variation {
	return availableVariations[rand.Intn(len(availableVariations))]
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
		canvas.Rect(0, 0, 160, 40, "fill:"+variation.Tincture1.Hexcode)
		canvas.Rect(0, 40, 160, 40, "fill:"+variation.Tincture2.Hexcode)
		canvas.Rect(0, 80, 160, 40, "fill:"+variation.Tincture1.Hexcode)
		canvas.Rect(0, 120, 160, 40, "fill:"+variation.Tincture2.Hexcode)
	case "bendy":
		canvas.Pattern(variation.Name, 0, 0, 70, 70, "user")
		canvas.Rect(0, 0, 70, 70, "fill:"+variation.Tincture1.Hexcode)
		canvas.Rotate(45.0)
		canvas.Rect(0, 0, 99, 25, "fill:"+variation.Tincture2.Hexcode)
		canvas.Rect(0, -50, 99, 25, "fill:"+variation.Tincture2.Hexcode)
		canvas.Gend()
	case "chequy":
		canvas.Pattern(variation.Name, 0, 0, 80, 80, "user")
		canvas.Rect(0, 0, 40, 40, "fill:"+variation.Tincture1.Hexcode)
		canvas.Rect(40, 0, 40, 40, "fill:"+variation.Tincture2.Hexcode)
		canvas.Rect(0, 40, 40, 40, "fill:"+variation.Tincture2.Hexcode)
		canvas.Rect(40, 40, 40, 40, "fill:"+variation.Tincture1.Hexcode)
	case "paly":
		canvas.Pattern(variation.Name, 0, 0, 160, 160, "user")
		canvas.Rect(0, 0, 40, 160, "fill:"+variation.Tincture1.Hexcode)
		canvas.Rect(40, 0, 40, 160, "fill:"+variation.Tincture2.Hexcode)
		canvas.Rect(80, 0, 40, 160, "fill:"+variation.Tincture1.Hexcode)
		canvas.Rect(120, 0, 40, 160, "fill:"+variation.Tincture2.Hexcode)
	}
	canvas.PatternEnd()
}
