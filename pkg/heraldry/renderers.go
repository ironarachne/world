package heraldry

import (
	"fmt"
	"image"
	"os"
	"time"

	"github.com/fogleman/gg"
	"github.com/ironarachne/world/pkg/save"
	"github.com/ironarachne/world/pkg/slices"
)

// RenderToBlazon renders a device as its blazon and returns it.
func (device Device) RenderToBlazon() string {
	blazon := device.Field.Division.Blazon
	if len(device.Field.ChargeGroups[0].Charges) > 0 {
		blazon += ", " + device.Field.ChargeGroups[0].RenderBlazon()
	}

	return blazon
}

// RenderToPNG renders a device as PNG and returns the image
func (device Device) RenderToPNG() (string, error) {
	dataPath := os.Getenv("WORLDAPI_DATA_PATH")

	var cg image.Image

	shield, err := gg.LoadPNG(dataPath + "/images/fields/" + device.FieldType.MaskFileName)
	if err != nil {
		err = fmt.Errorf("Could not load field mask image "+device.FieldType.MaskFileName+": %w", err)
		return "", err
	}
	shieldBorder, err := gg.LoadPNG(dataPath + "/images/fields/" + device.FieldType.Name + "-lines.png")
	if err != nil {
		err = fmt.Errorf("Could not load field border image "+device.FieldType.Name+"-lines.png: %w", err)
		return "", err
	}
	width := device.FieldType.ImageWidth
	height := device.FieldType.ImageHeight

	dc := gg.NewContext(width, height)
	dc.SetRGB(255, 255, 255)
	dc.Fill()

	field := device.Field.Division.Render(width, height, device.Field.Variations)

	shieldMask := gg.NewContextForImage(shield)
	err = dc.SetMask(shieldMask.AsMask())
	if err != nil {
		err = fmt.Errorf("Could not set shield mask: %w", err)
		return "", err
	}
	dc.DrawImage(field, 0, 0)

	for _, g := range device.Field.ChargeGroups {
		cg = g.RenderPNG(width, height)

		if slices.StringIn("full size", g.Charges[0].Tags) {
			dc.DrawImage(cg, 0, 0)
		} else {
			dc.DrawImageAnchored(cg, device.FieldType.CenterPoint.X, device.FieldType.CenterPoint.Y, 0.5, 0.5)
		}
	}
	fieldContents := dc.Image()

	finalContext := gg.NewContextForImage(shieldBorder)
	finalContext.DrawImage(fieldContents, 0, 0)

	finalImage := finalContext.Image()

	layout := "2006/01/02"
	now := time.Now()
	directory := now.Format(layout)

	imageURL, err := save.PNG("images/heraldry/devices/"+directory, device.FileName, finalImage)
	if err != nil {
		err = fmt.Errorf("Could not save heraldic device: %w", err)
		return "", err
	}

	return imageURL, nil
}
