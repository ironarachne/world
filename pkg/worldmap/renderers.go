package worldmap

import (
	"bytes"
	"github.com/ironarachne/world/pkg/heraldry/tincture"

	svg "github.com/ajstarks/svgo"
	"github.com/ironarachne/world/pkg/grid"
	"github.com/ironarachne/world/pkg/slices"
)

// RenderAsSVG renders a world map to SVG
func (worldMap WorldMap) RenderAsSVG() string {
	var labelMidpoint grid.Coordinate
	var line string
	var xs []int
	var ys []int
	var maxX int
	var maxY int
	var minX int
	var minY int

	tileWidth := worldMap.TileWidth
	tileHeight := worldMap.TileHeight

	existingPatterns := []string{}

	buffer := new(bytes.Buffer)
	canvas := svg.New(buffer)
	canvas.Start(worldMap.Width*tileWidth, worldMap.Height*tileHeight)
	canvas.Def()
	for _, b := range worldMap.Boundaries {
		if b.Pattern != "" && !slices.StringIn(b.Pattern, existingPatterns) {
			tincture.InsertErmine(canvas, b.Pattern)
			existingPatterns = append(existingPatterns, b.Pattern)
		}
	}
	canvas.DefEnd()
	for y, row := range worldMap.Tiles {
		for x, tile := range row {
			tile.renderSVG(canvas, tileWidth, tileHeight, x, y)
		}
	}
	for _, b := range worldMap.Boundaries {
		labelMidpoint = grid.Coordinate{X: 0, Y: 0}
		line = "M"
		xs = []int{}
		ys = []int{}
		minX = b.Vertices[0].X
		minY = b.Vertices[0].Y
		maxX = 0
		maxY = 0
		for _, v := range b.Vertices {
			if minX > v.X {
				minX = v.X
			}
			if minY > v.Y {
				minY = v.Y
			}
			if maxX < v.X {
				maxX = v.X
			}
			if maxY < v.Y {
				maxY = v.Y
			}
			xs = append(xs, v.X)
			ys = append(ys, v.Y)
		}
		labelMidpoint.X = int((minX + maxX) / 2)
		labelMidpoint.Y = int((minY + maxY) / 2)
		canvas.Polygon(xs, ys, b.Style)
		canvas.Path(line, b.Style)
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
		"ocean":             "#1744d6",
		"rainforest":        "#0eb514",
		"savanna":           "#bec697",
		"steppe":            "#96af83",
		"taiga":             "#0f6d44",
		"tundra":            "#a3d1bd",
	}

	canvas.Rect(x*width, y*height, width, height, "fill:"+fillValues[tile.TileType])
}
