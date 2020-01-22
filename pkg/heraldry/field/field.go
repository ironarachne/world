package field

import (
	"fmt"
	"github.com/ironarachne/world/pkg/heraldry/charge"
	"github.com/ironarachne/world/pkg/heraldry/division"
)

// Field is the field of a coat of arms
type Field struct {
	Division     division.Division `json:"division"`
	ChargeGroups []charge.Group    `json:"charge_groups"`
	FieldType    Type              `json:"field_type"`
}

// ByName returns a field by name
func ByName(name string) (Field, error) {
	var fieldType Type

	fieldTypes := allTypes()

	for _, t := range fieldTypes {
		if t.Name == name {
			fieldType = t
		}
	}

	d, err := division.Generate()
	if err != nil {
		err = fmt.Errorf("failed to generate heraldic field by name: %w", err)
		return Field{}, err
	}
	chargeGroup, err := charge.RandomGroup(d.Variations[0].Tinctures[0])
	if err != nil {
		err = fmt.Errorf("failed to generate heraldic field by name: %w", err)
		return Field{}, err
	}

	chargeGroups := []charge.Group{
		chargeGroup,
	}

	field := Field{
		Division:     d,
		ChargeGroups: chargeGroups,
		FieldType:    fieldType,
	}

	return field, nil
}

// Random returns a random field
func Random() (Field, error) {
	t := RandomType()

	d, err := division.Generate()
	if err != nil {
		err = fmt.Errorf("failed to generate random heraldic field: %w", err)
		return Field{}, err
	}

	chargeGroup, err := charge.RandomGroup(d.Variations[0].Tinctures[0])
	if err != nil {
		err = fmt.Errorf("failed to generate random heraldic field: %w", err)
		return Field{}, err
	}

	chargeGroups := []charge.Group{
		chargeGroup,
	}

	field := Field{
		Division:     d,
		ChargeGroups: chargeGroups,
		FieldType:    t,
	}

	return field, nil
}
