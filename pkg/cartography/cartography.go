package cartography

import (
	"github.com/fogleman/gg"
	"github.com/ironarachne/world/pkg/world"
	"image"
)

// RenderMap renders the given map and returns it as an image
func RenderMap(width int, height int) image.Image {
	dc := gg.NewContext(width, height)
	dc.SetHexColor("#FFFFFF")
	dc.Clear()

	w := world.Generate()

	tileHeight := height/len(w.Tiles)
	tileWidth := width/len(w.Tiles[0])

	rows := len(w.Tiles)
	columns := len(w.Tiles[0])

	for i:=0;i<rows;i++ {
		for j:=0;j<columns;j++ {
			dc.DrawRectangle(float64(j*tileWidth), float64(i*tileHeight), float64(tileWidth), float64(tileHeight))
			if w.Tiles[i][j].IsOcean {
				if w.Tiles[i][j].Altitude < 10 {
					dc.SetHexColor("#55AAEB")
				} else if w.Tiles[i][j].Altitude > 18 {
					dc.SetHexColor("#ABCDEF")
				} else {
					dc.SetHexColor("#73BFFA")
				}
			} else {
				if w.Tiles[i][j].Altitude < 35 {
					dc.SetHexColor("#00AA33")
				} else if w.Tiles[i][j].Altitude < 60 {
					dc.SetHexColor("#33BB55")
				} else if w.Tiles[i][j].Altitude < 80 {
					dc.SetHexColor("#AABBBB")
				} else {
					dc.SetHexColor("#EEFFFF")
				}
			}
			dc.Fill()
		}
	}

	m := dc.Image()

	return m
}
