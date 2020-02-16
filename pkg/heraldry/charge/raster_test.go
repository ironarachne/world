package charge

import (
	"testing"
)

func TestBlazon(t *testing.T) {
	r := Raster{
		Identifier: "charge-testing",
		Name:       "charge testing",
		Noun:       "charge",
		NounPlural: "charges",
		Descriptor: "testing",
		SingleOnly: false,
		Tags: []string{
			"charge-testing",
		},
	}

	blazon, err := r.Blazon(1, "Or")
	if err != nil {
		t.Error("failed to generate blazon for raster charge")
	}

	if blazon != "a charge testing Or" {
		t.Error("failed to properly format blazon for raster charge: got '" + blazon + "'")
	}
}
