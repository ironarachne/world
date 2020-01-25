package field

import (
	"github.com/ironarachne/world/pkg/grid"
	"math/rand"
)

// Type is a type of field
type Type struct {
	CenterPoint  grid.Coordinate `json:"center_point"`
	ImageWidth   int             `json:"image_width"`
	ImageHeight  int             `json:"image_height"`
	MaskWidth    int             `json:"mask_width"`
	MaskHeight   int             `json:"mask_height"`
	Name         string          `json:"name"`
	MaskFileName string          `json:"mask_file_name"`
}

func allTypes() []Type {
	fieldTypes := []Type{
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

// RandomType returns a random field type
func RandomType() Type {
	fieldTypes := allTypes()

	fieldType := fieldTypes[rand.Intn(len(fieldTypes))]

	return fieldType
}
