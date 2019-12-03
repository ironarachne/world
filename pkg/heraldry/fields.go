package heraldry

import (
	"fmt"
	"github.com/ironarachne/world/pkg/grid"
	"github.com/ironarachne/world/pkg/heraldry/charge"
	"math/rand"
)

// Field is the field of a coat of arms
type Field struct {
	Division     `json:"division"`
	ChargeGroups []charge.Group `json:"charge_groups"`
	FieldType    FieldType      `json:"field_type"`
}

// FieldType is a type of field
type FieldType struct {
	CenterPoint  grid.Coordinate `json:"center_point"`
	ImageWidth   int             `json:"image_width"`
	ImageHeight  int             `json:"image_height"`
	MaskWidth    int             `json:"mask_width"`
	MaskHeight   int             `json:"mask_height"`
	Name         string          `json:"name"`
	MaskFileName string          `json:"mask_file_name"`
}

func allFieldTypes() []FieldType {
	fieldTypes := []FieldType{
		{
			Name:         "banner",
			MaskFileName: "banner.png",
			ImageWidth:   320,
			ImageHeight:  450,
			MaskWidth:    320,
			MaskHeight:   450,
			CenterPoint: grid.Coordinate{
				X: 160,
				Y: 225,
			},
		},
		{
			Name:         "engrailed",
			MaskFileName: "engrailed.png",
			ImageWidth:   374,
			ImageHeight:  450,
			MaskWidth:    374,
			MaskHeight:   450,
			CenterPoint: grid.Coordinate{
				X: 187,
				Y: 215,
			},
		},
		{
			Name:         "heater",
			MaskFileName: "heater.png",
			ImageWidth:   364,
			ImageHeight:  436,
			MaskWidth:    364,
			MaskHeight:   436,
			CenterPoint: grid.Coordinate{
				X: 182,
				Y: 200,
			},
		},
		{
			Name:         "wedge",
			MaskFileName: "wedge.png",
			ImageWidth:   358,
			ImageHeight:  450,
			MaskWidth:    358,
			MaskHeight:   450,
			CenterPoint: grid.Coordinate{
				X: 179,
				Y: 215,
			},
		},
	}

	return fieldTypes
}

func fieldByName(name string) (Field, error) {
	var fieldType FieldType

	fieldTypes := allFieldTypes()

	for _, t := range fieldTypes {
		if t.Name == name {
			fieldType = t
		}
	}

	division, err := generateDivision()
	if err != nil {
		err = fmt.Errorf("Failed to generate heraldic field: %w", err)
		return Field{}, err
	}
	chargeGroup, err := charge.RandomGroup(division.Variations[0].Tinctures[0])
	if err != nil {
		err = fmt.Errorf("Failed to generate heraldic field: %w", err)
		return Field{}, err
	}

	chargeGroups := []charge.Group{
		chargeGroup,
	}

	field := Field{
		Division:     division,
		ChargeGroups: chargeGroups,
		FieldType:    fieldType,
	}

	return field, nil
}

func randomField() (Field, error) {
	fieldType := randomFieldType()

	division, err := generateDivision()
	if err != nil {
		err = fmt.Errorf("Failed to generate heraldic field: %w", err)
		return Field{}, err
	}
	chargeGroup, err := charge.RandomGroup(division.Variations[0].Tinctures[0])
	if err != nil {
		err = fmt.Errorf("Failed to generate heraldic field: %w", err)
		return Field{}, err
	}

	chargeGroups := []charge.Group{
		chargeGroup,
	}

	field := Field{
		Division:     division,
		ChargeGroups: chargeGroups,
		FieldType:    fieldType,
	}

	return field, nil
}

func randomFieldType() FieldType {
	fieldTypes := allFieldTypes()

	fieldType := fieldTypes[rand.Intn(len(fieldTypes))]

	return fieldType
}
