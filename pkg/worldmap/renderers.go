package worldmap

import (
	"bytes"

	svg "github.com/ajstarks/svgo"
)

// RenderAsSVG renders a world map to SVG
func (worldMap WorldMap) RenderAsSVG() string {
	tileWidth := 16
	tileHeight := 16

	buffer := new(bytes.Buffer)
	canvas := svg.New(buffer)
	canvas.Start(worldMap.Width*tileWidth, worldMap.Height*tileHeight)
	for y, row := range worldMap.Tiles {
		for x, tile := range row {
			tile.renderSVG(canvas, tileWidth, tileHeight, x, y)
		}
	}
	canvas.End()

	return buffer.String()
}

// RenderAsText renders a world map to text
func (worldMap WorldMap) RenderAsText() string {
	output := ""

	characters := map[string]string{
		"coniferous forest": "!",
		"deciduous forest":  "*",
		"desert":            "_",
		"grassland":         "-",
		"marshland":         ";",
		"tropical":          "@",
		"mountain":          "^",
		"rainforest":        "?",
		"savanna":           "%",
		"steppe":            "s",
		"taiga":             "t",
		"tundra":            "u",
	}

	for _, r := range worldMap.Tiles {
		for _, t := range r {
			if t.TileType == "ocean" {
				output += "~"
			} else {
				output += characters[t.TileType]
			}
		}
		output += "\n"
	}

	return output
}

func (tile Tile) renderSVG(canvas *svg.SVG, width int, height int, x int, y int) {
	fillValues := map[string]string{
		"coniferous forest": "#21840e",
		"deciduous forest":  "#267717",
		"desert":            "#cfd16e",
		"grassland":         "#629b4a",
		"marshland":         "#0c421d",
		"tropical":          "#3fa517",
		"mountain":          "#5a6768",
		"ocean":             "#cde5f4",
		"rainforest":        "#0eb514",
		"savanna":           "#bec697",
		"steppe":            "#96af83",
		"taiga":             "#0f6d44",
		"tundra":            "#a3d1bd",
	}

	canvas.Rect(x*width, y*height, width, height, "fill:"+fillValues[tile.TileType])
}
