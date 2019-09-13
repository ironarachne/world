package charge

import (
	"github.com/ironarachne/world/pkg/heraldry/tincture"
	"image"
)

func getHeavensCharges() []Charge {
	charges := []Charge{
		{
			Identifier: "estoille",
			Name:       "estoille",
			Noun:       "estoille",
			NounPlural: "estoilles",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"estoille",
				"heavens",
				"star",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("estoille", bodyTincture)

				return img
			},
		},
		{
			Identifier: "sun-in-splendor",
			Name:       "sun in splendor",
			Noun:       "sun",
			NounPlural: "suns",
			Descriptor: "in splendor",
			SingleOnly: false,
			Tags: []string{
				"sun",
				"heavens",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("sun-in-splendor", bodyTincture)

				return img
			},
		},
	}

	return charges
}
