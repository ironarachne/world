package charge

import (
	"github.com/ironarachne/world/pkg/heraldry/tincture"
	"image"
)

func getObjectCharges() []Charge {
	charges := []Charge{
		{
			Identifier: "anchor",
			Name:       "anchor",
			Noun:       "anchor",
			NounPlural: "anchors",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"anchor",
				"object",
				"water",
				"navy",
				"sailor",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("anchor", bodyTincture)

				return img
			},
		},
		{
			Identifier: "annulet",
			Name:       "annulet",
			Noun:       "annulet",
			NounPlural: "annulets",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"annulet",
				"object",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("annulet", bodyTincture)

				return img
			},
		},
		{
			Identifier: "axe",
			Name:       "axe",
			Noun:       "axe",
			NounPlural: "axes",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"axe",
				"object",
				"weapon",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("axe", bodyTincture)

				return img
			},
		},
		{
			Identifier: "galleon",
			Name:       "galleon",
			Noun:       "galleon",
			NounPlural: "galleons",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"galleon",
				"object",
				"water",
				"navy",
				"sailor",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("galleon", bodyTincture)

				return img
			},
		},
	}

	return charges
}
