package field

import (
	"context"

	"github.com/ironarachne/world/pkg/geometry"
	"github.com/ironarachne/world/pkg/random"
)

// Type is a type of field
type Type struct {
	CenterPoint  geometry.Point `json:"center_point"`
	ImageWidth   int            `json:"image_width"`
	ImageHeight  int            `json:"image_height"`
	MaskWidth    int            `json:"mask_width"`
	MaskHeight   int            `json:"mask_height"`
	Name         string         `json:"name"`
	MaskFileName string         `json:"mask_file_name"`
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
			CenterPoint: geometry.Point{
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
			CenterPoint: geometry.Point{
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
			CenterPoint: geometry.Point{
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
			CenterPoint: geometry.Point{
				X: 179,
				Y: 215,
			},
		},
	}

	return fieldTypes
}

// RandomType returns a random field type
func RandomType(ctx context.Context) Type {
	fieldTypes := allTypes()

	fieldType := fieldTypes[random.Intn(ctx, len(fieldTypes))]

	return fieldType
}
