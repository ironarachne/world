package charge

import (
	svg "github.com/ajstarks/svgo"
)

var (
	OrdinaryBend = Charge{
		Identifier: "bend",
		Name:       "bend",
		Noun:       "bend",
		NounPlural: "bends",
		Descriptor: "",
		SingleOnly: true,
		Tags: []string{
			"ordinary",
		},
		Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
			canvas.TranslateRotate(-100, 135, -45)
			canvas.Rect(canvasWidth/3, 0, canvasWidth/3, canvasHeight, "fill:"+hexCode)
			canvas.Gend()
		},
	}
	OrdinaryBordure = Charge{
		Identifier: "bordure",
		Name:       "bordure",
		Noun:       "bordure",
		NounPlural: "bordures",
		Descriptor: "",
		SingleOnly: true,
		Tags: []string{
			"ordinary",
		},
		Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
			pathData := "m10.273 21.598v151.22c0 96.872 89.031 194.34 146.44 240.09 57.414-45.758 146.44-143.22 146.44-240.09v-151.22h-292.89z"
			canvas.Path(pathData, "stroke:"+hexCode+";stroke-width:100;fill:none")
		},
	}
	OrdinaryChevron = Charge{
		Identifier: "chevron",
		Name:       "chevron",
		Noun:       "chevron",
		NounPlural: "chevrons",
		Descriptor: "",
		SingleOnly: true,
		Tags: []string{
			"ordinary",
		},
		Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
			chevronHalfWidth := 40
			centerX := canvasWidth / 2
			canvas.Polygon(
				[]int{0, centerX, canvasWidth, canvasWidth, centerX, 0},
				[]int{canvasHeight - canvasHeight/4, canvasHeight - canvasHeight/3 - (3 * chevronHalfWidth), canvasHeight - canvasHeight/4, canvasHeight - canvasHeight/4 + (2 * chevronHalfWidth), canvasHeight - canvasHeight/3 - chevronHalfWidth, canvasHeight - canvasHeight/4 + (2 * chevronHalfWidth)},
				"fill:"+hexCode)
		},
	}
	OrdinaryChief = Charge{
		Identifier: "chief",
		Name:       "chief",
		Noun:       "chief",
		NounPlural: "chiefs",
		Descriptor: "",
		SingleOnly: true,
		Tags: []string{
			"ordinary",
		},
		Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
			canvas.Rect(0, 0, canvasWidth, canvasHeight/3, "fill:"+hexCode)
		},
	}
	OrdinaryCross = Charge{
		Identifier: "cross",
		Name:       "cross",
		Noun:       "cross",
		NounPlural: "crosses",
		Descriptor: "",
		SingleOnly: true,
		Tags: []string{
			"ordinary",
		},
		Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
			crossHalfWidth := 40
			centerX := canvasWidth / 2
			centerY := canvasHeight / 2
			canvas.Polygon(
				[]int{centerX - crossHalfWidth, centerX + crossHalfWidth, centerX + crossHalfWidth, canvasWidth, canvasWidth, centerX + crossHalfWidth, centerX + crossHalfWidth, centerX - crossHalfWidth, centerX - crossHalfWidth, 0, 0, centerX - crossHalfWidth},
				[]int{0, 0, centerY - crossHalfWidth, centerY - crossHalfWidth, centerY + crossHalfWidth, centerY + crossHalfWidth, canvasHeight, canvasHeight, centerY + crossHalfWidth, centerY + crossHalfWidth, centerY - crossHalfWidth, centerY - crossHalfWidth},
				"fill:"+hexCode)
		},
	}
	OrdinaryFess = Charge{
		Identifier: "fess",
		Name:       "fess",
		Noun:       "fess",
		NounPlural: "fesses",
		Descriptor: "",
		SingleOnly: true,
		Tags: []string{
			"ordinary",
		},
		Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
			canvas.Rect(0, canvasHeight/3, canvasWidth, canvasHeight/3, "fill:"+hexCode)
		},
	}
	OrdinaryLozenge = Charge{
		Identifier: "lozenge",
		Name:       "lozenge",
		Noun:       "lozenge",
		NounPlural: "lozenges",
		Descriptor: "",
		SingleOnly: false,
		Tags: []string{
			"ordinary",
		},
		Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
			centerX := canvasWidth / 2
			centerY := canvasHeight / 2
			lozengeHalfWidth := 80
			canvas.Polygon(
				[]int{centerX, centerX + lozengeHalfWidth, centerX, centerX - lozengeHalfWidth},
				[]int{centerY - lozengeHalfWidth - lozengeHalfWidth/2, centerY, centerY + lozengeHalfWidth + lozengeHalfWidth/2, centerY},
				"fill:"+hexCode)
		},
	}
	OrdinaryPale = Charge{
		Identifier: "pale",
		Name:       "pale",
		Noun:       "pale",
		NounPlural: "pales",
		Descriptor: "",
		SingleOnly: false,
		Tags: []string{
			"ordinary",
		},
		Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
			if number > 1 {
				canvas.Rect(canvasWidth/3, 0, canvasWidth/3, canvasHeight*2, "fill:"+hexCode)
			} else {
				canvas.Rect(canvasWidth/3, 0, canvasWidth/3, canvasHeight, "fill:"+hexCode)
			}
		},
	}
	OrdinaryPall = Charge{
		Identifier: "pall",
		Name:       "pall",
		Noun:       "pall",
		NounPlural: "palls",
		Descriptor: "",
		SingleOnly: true,
		Tags: []string{
			"ordinary",
		},
		Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
			centerX := canvasWidth / 2
			centerY := canvasHeight / 2
			pallHalfWidth := 40
			canvas.Polygon(
				[]int{0, pallHalfWidth, centerX, canvasWidth - pallHalfWidth, canvasWidth, canvasWidth, centerX + pallHalfWidth - pallHalfWidth/3, centerX + pallHalfWidth - pallHalfWidth/3, centerX - pallHalfWidth + pallHalfWidth/3, centerX - pallHalfWidth + pallHalfWidth/3, 0},
				[]int{0, 0, centerY - pallHalfWidth, 0, 0, pallHalfWidth, centerY + pallHalfWidth - pallHalfWidth/3, canvasHeight, canvasHeight, centerY + pallHalfWidth - pallHalfWidth/3, pallHalfWidth},
				"fill:"+hexCode)
		},
	}
	OrdinaryPile = Charge{
		Identifier: "pile",
		Name:       "pile",
		Noun:       "pile",
		NounPlural: "piles",
		Descriptor: "",
		SingleOnly: true,
		Tags: []string{
			"ordinary",
		},
		Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
			canvas.Polygon(
				[]int{0, canvasWidth, canvasWidth / 2},
				[]int{0, 0, canvasHeight / 2},
				"fill:"+hexCode)
		},
	}
	OrdinaryRoundel = Charge{
		Identifier: "roundel",
		Name:       "roundel",
		Noun:       "roundel",
		NounPlural: "roundels",
		Descriptor: "",
		SingleOnly: false,
		Tags: []string{
			"ordinary",
		},
		Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
			roundelRadius := 100
			canvas.Circle(
				canvasWidth/2,
				canvasHeight/2,
				roundelRadius,
				"fill:"+hexCode)
		},
	}
	OrdinarySaltire = Charge{
		Identifier: "saltire",
		Name:       "saltire",
		Noun:       "saltire",
		NounPlural: "saltires",
		Descriptor: "",
		SingleOnly: true,
		Tags: []string{
			"ordinary",
		},
		Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
			centerX := canvasWidth / 2
			centerY := canvasHeight / 2
			saltireHalfWidth := 30
			canvas.Polygon(
				[]int{0, saltireHalfWidth, centerX, canvasWidth - saltireHalfWidth, canvasWidth, canvasWidth, centerX + saltireHalfWidth, canvasWidth, canvasWidth, canvasWidth - saltireHalfWidth, centerX, saltireHalfWidth, 0, 0, centerX - saltireHalfWidth, 0},
				[]int{0, 0, centerY - saltireHalfWidth, 0, 0, saltireHalfWidth, centerY, canvasHeight - saltireHalfWidth, canvasHeight, canvasHeight, centerY + saltireHalfWidth, canvasHeight, canvasHeight, canvasHeight - saltireHalfWidth, centerY, saltireHalfWidth},
				"fill:"+hexCode)
		},
	}
)
