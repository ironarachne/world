package heraldry

import (
	"context"
	"fmt"
	"image"
	"os"
	"time"

	"github.com/fogleman/gg"
	"github.com/ironarachne/world/pkg/save"
	"github.com/ironarachne/world/pkg/slices"
)

// RenderToBlazon renders a device as its blazon and returns it.
func (device Device) RenderToBlazon() (string, error) {
	blazon := device.Field.Division.Blazon
	if len(device.Field.ChargeGroups[0].Charges) > 0 {
		chargeBlazon, err := device.Field.ChargeGroups[0].RenderBlazon()
		if err != nil {
			err = fmt.Errorf("failed to render charge group to blazon: %w", err)
			return "", err
		}
		blazon += ", " + chargeBlazon
	}

	return blazon, nil
}

// RenderToPNG renders a device as PNG and returns the image
func (device Device) RenderToPNG(ctx context.Context) (string, error) {
	dataPath := os.Getenv("WORLDAPI_DATA_PATH")

	var cg image.Image

	shield, err := gg.LoadPNG(dataPath + "/images/fields/" + device.Field.FieldType.MaskFileName)
	if err != nil {
		err = fmt.Errorf("failed to load field mask image "+device.Field.FieldType.MaskFileName+": %w", err)
		return "", err
	}

	shieldBorder, err := gg.LoadPNG(dataPath + "/images/fields/" + device.Field.FieldType.Name + "-lines.png")
	if err != nil {
		err = fmt.Errorf("failed to load field border image "+device.Field.FieldType.Name+"-lines.png: %w", err)
		return "", err
	}
	width := device.Field.FieldType.ImageWidth
	height := device.Field.FieldType.ImageHeight

	dc := gg.NewContext(width, height)
	dc.SetRGB(255, 255, 255)
	dc.Fill()

	field := device.Field.Division.Render(width, height, device.Field.Division.Variations)

	shieldMask := gg.NewContextForImage(shield)
	err = dc.SetMask(shieldMask.AsMask())
	if err != nil {
		err = fmt.Errorf("failed to set shield mask: %w", err)
		return "", err
	}
	dc.DrawImage(field, 0, 0)

	for _, g := range device.Field.ChargeGroups {
		cg = g.RenderPNG(ctx, width, height)

		if slices.StringIn("full size", g.Charges[0].GetTags()) {
			dc.DrawImage(cg, 0, 0)
		} else {
			dc.DrawImageAnchored(cg, int(device.Field.FieldType.CenterPoint.X), int(device.Field.FieldType.CenterPoint.Y), 0.5, 0.5)
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
		err = fmt.Errorf("failed to save heraldic device image: %w", err)
		return "", err
	}

	return imageURL, nil
}
