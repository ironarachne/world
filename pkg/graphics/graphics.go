package graphics

import (
	"github.com/fogleman/gg"
	"image"
	"image/color"
)

// Pattern is a fill pattern
type Pattern struct {
	Type string
	Color color.RGBA
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
	rMinX := centerPoint.X - (width/2)
	rMinY := centerPoint.Y - (height/2)
	rMaxX := centerPoint.X + (width/2)
	rMaxY := centerPoint.Y + (height/2)
	centeredRectangle := image.Rectangle{
		Min: image.Point{X:rMinX, Y:rMinY},
		Max: image.Point{X:rMaxX, Y:rMaxY},
	}

	return centeredRectangle
}

// LoadPNG loads a PNG image and returns it as an Image
func LoadPNG(path string) image.Image {
	filePath := path
	im, err := gg.LoadPNG(filePath)
	if err != nil {
		panic(err)
	}
	return im
}

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
	dc.DrawRectangle(0,0, float64(dc.Width()), float64(dc.Height()))
	pattern.Fill(dc)

	newImage := dc.Image()
	return newImage
}
