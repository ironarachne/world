package heraldry

import (
	"github.com/fogleman/gg"
	"github.com/ironarachne/world/pkg/heraldry/tincture"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/words"
	"image"
	"math/rand"
)

// Division is a division of the field
type Division struct {
	Name   string
	Blazon string
	NumberOfSections int
	Variations []Variation
	Render func(width int, height int, variations []Variation) image.Image
	Commonality int
}

func allDivisions() []Division {
	divisions := []Division{
		{
			Name: "bend",
			Blazon: "Per bend ",
			NumberOfSections: 2,
			Commonality: 2,
			Render: func(width int, height int, variations []Variation) image.Image {
				dc := gg.NewContext(width, height)
				image1 := variations[0].Render(width, height, variations[0].Tinctures)
				image2 := variations[1].Render(width, height, variations[1].Tinctures)
				pattern1 := gg.NewSurfacePattern(image1, gg.RepeatBoth)
				pattern2 := gg.NewSurfacePattern(image2, gg.RepeatBoth)

				dc.DrawRectangle(0,0, float64(width), float64(height))
				dc.SetFillStyle(pattern1)
				dc.Fill()

				dc.SetLineWidth(3)
				dc.MoveTo(0, 0)
				dc.LineTo(float64(width),0)
				dc.LineTo(float64(width), float64(height))
				dc.LineTo(0,0)
				dc.ClosePath()
				dc.SetFillStyle(pattern2)
				dc.Fill()

				return dc.Image()
			},
		},
		{
			Name: "bend sinister",
			Blazon: "Per bend sinister ",
			NumberOfSections: 2,
			Commonality: 2,
			Render: func(width int, height int, variations []Variation) image.Image {
				dc := gg.NewContext(width, height)
				image1 := variations[0].Render(width, height, variations[0].Tinctures)
				image2 := variations[1].Render(width, height, variations[1].Tinctures)
				pattern1 := gg.NewSurfacePattern(image1, gg.RepeatBoth)
				pattern2 := gg.NewSurfacePattern(image2, gg.RepeatBoth)

				dc.DrawRectangle(0,0, float64(width), float64(height))
				dc.SetFillStyle(pattern1)
				dc.Fill()

				dc.SetLineWidth(3)
				dc.MoveTo(float64(width), 0)
				dc.LineTo(float64(width), float64(height))
				dc.LineTo(0, float64(height))
				dc.LineTo(float64(width), 0)
				dc.ClosePath()
				dc.SetFillStyle(pattern2)
				dc.Fill()

				return dc.Image()
			},
		},
		{
			Name: "fess",
			Blazon: "Per fess ",
			NumberOfSections: 2,
			Commonality: 2,
			Render: func(width int, height int, variations []Variation) image.Image {
				dc := gg.NewContext(width, height)
				image1 := variations[0].Render(width, height, variations[0].Tinctures)
				image2 := variations[1].Render(width, height, variations[1].Tinctures)
				pattern1 := gg.NewSurfacePattern(image1, gg.RepeatBoth)
				pattern2 := gg.NewSurfacePattern(image2, gg.RepeatBoth)

				dc.DrawRectangle(0,0, float64(width), float64(height))
				dc.SetFillStyle(pattern1)
				dc.Fill()

				dc.DrawRectangle(0, float64(height/2), float64(width), float64(height))
				dc.SetFillStyle(pattern2)
				dc.Fill()

				return dc.Image()
			},
		},
		{
			Name: "pale",
			Blazon: "Per pale ",
			NumberOfSections: 2,
			Commonality: 2,
			Render: func(width int, height int, variations []Variation) image.Image {
				dc := gg.NewContext(width, height)
				image1 := variations[0].Render(width, height, variations[0].Tinctures)
				image2 := variations[1].Render(width, height, variations[1].Tinctures)
				pattern1 := gg.NewSurfacePattern(image1, gg.RepeatBoth)
				pattern2 := gg.NewSurfacePattern(image2, gg.RepeatBoth)

				dc.DrawRectangle(0,0, float64(width), float64(height))
				dc.SetFillStyle(pattern1)
				dc.Fill()

				dc.DrawRectangle(float64(width/2), 0, float64(width), float64(height))
				dc.SetFillStyle(pattern2)
				dc.Fill()

				return dc.Image()
			},
		},
		{
			Name: "plain",
			Blazon: "",
			NumberOfSections: 1,
			Commonality: 2,
			Render: func(width int, height int, variations []Variation) image.Image {
				dc := gg.NewContext(width, height)
				image1 := variations[0].Render(width, height, variations[0].Tinctures)
				pattern1 := gg.NewSurfacePattern(image1, gg.RepeatBoth)

				dc.DrawRectangle(0,0, float64(width), float64(height))
				dc.SetFillStyle(pattern1)
				dc.Fill()

				return dc.Image()
			},
		},
		{
			Name: "quarterly",
			Blazon: "Quarterly ",
			NumberOfSections: 2,
			Commonality: 1,
			Render: func(width int, height int, variations []Variation) image.Image {
				dc := gg.NewContext(width, height)
				image1 := variations[0].Render(width, height, variations[0].Tinctures)
				image2 := variations[1].Render(width, height, variations[1].Tinctures)
				pattern1 := gg.NewSurfacePattern(image1, gg.RepeatBoth)
				pattern2 := gg.NewSurfacePattern(image2, gg.RepeatBoth)

				dc.DrawRectangle(0,0, float64(width), float64(height))
				dc.SetFillStyle(pattern1)
				dc.Fill()

				dc.DrawRectangle(float64(width/2), 0, float64(width), float64(height/2))
				dc.SetFillStyle(pattern2)
				dc.Fill()

				dc.DrawRectangle(0, float64(height/2), float64(width/2), float64(height))
				dc.SetFillStyle(pattern2)
				dc.Fill()

				return dc.Image()
			},
		},
		{
			Name: "saltire",
			Blazon: "Per saltire ",
			NumberOfSections: 2,
			Commonality: 1,
			Render: func(width int, height int, variations []Variation) image.Image {
				dc := gg.NewContext(width, height)
				image1 := variations[0].Render(width, height, variations[0].Tinctures)
				image2 := variations[1].Render(width, height, variations[1].Tinctures)
				pattern1 := gg.NewSurfacePattern(image1, gg.RepeatBoth)
				pattern2 := gg.NewSurfacePattern(image2, gg.RepeatBoth)

				dc.DrawRectangle(0,0, float64(width), float64(height))
				dc.SetFillStyle(pattern1)
				dc.Fill()

				dc.MoveTo(0, 0)
				dc.LineTo(float64(width/2), float64(height/2))
				dc.LineTo(float64(width), 0)
				dc.LineTo(0,0)
				dc.MoveTo(0, float64(height))
				dc.LineTo(float64(width/2), float64(height/2))
				dc.LineTo(float64(width), float64(height))
				dc.LineTo(0,float64(height))
				dc.SetFillStyle(pattern2)
				dc.Fill()

				return dc.Image()
			},
		},
		{
			Name: "chevron",
			Blazon: "Per chevron ",
			NumberOfSections: 2,
			Commonality: 1,
			Render: func(width int, height int, variations []Variation) image.Image {
				dc := gg.NewContext(width, height)
				image1 := variations[0].Render(width, height, variations[0].Tinctures)
				image2 := variations[1].Render(width, height, variations[1].Tinctures)
				pattern1 := gg.NewSurfacePattern(image1, gg.RepeatBoth)
				pattern2 := gg.NewSurfacePattern(image2, gg.RepeatBoth)

				dc.DrawRectangle(0,0, float64(width), float64(height))
				dc.SetFillStyle(pattern1)
				dc.Fill()

				dc.SetLineWidth(3)
				dc.MoveTo(0, float64(height))
				dc.LineTo(float64(width/2), float64(height*2/3))
				dc.LineTo(float64(width), float64(height))
				dc.LineTo(0,float64(height))
				dc.ClosePath()
				dc.SetFillStyle(pattern2)
				dc.Fill()

				return dc.Image()
			},
		},
	}

	return divisions
}

func randomDivision() Division {
	divisions := allDivisions()
	return divisions[rand.Intn(len(divisions))]
}

func randomWeightedDivision() Division {
	all := allDivisions()

	weights := map[string]int{}

	for _, c := range all {
		weights[c.Name] = c.Commonality
	}

	name := random.StringFromThresholdMap(weights)

	for _, c := range all {
		if c.Name == name {
			return c
		}
	}

	return Division{}
}

func generateDivision() Division {
	var divisionBlazons []string
	var possible []tincture.Tincture
	var variation Variation
	var tinc tincture.Tincture
	division := randomWeightedDivision()

	firstTincture := tincture.RandomAll()
	lastTincture := firstTincture

	for i:=0;i<division.NumberOfSections;i++ {
		possible = tincture.All()
		possible = tincture.Remove(lastTincture, possible)
		tinc = tincture.RandomWeighted(possible)
		variation = generateVariation(tinc)
		division.Variations = append(division.Variations, variation)
		divisionBlazons = append(divisionBlazons, variation.Blazon)
		lastTincture = tinc
	}

	division.Blazon += words.CombinePhrases(divisionBlazons)

	return division
}
