package heraldry

import (
	"bytes"
	"fmt"
	"github.com/ironarachne/world/pkg/heraldry/tincture"
	"github.com/ironarachne/world/pkg/save"
	"github.com/ironarachne/world/pkg/words"
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
				blazon += ", " + words.Pronoun(chargeGroup.Charges[0].Noun) + " " + chargeGroup.Charges[0].Noun + " " + chargeGroup.Charges[0].Descriptor + " " + chargeGroup.Tincture.Name
			}
		}
	}

	blazon = strings.Replace(blazon, "  ", " ", -1)

	return blazon
}

// RenderToSVG renders a device as SVG and outputs as a string
func (device Device) RenderToSVG(width int, height int) string {
	buffer := new(bytes.Buffer)

	lineColor := "#000000"

	blazon := device.RenderToBlazon()

	canvas := svg.New(buffer)
	canvas.Start(width, height+50)
	canvas.Title(blazon)
	canvas.Def()
	canvas.Mask("shieldmask", 0, 0, width, height)
	canvas.Path("m10.273 21.598v151.22c0 96.872 89.031 194.34 146.44 240.09 57.414-45.758 146.44-143.22 146.44-240.09v-151.22h-292.89z", "fill:#FFFFFF")
	canvas.MaskEnd()
	for _, t := range device.AllTinctures {
		if t.Type == "fur" {
			tincture.InsertErmine(canvas, t.Name)
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
		canvas.Rect(0, 0, width, height, "fill:"+device.Field.Tincture.HexCode)
	}

	switch device.Field.Division.Name {
	case "plain":
	case "pale":
		canvas.Rect(int(width/2), 0, int(width/2), height, "fill:"+device.Field.Division.Tincture.HexCode)
	case "fess":
		canvas.Rect(0, 0, width, int(height/2), "fill:"+device.Field.Division.Tincture.HexCode)
	case "bend":
		canvas.Polygon([]int{0, 0, width}, []int{0, height, height}, "fill:"+device.Field.Division.Tincture.HexCode)
	case "bendsinister":
		canvas.Polygon([]int{0, width, 0}, []int{0, 0, height}, "fill:"+device.Field.Division.Tincture.HexCode)
	case "chevron":
		canvas.Polygon([]int{0, int(width / 2), width}, []int{height, int(height / 2), height}, "fill:"+device.Field.Division.Tincture.HexCode)
	case "quarterly":
		canvas.Rect(int(width/2), 0, int(width/2), int(height/2), "fill:"+device.Field.Division.Tincture.HexCode)
		canvas.Rect(0, int(height/2), int(width/2), int(height/2), "fill:"+device.Field.Division.Tincture.HexCode)
	case "saltire":
		canvas.Polygon([]int{0, int(width / 2), 0}, []int{0, int(height / 2), height}, "fill:"+device.Field.Division.Tincture.HexCode)
		canvas.Polygon([]int{width, int(width / 2), width}, []int{0, int(height / 2), height}, "fill:"+device.Field.Division.Tincture.HexCode)
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
			charge.Render(canvas, chargeGroup.Tincture.HexCode, lineColor, width, height, 1)

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

// RenderToRemoteSVG renders a device to SVG and saves it remotely
func RenderToRemoteSVG(device Device, fileName string, width int, height int) {
	contents := device.RenderToSVG(width, height)

	save.SVG("heraldry/devices/"+fileName, contents)
}
