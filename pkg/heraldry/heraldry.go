package heraldry

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
)

// Tincture is a tincture
type Tincture struct {
	Type    string
	Name    string
	Hexcode string
}

// Charge is a charge
type Charge struct {
	Identifier string
	Name       string
	Noun       string
	NounPlural string
	Descriptor string
	Article    string
	SingleOnly bool
	Tags       []string
}

// ChargeGroup is a group of charges with a common tincture
type ChargeGroup struct {
	Charges []Charge
	Tincture
}

// Division is a division of the field
type Division struct {
	Name   string
	Blazon string
	Tincture
}

// Variation is a variation of the field
type Variation struct {
	Name      string
	Tincture1 Tincture
	Tincture2 Tincture
}

// Field is the field of a coat of arms
type Field struct {
	Division
	Tincture
	HasVariation bool
	Variation
}

// Device is the entire coat of arms
type Device struct {
	Field
	ChargeGroups []ChargeGroup
	AllTinctures []Tincture
}

// Heraldry is a rendered version of a device
type Heraldry struct {
	Blazon string
	Device string
}

func randomCharge() Charge {
	return AvailableCharges[rand.Intn(len(AvailableCharges))]
}

func randomDivision() Division {
	return AvailableDivisions[rand.Intn(len(AvailableDivisions))]
}

func randomTincture() Tincture {
	typeOfTincture := random.StringFromThresholdMap(tinctureChances)

	if typeOfTincture == "metal" {
		return randomTinctureMetal()
	} else if typeOfTincture == "color" {
		return randomTinctureColor()
	} else if typeOfTincture == "stain" {
		return randomTinctureStain()
	}

	return randomTinctureFur()
}

func randomTinctureColor() Tincture {
	t := Colors[rand.Intn(len(Colors))]
	return t
}

func randomTinctureFur() Tincture {
	t := Furs[rand.Intn(len(Furs))]
	return t
}

func randomTinctureMetal() Tincture {
	t := Metals[rand.Intn(len(Metals))]
	return t
}

func randomTinctureStain() Tincture {
	t := Stains[rand.Intn(len(Stains))]
	return t
}

func randomComplementaryTincture(t Tincture) Tincture {
	var availableTinctures []Tincture
	if t.Type == "color" || t.Type == "stain" {
		typeOfTincture := random.StringFromThresholdMap(colorOrStainChances)
		if typeOfTincture == "color" {
			for _, color := range Colors {
				if color.Name != t.Name {
					availableTinctures = append(availableTinctures, color)
				}
			}
		} else {
			for _, stain := range Stains {
				if stain.Name != t.Name {
					availableTinctures = append(availableTinctures, stain)
				}
			}
		}
	} else {
		for _, metal := range Metals {
			if metal.Name != t.Name {
				availableTinctures = append(availableTinctures, metal)
			}
		}
	}
	t2 := availableTinctures[rand.Intn(len(availableTinctures))]

	return t2
}

func randomContrastingTincture(t Tincture) Tincture {
	t2 := Tincture{}
	if t.Type == "metal" {
		typeOfTincture := random.StringFromThresholdMap(colorOrStainChances)
		if typeOfTincture == "color" {
			t2 = randomTinctureColor()
		} else {
			t2 = randomTinctureStain()
		}
	} else {
		t2 = randomTinctureMetal()
	}

	return t2
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
	fieldTincture1 := randomTincture()
	fieldTincture2 := randomComplementaryTincture(fieldTincture1)
	chargeTincture := randomContrastingTincture(fieldTincture1)

	var tinctures []Tincture

	fieldHasContrastingTinctures := false

	if rand.Intn(10) > 1 {
		fieldHasContrastingTinctures = true
	}

	if fieldHasContrastingTinctures {
		fieldTincture2 = randomContrastingTincture(fieldTincture1)
	}

	division := randomDivision()
	division.Tincture = fieldTincture2

	field := Field{Division: division, Tincture: fieldTincture1, HasVariation: false, Variation: Variation{}}

	if shallWeHaveAVariation() {
		variation := randomVariation()
		variation.Tincture1 = randomTincture()
		variation.Tincture2 = randomContrastingTincture(variation.Tincture1)
		tinctures = append(tinctures, variation.Tincture1, variation.Tincture2)

		field.HasVariation = true
		field.Variation = variation
	}

	tinctures = append(tinctures, fieldTincture1)
	tinctures = append(tinctures, fieldTincture2)

	var charges []Charge
	var chargeGroups []ChargeGroup

	if shallWeIncludeCharges() {
		charge := randomCharge()

		chargeCountRanking := rand.Intn(10)
		countOfCharges := 1
		if chargeCountRanking >= 9 && !charge.SingleOnly {
			countOfCharges = 3
		} else if chargeCountRanking >= 7 && chargeCountRanking < 9 && !charge.SingleOnly {
			countOfCharges = 2
		}
		for i := 0; i < countOfCharges; i++ {
			charges = append(charges, charge)
		}
		chargeGroup := ChargeGroup{charges, chargeTincture}
		tinctures = append(tinctures, chargeTincture)
		chargeGroups = append(chargeGroups, chargeGroup)
	}

	d := Device{Field: field, ChargeGroups: chargeGroups, AllTinctures: tinctures}

	return d
}

// GenerateHeraldry renders a heraldic device and returns it
func GenerateHeraldry() Heraldry {
	device := Generate()

	heraldry := Heraldry{
		Blazon: device.RenderToBlazon(),
		Device: device.RenderToSVG(320, 420),
	}

	return heraldry
}
