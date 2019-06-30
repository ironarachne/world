package buildings

import (
	"github.com/ironarachne/world/pkg/random"
)

func getRandomDoorMaterial() string {
	materials := []string{
		"light wood",
		"thick wood",
		"thin wood",
		"heavy wood",
	}

	return random.String(materials)
}

func getRandomDoorStyle() string {
	style := "made of " + getRandomDoorMaterial() + " and " + getRandomDoorShape()

	return style
}

func getRandomDoorShape() string {
	shapes := []string{
		"braced with metal",
		"broad",
		"carved with simple shapes",
		"curved on top",
		"narrow",
		"often possessed of a window in the middle",
		"onion-shaped",
		"ornately carved",
		"rectangular",
		"rounded on top",
		"tall",
	}

	return random.String(shapes)
}
