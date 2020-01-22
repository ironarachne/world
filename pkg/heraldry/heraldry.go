/*
Package heraldry implements methods for randomly generating heraldic coats of arms
and their associated blazons.
*/
package heraldry

import (
	"fmt"

	"github.com/google/uuid"
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
func Generate() (Device, error) {
	f, err := field.Random()
	if err != nil {
		err = fmt.Errorf("Failed to generate heraldic device: %w", err)
		return Device{}, err
	}

	d := Device{
		Field: f,
	}

	guid := uuid.New()
	d.GUID = guid.String()
	d.FileName = d.GUID + ".png"

	blazon, err := d.RenderToBlazon()
	if err != nil {
		err = fmt.Errorf("Failed to generate heraldic device: %w", err)
		return Device{}, err
	}
	d.Blazon = words.CapitalizeFirst(blazon)

	imageURL, err := d.RenderToPNG()
	if err != nil {
		err = fmt.Errorf("Failed to generate heraldic device: %w", err)
		return Device{}, err
	}
	d.ImageURL = imageURL

	return d, nil
}

// GenerateByFieldName generates a random heraldic device with a given field shape.
func GenerateByFieldName(name string) (Device, error) {
	f, err := field.ByName(name)
	if err != nil {
		err = fmt.Errorf("Failed to generate heraldic device: %w", err)
		return Device{}, err
	}

	d := Device{
		Field: f,
	}

	guid := uuid.New()
	d.GUID = guid.String()
	d.FileName = d.GUID + ".png"

	blazon, err := d.RenderToBlazon()
	if err != nil {
		err = fmt.Errorf("Failed to generate heraldic device: %w", err)
		return Device{}, err
	}
	d.Blazon = words.CapitalizeFirst(blazon)

	imageURL, err := d.RenderToPNG()
	if err != nil {
		err = fmt.Errorf("Failed to generate heraldic device: %w", err)
		return Device{}, err
	}
	d.ImageURL = imageURL

	return d, nil
}
