package heraldry

import (
	"github.com/ironarachne/world/pkg/grid"
	"github.com/ironarachne/world/pkg/heraldry/charge"
	"math/rand"
)

// Field is the field of a coat of arms
type Field struct {
	Division
	ChargeGroups []charge.Group
	FieldType FieldType
}

// FieldType is a type of field
type FieldType struct {
	CenterPoint grid.Coordinate
	ImageWidth int
	ImageHeight int
	MaskWidth int
	MaskHeight int
	Name string
	MaskFileName string
}

func allFieldTypes() []FieldType {
	fieldTypes := []FieldType{
		{
			Name: "banner",
			MaskFileName: "banner.png",
			ImageWidth: 320,
			ImageHeight: 450,
			MaskWidth: 320,
			MaskHeight: 450,
			CenterPoint: grid.Coordinate{
				X: 160,
				Y: 225,
			},
		},
		{
			Name: "engrailed",
			MaskFileName: "engrailed.png",
			ImageWidth: 374,
			ImageHeight: 450,
			MaskWidth: 374,
			MaskHeight: 450,
			CenterPoint: grid.Coordinate{
				X: 187,
				Y: 215,
			},
		},
		{
			Name: "heater",
			MaskFileName: "heater.png",
			ImageWidth: 364,
			ImageHeight: 436,
			MaskWidth: 364,
			MaskHeight: 436,
			CenterPoint: grid.Coordinate{
				X: 182,
				Y: 200,
			},
		},
		{
			Name: "wedge",
			MaskFileName: "wedge.png",
			ImageWidth: 358,
			ImageHeight: 450,
			MaskWidth: 358,
			MaskHeight: 450,
			CenterPoint: grid.Coordinate{
				X: 179,
				Y: 215,
			},
		},
	}

	return fieldTypes
}

func fieldByName(name string) Field {
	var fieldType FieldType

	fieldTypes := allFieldTypes()

	for _, t := range fieldTypes {
		if t.Name == name {
			fieldType = t
		}
	}

	division := generateDivision()
	chargeGroup := charge.RandomGroup(division.Variations[0].Tinctures[0])

	chargeGroups := []charge.Group{
		chargeGroup,
	}

	field := Field{
		Division: division,
		ChargeGroups: chargeGroups,
		FieldType: fieldType,
	}

	return field
}

func randomField() Field {
	fieldType := randomFieldType()

	division := generateDivision()
	chargeGroup := charge.RandomGroup(division.Variations[0].Tinctures[0])

	chargeGroups := []charge.Group{
		chargeGroup,
	}

	field := Field{
		Division: division,
		ChargeGroups: chargeGroups,
		FieldType: fieldType,
	}

	return field
}

func randomFieldType() FieldType {
	fieldTypes := allFieldTypes()

	fieldType := fieldTypes[rand.Intn(len(fieldTypes))]

	return fieldType
}

