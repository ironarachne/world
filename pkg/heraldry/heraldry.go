package heraldry

import (
	"github.com/ironarachne/world/pkg/heraldry/charge"
	"github.com/ironarachne/world/pkg/heraldry/tincture"
	"math/rand"
)

// Variation is a variation of the field
type Variation struct {
	Name      string
	Tincture1 tincture.Tincture
	Tincture2 tincture.Tincture
}

// Field is the field of a coat of arms
type Field struct {
	Division
	tincture.Tincture
	HasVariation bool
	Variation
}

// Device is the entire coat of arms
type Device struct {
	Field
	ChargeGroups []charge.Group
	AllTinctures []tincture.Tincture
}

// Heraldry is a rendered version of a device
type Heraldry struct {
	Blazon    string
	Device    string
	Tinctures []string
	Patterns  []string
}

func shallWeIncludeCharges() bool {
	someRandomValue := rand.Intn(10)
	if someRandomValue > 2 {
		return true
	}
	return false
}

// Generate procedurally generates a random heraldic device and returns it.
func Generate() Device {
	fieldTincture1 := tincture.Random()
	fieldTincture2 := tincture.RandomComplementaryTincture(fieldTincture1)
	chargeTincture := tincture.RandomContrastingTincture(fieldTincture1)

	var tinctures []tincture.Tincture

	fieldHasContrastingTinctures := false

	if rand.Intn(10) > 1 {
		fieldHasContrastingTinctures = true
	}

	if fieldHasContrastingTinctures {
		fieldTincture2 = tincture.RandomContrastingTincture(fieldTincture1)
	}

	division := randomDivision()
	division.Tincture = fieldTincture2

	field := Field{Division: division, Tincture: fieldTincture1, HasVariation: false, Variation: Variation{}}

	if shallWeHaveAVariation() {
		variation := randomVariation()
		variation.Tincture1 = tincture.Random()
		variation.Tincture2 = tincture.RandomContrastingTincture(variation.Tincture1)
		tinctures = append(tinctures, variation.Tincture1, variation.Tincture2)

		field.HasVariation = true
		field.Variation = variation
	}

	tinctures = append(tinctures, fieldTincture1)
	tinctures = append(tinctures, fieldTincture2)

	var charges []charge.Charge
	var chargeGroups []charge.Group

	if shallWeIncludeCharges() {
		c := charge.Random()

		chargeCountRanking := rand.Intn(10)
		countOfCharges := 1
		if chargeCountRanking >= 9 && !c.SingleOnly {
			countOfCharges = 3
		} else if chargeCountRanking >= 7 && chargeCountRanking < 9 && !c.SingleOnly {
			countOfCharges = 2
		}
		for i := 0; i < countOfCharges; i++ {
			charges = append(charges, c)
		}
		chargeGroup := charge.Group{
			Charges: charges,
			Tincture: chargeTincture,
		}
		tinctures = append(tinctures, chargeTincture)
		chargeGroups = append(chargeGroups, chargeGroup)
	}

	d := Device{
		Field: field,
		ChargeGroups: chargeGroups,
		AllTinctures: tinctures,
	}

	return d
}

// GenerateHeraldry renders a heraldic device and returns it
func GenerateHeraldry() Heraldry {
	patterns := []string{}
	tinctures := []string{}
	device := Generate()

	for _, t := range device.AllTinctures {
		tinctures = append(tinctures, t.HexCode)
		if t.Type == "fur" {
			patterns = append(patterns, t.Name)
		}
	}

	heraldry := Heraldry{
		Blazon:    device.RenderToBlazon(),
		Device:    device.RenderToSVG(320, 420),
		Tinctures: tinctures,
		Patterns:  patterns,
	}

	return heraldry
}
