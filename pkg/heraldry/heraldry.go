package heraldry

import (
	"github.com/google/uuid"
	"github.com/ironarachne/world/pkg/words"
)

// Device is the entire coat of arms
type Device struct {
	Blazon string
	Field
	ImageURL string
	FileName string
	GUID string
}

// Generate procedurally generates a random heraldic device and returns it.
func Generate() Device {
	field := randomField()

	d := Device{
		Field: field,
	}

	guid := uuid.New()
	d.GUID = guid.String()
	d.FileName = d.GUID + ".png"
	d.Blazon = d.RenderToBlazon()
	d.Blazon = words.CapitalizeFirst(d.Blazon)

	return d
}

// GenerateByFieldName generates a random heraldic device with a given field shape.
func GenerateByFieldName(name string) Device {
	field := fieldByName(name)

	d := Device{
		Field: field,
	}

	guid := uuid.New()
	d.GUID = guid.String()
	d.FileName = d.GUID + ".png"
	d.Blazon = d.RenderToBlazon()
	d.Blazon = words.CapitalizeFirst(d.Blazon)

	return d
}

