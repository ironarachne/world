/*
Package heraldry implements methods for randomly generating heraldic coats of arms
and their associated blazons.
*/
package heraldry

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/ironarachne/world/pkg/heraldry/charge"
	"github.com/ironarachne/world/pkg/heraldry/field"
	"github.com/ironarachne/world/pkg/words"
)

// Device is the entire coat of arms
type Device struct {
	Blazon   string      `json:"blazon"`
	Field    field.Field `json:"field"`
	ImageURL string      `json:"image_url"`
	FileName string      `json:"file_name"`
	GUID     string      `json:"guid"`
}

// Generate procedurally generates a random heraldic device and returns it.
func Generate(ctx context.Context) (Device, error) {
	d, err := GenerateByParameters(ctx, "", "")
	if err != nil {
		err = fmt.Errorf("failed to generate random heraldic device: %w", err)
		return Device{}, err
	}

	return d, nil
}

// GenerateByParameters generates a random heraldic device fitting certain criteria.
func GenerateByParameters(ctx context.Context, fieldName string, chargeTag string) (Device, error) {
	var chargeGroup charge.Group
	var f field.Field
	var err error

	if fieldName == "" {
		f, err = field.Random(ctx)
	} else {
		f, err = field.ByName(ctx, fieldName)
	}
	if err != nil {
		err = fmt.Errorf("failed to generate heraldic device: %w", err)
		return Device{}, err
	}

	if chargeTag == "" {
		chargeGroup, err = charge.RandomGroup(ctx, f.Division.Variations[0].Tinctures[0])
	} else {
		chargeGroup, err = charge.RandomGroupByParameters(ctx, f.Division.Variations[0].Tinctures[0], chargeTag, 0)
	}
	if err != nil {
		err = fmt.Errorf("failed to generate heraldic device: %w", err)
		return Device{}, err
	}

	chargeGroups := []charge.Group{
		chargeGroup,
	}

	f.ChargeGroups = chargeGroups

	d := Device{
		Field: f,
	}

	d, err = d.finalize(ctx)
	if err != nil {
		err = fmt.Errorf("failed to generate heraldic device: %w", err)
		return Device{}, err
	}

	return d, nil
}

func (d Device) finalize(ctx context.Context) (Device, error) {
	guid := uuid.New()
	d.GUID = guid.String()
	d.FileName = d.GUID + ".png"

	blazon, err := d.RenderToBlazon()
	if err != nil {
		err = fmt.Errorf("failed to finalize heraldic device: %w", err)
		return Device{}, err
	}
	d.Blazon = words.CapitalizeFirst(blazon)

	imageURL, err := d.RenderToPNG(ctx)
	if err != nil {
		err = fmt.Errorf("failed to finalize heraldic device: %w", err)
		return Device{}, err
	}
	d.ImageURL = imageURL

	return d, nil
}
