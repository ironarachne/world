package mask

import (
	"fmt"
	"github.com/ironarachne/world/pkg/math"
	"image"
	"image/color"
)

type Triangle struct {
	P1 image.Point
	P2 image.Point
	P3 image.Point
}

func (t *Triangle) ColorModel() color.Model {
	return color.AlphaModel
}

func (t *Triangle) Bounds() image.Rectangle {
	minX := math.Min(t.P1.X, t.P2.X)
	minX = math.Min(minX, t.P3.X)
	minY := math.Min(t.P1.Y, t.P2.Y)
	minY = math.Min(minY, t.P3.Y)
	maxX := math.Max(t.P1.X, t.P2.X)
	maxX = math.Max(maxX, t.P3.X)
	maxY := math.Max(t.P1.Y, t.P2.Y)
	maxY = math.Max(maxY, t.P3.Y)

	return image.Rect(minX, minY, maxX, maxY)
}

func (t *Triangle) At(x, y int) color.Color {
	a := t.Area()
	a1 := Triangle{P1: image.Point{X: x, Y: y}, P2: image.Point{X: t.P1.X, Y: t.P1.Y}, P3: image.Point{X: t.P2.X, Y: t.P2.Y}}.Area()
	a2 := Triangle{P1: image.Point{X: x, Y: y}, P2: image.Point{X: t.P2.X, Y: t.P2.Y}, P3: image.Point{X: t.P3.X, Y: t.P3.Y}}.Area()
	a3 := Triangle{P1: image.Point{X: x, Y: y}, P2: image.Point{X: t.P3.X, Y: t.P3.Y}, P3: image.Point{X: t.P1.X, Y: t.P1.Y}}.Area()

	fmt.Println("Area of combined triangles: ", a1+a2+a3)
	fmt.Println("Area of triangle: ", a)

	for i := -50; i < 50; i++ {
		if a+i == a1+a2+a3 {
			return color.Alpha{255}
		}
	}

	return color.Alpha{0}
}

func (t Triangle) Area() int {
	return (t.P1.X*(t.P2.Y-t.P3.Y) + t.P2.X*(t.P3.Y-t.P1.Y) + t.P3.X*(t.P1.Y-t.P2.Y)) / 2
}
