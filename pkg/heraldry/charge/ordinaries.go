package charge

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/ironarachne/world/pkg/heraldry/tincture"
	"image"
)

// Cross is a cross ordinary
type Cross struct{}

// Fess is a fess ordinary
type Fess struct{}

// Pale is a pale ordinary
type Pale struct{}

// Roundel is a roundel ordinary
type Roundel struct{}

// Blazon renders the blazon for a cross
func (c Cross) Blazon(count int, tincture string) (string, error) {
	blazon, err := blazonForCharge(count, "cross", "crosses", "", tincture)
	if err != nil {
		err = fmt.Errorf("failed to generate blazon for cross: %w", err)
		return "", err
	}

	return blazon, nil
}

// Blazon renders the blazon for a fess
func (f Fess) Blazon(count int, tincture string) (string, error) {
	blazon, err := blazonForCharge(count, "fess", "fesses", "", tincture)
	if err != nil {
		err = fmt.Errorf("failed to generate blazon for fess: %w", err)
		return "", err
	}

	return blazon, nil
}

// Blazon renders the blazon for a pale
func (p Pale) Blazon(count int, tincture string) (string, error) {
	blazon, err := blazonForCharge(count, "pale", "pales", "", tincture)
	if err != nil {
		err = fmt.Errorf("failed to generate blazon for pale: %w", err)
		return "", err
	}

	return blazon, nil
}

// Blazon renders the blazon for a roundel
func (r Roundel) Blazon(count int, tincture string) (string, error) {
	blazon, err := blazonForCharge(count, "roundel", "roundels", "", tincture)
	if err != nil {
		err = fmt.Errorf("failed to generate blazon for roundel: %w", err)
		return "", err
	}

	return blazon, nil
}

// GetTags returns all tags for the given cross
func (c Cross) GetTags() []string {
	tags := []string{
		"full size",
		"ordinary",
	}
	return tags
}

// GetTags returns all tags for the given fess
func (f Fess) GetTags() []string {
	tags := []string{
		"full size",
		"ordinary",
	}
	return tags
}

// GetTags returns all tags for the given pale
func (p Pale) GetTags() []string {
	tags := []string{
		"full size",
		"ordinary",
	}
	return tags
}

// GetTags returns all tags for the given roundel
func (r Roundel) GetTags() []string {
	tags := []string{
		"ordinary",
	}
	return tags
}

func (c Cross) Render(bodyTincture tincture.Tincture) image.Image {
	width := 350
	height := 450

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(110, 0, 130, float64(height))
	dc.DrawRectangle(0, 140, float64(width), 160)
	bodyTincture.Fill(dc)

	img := dc.Image()

	return img
}

func (f Fess) Render(bodyTincture tincture.Tincture) image.Image {
	width := 350
	height := 450

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(0, 140, float64(width), 160)
	bodyTincture.Fill(dc)

	img := dc.Image()

	return img
}

func (p Pale) Render(bodyTincture tincture.Tincture) image.Image {
	width := 350
	height := 450

	dc := gg.NewContext(width, height)

	dc.DrawRectangle(110, 0, 130, float64(height))
	bodyTincture.Fill(dc)

	img := dc.Image()

	return img
}

func (r Roundel) Render(bodyTincture tincture.Tincture) image.Image {
	width := 400
	height := 400

	dc := gg.NewContext(width, height)

	dc.DrawCircle(float64(width/2), float64(height/2), float64(width/3))
	bodyTincture.Fill(dc)

	img := dc.Image()

	return img
}

func getOrdinaryCharges() []Charge {
	c := Cross{}
	f := Fess{}
	p := Pale{}
	r := Roundel{}

	charges := []Charge{
		c,
		f,
		p,
		r,
	}

	return charges
}
