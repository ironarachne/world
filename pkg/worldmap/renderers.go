package worldmap

import (
	"bytes"

	svg "github.com/ajstarks/svgo"
	"github.com/ironarachne/world/pkg/grid"
)

// RenderAsSVG renders a world map to SVG
func (worldMap WorldMap) RenderAsSVG() string {
	var labelMidpoint grid.Coordinate
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
	for _, b := range worldMap.Boundaries {
		labelMidpoint = grid.Coordinate{X: 0, Y: 0}
		for _, c := range b.TileCoordinates {
			labelMidpoint.X += c.X * tileWidth
			labelMidpoint.Y += c.Y * tileHeight
			canvas.Rect(c.X*tileWidth, c.Y*tileHeight, tileWidth, tileHeight, "fill:#FF0000;fill-opacity:0.5")
		}
		labelMidpoint.X = int(labelMidpoint.X / len(b.TileCoordinates))
		labelMidpoint.Y = int(labelMidpoint.Y/len(b.TileCoordinates)) + tileHeight
		canvas.Text(labelMidpoint.X, labelMidpoint.Y, b.Label, "fill:#000000;font-size:20px")
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
