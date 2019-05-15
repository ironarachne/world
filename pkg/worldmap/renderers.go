package worldmap

import (
	"bytes"

	svg "github.com/ajstarks/svgo"
)

// RenderAsSVG renders a world map to SVG
func (worldMap WorldMap) RenderAsSVG() string {
	tileWidth := 10
	tileHeight := 10
	tileColor := ""
	buffer := new(bytes.Buffer)
	canvas := svg.New(buffer)
	canvas.Start(worldMap.Width*tileWidth, worldMap.Height*tileHeight)
	for y, row := range worldMap.Tiles {
		for x, tile := range row {
			if tile.IsOcean {
				tileColor = "#CCFFFF"
			} else {
				tileColor = "#55FF55"
			}
			canvas.Rect(x*tileWidth, y*tileHeight, tileWidth, tileHeight, "fill:"+tileColor)
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
