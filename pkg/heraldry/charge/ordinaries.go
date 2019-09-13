package charge

import (
	"github.com/fogleman/gg"
	"github.com/ironarachne/world/pkg/heraldry/tincture"
	"image"
)

func getOrdinaryCharges() []Charge {
	charges := []Charge{
		{
			Identifier: "cross",
			Name:       "cross",
			Noun:       "cross",
			NounPlural: "crosses",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"full size",
				"ordinary",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				width := 350
				height := 450

				dc := gg.NewContext(width, height)

				dc.DrawRectangle(110, 0, 130, float64(height))
				dc.DrawRectangle(0, 140, float64(width), 160)
				bodyTincture.Fill(dc)

				img := dc.Image()

				return img
			},
		},
		{
			Identifier: "fess",
			Name:       "fess",
			Noun:       "fess",
			NounPlural: "fesses",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"full size",
				"ordinary",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				width := 350
				height := 450

				dc := gg.NewContext(width, height)

				dc.DrawRectangle(0, 140, float64(width), 160)
				bodyTincture.Fill(dc)

				img := dc.Image()

				return img
			},
		},
		{
			Identifier: "pale",
			Name:       "pale",
			Noun:       "pale",
			NounPlural: "pales",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"full size",
				"ordinary",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				width := 350
				height := 450

				dc := gg.NewContext(width, height)

				dc.DrawRectangle(110, 0, 130, float64(height))
				bodyTincture.Fill(dc)

				img := dc.Image()

				return img
			},
		},
		{
			Identifier: "roundel",
			Name:       "roundel",
			Noun:       "roundel",
			NounPlural: "roundels",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"ordinary",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				width := 400
				height := 400

				dc := gg.NewContext(width, height)

				dc.DrawCircle(float64(width/2), float64(height/2), float64(width/3))
				bodyTincture.Fill(dc)

				img := dc.Image()

				return img
			},
		},
	}

	return charges
}
