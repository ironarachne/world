package heraldry

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	svg "github.com/ajstarks/svgo"
	"github.com/divan/num2words"
)

// RenderToBlazon renders a device as its blazon and returns it.
func (device Device) RenderToBlazon() string {
	blazon := ""
	if device.Field.Division.Name == "plain" {
		if device.Field.HasVariation {
			blazon = strings.Title(device.Field.Variation.Name)
			blazon += " " + device.Field.Variation.Tincture1.Name + " and " + device.Field.Variation.Tincture2.Name
		} else {
			blazon = strings.Title(device.Field.Tincture.Name)
		}
	} else {
		blazon = device.Field.Division.Blazon
		if device.Field.HasVariation {
			blazon += device.Field.Variation.Name
			blazon += " " + device.Field.Variation.Tincture1.Name + " and " + device.Field.Variation.Tincture2.Name
		} else {
			blazon += device.Field.Tincture.Name
		}
	}

	if device.Field.Division.Name != "plain" {
		blazon += " and " + device.Field.Division.Tincture.Name
	}

	if len(device.ChargeGroups) > 0 {
		numberOfCharges := ""
		for _, chargeGroup := range device.ChargeGroups {
			numberOfCharges = num2words.Convert(len(chargeGroup.Charges))
			if len(chargeGroup.Charges) > 1 {
				blazon += ", " + numberOfCharges + " " + chargeGroup.Charges[0].NounPlural + " " + chargeGroup.Charges[0].Descriptor + " " + chargeGroup.Tincture.Name
			} else {
				blazon += ", " + chargeGroup.Charges[0].Article + " " + chargeGroup.Charges[0].Noun + " " + chargeGroup.Charges[0].Descriptor + " " + chargeGroup.Tincture.Name
			}
		}
	}

	blazon = strings.Replace(blazon, "  ", " ", -1)

	return blazon
}

// RenderToSVG renders a device as SVG and outputs as a string
func (device Device) RenderToSVG(width int, height int) string {
	buffer := new(bytes.Buffer)

	centerX := int(width / 2)
	centerY := int(height / 2)

	lineColor := "#000000"

	blazon := device.RenderToBlazon()

	canvas := svg.New(buffer)
	canvas.Start(width, height+50)
	canvas.Title(blazon)
	canvas.Def()
	canvas.Mask("shieldmask", 0, 0, width, height)
	canvas.Path("m10.273 21.598v151.22c0 96.872 89.031 194.34 146.44 240.09 57.414-45.758 146.44-143.22 146.44-240.09v-151.22h-292.89z", "fill:#FFFFFF")
	canvas.MaskEnd()
	for _, tincture := range device.AllTinctures {
		if tincture.Name == "erminois" {
			insertErmine(canvas, "erminois")
		} else if tincture.Name == "ermine" {
			insertErmine(canvas, "ermine")
		} else if tincture.Name == "ermines" {
			insertErmine(canvas, "ermines")
		} else if tincture.Name == "pean" {
			insertErmine(canvas, "pean")
		}
	}
	if device.Field.HasVariation {
		insertVariationPattern(canvas, device.Field.Variation)
	}
	canvas.DefEnd()
	canvas.Group("mask='url(#shieldmask)'")

	if device.Field.HasVariation {
		canvas.Rect(0, 0, width, height, "fill:url(#"+device.Field.Variation.Name+")")
	} else {
		canvas.Rect(0, 0, width, height, "fill:"+device.Field.Tincture.Hexcode)
	}

	switch device.Field.Division.Name {
	case "plain":
	case "pale":
		canvas.Rect(int(width/2), 0, int(width/2), height, "fill:"+device.Field.Division.Tincture.Hexcode)
	case "fess":
		canvas.Rect(0, 0, width, int(height/2), "fill:"+device.Field.Division.Tincture.Hexcode)
	case "bend":
		canvas.Polygon([]int{0, 0, width}, []int{0, height, height}, "fill:"+device.Field.Division.Tincture.Hexcode)
	case "bendsinister":
		canvas.Polygon([]int{0, width, 0}, []int{0, 0, height}, "fill:"+device.Field.Division.Tincture.Hexcode)
	case "chevron":
		canvas.Polygon([]int{0, int(width / 2), width}, []int{height, int(height / 2), height}, "fill:"+device.Field.Division.Tincture.Hexcode)
	case "quarterly":
		canvas.Rect(int(width/2), 0, int(width/2), int(height/2), "fill:"+device.Field.Division.Tincture.Hexcode)
		canvas.Rect(0, int(height/2), int(width/2), int(height/2), "fill:"+device.Field.Division.Tincture.Hexcode)
	case "saltire":
		canvas.Polygon([]int{0, int(width / 2), 0}, []int{0, int(height / 2), height}, "fill:"+device.Field.Division.Tincture.Hexcode)
		canvas.Polygon([]int{width, int(width / 2), width}, []int{0, int(height / 2), height}, "fill:"+device.Field.Division.Tincture.Hexcode)
	}
	for _, chargeGroup := range device.ChargeGroups {
		if len(chargeGroup.Charges) == 2 {
			canvas.Scale(0.5)
		} else if len(chargeGroup.Charges) == 3 {
			canvas.Scale(0.5)
		}
		for index, charge := range chargeGroup.Charges {
			if len(chargeGroup.Charges) == 2 {
				if charge.Identifier == "pale" {
					if index == 0 {
						canvas.Translate(20, 0)
					} else if index == 1 {
						canvas.Translate(300, 0)
					}
				} else {
					if index == 0 {
						canvas.Translate(20, 200)
					} else if index == 1 {
						canvas.Translate(300, 200)
					}
				}
			} else if len(chargeGroup.Charges) == 3 {
				if charge.Identifier == "pale" {
					if index == 0 {
						canvas.Translate(-20, 0)
					} else if index == 1 {
						canvas.Translate(160, 0)
					} else if index == 2 {
						canvas.Translate(340, 00)
					}
				} else {
					if index == 0 {
						canvas.Translate(20, 80)
					} else if index == 1 {
						canvas.Translate(300, 80)
					} else if index == 2 {
						canvas.Translate(160, 400)
					}
				}
			}
			if chargeGroup.Tincture.Name == "sable" {
				lineColor = "#FFFFFF"
			}
			switch charge.Identifier {
			case "pale":
				if len(chargeGroup.Charges) > 1 {
					canvas.Rect(int(width/3), 0, int(width/3), height*2, "fill:"+chargeGroup.Tincture.Hexcode)
				} else {
					canvas.Rect(int(width/3), 0, int(width/3), height, "fill:"+chargeGroup.Tincture.Hexcode)
				}
			case "fess":
				canvas.Rect(0, int(height/3), width, int(height/3), "fill:"+chargeGroup.Tincture.Hexcode)
			case "cross":
				crossHalfWidth := 40
				canvas.Polygon(
					[]int{centerX - crossHalfWidth, centerX + crossHalfWidth, centerX + crossHalfWidth, width, width, centerX + crossHalfWidth, centerX + crossHalfWidth, centerX - crossHalfWidth, centerX - crossHalfWidth, 0, 0, centerX - crossHalfWidth},
					[]int{0, 0, centerY - crossHalfWidth, centerY - crossHalfWidth, centerY + crossHalfWidth, centerY + crossHalfWidth, height, height, centerY + crossHalfWidth, centerY + crossHalfWidth, centerY - crossHalfWidth, centerY - crossHalfWidth},
					"fill:"+chargeGroup.Tincture.Hexcode)
			case "bend":
				canvas.TranslateRotate(-100, 135, -45)
				canvas.Rect(int(width/3), 0, int(width/3), height, "fill:"+chargeGroup.Tincture.Hexcode)
				canvas.Gend()
			case "saltire":
				saltireHalfWidth := 30
				canvas.Polygon(
					[]int{0, saltireHalfWidth, centerX, width - saltireHalfWidth, width, width, centerX + saltireHalfWidth, width, width, width - saltireHalfWidth, centerX, saltireHalfWidth, 0, 0, centerX - saltireHalfWidth, 0},
					[]int{0, 0, centerY - saltireHalfWidth, 0, 0, saltireHalfWidth, centerY, height - saltireHalfWidth, height, height, centerY + saltireHalfWidth, height, height, height - saltireHalfWidth, centerY, saltireHalfWidth},
					"fill:"+chargeGroup.Tincture.Hexcode)
			case "chevron":
				chevronHalfWidth := 40
				canvas.Polygon(
					[]int{0, centerX, width, width, centerX, 0},
					[]int{height - int(height/4), height - int(height/3) - (3 * chevronHalfWidth), height - int(height/4), height - int(height/4) + (2 * chevronHalfWidth), height - int(height/3) - chevronHalfWidth, height - int(height/4) + (2 * chevronHalfWidth)},
					"fill:"+chargeGroup.Tincture.Hexcode)
			case "chief":
				canvas.Rect(0, 0, width, int(height/3), "fill:"+chargeGroup.Tincture.Hexcode)
			case "pile":
				canvas.Polygon(
					[]int{0, width, int(width / 2)},
					[]int{0, 0, int(height / 2)},
					"fill:"+chargeGroup.Tincture.Hexcode)
			case "pall":
				pallHalfWidth := 40
				canvas.Polygon(
					[]int{0, pallHalfWidth, centerX, width - pallHalfWidth, width, width, centerX + pallHalfWidth - int(pallHalfWidth/3), centerX + pallHalfWidth - int(pallHalfWidth/3), centerX - pallHalfWidth + int(pallHalfWidth/3), centerX - pallHalfWidth + int(pallHalfWidth/3), 0},
					[]int{0, 0, centerY - pallHalfWidth, 0, 0, pallHalfWidth, centerY + pallHalfWidth - int(pallHalfWidth/3), height, height, centerY + pallHalfWidth - int(pallHalfWidth/3), pallHalfWidth},
					"fill:"+chargeGroup.Tincture.Hexcode)
			case "bordure":
				renderBordureToSvg(canvas, chargeGroup.Tincture.Hexcode)
			case "lozenge":
				lozengeHalfWidth := 80
				canvas.Polygon(
					[]int{centerX, centerX + lozengeHalfWidth, centerX, centerX - lozengeHalfWidth},
					[]int{centerY - lozengeHalfWidth - int(lozengeHalfWidth/2), centerY, centerY + lozengeHalfWidth + int(lozengeHalfWidth/2), centerY},
					"fill:"+chargeGroup.Tincture.Hexcode)
			case "roundel":
				roundelRadius := 100
				canvas.Circle(
					int(width/2),
					int(height/2),
					roundelRadius,
					"fill:"+chargeGroup.Tincture.Hexcode)
			case "eagle-displayed":
				renderEagleDisplayedToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "dragon-passant":
				renderDragonPassantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "gryphon-passant":
				renderGryphonPassantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "fox-passant":
				renderFoxPassantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "antelope-passant":
				renderAntelopePassantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "antelope-rampant":
				renderAntelopeRampantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "bat-volant":
				renderBatVolantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "battleaxe":
				renderBattleaxeToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "bear-head-couped":
				renderBearHeadCoupedToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "bear-rampant":
				renderBearRampantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "bear-statant":
				renderBearStatantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "bee-volant":
				renderBeeVolantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "bell":
				renderBellToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "boar-head-erased":
				renderBoarHeadErasedToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "boar-passant":
				renderBoarPassantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "boar-rampant":
				renderBoarRampantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "bugle-horn":
				renderBugleHornToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "bull-passant":
				renderBullPassantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "bull-rampant":
				renderBullRampantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "castle":
				renderCastleToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "cock":
				renderCockToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "cockatrice":
				renderCockatriceToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "crown":
				renderCrownToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "dolphin-hauriant":
				renderDolphinHauriantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "double-headed-eagle-displayed":
				renderDoubleHeadedEagleDisplayedToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "dragon-rampant":
				renderDragonRampantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "eagles-head-erased":
				renderEaglesHeadErasedToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "fox-sejant":
				renderFoxSejantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "gryphon-segreant":
				renderGryphonSegreantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "hare-salient":
				renderHareSalientToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "hare":
				renderHareToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "heron":
				renderHeronToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "horse-passant":
				renderHorsePassantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "horse-rampant":
				renderHorseRampantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "leopard-passant":
				renderLeopardPassantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "lion-passant":
				renderLionPassantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "lion-rampant":
				renderLionRampantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "lions-head-erased":
				renderLionsHeadErasedToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "owl":
				renderOwlToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "pegasus-passant":
				renderPegasusPassantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "pegasus-rampant":
				renderPegasusRampantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "ram-rampant":
				renderRamRampantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "ram-statant":
				renderRamStatantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "rose":
				renderRoseToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "sea-horse":
				renderSeaHorseToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "serpent-nowed":
				renderSerpentNowedToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "squirrel":
				renderSquirrelToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "stag-lodged":
				renderStagLodgedToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "stag-statant":
				renderStagStatantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "sun-in-splendor":
				renderSunInSplendorToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "tiger-passant":
				renderTigerPassantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "tiger-rampant":
				renderTigerRampantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "tower":
				renderTowerToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "two-axes-in-saltire":
				renderTwoAxesInSaltireToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "two-bones-in-saltire":
				renderTwoBonesInSaltireToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "unicorn-rampant":
				renderUnicornRampantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "unicorn-statant":
				renderUnicornStatantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "wolf-passant":
				renderWolfPassantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "wolf-rampant":
				renderWolfRampantToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			case "wyvern":
				renderWyvernToSvg(canvas, chargeGroup.Tincture.Hexcode, lineColor)
			}
			if len(chargeGroup.Charges) > 1 {
				canvas.Gend()
			}
		}
		if len(chargeGroup.Charges) > 1 {
			canvas.Gend()
		}
	}

	canvas.Path("m10.273 21.598v151.22c0 96.872 89.031 194.34 146.44 240.09 57.414-45.758 146.44-143.22 146.44-240.09v-151.22h-292.89z", "stroke:#000000;stroke-width:4;fill:none")
	canvas.Gend()
	// canvas.Text(centerX, height+25, blazon, "font-size:"+blazonSize+"px;text-anchor:middle")
	canvas.End()

	return buffer.String()
}

// RenderToSVGFile renders a device as SVG and writes to fileName
func RenderToSVGFile(device Device, fileName string, width int, height int) {
	var writer io.WriteCloser
	writer, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	contents := device.RenderToSVG(width, height)

	io.WriteString(writer, contents)

	defer writer.Close()
}
