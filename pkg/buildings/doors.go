package buildings

import (
	"github.com/ironarachne/world/pkg/random"
)

func getRandomDoorStyle() string {
	style := getRandomDoorShape()

	return style
}

func getRandomDoorShape() string {
	shapes := []string{
		"curved on top",
		"narrow",
		"often have a window in the middle",
		"onion-shaped",
		"ornately carved",
		"rectangular",
		"tall",
	}

	return random.String(shapes)
}