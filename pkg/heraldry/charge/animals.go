package charge

import (
	"github.com/ironarachne/world/pkg/heraldry/tincture"
	"image"
)

func getAnimalCharges() []Charge {
	charges := []Charge{
		{
			Identifier: "allocamelus",
			Name:       "allocamelus",
			Noun:       "allocamelus",
			NounPlural: "allocameluses",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"allocamelus",
				"animal",
				"desert",
				"standing",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("allocamelus", bodyTincture)

				return img
			},
		},
		{
			Identifier: "alphyn-salient",
			Name:       "alphyn salient",
			Noun:       "alphyn",
			NounPlural: "alphyns",
			Descriptor: "salient",
			SingleOnly: false,
			Tags: []string{
				"alphyn",
				"monster",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("alphyn-salient", bodyTincture)

				return img
			},
		},
		{
			Identifier: "antelope-passant",
			Name:       "antelope passant",
			Noun:       "antelope",
			NounPlural: "antelopes",
			Descriptor: "passant",
			SingleOnly: false,
			Tags: []string{
				"animal",
				"antelope",
				"hunting",
				"passive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("antelope-passant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "antelope-rampant",
			Name:       "antelope rampant",
			Noun:       "antelope",
			NounPlural: "antelopes",
			Descriptor: "rampant",
			SingleOnly: false,
			Tags: []string{
				"animal",
				"antelope",
				"hunting",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("antelope-rampant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "antelope-statant",
			Name:       "antelope statant",
			Noun:       "antelope",
			NounPlural: "antelopes",
			Descriptor: "statant",
			SingleOnly: false,
			Tags: []string{
				"animal",
				"antelope",
				"hunting",
				"standing",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("antelope-statant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "ape",
			Name:       "ape",
			Noun:       "ape",
			NounPlural: "apes",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"ape",
				"animal",
				"passive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("ape", bodyTincture)

				return img
			},
		},
		{
			Identifier: "ape-gorged-and-chained",
			Name:       "ape gorged and chained",
			Noun:       "ape",
			NounPlural: "apes",
			Descriptor: "gorged and chained",
			SingleOnly: false,
			Tags: []string{
				"ape",
				"animal",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("ape-gorged-and-chained", bodyTincture)

				return img
			},
		},
		{
			Identifier: "bat-volant",
			Name:       "bat volant",
			Noun:       "bat",
			NounPlural: "bats",
			Descriptor: "volant",
			SingleOnly: false,
			Tags: []string{
				"bat",
				"animal",
				"passive",
				"hunting",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("bat-volant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "bear-passant",
			Name:       "bear passant",
			Noun:       "bear",
			NounPlural: "bears",
			Descriptor: "passant",
			SingleOnly: false,
			Tags: []string{
				"bear",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("bear-passant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "bear-rampant",
			Name:       "bear rampant",
			Noun:       "bear",
			NounPlural: "bears",
			Descriptor: "rampant",
			SingleOnly: false,
			Tags: []string{
				"bear",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("bear-rampant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "bear-sejant-erect",
			Name:       "bear sejant erect",
			Noun:       "bear",
			NounPlural: "bears",
			Descriptor: "sejant erect",
			SingleOnly: false,
			Tags: []string{
				"bear",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("bear-sejant-erect", bodyTincture)

				return img
			},
		},
		{
			Identifier: "bear-statant",
			Name:       "bear statant",
			Noun:       "bear",
			NounPlural: "bears",
			Descriptor: "statant",
			SingleOnly: false,
			Tags: []string{
				"bear",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("bear-statant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "bears-head-couped",
			Name:       "bear's head couped",
			Noun:       "bear's head",
			NounPlural: "bear's heads",
			Descriptor: "couped",
			SingleOnly: false,
			Tags: []string{
				"bear",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("bears-head-couped", bodyTincture)

				return img
			},
		},
		{
			Identifier: "boar-passant",
			Name:       "boar passant",
			Noun:       "boar",
			NounPlural: "boars",
			Descriptor: "passant",
			SingleOnly: false,
			Tags: []string{
				"boar",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("boar-passant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "boar-rampant",
			Name:       "boar rampant",
			Noun:       "boar",
			NounPlural: "boars",
			Descriptor: "rampant",
			SingleOnly: false,
			Tags: []string{
				"boar",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("boar-rampant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "boar-statant",
			Name:       "boar statant",
			Noun:       "boar",
			NounPlural: "boars",
			Descriptor: "statant",
			SingleOnly: false,
			Tags: []string{
				"boar",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("boar-statant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "boars-head-erased",
			Name:       "boar's head erased",
			Noun:       "boar's head",
			NounPlural: "boar's heads",
			Descriptor: "erased",
			SingleOnly: false,
			Tags: []string{
				"boar",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("boars-head-erased", bodyTincture)

				return img
			},
		},
		{
			Identifier: "bull-passant",
			Name:       "bull passant",
			Noun:       "bull",
			NounPlural: "bulls",
			Descriptor: "passant",
			SingleOnly: false,
			Tags: []string{
				"bull",
				"animal",
				"aggressive",
				"strength",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("bull-passant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "bull-rampant",
			Name:       "bull rampant",
			Noun:       "bull",
			NounPlural: "bulls",
			Descriptor: "rampant",
			SingleOnly: false,
			Tags: []string{
				"bull",
				"animal",
				"aggressive",
				"strength",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("bull-rampant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "escallop",
			Name:       "escallop",
			Noun:       "escallop",
			NounPlural: "escallops",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"escallop",
				"animal",
				"water",
				"navy",
				"sailor",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("escallop", bodyTincture)

				return img
			},
		},
		{
			Identifier: "fox-passant",
			Name:       "fox passant",
			Noun:       "fox",
			NounPlural: "foxes",
			Descriptor: "passant",
			SingleOnly: false,
			Tags: []string{
				"fox",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("fox-passant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "fox-sejant",
			Name:       "fox sejant",
			Noun:       "fox",
			NounPlural: "foxes",
			Descriptor: "sejant",
			SingleOnly: false,
			Tags: []string{
				"fox",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("fox-sejant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "goat-salient",
			Name:       "goat salient",
			Noun:       "goat",
			NounPlural: "goats",
			Descriptor: "salient",
			SingleOnly: false,
			Tags: []string{
				"goat",
				"animal",
				"hunting",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("goat-salient", bodyTincture)

				return img
			},
		},
		{
			Identifier: "hare",
			Name:       "hare",
			Noun:       "hare",
			NounPlural: "hares",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"hare",
				"animal",
				"passive",
				"hunting",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("hare", bodyTincture)

				return img
			},
		},
		{
			Identifier: "hare-salient",
			Name:       "hare salient",
			Noun:       "hare",
			NounPlural: "hares",
			Descriptor: "salient",
			SingleOnly: false,
			Tags: []string{
				"hare",
				"animal",
				"passive",
				"hunting",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("hare-salient", bodyTincture)

				return img
			},
		},
		{
			Identifier: "hind-statant",
			Name:       "hind statant",
			Noun:       "hind",
			NounPlural: "hinds",
			Descriptor: "statant",
			SingleOnly: false,
			Tags: []string{
				"hind",
				"animal",
				"passive",
				"hunting",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("hind-statant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "horse-passant",
			Name:       "horse passant",
			Noun:       "horse",
			NounPlural: "horses",
			Descriptor: "passant",
			SingleOnly: false,
			Tags: []string{
				"horse",
				"animal",
				"passive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("horse-passant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "horse-rampant",
			Name:       "horse rampant",
			Noun:       "horse",
			NounPlural: "horses",
			Descriptor: "rampant",
			SingleOnly: false,
			Tags: []string{
				"horse",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("horse-rampant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "leopards-head-erased",
			Name:       "leopard's head erased",
			Noun:       "leopard's head",
			NounPlural: "leopard's heads",
			Descriptor: "erased",
			SingleOnly: false,
			Tags: []string{
				"leopard",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("leopards-head-erased", bodyTincture)

				return img
			},
		},
		{
			Identifier: "otter",
			Name:       "otter",
			Noun:       "otter",
			NounPlural: "otters",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"otter",
				"animal",
				"passive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("otter", bodyTincture)

				return img
			},
		},
		{
			Identifier: "pegasus-passant",
			Name:       "pegasus passant",
			Noun:       "pegasus",
			NounPlural: "pegasi",
			Descriptor: "passant",
			SingleOnly: false,
			Tags: []string{
				"pegasus",
				"monster",
				"passive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("pegasus-passant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "pegasus-rampant",
			Name:       "pegasus rampant",
			Noun:       "pegasus",
			NounPlural: "pegasi",
			Descriptor: "rampant",
			SingleOnly: false,
			Tags: []string{
				"pegasus",
				"monster",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("pegasus-rampant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "ram-rampant",
			Name:       "ram rampant",
			Noun:       "ram",
			NounPlural: "rams",
			Descriptor: "rampant",
			SingleOnly: false,
			Tags: []string{
				"ram",
				"animal",
				"hunting",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("ram-rampant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "ram-statant",
			Name:       "ram statant",
			Noun:       "ram",
			NounPlural: "rams",
			Descriptor: "statant",
			SingleOnly: false,
			Tags: []string{
				"ram",
				"animal",
				"hunting",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("ram-statant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "rams-head-caboshed",
			Name:       "ram's head caboshed",
			Noun:       "ram's head",
			NounPlural: "ram's heads",
			Descriptor: "caboshed",
			SingleOnly: false,
			Tags: []string{
				"ram",
				"animal",
				"hunting",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("rams-head-caboshed", bodyTincture)

				return img
			},
		},
		{
			Identifier: "sea-horse",
			Name:       "sea horse",
			Noun:       "sea horse",
			NounPlural: "sea horses",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"sea horse",
				"animal",
				"water",
				"navy",
				"sailor",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("sea-horse", bodyTincture)

				return img
			},
		},
		{
			Identifier: "sheep-passant",
			Name:       "sheep passant",
			Noun:       "sheep",
			NounPlural: "sheeps",
			Descriptor: "passant",
			SingleOnly: false,
			Tags: []string{
				"sheep",
				"animal",
				"passive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("sheep-passant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "squirrel",
			Name:       "squirrel",
			Noun:       "squirrel",
			NounPlural: "squirrels",
			Descriptor: "",
			SingleOnly: false,
			Tags: []string{
				"squirrel",
				"animal",
				"passive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("squirrel", bodyTincture)

				return img
			},
		},
		{
			Identifier: "stag-at-gaze",
			Name:       "stag at gaze",
			Noun:       "stag",
			NounPlural: "stags",
			Descriptor: "at gaze",
			SingleOnly: false,
			Tags: []string{
				"stag",
				"animal",
				"hunting",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("stag-at-gaze", bodyTincture)

				return img
			},
		},
		{
			Identifier: "stag-lodged",
			Name:       "stag lodged",
			Noun:       "stag",
			NounPlural: "stags",
			Descriptor: "lodged",
			SingleOnly: false,
			Tags: []string{
				"stag",
				"animal",
				"hunting",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("stag-lodged", bodyTincture)

				return img
			},
		},
		{
			Identifier: "stag-springing",
			Name:       "stag springing",
			Noun:       "stag",
			NounPlural: "stags",
			Descriptor: "springing",
			SingleOnly: false,
			Tags: []string{
				"stag",
				"animal",
				"hunting",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("stag-springing", bodyTincture)

				return img
			},
		},
		{
			Identifier: "stag-statant",
			Name:       "stag statant",
			Noun:       "stag",
			NounPlural: "stags",
			Descriptor: "statant",
			SingleOnly: false,
			Tags: []string{
				"stag",
				"animal",
				"hunting",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("stag-statant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "stag-trippant",
			Name:       "stag trippant",
			Noun:       "stag",
			NounPlural: "stags",
			Descriptor: "trippant",
			SingleOnly: false,
			Tags: []string{
				"stag",
				"animal",
				"hunting",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("stag-trippant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "stags-head-couped",
			Name:       "stag's head couped'",
			Noun:       "stag's head",
			NounPlural: "stag's heads",
			Descriptor: "couped",
			SingleOnly: false,
			Tags: []string{
				"stag",
				"animal",
				"hunting",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("stags-head-couped", bodyTincture)

				return img
			},
		},
		{
			Identifier: "talbot-passant",
			Name:       "talbot passant",
			Noun:       "talbot",
			NounPlural: "talbots",
			Descriptor: "passant",
			SingleOnly: false,
			Tags: []string{
				"talbot",
				"animal",
				"passive",
				"hunting",
				"dog",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("talbot-passant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "talbot-rampant",
			Name:       "talbot rampant",
			Noun:       "talbot",
			NounPlural: "talbots",
			Descriptor: "rampant",
			SingleOnly: false,
			Tags: []string{
				"talbot",
				"animal",
				"aggressive",
				"hunting",
				"dog",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("talbot-rampant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "talbot-sejant",
			Name:       "talbot sejant",
			Noun:       "talbot",
			NounPlural: "talbots",
			Descriptor: "sejant",
			SingleOnly: false,
			Tags: []string{
				"talbot",
				"animal",
				"passive",
				"hunting",
				"dog",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("talbot-sejant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "talbot-statant",
			Name:       "talbot statant",
			Noun:       "talbot",
			NounPlural: "talbots",
			Descriptor: "statant",
			SingleOnly: false,
			Tags: []string{
				"talbot",
				"animal",
				"passive",
				"hunting",
				"dog",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("talbot-statant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "tiger-passant",
			Name:       "tiger passant",
			Noun:       "tiger",
			NounPlural: "tigers",
			Descriptor: "passant",
			SingleOnly: false,
			Tags: []string{
				"tiger",
				"animal",
				"cat",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("tiger-passant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "tiger-rampant",
			Name:       "tiger rampant",
			Noun:       "tiger",
			NounPlural: "tigers",
			Descriptor: "rampant",
			SingleOnly: false,
			Tags: []string{
				"tiger",
				"animal",
				"cat",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("tiger-rampant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "wolf-courant",
			Name:       "wolf courant",
			Noun:       "wolf",
			NounPlural: "wolfs",
			Descriptor: "courant",
			SingleOnly: false,
			Tags: []string{
				"wolf",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("wolf-courant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "wolf-passant",
			Name:       "wolf passant",
			Noun:       "wolf",
			NounPlural: "wolfs",
			Descriptor: "passant",
			SingleOnly: false,
			Tags: []string{
				"wolf",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("wolf-passant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "wolf-rampant",
			Name:       "wolf rampant",
			Noun:       "wolf",
			NounPlural: "wolfs",
			Descriptor: "rampant",
			SingleOnly: false,
			Tags: []string{
				"wolf",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("wolf-rampant", bodyTincture)

				return img
			},
		},
		{
			Identifier: "wolf-salient",
			Name:       "wolf salient",
			Noun:       "wolf",
			NounPlural: "wolfs",
			Descriptor: "salient",
			SingleOnly: false,
			Tags: []string{
				"wolf",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("wolf-salient", bodyTincture)

				return img
			},
		},
		{
			Identifier: "wolf-statant",
			Name:       "wolf statant",
			Noun:       "wolf",
			NounPlural: "wolfs",
			Descriptor: "statant",
			SingleOnly: false,
			Tags: []string{
				"wolf",
				"animal",
				"aggressive",
			},
			Render: func(bodyTincture tincture.Tincture) image.Image {
				img := RenderChargeFromFile("wolf-statant", bodyTincture)

				return img
			},
		},
	}

	return charges
}
