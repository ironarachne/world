package charge

import (
	"github.com/ironarachne/world/pkg/heraldry/tincture"
	"image"
)

func getPeopleCharges() []Charge {
	charges := []Charge{
		{
			Identifier: "chevalier-on-horseback",
			Name:       "chevalier on horseback",
			Noun:       "chevalier",
			NounPlural: "chevaliers",
			Descriptor: "on horseback",
			SingleOnly: false,
			Tags: []string{
				"chevalier",
				"knight",
				"people",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("chevalier-on-horseback", bodyTincture)

				return img
			},
		},
	}

	return charges
}
