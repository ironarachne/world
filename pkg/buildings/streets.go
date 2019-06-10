package buildings

import (
	"github.com/ironarachne/world/pkg/random"
)

func getRandomStreetStyle() string {
	styles := []string{
		"narrow",
		"wide",
	}

	return random.String(styles)
}