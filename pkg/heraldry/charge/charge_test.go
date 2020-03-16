package charge

import (
	"testing"
)

func TestBlazonForCharge(t *testing.T) {
	blazon, err := blazonForCharge(1, "charge", "charges", "testing", "Or")
	if err != nil {
		t.Error("failed to generate blazon for singular charge")
	}
	if blazon != "a charge testing Or" {
		t.Error("failed to parse blazon for singular charge: got '" + blazon + "'")
	}

	blazon, err = blazonForCharge(2, "charge", "charges", "testing", "Or")
	if err != nil {
		t.Error("failed to generate blazon for two charges")
	}
	if blazon != "two charges testing Or" {
		t.Error("failed to parse blazon for two charges: got '" + blazon + "'")
	}
}
