/*
Package graphics provides tools for generating images.
*/
package graphics

import (
	"github.com/ironarachne/world/pkg/geometry"
	"image"
	"image/color"
	"os"

	"github.com/fogleman/gg"
)

// Pattern is a fill pattern
type Pattern struct {
	Type            string
	Color           color.RGBA
	PatternFileName string
}

// Center returns the center point of a given rectangle
func Center(rect image.Rectangle) image.Point {
	centerX := (rect.Max.X - rect.Min.X) / 2
	centerY := (rect.Max.Y - rect.Min.Y) / 2

	point := image.Point{
		X: centerX,
		Y: centerY,
	}

	return point
}

// CenteredRectangle returns a rectangle of the given dimensions centered on the given point
func CenteredRectangle(centerPoint image.Point, width int, height int) image.Rectangle {
	rMinX := centerPoint.X - (width / 2)
	rMinY := centerPoint.Y - (height / 2)
	rMaxX := centerPoint.X + (width / 2)
	rMaxY := centerPoint.Y + (height / 2)
	centeredRectangle := image.Rectangle{
		Min: image.Point{X: rMinX, Y: rMinY},
		Max: image.Point{X: rMaxX, Y: rMaxY},
	}

	return centeredRectangle
}

// DrawCircle returns an image with a circle drawn matching the given input on the given image
func DrawCircle(circle geometry.Circle, color string, canvas image.Image) image.Image {
	c := gg.NewContextForImage(canvas)
	c.DrawCircle(circle.Center.X, circle.Center.Y, circle.Radius)
	c.SetHexColor(color)
	c.SetLineWidth(1)
	c.Stroke()

	newImage := c.Image()

	return newImage
}

// DrawEdges returns an image with all of the given edges rendered on the given image
func DrawEdges(edges []geometry.Edge, color string, canvas image.Image) image.Image {
	c := gg.NewContextForImage(canvas)
	for _, e := range edges {
		c.DrawLine(e.A.X, e.A.Y, e.B.X, e.B.Y)
		c.SetHexColor(color)
		c.Stroke()
	}

	newImage := c.Image()

	return newImage
}

// DrawEdge returns an image with all of the given edges rendered on the given image
func DrawEdge(edge geometry.Edge, color string, canvas image.Image) image.Image {
	im := DrawEdges([]geometry.Edge{edge}, color, canvas)

	return im
}

// DrawPoints returns an image with all of the given points rendered on the given image
func DrawPoints(points []geometry.Point, color string, canvas image.Image) image.Image {
	c := gg.NewContextForImage(canvas)
	for _, p := range points {
		c.DrawPoint(p.X, p.Y, 1)
		c.SetHexColor(color)
		c.Fill()
	}

	newImage := c.Image()

	return newImage
}

// DrawPolygon returns an image with the given polygon drawn on the given image
func DrawPolygon(polygon geometry.Polygon, color string, canvas image.Image) image.Image {
	newImage := DrawEdges(polygon.Edges, color, canvas)

	return newImage
}

// DrawTriangle returns an image with the given triangle drawn on the given image
func DrawTriangle(triangle geometry.Triangle, color string, canvas image.Image) image.Image {
	var edges []geometry.Edge
	edges = append(edges, triangle.AB)
	edges = append(edges, triangle.BC)
	edges = append(edges, triangle.CA)

	newImage := DrawEdges(edges, color, canvas)

	return newImage
}

// LoadPNG loads a PNG image and returns it as an Image. It relies on the WORLDAPI_DATA_PATH
// environment variable to determine where to load the file from.
func LoadPNG(path string) image.Image {
	dataPath := os.Getenv("WORLDAPI_DATA_PATH")

	filePath := dataPath + "/" + path
	im, err := gg.LoadPNG(filePath)
	if err != nil {
		panic(err)
	}
	return im
}

// Fill fills an image with a pattern
func (pattern Pattern) Fill(dc *gg.Context) {
	if pattern.Type == "image" {
		filePath := "images/patterns/" + pattern.PatternFileName
		im, err := gg.LoadPNG(filePath)
		if err != nil {
			panic(err)
		}
		pattern := gg.NewSurfacePattern(im, gg.RepeatBoth)
		dc.SetFillStyle(pattern)
	} else {
		dc.SetColor(pattern.Color)
	}
	dc.Fill()
}

// RenderFilledImage returns an image filled with a particular pattern
func RenderFilledImage(path string, pattern Pattern) image.Image {
	maskImage := LoadPNG(path)

	mc := gg.NewContextForImage(maskImage)
	imageMask := mc.AsMask()

	dc := gg.NewContextForImage(maskImage)
	err := dc.SetMask(imageMask)
	if err != nil {
		panic("Could not set mask for pattern image")
	}
	dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
	pattern.Fill(dc)

	newImage := dc.Image()
	return newImage
}
