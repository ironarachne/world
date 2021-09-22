package cartography

import (
	"context"
	"fmt"
	"github.com/ironarachne/world/pkg/world"
	"image"
	"image/color"
)

// RenderMap renders the given map and returns it as an image
func RenderMap(ctx context.Context, width int, height int) (image.Image, error) {
	var c color.RGBA
	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	w, err := world.Generate(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to generate world map: %w", err)
	}

	rows := len(w.Tiles)
	columns := len(w.Tiles[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			c = getColor(w.Tiles[i][j])
			img.Set(j, i, c)
		}
	}

	return img, nil
}

func getColor(tile world.Tile) color.RGBA {
	if tile.IsOcean {
		if tile.Altitude < -20 {
			return color.RGBA{R: 85, G: 170, B: 235, A: 255} // Open Ocean
		} else {
			return color.RGBA{R: 115, G: 192, B: 250, A: 255} // Coastal Shelf
		}
	}

	if tile.Altitude > 80 {
		return color.RGBA{R: 238, G: 255, B: 255, A: 255} // Mountains
	}

	if tile.Temperature > 80 && tile.Humidity > 80 {
		return color.RGBA{R: 39, G: 145, B: 17, A: 255} // Rainforest
	}

	if tile.Temperature > 80 && tile.Humidity < 30 {
		return color.RGBA{R: 207, G: 205, B: 135, A: 255} // Desert
	}

	if tile.Temperature < 30 && tile.Humidity < 30 {
		return color.RGBA{R: 139, G: 173, B: 151, A: 255} // Tundra
	}

	if tile.Humidity > 60 {
		return color.RGBA{R: 41, G: 110, B: 59, A: 255} // Swamps and Marshes
	}

	if tile.Humidity > 30 {
		return color.RGBA{R: 115, G: 156, B: 76, A: 255} // Plains
	}

	return color.RGBA{R: 108, G: 153, B: 73, A: 255} // Default to warm green
}
