package heraldry

import svg "github.com/ajstarks/svgo"

func renderBordureToSvg(canvas *svg.SVG, hexcode string) {
	pathData := "m10.273 21.598v151.22c0 96.872 89.031 194.34 146.44 240.09 57.414-45.758 146.44-143.22 146.44-240.09v-151.22h-292.89z"
	canvas.Path(pathData, "stroke:"+hexcode+";stroke-width:100;fill:none")
}
