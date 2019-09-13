package charge

import (
	"github.com/ironarachne/world/pkg/heraldry/tincture"
	"image"
)

func getPlantCharges() []Charge {
	charges := []Charge{
		{
			Identifier: "acorn-slipped-and-leaved",
			Name:       "acorn slipped and leaved",
			Noun:       "acorn",
			NounPlural: "acorns",
			Descriptor: "slipped and leaved",
			SingleOnly: false,
			Tags: []string{
				"acorn",
				"plant",
				"harvest",
				"farmer",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("acorn-slipped-and-leaved", bodyTincture)

				return img
			},
		},
	}

	return charges
}
